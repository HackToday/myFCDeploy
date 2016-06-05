// Copyright © 2016 Kai Qiang Wu <wkq5325@gmail.com>
// This file is part of myFCDeploy

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "This is used to scale cloud",
	Long: `This is used to scale cloud apps quickly to make cloud native
service to respond to emergent changing requirements`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("scale called")
	},
}

func init() {
	RootCmd.AddCommand(scaleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scaleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scaleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
