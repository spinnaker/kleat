// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: storage/oracle.proto

package storage

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

// Configuration for an Oracle persistent store.
type Oracle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether this persistent store is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The bucket name to store persistent state object in.
	BucketName string `protobuf:"bytes,2,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	// The namespace the bucket and objects should be created in.
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// An Oracle region (e.g., us-phoenix-1).
	Region string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	// The OCID of the Oracle User you’re authenticating as.
	UserId string `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	// Fingerprint of the public key.
	Fingerprint string `protobuf:"bytes,6,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Path to the private key in PEM format.
	SshPrivateKeyFilePath string `protobuf:"bytes,7,opt,name=sshPrivateKeyFilePath,proto3" json:"sshPrivateKeyFilePath,omitempty"`
	// Passphrase used for the private key, if it is encrypted.
	PrivateKeyPassphrase string `protobuf:"bytes,8,opt,name=privateKeyPassphrase,proto3" json:"privateKeyPassphrase,omitempty"`
	// The OCID of the Oracle Tenancy to use.
	TenancyId string `protobuf:"bytes,9,opt,name=tenancyId,proto3" json:"tenancyId,omitempty"`
}

func (x *Oracle) Reset() {
	*x = Oracle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_oracle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oracle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oracle) ProtoMessage() {}

func (x *Oracle) ProtoReflect() protoreflect.Message {
	mi := &file_storage_oracle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oracle.ProtoReflect.Descriptor instead.
func (*Oracle) Descriptor() ([]byte, []int) {
	return file_storage_oracle_proto_rawDescGZIP(), []int{0}
}

func (x *Oracle) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Oracle) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *Oracle) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Oracle) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Oracle) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Oracle) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *Oracle) GetSshPrivateKeyFilePath() string {
	if x != nil {
		return x.SshPrivateKeyFilePath
	}
	return ""
}

func (x *Oracle) GetPrivateKeyPassphrase() string {
	if x != nil {
		return x.PrivateKeyPassphrase
	}
	return ""
}

func (x *Oracle) GetTenancyId() string {
	if x != nil {
		return x.TenancyId
	}
	return ""
}

var File_storage_oracle_proto protoreflect.FileDescriptor

var file_storage_oracle_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0xba, 0x02, 0x0a, 0x06, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x73,
	0x68, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x73, 0x68, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x32, 0x0a, 0x14, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61,
	0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68,
	0x72, 0x61, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x49,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79,
	0x49, 0x64, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_oracle_proto_rawDescOnce sync.Once
	file_storage_oracle_proto_rawDescData = file_storage_oracle_proto_rawDesc
)

func file_storage_oracle_proto_rawDescGZIP() []byte {
	file_storage_oracle_proto_rawDescOnce.Do(func() {
		file_storage_oracle_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_oracle_proto_rawDescData)
	})
	return file_storage_oracle_proto_rawDescData
}

var file_storage_oracle_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_oracle_proto_goTypes = []interface{}{
	(*Oracle)(nil), // 0: proto.storage.Oracle
}
var file_storage_oracle_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_storage_oracle_proto_init() }
func file_storage_oracle_proto_init() {
	if File_storage_oracle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_oracle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oracle); i {
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
			RawDescriptor: file_storage_oracle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_oracle_proto_goTypes,
		DependencyIndexes: file_storage_oracle_proto_depIdxs,
		MessageInfos:      file_storage_oracle_proto_msgTypes,
	}.Build()
	File_storage_oracle_proto = out.File
	file_storage_oracle_proto_rawDesc = nil
	file_storage_oracle_proto_goTypes = nil
	file_storage_oracle_proto_depIdxs = nil
}
