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

import (
	"github.com/spinnaker/kleat/api/client/config"
)

// HalToIgor generates the igor config for the supplied config.Hal h.
func HalToIgor(h *config.Hal) *config.Igor {
	return &config.Igor{
		Artifactory:    h.GetRepository().GetArtifactory(),
		Artifacts:      getIgorArtifacts(h),
		DockerRegistry: getIgorDockerRegistry(h),
		Gcb:            h.GetCi().GetGcb(),
		Codebuild:      h.GetCi().GetCodebuild(),
		Concourse:      h.GetCi().GetConcourse(),
		Jenkins:        h.GetCi().GetJenkins(),
		Travis:         h.GetCi().GetTravis(),
		Wercker:        h.GetCi().GetWercker(),
	}
}

func getIgorArtifacts(h *config.Hal) *config.Igor_Artifacts {
	if h.GetArtifacts().GetTemplates() == nil {
		return nil
	}
	return &config.Igor_Artifacts{
		Templates: h.GetArtifacts().GetTemplates(),
	}
}

func getIgorDockerRegistry(h *config.Hal) *config.Igor_DockerRegistry {
	if h.GetProviders().GetDockerRegistry() == nil {
		return nil
	}
	return &config.Igor_DockerRegistry{
		Enabled: h.GetProviders().GetDockerRegistry().GetEnabled(),
	}
}
