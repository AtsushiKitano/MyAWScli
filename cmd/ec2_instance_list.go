package cmd

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	ec2Cmd.AddCommand(instanceListCmd)
}

var instanceListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		makeEc2InstancesList(getEc2InstancesList())
	},
}

func getEc2InstancesList() [][]string {
	var ec2InstaceList [][]string
	sess := session.Must(session.NewSession())
	input := &ec2.DescribeInstancesInput{}
	svc := ec2.New(
		sess,
		aws.NewConfig().WithRegion(os.Getenv("AWS_DEFAULT_REGION")),
	)
	result, err := svc.DescribeInstances(input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			var tmp_infos []string
			tmp_infos = append(tmp_infos, *i.Tags[0].Value)
			tmp_infos = append(tmp_infos, *i.PublicIpAddress)
			tmp_infos = append(tmp_infos, *i.PrivateIpAddress)
			tmp_infos = append(tmp_infos, *i.State.Name)
			tmp_infos = append(tmp_infos, *i.InstanceType)
			tmp_infos = append(tmp_infos, *i.Placement.AvailabilityZone)
			tmp_infos = append(tmp_infos, *i.RootDeviceType)
			ec2InstaceList = append(ec2InstaceList, tmp_infos)
		}
	}
	return ec2InstaceList
}

func makeEc2InstancesList(ec2List [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "PublicIP", "PrivateIP", "State", "InstanceType", "AZ", "DeviceType"})

	for _, v := range ec2List {
		table.Append(v)
	}
	table.Render()
}
