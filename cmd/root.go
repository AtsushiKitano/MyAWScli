package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use: "aws",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aws command")
	},
}

func Execute() {
	if err := awsCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
