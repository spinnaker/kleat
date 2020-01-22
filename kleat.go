package main

import (
	"errors"
	"github.com/ezimanyi/kleat/proto"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	h := parseHalConfig()
	err := validateHalConfig(h)
	if err != nil {
		panic(err)
	}
	printHalConfig(*h)
}

func validateHalConfig(h *proto.HalConfig) error {
	err := validateKindsAndOmitKinds(h)
	return err
}

func validateKindsAndOmitKinds(h *proto.HalConfig) error {
	for _, a := range h.Providers.Kubernetes.Accounts {
		if (!(len(a.Kinds) == 0) && !(len(a.OmitKinds) == 0)) {
			return errors.New("Cannot specify both kinds and omitKinds.")
		}
	}
	return nil
}

func parseHalConfig() *proto.HalConfig {
	fn := os.Args[1]

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
				Enabled:              false,
				Accounts:             []*proto.Kubernetes_Account{
					&proto.Kubernetes_Account{
						Name:                 "hal",
						ProviderVersion:      proto.Kubernetes_V2,
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
