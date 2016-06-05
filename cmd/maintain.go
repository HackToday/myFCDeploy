// Copyright © 2016 Kai Qiang Wu <wkq5325@gmail.com>
// This file is part of {{ .appName }}.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// maintainCmd represents the maintain command
var maintainCmd = &cobra.Command{
	Use:   "maintain",
	Short: "This is used to maintain the cloud",
	Long:  `This is used to maintain the cloud native services to help upgrade case`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("maintain called")
	},
}

func init() {
	RootCmd.AddCommand(maintainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// maintainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// maintainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
