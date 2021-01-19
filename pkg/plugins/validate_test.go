package plugins_test

import (
	"fmt"
	"github.com/armory/plugin-tools/pkg/validate"
	"github.com/spinnaker/kleat/internal/fileio"
	"github.com/spinnaker/kleat/pkg/plugins"
	"path/filepath"
	"testing"
)

const dataDir = "../../testdata"

type validatorMock struct{}

func (v *validatorMock) IsPluginCompatibleWithPlatform(platformVersion string, pluginId string, pluginVersion string, repos []validate.PluginRepository) (validate.Verdict, error) {
	return v.mockResultByPluginId(pluginId)
}

func (v *validatorMock) IsPluginCompatibleWithService(serviceName string, serviceVersion string, pluginId string, pluginVersion string, repos []validate.PluginRepository) (validate.Verdict, error) {
	return v.mockResultByPluginId(pluginId)
}

func (v *validatorMock) mockResultByPluginId(pluginId string) (validate.Verdict, error) {
	switch pluginId {
	case "Test.Compatible":
		return validate.Compatible, nil
	case "Test.Incompatible":
		return validate.NotCompatible, nil
	default:
		return validate.Unknown, fmt.Errorf("mock error")
	}
}

func TestOnlyIncompatibleMessagesAreReturned(t *testing.T) {
	h, err := fileio.ParseHalConfig(filepath.Join(dataDir, "halconfig_with_plugins.yml"))
	if err != nil {
		t.Fatal(err)
	}
	mock := &validatorMock{}
	errMsg := plugins.CollectCompatibilityErrors(h, mock)
	if len(errMsg) != 1 {
		t.Errorf("A single error message was expected but found %d", len(errMsg))
	}
}

func TestIncompatibleAndUnknownMessagesAreReturned(t *testing.T) {
	h, err := fileio.ParseHalConfig(filepath.Join(dataDir, "halconfig_with_plugins.yml"))
	if err != nil {
		t.Fatal(err)
	}
	mock := &validatorMock{}
	errMsg := plugins.CollectCompatibilityIssues(h, mock)
	if len(errMsg) != 2 {
		t.Errorf("Two error messages were expected but found %d", len(errMsg))
	}
}

func TestAllMessagesAreReturned(t *testing.T) {
	h, err := fileio.ParseHalConfig(filepath.Join(dataDir, "halconfig_with_plugins.yml"))
	if err != nil {
		t.Fatal(err)
	}
	mock := &validatorMock{}
	errMsg := plugins.CollectCompatibilityMessages(h, mock)
	if len(errMsg) != 3 {
		t.Errorf("All messages were expected but found %d", len(errMsg))
	}
}
