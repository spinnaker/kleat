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

package convert_test

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/spinnaker/kleat/api/client/artifact"
	"github.com/spinnaker/kleat/api/client/ci"
	"github.com/spinnaker/kleat/api/client/cloudprovider"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/repository"
	"github.com/spinnaker/kleat/internal/convert"
)

var igorTests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToIgor(h) },
	tests: []testCase{
		{
			"Empty hal config",
			&config.Hal{},
			&config.Igor{},
		},
		{
			"Docker registry enabled",
			&config.Hal{
				Providers: &cloudprovider.Providers{
					DockerRegistry: &cloudprovider.DockerRegistry{
						Enabled: wrapperspb.Bool(true),
					},
				},
			},
			&config.Igor{
				DockerRegistry: &config.Igor_DockerRegistry{
					Enabled: wrapperspb.Bool(true),
				},
			},
		},
		{
			"Artifact template configured",
			&config.Hal{
				Artifacts: &artifact.Artifacts{
					Templates: []*artifact.Template{
						{
							Name:         "my-template",
							TemplatePath: "/var/secrets/my-template",
						},
					},
				},
			},
			&config.Igor{
				Artifacts: &config.Igor_Artifacts{
					Templates: []*artifact.Template{
						{
							Name:         "my-template",
							TemplatePath: "/var/secrets/my-template",
						},
					},
				},
			},
		},
		{
			"Artifactory enabled",
			&config.Hal{
				Repository: &repository.Repository{
					Artifactory: &repository.Artifactory{
						Enabled: wrapperspb.Bool(true),
						Searches: []*repository.Artifactory_Search{
							{
								Name:     "my-search",
								BaseUrl:  "https://my-artifactory",
								Repo:     "my-repo",
								GroupId:  "abc",
								Username: "my-un",
								Password: "my-pw",
							},
						},
					},
				},
			},
			&config.Igor{
				Artifactory: &repository.Artifactory{
					Enabled: wrapperspb.Bool(true),
					Searches: []*repository.Artifactory_Search{
						{
							Name:     "my-search",
							BaseUrl:  "https://my-artifactory",
							Repo:     "my-repo",
							GroupId:  "abc",
							Username: "my-un",
							Password: "my-pw",
						},
					},
				},
			},
		},
		{
			"CI providers configured",
			&config.Hal{
				Ci: &ci.Ci{
					Travis: &ci.Travis{
						Enabled: wrapperspb.Bool(true),
						Masters: []*ci.TravisAccount{
							{
								Name: "my-travis-account",
							},
						},
					},
					Gcb: &ci.GoogleCloudBuild{
						Enabled: wrapperspb.Bool(true),
						Accounts: []*ci.GoogleCloudBuildAccount{
							{
								Name: "my-gcb-account",
							},
						},
					},
				},
			},
			&config.Igor{
				Travis: &ci.Travis{
					Enabled: wrapperspb.Bool(true),
					Masters: []*ci.TravisAccount{
						{
							Name: "my-travis-account",
						},
					},
				},
				Gcb: &ci.GoogleCloudBuild{
					Enabled: wrapperspb.Bool(true),
					Accounts: []*ci.GoogleCloudBuildAccount{
						{
							Name: "my-gcb-account",
						},
					},
				},
			},
		},
	},
}

func TestHalToIgor(t *testing.T) {
	runConfigTest(t, igorTests)
}
