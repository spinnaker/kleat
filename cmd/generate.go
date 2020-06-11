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
	"github.com/spinnaker/kleat/internal/fileio"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Spinnaker microservice config files",
	Long: `Given a path to your top-level Spinnaker config (halconfig) and an
output directory, writes each generated Spinnaker microservice config file to
the output directory.

Example usage:

kleat /path/to/halconfig /path/to/output`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		hal := args[0]
		dir := args[1]
		return writeServiceConfigs(hal, dir)
	},
}

func writeServiceConfigs(halPath string, dir string) error {
	h, err := fileio.ParseHalConfig(halPath)
	if err != nil {
		return err
	}

	return fileio.WriteConfigs(h, dir)
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
