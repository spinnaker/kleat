/*
 * Copyright The Spinnaker Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"errors"
	"fmt"
	"github.com/ezimanyi/kleat/api/client"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"path/filepath"

	"github.com/ghodss/yaml"
)

func printConfig() {
	fn := os.Args[1]
	h := parseHalConfig(fn)
	if err := validateHalConfig(h); err != nil {
		panic(err)
	}
	if err := printHalConfig(*h); err != nil {
		panic(err)
	}
}

func main() {
	hal := os.Args[1]
	dir := os.Args[2]
	if err := printConfigs(hal, dir); err != nil {
		panic(err)
	}
}

func printConfigs(hal string, d string) error {
	if err := ensureFile(hal); err != nil {
		return err
	}
	h := parseHalConfig(hal)
	if err := validateHalConfig(h); err != nil {
		panic(err)
	}
	if err := ensureDirectory(d); err != nil {
		return err
	}

	c, _ := halToClouddriver(*h)
	f, err := os.Create(filepath.Join(d, "clouddriver.yml"))
	if err != nil {
		return err
	}
	if err = printObject(c, f); err != nil {
		return err
	}
	return nil
}

func ensureDirectory(d string) error {
	stat, err := os.Stat(d)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return errors.New(fmt.Sprintf("%s is not a directory", d))
	}
	return nil
}

func ensureFile(f string) error {
	stat, err := os.Stat(f)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return errors.New(fmt.Sprintf("%s is a directory, not a file", f))
	}
	return nil
}

type HalConfigValidator func(*proto.HalConfig) []ValidationFailure

func getValidators() []HalConfigValidator {
	return []HalConfigValidator{
		validateKindsAndOmitKinds,
	}
}

func validateHalConfig(h *proto.HalConfig) error {
	messages := getValidationMessages(h, getValidators())
	if len(messages) > 0 {
		msg := strings.Join(messages, "\n")
		return errors.New(msg)
	}
	return nil
}

func getValidationMessages(h *proto.HalConfig, fa []HalConfigValidator) []string {
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
	msg string
}

func fatalResult(msg string) ValidationFailure {
	return ValidationFailure{
		msg: msg,
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

func printHalConfig(h proto.HalConfig) error {
	c, _ := halToClouddriver(h)
	if err := printObject(c, os.Stdout); err != nil {
		return err
	}
	return nil
	//todo: test whether enum providerVersion marshals correctly
}

func printObject(i interface{}, w io.Writer) error {
	bytes, err := yaml.Marshal(i)
	if err != nil {
		return err
	}
	if _, err = w.Write(bytes); err != nil {
		return err
	}
	return nil
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
		Kubernetes:     h.Providers.Kubernetes,
		Google:         h.Providers.Google,
		Appengine:      h.Providers.Appengine,
		Aws:            h.Providers.Aws,
		Azure:          h.Providers.Azure,
		Cloudfoundry:   h.Providers.Cloudfoundry,
		Dcos:           h.Providers.Dcos,
		DockerRegistry: h.Providers.DockerRegistry,
		Ecs:            h.Providers.Ecs,
		Huaweicloud:    h.Providers.Huaweicloud,
		Oracle:         h.Providers.Oracle,
		Artifacts:      h.Artifacts,
	}
	return c, nil
}
