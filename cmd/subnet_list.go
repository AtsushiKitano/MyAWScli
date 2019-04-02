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
	subnetCmd.AddCommand(subnetListCmd)
}

var subnetListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		createSubnetTable(getSubnets())
	},
}

func getSubnets() [][]string {
	var subnetLists [][]string
	sess := session.Must(session.NewSession())
	input := &ec2.DescribeSubnetsInput{}
	svc := ec2.New(
		sess,
		aws.NewConfig().WithRegion(os.Getenv("AWS_DEFAULT_REGION")),
	)

	result, err := svc.DescribeSubnets(input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, r := range result.Subnets {
		var tmp_infos []string
		if r.Tags != nil {
			tmp_infos = append(tmp_infos, *r.Tags[0].Value)
		} else {
			tmp_infos = append(tmp_infos, "-")
		}
		tmp_infos = append(tmp_infos, *r.AvailabilityZone)
		tmp_infos = append(tmp_infos, *r.CidrBlock)
		tmp_infos = append(tmp_infos, *r.VpcId)
		subnetLists = append(subnetLists, tmp_infos)
	}

	return subnetLists
}

func createSubnetTable(subnetLists [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "AZ", "CIDR", "VPC"})

	for _, t := range subnetLists {
		table.Append(t)
	}

	table.Render()
}
