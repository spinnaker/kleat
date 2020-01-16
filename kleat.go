package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

var data = `
persistentStorage:
  gcs:
    jsonPath: hello
`

func main() {
	h := HalConfig{}

	err := yaml.Unmarshal([]byte(data), &h)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	bytes, err := json.Marshal(h)
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(bytes)
	if err != nil {
		panic(err)
	}
}
