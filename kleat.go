package main

import (
	"encoding/json"
	"os"
)

func main() {
	p := HalConfig{}

	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(bytes)
	if err != nil {
		panic(err)
	}
}
