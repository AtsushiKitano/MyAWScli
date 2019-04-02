package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	awsCmd.AddCommand(ec2Cmd)
	awsCmd.AddCommand(vpcCmd)
	awsCmd.AddCommand(subnetCmd)
}

var ec2Cmd = &cobra.Command{
	Use: "ec2",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ec2 command")
	},
}

var vpcCmd = &cobra.Command{
	Use: "vpc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vpc command")
	},
}

var subnetCmd = &cobra.Command{
	Use: "subnet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subnet command")
	},
}
