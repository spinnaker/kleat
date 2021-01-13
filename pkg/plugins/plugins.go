package plugins

import (
	"fmt"
	"github.com/armory/plugin-tools/pkg/validate"
	"github.com/spinnaker/kleat/api/client/config"
)

type compatibilitySummary struct {
	verdict validate.Verdict
	msg     string
}

//ValidatePluginsCompatibility validates the supplied *config.Hal, returning any Incompatible plugin error.
func ValidatePluginsCompatibility(h *config.Hal) []string {
	results := getCompatibilityResults(h)
	var errMsg []string
	for _, v := range results {
		if v.verdict == validate.NotCompatible { //summary with Status NotCompatible are treated as error
			errMsg = append(errMsg, v.msg)
		}
		//TODO: implement action for Unknown status
	}
	return errMsg
}

//GetCompatibilityIssues validates the supplied *config.Hal, returning any issue found with configured plugin.
func GetCompatibilityIssues(h *config.Hal) []string {
	results := getCompatibilityResults(h)
	var issues []string
	for _, v := range results {
		if v.verdict != validate.Compatible {
			issues = append(issues, v.msg)
		}
	}
	return issues
}

func getCompatibilityResults(h *config.Hal) []compatibilitySummary {
	validator := validate.NewAstrolabeValidator()
	repos := make([]validate.PluginRepository, 0)
	var summary []compatibilitySummary
	for _, v := range h.Spinnaker.Extensibility.Repositories {
		repos = append(repos, validate.PluginRepository{v.Id, v.Url})
	}

	for k, v := range h.Spinnaker.Extensibility.Plugins {
		verdict, err := validator.IsPluginCompatibleWithPlatform(h.Version, k, v.Version, repos)
		if err != nil {
			summary = append(summary, mapVerdictMessage(verdict, k, v.Version, h.Version, err.Error()))
		} else {
			summary = append(summary, mapVerdictMessage(verdict, k, v.Version, h.Version, ""))
		}
	}
	return summary
}

func mapVerdictMessage(verdict validate.Verdict, pluginName string, pluginVersion string, spinVersion string, reason string) compatibilitySummary {
	switch verdict {
	case validate.Compatible:
		return compatibilitySummary{
			verdict,
			fmt.Sprintf("Plugin %s:%s is Compatible with Spinnaker %s\n", pluginName, pluginVersion, spinVersion),
		}
	case validate.NotCompatible:
		return compatibilitySummary{
			verdict,
			fmt.Sprintf("Plugin %s:%s is Incompatible with Spinnaker %s", pluginName, pluginVersion, spinVersion),
		}
	default:
		return compatibilitySummary{
			verdict,
			fmt.Sprintf("Unable to determine if plugin %s:%s is compatible with Spinnaker %s, Reason: %s", pluginName, pluginVersion, spinVersion, reason),
		}
	}
}
