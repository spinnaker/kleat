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

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/storage"
	"github.com/spinnaker/kleat/internal/convert"
	"google.golang.org/protobuf/proto"
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
						Enabled:  true,
						JsonPath: "/var/secrets/my-gcs-key.json",
						Project:  "my-project",
						Bucket:   "my-gcs-bucket",
					},
					S3: &storage.S3{
						Enabled:              false,
						Bucket:               "my-s3-bucket",
						PathStyleAccess:      &wrappers.BoolValue{Value: true},
						ServerSideEncryption: storage.S3ServerSideEncryption_AES256,
					},
				},
			},
			&config.Front50{
				Spinnaker: &config.Front50_Spinnaker{
					Gcs: &storage.Gcs{
						Enabled:  true,
						JsonPath: "/var/secrets/my-gcs-key.json",
						Project:  "my-project",
						Bucket:   "my-gcs-bucket",
					},
					S3: &storage.S3{
						Enabled:              false,
						Bucket:               "my-s3-bucket",
						PathStyleAccess:      &wrappers.BoolValue{Value: true},
						ServerSideEncryption: storage.S3ServerSideEncryption_AES256,
					},
				},
			},
		},
	},
}

func TestHalToFront50(t *testing.T) {
	runConfigTest(t, front50Tests)
}
