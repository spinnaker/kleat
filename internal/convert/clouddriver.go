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

package convert

import "github.com/spinnaker/kleat/api/client/config"

// HalToClouddriver generates the clouddriver config for the supplied config.Hal h.
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
