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

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/storage"
	"github.com/spinnaker/kleat/internal/convert"
)

var front50Tests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToFront50(h) },
	tests: []testCase{
		{
			"Empty hal config",
			&config.Hal{},
			&config.Front50{
				Spinnaker: &config.Front50_Spinnaker{},
			},
		},
		{
			"Empty persistent storage",
			&config.Hal{
				PersistentStorage: &storage.PersistentStorage{},
			},
			&config.Front50{
				Spinnaker: &config.Front50_Spinnaker{},
			},
		},
		{
			"Empty persistent storage configs",
			&config.Hal{
				PersistentStorage: &storage.PersistentStorage{
					Gcs: &storage.Gcs{},
					S3:  &storage.S3{},
				},
			},
			&config.Front50{
				Spinnaker: &config.Front50_Spinnaker{
					Gcs: &storage.Gcs{},
					S3:  &storage.S3{},
				},
			},
		},
		{
			"GCS and S3 storage configured",
			&config.Hal{
				PersistentStorage: &storage.PersistentStorage{
					Gcs: &storage.Gcs{
						Enabled:  wrapperspb.Bool(true),
						JsonPath: "/var/secrets/my-gcs-key.json",
						Project:  "my-project",
						Bucket:   "my-gcs-bucket",
					},
					S3: &storage.S3{
						Enabled:              wrapperspb.Bool(false),
						Bucket:               "my-s3-bucket",
						PathStyleAccess:      wrapperspb.Bool(true),
						ServerSideEncryption: storage.S3ServerSideEncryption_AES256,
					},
				},
			},
			&config.Front50{
				Spinnaker: &config.Front50_Spinnaker{
					Gcs: &storage.Gcs{
						Enabled:  wrapperspb.Bool(true),
						JsonPath: "/var/secrets/my-gcs-key.json",
						Project:  "my-project",
						Bucket:   "my-gcs-bucket",
					},
					S3: &storage.S3{
						Enabled:              wrapperspb.Bool(false),
						Bucket:               "my-s3-bucket",
						PathStyleAccess:      wrapperspb.Bool(true),
						ServerSideEncryption: storage.S3ServerSideEncryption_AES256,
					},
				},
			},
		},
		{
			"Enable ManagedDelivery",
			&config.Hal{
				ManagedDelivery: &config.ManagedDelivery{
					Enabled: wrapperspb.Bool(true),
				},
			},
			&config.Front50{
				Spinnaker: &config.Front50_Spinnaker{
					Delivery: &config.Front50_Delivery{
						Enabled: wrapperspb.Bool(true),
					},
				},
			},
		},
	},
}

func TestHalToFront50(t *testing.T) {
	runConfigTest(t, front50Tests)
}
