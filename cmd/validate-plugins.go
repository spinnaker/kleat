package cmd

import (
	"github.com/armory/plugin-tools/pkg/validate"
	"github.com/spf13/cobra"
	"github.com/spinnaker/kleat/internal/fileio"
	"github.com/spinnaker/kleat/pkg/plugins"
)

var validatePlugCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate compatibility of configured plugins.",
	Long: `Given a path to your top-level Spinnaker config (halconfig), takes
every plugin configured in the Spinnaker config and the top-level version of
Spinnaker, it verify if there is a compatibility issue between the plugin version
and Spinnaker version and output the result.

Example usage:

kleat plugins validate /path/to/halconfig`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hal := args[0]
		issues, err := validateCompatibility(hal)
		if err != nil {
			return err
		}
		if len(issues) > 0 {
			for _, m := range issues {
				cmd.Println(m)
			}
		} else {
			cmd.Println("No plugin compatibility issue was found")
		}
		return nil
	},
}

func validateCompatibility(halPath string) ([]string, error) {
	h, err := fileio.ParseHalConfig(halPath)
	if err != nil {
		return nil, err
	}
	validator := validate.NewAstrolabeValidator()
	return plugins.GetCompatibilityIssues(h, validator), nil
}

func init() {
	pluginsCmd.AddCommand(validatePlugCmd)
}
