package main

import (
	"github.com/ezimanyi/kleat/proto"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	fn := os.Args[1]

	dat, err := ioutil.ReadFile(fn)

	h := proto.HalConfig{}
	err = yaml.Unmarshal([]byte(dat), &h)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Test converting halconfig to front50config
	f, err := halToFront50(h)
	printObject(f)

	c, err := halToClouddriver(h)
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

func halToFront50(h proto.HalConfig) (proto.Front50Config, error) {
	f := proto.Front50Config{
		Spinnaker: &proto.Front50Config_Spinnaker{
			PersistentStoreType: h.PersistentStorage.PersistentStoreType,
			Gcs: h.PersistentStorage.Gcs,
		},
	}
	return f, nil
}

func halToClouddriver(h proto.HalConfig) (proto.ClouddriverConfig, error) {
	c := proto.ClouddriverConfig{
		Kubernetes:           h.Providers.Kubernetes,
	}
	return c, nil
}
