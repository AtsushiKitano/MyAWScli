package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	ec2Cmd.AddCommand(volumeListCmd)
}

var volumeListCmd = &cobra.Command{
	Use: "volume-list",
	Run: func(cmd *cobra.Command, args []string) {
		createVolumeList(getVolumeList())
	},
}

func getVolumeList() [][]string {
	var volumeList [][]string
	sess := session.Must(session.NewSession())
	input := &ec2.DescribeVolumesInput{}
	svc := ec2.New(
		sess,
		aws.NewConfig().WithRegion(os.Getenv("AWS_DEFAULT_REGION")),
	)

	result, err := svc.DescribeVolumes(input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, r := range result.Volumes {
		var tmp_infos []string
		if r.Tags == nil {
			tmp_infos = append(tmp_infos, "-")
		}
		for _, i := range r.Tags {
			if *i.Key == "Name" {
				tmp_infos = append(tmp_infos, *i.Value)
			}
		}
		tmp_infos = append(tmp_infos, strconv.FormatInt(*r.Size, 10))
		tmp_infos = append(tmp_infos, *r.VolumeType)
		tmp_infos = append(tmp_infos, *r.State)
		tmp_infos = append(tmp_infos, *r.AvailabilityZone)
		if *r.State == "in-use" {
			tmp_infos = append(tmp_infos, Ec2Id2Name(*r.Attachments[0].InstanceId, svc))
		} else {
			tmp_infos = append(tmp_infos, "-")
		}

		if r.Iops != nil {
			tmp_infos = append(tmp_infos, strconv.FormatInt(*r.Iops, 10))
		} else {
			tmp_infos = append(tmp_infos, "-")
		}
		volumeList = append(volumeList, tmp_infos)
	}

	return volumeList
}

func createVolumeList(volumesList [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Size", "VolumeType", "State", "AZ", "Attached Instance", "IOPS"})

	for _, v := range volumesList {
		table.Append(v)
	}
	table.Render()
}

func Ec2Id2Name(ec2Id string, svc *ec2.EC2) string {
	input := &ec2.DescribeInstancesInput{}
	result, err := svc.DescribeInstances(input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ec2Name string
	ec2Name = ""
	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			if *i.InstanceId == ec2Id {
				for _, t := range i.Tags {
					if *t.Key == "Name" {
						ec2Name = *t.Value
					}
				}

			}
		}
	}
	if ec2Name == "" {
		ec2Name = "-"
	}

	return ec2Name
}
