package plugins

import (
	"fmt"
	"github.com/armory/plugin-tools/pkg/validate"
	"github.com/spinnaker/kleat/api/client/config"
)

// CollectCompatibilityErrors collects the supplied *config.Hal using validate.PluginCompatibilityValidator,
//returning only validate.NotCompatible messages.
func CollectCompatibilityErrors(h *config.Hal, validator validate.PluginCompatibilityValidator) []string {
	repos := make([]validate.PluginRepository, 0)
	var messages []string
	for _, v := range h.Spinnaker.Extensibility.Repositories {
		repos = append(repos, validate.PluginRepository{v.Id, v.Url})
	}

	for k, v := range h.Spinnaker.Extensibility.Plugins {
		verdict, _ := validator.IsPluginCompatibleWithPlatform(h.Version, k, v.Version, repos)
		if verdict == validate.NotCompatible {
			messages = append(messages, generateCompatibilityMessage(verdict, k, v.Version, h.Version, nil))
		}
	}
	return messages
}

// CollectCompatibilityIssues collects the supplied *config.Hal using validate.PluginCompatibilityValidator,
//returning any validate.NotCompatible & validate.Unknown messages.
func CollectCompatibilityIssues(h *config.Hal, validator validate.PluginCompatibilityValidator) []string {
	repos := make([]validate.PluginRepository, 0)
	var messages []string
	for _, v := range h.Spinnaker.Extensibility.Repositories {
		repos = append(repos, validate.PluginRepository{v.Id, v.Url})
	}

	for k, v := range h.Spinnaker.Extensibility.Plugins {
		verdict, err := validator.IsPluginCompatibleWithPlatform(h.Version, k, v.Version, repos)
		if err != nil {
			messages = append(messages, generateCompatibilityMessage(verdict, k, v.Version, h.Version, err))
		}
		if verdict == validate.NotCompatible {
			messages = append(messages, generateCompatibilityMessage(verdict, k, v.Version, h.Version, nil))
		}
	}
	return messages
}

// CollectCompatibilityMessages collects the supplied *config.Hal using validate.PluginCompatibilityValidator,
//returning all messages.
func CollectCompatibilityMessages(h *config.Hal, validator validate.PluginCompatibilityValidator) []string {
	repos := make([]validate.PluginRepository, 0)
	var messages []string
	for _, v := range h.Spinnaker.Extensibility.Repositories {
		repos = append(repos, validate.PluginRepository{v.Id, v.Url})
	}

	for k, v := range h.Spinnaker.Extensibility.Plugins {
		verdict, err := validator.IsPluginCompatibleWithPlatform(h.Version, k, v.Version, repos)
		messages = append(messages, generateCompatibilityMessage(verdict, k, v.Version, h.Version, err))
	}
	return messages
}

func generateCompatibilityMessage(verdict validate.Verdict, pluginName string, pluginVersion string, spinVersion string, err error) string {
	switch verdict {
	case validate.Compatible:
		return fmt.Sprintf("Plugin %s:%s is Compatible with Spinnaker %s\n", pluginName, pluginVersion, spinVersion)
	case validate.NotCompatible:
		return fmt.Sprintf("Plugin %s:%s is Incompatible with Spinnaker %s", pluginName, pluginVersion, spinVersion)
	default:
		return fmt.Sprintf("Unable to determine if plugin %s:%s is compatible with Spinnaker %s, Reason: %s", pluginName, pluginVersion, spinVersion, err.Error())
	}
}
