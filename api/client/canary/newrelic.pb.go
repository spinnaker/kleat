// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: canary/newrelic.proto

package canary

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration for the New Relic canary integration.
type NewRelic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the integration is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*NewRelicAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *NewRelic) Reset() {
	*x = NewRelic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_canary_newrelic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRelic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRelic) ProtoMessage() {}

func (x *NewRelic) ProtoReflect() protoreflect.Message {
	mi := &file_canary_newrelic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRelic.ProtoReflect.Descriptor instead.
func (*NewRelic) Descriptor() ([]byte, []int) {
	return file_canary_newrelic_proto_rawDescGZIP(), []int{0}
}

func (x *NewRelic) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *NewRelic) GetAccounts() []*NewRelicAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

// Configuration for a New Relic account.
type NewRelicAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Configuration for the New Relic Insights server endpoint.
	Endpoint *NewRelicAccount_Endpoint `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// (Required) Your account's unique New Relic Insights API key. See
	// https://docs.newrelic.com/docs/insights/insights-api/get-data/query-insights-event-data-api.
	ApiKey string `protobuf:"bytes,3,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	// (Required) Your New Relic account ID. See
	// https://docs.newrelic.com/docs/accounts/install-new-relic/account-setup/account-id.
	ApplicationKey string `protobuf:"bytes,4,opt,name=applicationKey,proto3" json:"applicationKey,omitempty"`
}

func (x *NewRelicAccount) Reset() {
	*x = NewRelicAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_canary_newrelic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRelicAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRelicAccount) ProtoMessage() {}

func (x *NewRelicAccount) ProtoReflect() protoreflect.Message {
	mi := &file_canary_newrelic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRelicAccount.ProtoReflect.Descriptor instead.
func (*NewRelicAccount) Descriptor() ([]byte, []int) {
	return file_canary_newrelic_proto_rawDescGZIP(), []int{1}
}

func (x *NewRelicAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewRelicAccount) GetEndpoint() *NewRelicAccount_Endpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *NewRelicAccount) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *NewRelicAccount) GetApplicationKey() string {
	if x != nil {
		return x.ApplicationKey
	}
	return ""
}

// Configuration for the New Relic Insights server endpoint.
type NewRelicAccount_Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The base URL to the New Relic Insights server.
	BaseUrl string `protobuf:"bytes,1,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
}

func (x *NewRelicAccount_Endpoint) Reset() {
	*x = NewRelicAccount_Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_canary_newrelic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRelicAccount_Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRelicAccount_Endpoint) ProtoMessage() {}

func (x *NewRelicAccount_Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_canary_newrelic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRelicAccount_Endpoint.ProtoReflect.Descriptor instead.
func (*NewRelicAccount_Endpoint) Descriptor() ([]byte, []int) {
	return file_canary_newrelic_proto_rawDescGZIP(), []int{1, 0}
}

func (x *NewRelicAccount_Endpoint) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

var File_canary_newrelic_proto protoreflect.FileDescriptor

var file_canary_newrelic_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x61, 0x6e, 0x61, 0x72, 0x79, 0x22, 0x5f, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x6c, 0x69,
	0x63, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x4e, 0x65, 0x77,
	0x52, 0x65, 0x6c, 0x69, 0x63, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x52, 0x65,
	0x6c, 0x69, 0x63, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e,
	0x4e, 0x65, 0x77, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x1a, 0x24, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72,
	0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_canary_newrelic_proto_rawDescOnce sync.Once
	file_canary_newrelic_proto_rawDescData = file_canary_newrelic_proto_rawDesc
)

func file_canary_newrelic_proto_rawDescGZIP() []byte {
	file_canary_newrelic_proto_rawDescOnce.Do(func() {
		file_canary_newrelic_proto_rawDescData = protoimpl.X.CompressGZIP(file_canary_newrelic_proto_rawDescData)
	})
	return file_canary_newrelic_proto_rawDescData
}

var file_canary_newrelic_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_canary_newrelic_proto_goTypes = []interface{}{
	(*NewRelic)(nil),                 // 0: proto.canary.NewRelic
	(*NewRelicAccount)(nil),          // 1: proto.canary.NewRelicAccount
	(*NewRelicAccount_Endpoint)(nil), // 2: proto.canary.NewRelicAccount.Endpoint
}
var file_canary_newrelic_proto_depIdxs = []int32{
	1, // 0: proto.canary.NewRelic.accounts:type_name -> proto.canary.NewRelicAccount
	2, // 1: proto.canary.NewRelicAccount.endpoint:type_name -> proto.canary.NewRelicAccount.Endpoint
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_canary_newrelic_proto_init() }
func file_canary_newrelic_proto_init() {
	if File_canary_newrelic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_canary_newrelic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRelic); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_canary_newrelic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRelicAccount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_canary_newrelic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRelicAccount_Endpoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_canary_newrelic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_canary_newrelic_proto_goTypes,
		DependencyIndexes: file_canary_newrelic_proto_depIdxs,
		MessageInfos:      file_canary_newrelic_proto_msgTypes,
	}.Build()
	File_canary_newrelic_proto = out.File
	file_canary_newrelic_proto_rawDesc = nil
	file_canary_newrelic_proto_goTypes = nil
	file_canary_newrelic_proto_depIdxs = nil
}
