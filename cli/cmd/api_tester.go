package cmd

import (
	"github.com/spf13/cobra"
)

var (
	apiCmd = &cobra.Command{
		Use:       "api",
		Short:     "Start testing api server",
		Long:      `Start testing api server`,
		ValidArgs: validArgs,
		Run:       testingAPI,
	}
)

func testingAPI(cmd *cobra.Command, args []string) {

}
