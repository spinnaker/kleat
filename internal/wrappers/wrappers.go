package wrappers

import "github.com/golang/protobuf/ptypes/wrappers"

func True() *wrappers.BoolValue {
	return Bool(true)
}

func False() *wrappers.BoolValue {
	return Bool(false)
}

func Bool(value bool) *wrappers.BoolValue {
	return &wrappers.BoolValue{Value: value}
}
