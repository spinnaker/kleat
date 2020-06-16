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

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kleat",
	Short: "Kleat is a tool for managing Spinnaker configuration files.",
	Long: `Kleat is a tool for generating configuration files for each Spinnaker
microservice from one top-level Spinnaker configuration file. It is intended
for use with the Spinnaker Kustomization base
(https://github.com/spinnaker/kustomization-base)
as an alternative to the Halyard-managed configuration and deployment
experience. Please check out the README at https://github.com/spinnaker/kleat
for more details.`,
}

// Execute runs the kleat binary.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
