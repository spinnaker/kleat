package main

import (
	"github.com/ezimanyi/kleat/proto"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ghodss/yaml"
)

func main() {
	fn := os.Args[1]
	h := parseHalConfig(fn)
	messages := validateHalConfig(h, getValidators())
	if len(messages) > 0 {
		msg := strings.Join(messages, "\n")
		panic(msg)
	}
	printHalConfig(*h)
}

type HalConfigValidator func(*proto.HalConfig) []ValidationFailure

func getValidators() []HalConfigValidator {
	return []HalConfigValidator{
		validateKindsAndOmitKinds,
	}
}

func validateHalConfig(h *proto.HalConfig, fa []HalConfigValidator) []string {
	var messages []string
	for _, f := range fa {
		rs := f(h)
		for _, r := range rs {
			messages = append(messages, r.msg)
		}
	}
	return messages
}

func validateKindsAndOmitKinds(h *proto.HalConfig) []ValidationFailure {
	var messages []ValidationFailure
	for _, a := range h.Providers.Kubernetes.Accounts {
		if !(len(a.Kinds) == 0) && !(len(a.OmitKinds) == 0) {
			messages = append(messages, fatalResult("Cannot specify both kinds and omitKinds."))
		}
	}
	return messages
}

type ValidationFailure struct {
	msg   string
}

func fatalResult(msg string) ValidationFailure {
	return ValidationFailure{
		msg:   msg,
	}
}

func parseHalConfig(fn string) *proto.HalConfig {
	dat, err := ioutil.ReadFile(fn)

	h := proto.HalConfig{}
	err = yaml.Unmarshal([]byte(dat), &h)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &h
}

func printHalConfig(h proto.HalConfig) {
	//Test converting halconfig to front50config
	f, _ := halToFront50(h)
	printObject(f)

	c, _ := halToClouddriver(h)
	printObject(c)

	//todo: test whether enum providerVersion marshals correctly
}

func printObject(i interface{}) {
	bytes, err := yaml.Marshal(i)
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func getTestHalConfig() proto.HalConfig {
	h := proto.HalConfig{
		Providers: &proto.HalConfig_Providers{
			Kubernetes: &proto.Kubernetes{
				Enabled: false,
				Accounts: []*proto.Kubernetes_Account{
					{
						Name:            "hal",
						ProviderVersion: proto.Kubernetes_V2,
					},
				},
			},
		},
	}
	return h
}

func halToFront50(h proto.HalConfig) (proto.Front50Config, error) {
	f := proto.Front50Config{
		Spinnaker: &proto.Front50Config_Spinnaker{
			PersistentStoreType: h.PersistentStorage.PersistentStoreType,
			Gcs:                 h.PersistentStorage.Gcs,
		},
	}
	return f, nil
}

func extractPersistentStoreType(h proto.HalConfig) *string {
	if h.PersistentStorage == nil {
		return nil
	}
	return &h.PersistentStorage.PersistentStoreType
}

func halToClouddriver(h proto.HalConfig) (proto.ClouddriverConfig, error) {
	c := proto.ClouddriverConfig{
		Kubernetes: h.Providers.Kubernetes,
	}
	return c, nil
}
