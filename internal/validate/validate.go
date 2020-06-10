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

// Package validate supports validating a *config.Hal and reporting any
// validation failures.
package validate

import (
	"errors"
	"strings"

	"github.com/spinnaker/kleat/api/client/config"
)

type validationFailure struct {
	msg string
}

type halConfigValidator func(*config.Hal) []validationFailure

func getValidators() []halConfigValidator {
	return []halConfigValidator{
		validateKindsAndOmitKinds,
	}
}

// HalConfig validates the supplied *config.Hal, returning any errors encountered.
func HalConfig(h *config.Hal) error {
	messages := getValidationMessages(h, getValidators())
	if len(messages) > 0 {
		msg := strings.Join(messages, "\n")
		return errors.New(msg)
	}
	return nil
}

func getValidationMessages(h *config.Hal, fa []halConfigValidator) []string {
	var messages []string
	for _, f := range fa {
		rs := f(h)
		for _, r := range rs {
			messages = append(messages, r.msg)
		}
	}
	return messages
}

func validateKindsAndOmitKinds(h *config.Hal) []validationFailure {
	var messages []validationFailure
	for _, a := range h.GetProviders().GetKubernetes().GetAccounts() {
		if !(len(a.GetKinds()) == 0) && !(len(a.GetOmitKinds()) == 0) {
			messages = append(messages, fatalResult("Cannot specify both kinds and omitKinds."))
		}
	}
	return messages
}

func fatalResult(msg string) validationFailure {
	return validationFailure{
		msg: msg,
	}
}
