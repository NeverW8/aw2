# Aw2

![build workflow](https://github.com/NeverW8/aw2/actions/workflows/go.yml/badge.svg)
![Code scan workflow](https://github.com/NeverW8/aw2/actions/workflows/code_checks.yml/badge.svg)

This is a command-line tool for finding EC2 instances in your AWS account.

## How to Use

Download from the release tab, follow those instructions and start using it!

* When aw2 is used without an argument, it will print every ec2 instance on your aws account.
* When aw2 is used with an argument, it will filter for what you've supplied. (IP adress or ec2 name i.e)

## How It Works

This tool uses the AWS SDK for Go to interact with the EC2 API. When you run the `aw2` command, it makes a DescribeInstances API call with optional filters based on the command-line arguments. It then parses the response to extract details for each instance, such as its public and private IP addresses, instance type, and status. Finally, it prints the instance details in a pretty table using the `tablewriter` package.


## Contributers
<a href="https://github.com/neverw8/aw2/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=neverw8/aw2" />
</a>

- [NeverW8](https://github.com/neverw8)
- [Lapponiandevil](https://github.com/Lapponiandevil)

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/NeverW8/aw2/blob/main/LICENSE) file for details.
