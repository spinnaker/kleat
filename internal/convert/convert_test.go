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

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/prototest"
)

// testCase represents a test case for config generation, indicating
// that we expect the supplied *config.Hal hal to be converted into the supplied
// proto.Message want.
type testCase struct {
	name string
	hal  *config.Hal
	want proto.Message
}

// configTest represents a set of test on a config generator. The supplied
// generator will be used for each of the supplied tests.
type configTest struct {
	generator func(h *config.Hal) proto.Message
	tests     []testCase
}

func runConfigTest(t *testing.T, test configTest) {
	for _, tt := range test.tests {
		t.Run(tt.name, func(t *testing.T) {
			if equal, diff := prototest.Equal(test.generator(tt.hal), tt.want); !equal {
				t.Errorf("Incorrect config generated:\n%s", diff)
			}
		})
	}
}
