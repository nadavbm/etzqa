package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of etzqa",
	Long:  `Print the version number of the application from git tags`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 0.0.1")
	},
}
