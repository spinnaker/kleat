// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: echo.proto

package client

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

type EchoConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slack        *SlackNotification        `protobuf:"bytes,1,opt,name=slack,proto3" json:"slack,omitempty"`
	Twilio       *TwilioNotification       `protobuf:"bytes,2,opt,name=twilio,proto3" json:"twilio,omitempty"`
	GithubStatus *GithubStatusNotification `protobuf:"bytes,3,opt,name=githubStatus,proto3" json:"githubStatus,omitempty"`
	Artifacts    *ArtifactProviders        `protobuf:"bytes,4,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
	Pubsub       *PubsubProviders          `protobuf:"bytes,5,opt,name=pubsub,proto3" json:"pubsub,omitempty"`
	Gcb          *GoogleCloudBuildProvider `protobuf:"bytes,6,opt,name=gcb,proto3" json:"gcb,omitempty"`
}

func (x *EchoConfig) Reset() {
	*x = EchoConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoConfig) ProtoMessage() {}

func (x *EchoConfig) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoConfig.ProtoReflect.Descriptor instead.
func (*EchoConfig) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoConfig) GetSlack() *SlackNotification {
	if x != nil {
		return x.Slack
	}
	return nil
}

func (x *EchoConfig) GetTwilio() *TwilioNotification {
	if x != nil {
		return x.Twilio
	}
	return nil
}

func (x *EchoConfig) GetGithubStatus() *GithubStatusNotification {
	if x != nil {
		return x.GithubStatus
	}
	return nil
}

func (x *EchoConfig) GetArtifacts() *ArtifactProviders {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

func (x *EchoConfig) GetPubsub() *PubsubProviders {
	if x != nil {
		return x.Pubsub
	}
	return nil
}

func (x *EchoConfig) GetGcb() *GoogleCloudBuildProvider {
	if x != nil {
		return x.Gcb
	}
	return nil
}

var File_echo_proto protoreflect.FileDescriptor

var file_echo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x63, 0x62, 0x5f, 0x63, 0x69, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf,
	0x02, 0x0a, 0x0a, 0x45, 0x63, 0x68, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2e, 0x0a,
	0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x31, 0x0a,
	0x06, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f,
	0x12, 0x43, 0x0a, 0x0c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x12, 0x2e, 0x0a,
	0x06, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x06, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x12, 0x31, 0x0a,
	0x03, 0x67, 0x63, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x03, 0x67, 0x63, 0x62,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_echo_proto_rawDescOnce sync.Once
	file_echo_proto_rawDescData = file_echo_proto_rawDesc
)

func file_echo_proto_rawDescGZIP() []byte {
	file_echo_proto_rawDescOnce.Do(func() {
		file_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_proto_rawDescData)
	})
	return file_echo_proto_rawDescData
}

var file_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_echo_proto_goTypes = []interface{}{
	(*EchoConfig)(nil),               // 0: proto.EchoConfig
	(*SlackNotification)(nil),        // 1: proto.SlackNotification
	(*TwilioNotification)(nil),       // 2: proto.TwilioNotification
	(*GithubStatusNotification)(nil), // 3: proto.GithubStatusNotification
	(*ArtifactProviders)(nil),        // 4: proto.ArtifactProviders
	(*PubsubProviders)(nil),          // 5: proto.PubsubProviders
	(*GoogleCloudBuildProvider)(nil), // 6: proto.GoogleCloudBuildProvider
}
var file_echo_proto_depIdxs = []int32{
	1, // 0: proto.EchoConfig.slack:type_name -> proto.SlackNotification
	2, // 1: proto.EchoConfig.twilio:type_name -> proto.TwilioNotification
	3, // 2: proto.EchoConfig.githubStatus:type_name -> proto.GithubStatusNotification
	4, // 3: proto.EchoConfig.artifacts:type_name -> proto.ArtifactProviders
	5, // 4: proto.EchoConfig.pubsub:type_name -> proto.PubsubProviders
	6, // 5: proto.EchoConfig.gcb:type_name -> proto.GoogleCloudBuildProvider
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_echo_proto_init() }
func file_echo_proto_init() {
	if File_echo_proto != nil {
		return
	}
	file_notifications_proto_init()
	file_artifacts_proto_init()
	file_pubsub_proto_init()
	file_gcb_ci_provider_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoConfig); i {
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
			RawDescriptor: file_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_echo_proto_goTypes,
		DependencyIndexes: file_echo_proto_depIdxs,
		MessageInfos:      file_echo_proto_msgTypes,
	}.Build()
	File_echo_proto = out.File
	file_echo_proto_rawDesc = nil
	file_echo_proto_goTypes = nil
	file_echo_proto_depIdxs = nil
}
