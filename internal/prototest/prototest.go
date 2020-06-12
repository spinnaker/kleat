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

// Package prototest contains utilities for testing protocol buffer messages.
package prototest

import (
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

// Equal compares two protocol buffer messages.
// If the messages are equal, the returned boolean will be true and the returned
// string will not be meaningful. If they are not equal, the returned boolean
// will be false and the returned string will contain a human-readable report of
// the differences between the messages.
func Equal(a proto.Message, b proto.Message) (bool, string) {
	if !proto.Equal(a, b) {
		diff := cmp.Diff(a, b, protocmp.Transform())
		return false, diff
	}
	return true, ""
}
