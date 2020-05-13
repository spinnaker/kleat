// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: ci/gcb.proto

package ci

import (
	proto "github.com/golang/protobuf/proto"
	client "github.com/spinnaker/kleat/api/client"
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

// Configuration for the Google Cloud Build Provider.
type GoogleCloudBuild struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the Google Cloud Build provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured Google Cloud Build accounts.
	Accounts []*GoogleCloudBuildAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *GoogleCloudBuild) Reset() {
	*x = GoogleCloudBuild{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_gcb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleCloudBuild) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleCloudBuild) ProtoMessage() {}

func (x *GoogleCloudBuild) ProtoReflect() protoreflect.Message {
	mi := &file_ci_gcb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleCloudBuild.ProtoReflect.Descriptor instead.
func (*GoogleCloudBuild) Descriptor() ([]byte, []int) {
	return file_ci_gcb_proto_rawDescGZIP(), []int{0}
}

func (x *GoogleCloudBuild) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *GoogleCloudBuild) GetAccounts() []*GoogleCloudBuildAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

// Configuration for a Google Cloud Build account.
type GoogleCloudBuildAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The name of the Google Cloud Platform project in which to trigger and monitor builds.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// The name of the Pub/Sub subscription on which to listen for build changes.
	SubscriptionName string `protobuf:"bytes,3,opt,name=subscriptionName,proto3" json:"subscriptionName,omitempty"`
	// The path to a JSON service account that Spinnaker will use as
	// credentials. This is only needed if Spinnaker is not deployed on a
	// Google Compute Engine VM, or needs permissions not afforded to the VM
	// it is running on.
	JsonKey string `protobuf:"bytes,4,opt,name=jsonKey,proto3" json:"jsonKey,omitempty"`
	// Fiat permissions configuration.
	Permissions *client.Permissions `protobuf:"bytes,5,opt,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *GoogleCloudBuildAccount) Reset() {
	*x = GoogleCloudBuildAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_gcb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleCloudBuildAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleCloudBuildAccount) ProtoMessage() {}

func (x *GoogleCloudBuildAccount) ProtoReflect() protoreflect.Message {
	mi := &file_ci_gcb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleCloudBuildAccount.ProtoReflect.Descriptor instead.
func (*GoogleCloudBuildAccount) Descriptor() ([]byte, []int) {
	return file_ci_gcb_proto_rawDescGZIP(), []int{1}
}

func (x *GoogleCloudBuildAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoogleCloudBuildAccount) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GoogleCloudBuildAccount) GetSubscriptionName() string {
	if x != nil {
		return x.SubscriptionName
	}
	return ""
}

func (x *GoogleCloudBuildAccount) GetJsonKey() string {
	if x != nil {
		return x.JsonKey
	}
	return ""
}

func (x *GoogleCloudBuildAccount) GetPermissions() *client.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_ci_gcb_proto protoreflect.FileDescriptor

var file_ci_gcb_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x69, 0x2f, 0x67, 0x63, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x1a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x10, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3d, 0x0a, 0x08, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0xc3, 0x01, 0x0a, 0x17, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6a, 0x73, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6a, 0x73, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69,
	0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ci_gcb_proto_rawDescOnce sync.Once
	file_ci_gcb_proto_rawDescData = file_ci_gcb_proto_rawDesc
)

func file_ci_gcb_proto_rawDescGZIP() []byte {
	file_ci_gcb_proto_rawDescOnce.Do(func() {
		file_ci_gcb_proto_rawDescData = protoimpl.X.CompressGZIP(file_ci_gcb_proto_rawDescData)
	})
	return file_ci_gcb_proto_rawDescData
}

var file_ci_gcb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ci_gcb_proto_goTypes = []interface{}{
	(*GoogleCloudBuild)(nil),        // 0: proto.ci.GoogleCloudBuild
	(*GoogleCloudBuildAccount)(nil), // 1: proto.ci.GoogleCloudBuildAccount
	(*client.Permissions)(nil),      // 2: proto.Permissions
}
var file_ci_gcb_proto_depIdxs = []int32{
	1, // 0: proto.ci.GoogleCloudBuild.accounts:type_name -> proto.ci.GoogleCloudBuildAccount
	2, // 1: proto.ci.GoogleCloudBuildAccount.permissions:type_name -> proto.Permissions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ci_gcb_proto_init() }
func file_ci_gcb_proto_init() {
	if File_ci_gcb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ci_gcb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleCloudBuild); i {
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
		file_ci_gcb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleCloudBuildAccount); i {
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
			RawDescriptor: file_ci_gcb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ci_gcb_proto_goTypes,
		DependencyIndexes: file_ci_gcb_proto_depIdxs,
		MessageInfos:      file_ci_gcb_proto_msgTypes,
	}.Build()
	File_ci_gcb_proto = out.File
	file_ci_gcb_proto_rawDesc = nil
	file_ci_gcb_proto_goTypes = nil
	file_ci_gcb_proto_depIdxs = nil
}
