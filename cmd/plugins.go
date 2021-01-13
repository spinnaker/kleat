package cmd

import (
	"github.com/spf13/cobra"
)

var pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "Show Spinnaker's configured plugins.",
	Long:  "Provide subcommands to configure and validate Spinnaker plugins.",
}

func init() {
	rootCmd.AddCommand(pluginsCmd)
}
