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
	"github.com/spinnaker/kleat/api/client/cloudprovider"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/convert"
)

var clouddriverTests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToClouddriver(h) },
	tests: []testCase{
		{
			"Empty hal config",
			&config.Hal{},
			&config.Clouddriver{},
		},
		{
			"Empty providers",
			&config.Hal{
				Providers: &cloudprovider.Providers{},
			},
			&config.Clouddriver{},
		},
		{
			"Empty Kubernetes provider",
			&config.Hal{
				Providers: &cloudprovider.Providers{
					Kubernetes: &cloudprovider.Kubernetes{},
				},
			},
			&config.Clouddriver{
				Kubernetes: &cloudprovider.Kubernetes{},
			},
		},
		{
			"Kubernetes account",
			&config.Hal{
				Providers: &cloudprovider.Providers{
					Kubernetes: &cloudprovider.Kubernetes{
						Enabled: wrapperspb.Bool(true),
						Accounts: []*cloudprovider.KubernetesAccount{
							{
								Name:           "my-account",
								Kinds:          []string{"deployment"},
								OmitNamespaces: []string{"kube-system"},
							},
						},
						PrimaryAccount: "my-account",
					},
				},
			},
			&config.Clouddriver{
				Kubernetes: &cloudprovider.Kubernetes{
					Enabled: wrapperspb.Bool(true),
					Accounts: []*cloudprovider.KubernetesAccount{
						{
							Name:           "my-account",
							Kinds:          []string{"deployment"},
							OmitNamespaces: []string{"kube-system"},
						},
					},
					PrimaryAccount: "my-account",
				},
			},
		},
		{
			"Empty artifacts",
			&config.Hal{
				Artifacts: &artifact.Artifacts{},
			},
			&config.Clouddriver{
				Artifacts: &artifact.Artifacts{},
			},
		},
		{
			"Empty GCS artifacts",
			&config.Hal{
				Artifacts: &artifact.Artifacts{
					Gcs: &artifact.Gcs{},
				},
			},
			&config.Clouddriver{
				Artifacts: &artifact.Artifacts{
					Gcs: &artifact.Gcs{},
				},
			},
		},
		{
			"GCS artifact account",
			&config.Hal{
				Artifacts: &artifact.Artifacts{
					Gcs: &artifact.Gcs{
						Enabled: wrapperspb.Bool(true),
						Accounts: []*artifact.GcsAccount{
							{
								Name:     "my-account",
								JsonPath: "/var/secrets/my-key.json",
							},
						},
					},
				},
			},
			&config.Clouddriver{
				Artifacts: &artifact.Artifacts{
					Gcs: &artifact.Gcs{
						Enabled: wrapperspb.Bool(true),
						Accounts: []*artifact.GcsAccount{
							{
								Name:     "my-account",
								JsonPath: "/var/secrets/my-key.json",
							},
						},
					},
				},
			},
		},
	},
}

func TestHalToClouddriver(t *testing.T) {
	runConfigTest(t, clouddriverTests)
}
