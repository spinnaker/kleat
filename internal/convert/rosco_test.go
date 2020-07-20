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

	"github.com/spinnaker/kleat/api/client/cloudprovider"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/convert"
)

var roscoTests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToRosco(h) },
	tests: []testCase{
		{
			"Empty hal config",
			&config.Hal{},
			&config.Rosco{},
		},
		{
			"Empty providers",
			&config.Hal{
				Providers: &cloudprovider.Providers{},
			},
			&config.Rosco{},
		},
		{
			"Empty Google provider",
			&config.Hal{
				Providers: &cloudprovider.Providers{
					Google: &cloudprovider.GoogleComputeEngine{},
				},
			},
			&config.Rosco{
				Google: &cloudprovider.GoogleComputeEngine{},
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
			&config.Rosco{},
		},
		{
			"AWS account",
			&config.Hal{
				Providers: &cloudprovider.Providers{
					Aws: &cloudprovider.Aws{
						Enabled: wrapperspb.Bool(true),
						Accounts: []*cloudprovider.AwsAccount{
							{
								Name: "my-account",
							},
						},
						BakeryDefaults: &cloudprovider.AwsBakeryDefaults{
							AwsAccessKey:                "my-key",
							AwsSecretKey:                "my-secret-key",
							AwsAssociatePublicIpAddress: wrapperspb.Bool(false),
							BaseImages: []*cloudprovider.AwsBaseImageSettings{
								{
									BaseImage: &cloudprovider.AwsBaseImage{
										Id:               "my-image",
										ShortDescription: "this is an image",
									},
									VirtualizationSettings: &cloudprovider.AwsVirtualizationSettings{
										Region:             "us-east-1",
										VirtualizationType: "hvm",
										InstanceType:       "t2-micro",
									},
								},
							},
						},
						PrimaryAccount: "my-account",
					},
				},
			},
			&config.Rosco{
				Aws: &cloudprovider.Aws{
					Enabled: wrapperspb.Bool(true),
					Accounts: []*cloudprovider.AwsAccount{
						{
							Name: "my-account",
						},
					},
					BakeryDefaults: &cloudprovider.AwsBakeryDefaults{
						AwsAccessKey:                "my-key",
						AwsSecretKey:                "my-secret-key",
						AwsAssociatePublicIpAddress: wrapperspb.Bool(false),
						BaseImages: []*cloudprovider.AwsBaseImageSettings{
							{
								BaseImage: &cloudprovider.AwsBaseImage{
									Id:               "my-image",
									ShortDescription: "this is an image",
								},
								VirtualizationSettings: &cloudprovider.AwsVirtualizationSettings{
									Region:             "us-east-1",
									VirtualizationType: "hvm",
									InstanceType:       "t2-micro",
								},
							},
						},
					},
					PrimaryAccount: "my-account",
				},
			},
		},
	},
}

func TestHalToRosco(t *testing.T) {
	runConfigTest(t, roscoTests)
}
