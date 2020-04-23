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
package parse_hal

import (
	"github.com/spinnaker/kleat/api/client"
	"io/ioutil"
	"sigs.k8s.io/yaml"
)

func ParseHalConfig(fn string) (*client.HalConfig, error) {
	dat, err := ioutil.ReadFile(fn)

	h := client.HalConfig{}
	err = yaml.Unmarshal([]byte(dat), &h)
	if err != nil {
		return nil, err
	}

	return &h, nil
}

func HalToFront50(h client.HalConfig) (client.Front50Config, error) {
	f := client.Front50Config{
		Spinnaker: &client.Front50Config_Spinnaker{
			PersistentStoreType: h.PersistentStorage.PersistentStoreType,
			Gcs:                 h.PersistentStorage.Gcs,
		},
	}
	return f, nil
}

func HalToClouddriver(h client.HalConfig) (client.ClouddriverConfig, error) {
	c := client.ClouddriverConfig{
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

func HalToEcho(h client.HalConfig) (client.EchoConfig, error) {
	c := client.EchoConfig{}
	return c, nil
}
