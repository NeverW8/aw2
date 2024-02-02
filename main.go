package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/charmbracelet/lipgloss"
	"github.com/olekukonko/tablewriter"
)

func main() {
	sess := session.Must(session.NewSession())
	svc := ec2.New(sess)

	filters := []*ec2.Filter{}
	if len(os.Args) > 1 {
		filterText := os.Args[1]
		filterKey := "tag:Name"
		useIpFilter, _ := regexp.MatchString(`^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){1,4})`, filterText)
		if useIpFilter {
			filterKey = "private-ip-address"
		}

		filter := &ec2.Filter{
			Name: aws.String(filterKey),
			Values: []*string{
				aws.String("*" + filterText + "*"),
			},
		}
		filters = append(filters, filter)
	}

	var resp *ec2.DescribeInstancesOutput
	var err error
	if len(filters) > 0 {
		resp, err = svc.DescribeInstances(&ec2.DescribeInstancesInput{
			Filters: filters,
		})
	} else {
		resp, err = svc.DescribeInstances(nil)
	}

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	instances := []*InstanceDetails{}
	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			instanceDetails := &InstanceDetails{
				PublicIP:  aws.StringValue(instance.PublicIpAddress),
				PrivateIP: aws.StringValue(instance.PrivateIpAddress),
				Name:      getNameTagValue(instance.Tags),
				Type:      aws.StringValue(instance.InstanceType),
				Status:    aws.StringValue(instance.State.Name),
				VpcId:     aws.StringValue(instance.VpcId),
			}
			instances = append(instances, instanceDetails)
		}
	}

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF8800")).
		Background(lipgloss.Color("#000000")).
		Padding(0, 1)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		style.Render("Public IP"),
		style.Render("Private IP"),
		style.Render("Name"),
		style.Render("Type"),
		style.Render("Status"),
		style.Render("VpcId"),
	})

	for _, instance := range instances {
		row := []string{
			style.Render(instance.PublicIP),
			style.Render(instance.PrivateIP),
			style.Render(instance.Name),
			style.Render(instance.Type),
			style.Render(instance.Status),
			style.Render(instance.VpcId),
		}
		table.Append(row)
	}
	table.Render()
}

type InstanceDetails struct {
	PublicIP  string
	PrivateIP string
	Name      string
	Type      string
	Status    string
	VpcId     string
}

func getNameTagValue(tags []*ec2.Tag) string {
	for _, tag := range tags {
		if aws.StringValue(tag.Key) == "Name" {
			return aws.StringValue(tag.Value)
		}
	}
	return ""
}
