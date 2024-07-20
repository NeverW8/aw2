package main

import (
    "fmt"
    "os"
    "regexp"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/olekukonko/tablewriter"
)

const version = "v1.0.2"

func main() {
    if len(os.Args) > 1 && os.Args[1] == "--version" {
        fmt.Println("aw2 version:", version)
        return
    }
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
                PublicIP:   aws.StringValue(instance.PublicIpAddress),
                PrivateIP:  aws.StringValue(instance.PrivateIpAddress),
                Name:       getNameTagValue(instance.Tags),
                Type:       aws.StringValue(instance.InstanceType),
                Status:     aws.StringValue(instance.State.Name),
                VpcId:      aws.StringValue(instance.VpcId),
            }
            instances = append(instances, instanceDetails)
        }
    }

    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"Public IP", "Private IP", "Name", "Type", "Status", "VpcId"})
    for _, instance := range instances {
        row := []string{instance.PublicIP, instance.PrivateIP, instance.Name, instance.Type, instance.Status, instance.VpcId}
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
