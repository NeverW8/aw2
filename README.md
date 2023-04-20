# Aw2

This is a command-line tool for finding EC2 instances in your AWS account.

## How to Use


## How It Works

This tool uses the AWS SDK for Go to interact with the EC2 API. When you run the `aw2` command, it makes a DescribeInstances API call with optional filters based on the command-line arguments. It then parses the response to extract details for each instance, such as its public and private IP addresses, instance type, and status. Finally, it prints the instance details in a pretty table using the `tablewriter` package.


## Contributions

Feel free to submit a pr

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/example/ec2-instance-finder/blob/main/LICENSE) file for details.
