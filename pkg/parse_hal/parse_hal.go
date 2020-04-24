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
	"io/ioutil"

	"github.com/spinnaker/kleat/api/client/config"
	"sigs.k8s.io/yaml"
)

func ParseHalConfig(fn string) (*config.Hal, error) {
	dat, err := ioutil.ReadFile(fn)

	h := config.Hal{}
	err = yaml.Unmarshal([]byte(dat), &h)
	if err != nil {
		return nil, err
	}

	return &h, nil
}

func HalToFront50(h *config.Hal) *config.Front50 {
	return &config.Front50{
		Spinnaker: &config.Front50Spinnaker{
			Gcs:    h.GetPersistentStorage().GetGcs(),
			Azs:    h.GetPersistentStorage().GetAzs(),
			Oracle: h.GetPersistentStorage().GetOracle(),
		},
	}
}

func HalToClouddriver(h *config.Hal) *config.Clouddriver {
	return &config.Clouddriver{
		Kubernetes:     h.GetProviders().GetKubernetes(),
		Google:         h.GetProviders().GetGoogle(),
		Appengine:      h.GetProviders().GetAppengine(),
		Aws:            h.GetProviders().GetAws(),
		Azure:          h.GetProviders().GetAzure(),
		Cloudfoundry:   h.GetProviders().GetCloudfoundry(),
		Dcos:           h.GetProviders().GetDcos(),
		DockerRegistry: h.GetProviders().GetDockerRegistry(),
		Ecs:            h.GetProviders().GetEcs(),
		Huaweicloud:    h.GetProviders().GetHuaweicloud(),
		Oracle:         h.GetProviders().GetOracle(),
		Artifacts:      h.GetArtifacts(),
	}
}

func HalToEcho(h *config.Hal) *config.Echo {
	return &config.Echo{
		Slack:        h.GetNotifications().GetSlack(),
		Twilio:       h.GetNotifications().GetTwilio(),
		GithubStatus: h.GetNotifications().GetGithubStatus(),
		Artifacts:    h.GetArtifacts(),
		Pubsub:       h.GetPubsub(),
		Gcb:          h.GetCi().GetGcb(),
		Stats:        h.GetStats(),
	}
}

func HalToServiceConfigs(h *config.Hal) *config.Services {
	return &config.Services{
		Clouddriver: HalToClouddriver(h),
		Echo:        HalToEcho(h),
		Front50:     HalToFront50(h),
	}
}
