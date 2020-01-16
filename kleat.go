package main

import (
	"github.com/ezimanyi/kleat/proto"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	h := proto.HalConfig{}
	fn := os.Args[1]

	dat, err := ioutil.ReadFile(fn)

	err = yaml.Unmarshal([]byte(dat), &h)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	bytes, err := yaml.Marshal(h)
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(bytes)
	if err != nil {
		panic(err)
	}
}
