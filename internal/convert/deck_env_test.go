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
	"github.com/spinnaker/kleat/api/client/security"
	"github.com/spinnaker/kleat/internal/convert"
)

var deckEnvTests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToDeckEnv(h) },
	tests: []testCase{
		{
			"Empty hal config",
			&config.Hal{},
			&config.DeckEnv{},
		},
		{
			"SSL disabled",
			&config.Hal{
				Security: &security.Security{
					UiSecurity: &security.UiSecurity{
						Ssl: &security.UiSsl{
							Enabled:                  wrapperspb.Bool(false),
							SslCertificateFile:       "/var/secrets/cert.crt",
							SslCertificateKeyFile:    "/var/secrets/cert.key",
							SslCertificatePassphrase: "passw0rd",
						},
					},
				},
			},
			&config.DeckEnv{},
		},
		{
			"SSL configured",
			&config.Hal{
				Security: &security.Security{
					UiSecurity: &security.UiSecurity{
						Ssl: &security.UiSsl{
							Enabled:                  wrapperspb.Bool(true),
							SslCertificateFile:       "/var/secrets/cert.crt",
							SslCertificateKeyFile:    "/var/secrets/cert.key",
							SslCertificatePassphrase: "passw0rd",
						},
					},
				},
			},
			&config.DeckEnv{
				DeckCert:   "/var/secrets/cert.crt",
				DeckKey:    "/var/secrets/cert.key",
				Passphrase: "passw0rd",
			},
		},
	},
}

func TestHalToDeckEnv(t *testing.T) {
	runConfigTest(t, deckEnvTests)
}
