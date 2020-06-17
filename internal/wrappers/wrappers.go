// Package wrappers provides functions for working with wrapper messages
// in the google.protobuf package.
package wrappers

import "google.golang.org/protobuf/types/known/wrapperspb"

// True returns a wrapper BoolValue containing the value true.
func True() *wrapperspb.BoolValue {
	return Bool(true)
}

// False returns a wrapper BoolValue containing the value false.
func False() *wrapperspb.BoolValue {
	return Bool(false)
}

// Bool returns a wrapper BoolValue containing the specified value.
func Bool(value bool) *wrapperspb.BoolValue {
	return &wrapperspb.BoolValue{Value: value}
}
