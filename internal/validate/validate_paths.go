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
package validate

import (
	"errors"
	"fmt"
	"os"
)

func EnsureDirectory(d string) error {
	stat, err := os.Stat(d)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return errors.New(fmt.Sprintf("%s is not a directory", d))
	}
	return nil
}

func EnsureFile(f string) error {
	stat, err := os.Stat(f)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return errors.New(fmt.Sprintf("%s is a directory, not a file", f))
	}
	return nil
}
