// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: config/kayenta.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	canary "github.com/spinnaker/kleat/api/client/canary"
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

// Configuration for the Kayenta microservice.
type Kayenta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kayenta *Kayenta_ServiceIntegrations `protobuf:"bytes,1,opt,name=kayenta,proto3" json:"kayenta,omitempty"`
}

func (x *Kayenta) Reset() {
	*x = Kayenta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_kayenta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kayenta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kayenta) ProtoMessage() {}

func (x *Kayenta) ProtoReflect() protoreflect.Message {
	mi := &file_config_kayenta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kayenta.ProtoReflect.Descriptor instead.
func (*Kayenta) Descriptor() ([]byte, []int) {
	return file_config_kayenta_proto_rawDescGZIP(), []int{0}
}

func (x *Kayenta) GetKayenta() *Kayenta_ServiceIntegrations {
	if x != nil {
		return x.Kayenta
	}
	return nil
}

type Kayenta_ServiceIntegrations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Google      *Kayenta_ServiceIntegrations_Google `protobuf:"bytes,1,opt,name=google,proto3" json:"google,omitempty"`
	Stackdriver *canary.Stackdriver                 `protobuf:"bytes,2,opt,name=stackdriver,proto3" json:"stackdriver,omitempty"`
	Gcs         *canary.Gcs                         `protobuf:"bytes,3,opt,name=gcs,proto3" json:"gcs,omitempty"`
	Prometheus  *canary.Prometheus                  `protobuf:"bytes,4,opt,name=prometheus,proto3" json:"prometheus,omitempty"`
	Datadog     *canary.Datadog                     `protobuf:"bytes,5,opt,name=datadog,proto3" json:"datadog,omitempty"`
	Aws         *Kayenta_ServiceIntegrations_Aws    `protobuf:"bytes,6,opt,name=aws,proto3" json:"aws,omitempty"`
	S3          *canary.S3                          `protobuf:"bytes,7,opt,name=s3,proto3" json:"s3,omitempty"`
	Signalfx    *canary.SignalFx                    `protobuf:"bytes,8,opt,name=signalfx,proto3" json:"signalfx,omitempty"`
	Newrelic    *canary.NewRelic                    `protobuf:"bytes,9,opt,name=newrelic,proto3" json:"newrelic,omitempty"`
}

func (x *Kayenta_ServiceIntegrations) Reset() {
	*x = Kayenta_ServiceIntegrations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_kayenta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kayenta_ServiceIntegrations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kayenta_ServiceIntegrations) ProtoMessage() {}

func (x *Kayenta_ServiceIntegrations) ProtoReflect() protoreflect.Message {
	mi := &file_config_kayenta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kayenta_ServiceIntegrations.ProtoReflect.Descriptor instead.
func (*Kayenta_ServiceIntegrations) Descriptor() ([]byte, []int) {
	return file_config_kayenta_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Kayenta_ServiceIntegrations) GetGoogle() *Kayenta_ServiceIntegrations_Google {
	if x != nil {
		return x.Google
	}
	return nil
}

func (x *Kayenta_ServiceIntegrations) GetStackdriver() *canary.Stackdriver {
	if x != nil {
		return x.Stackdriver
	}
	return nil
}

func (x *Kayenta_ServiceIntegrations) GetGcs() *canary.Gcs {
	if x != nil {
		return x.Gcs
	}
	return nil
}

func (x *Kayenta_ServiceIntegrations) GetPrometheus() *canary.Prometheus {
	if x != nil {
		return x.Prometheus
	}
	return nil
}

func (x *Kayenta_ServiceIntegrations) GetDatadog() *canary.Datadog {
	if x != nil {
		return x.Datadog
	}
	return nil
}

func (x *Kayenta_ServiceIntegrations) GetAws() *Kayenta_ServiceIntegrations_Aws {
	if x != nil {
		return x.Aws
	}
	return nil
}

func (x *Kayenta_ServiceIntegrations) GetS3() *canary.S3 {
	if x != nil {
		return x.S3
	}
	return nil
}

func (x *Kayenta_ServiceIntegrations) GetSignalfx() *canary.SignalFx {
	if x != nil {
		return x.Signalfx
	}
	return nil
}

func (x *Kayenta_ServiceIntegrations) GetNewrelic() *canary.NewRelic {
	if x != nil {
		return x.Newrelic
	}
	return nil
}

type Kayenta_ServiceIntegrations_Google struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled  bool                    `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Accounts []*canary.GoogleAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *Kayenta_ServiceIntegrations_Google) Reset() {
	*x = Kayenta_ServiceIntegrations_Google{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_kayenta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kayenta_ServiceIntegrations_Google) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kayenta_ServiceIntegrations_Google) ProtoMessage() {}

func (x *Kayenta_ServiceIntegrations_Google) ProtoReflect() protoreflect.Message {
	mi := &file_config_kayenta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kayenta_ServiceIntegrations_Google.ProtoReflect.Descriptor instead.
func (*Kayenta_ServiceIntegrations_Google) Descriptor() ([]byte, []int) {
	return file_config_kayenta_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Kayenta_ServiceIntegrations_Google) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Kayenta_ServiceIntegrations_Google) GetAccounts() []*canary.GoogleAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type Kayenta_ServiceIntegrations_Aws struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled  bool                 `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Accounts []*canary.AwsAccount `protobuf:"bytes,3,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *Kayenta_ServiceIntegrations_Aws) Reset() {
	*x = Kayenta_ServiceIntegrations_Aws{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_kayenta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kayenta_ServiceIntegrations_Aws) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kayenta_ServiceIntegrations_Aws) ProtoMessage() {}

func (x *Kayenta_ServiceIntegrations_Aws) ProtoReflect() protoreflect.Message {
	mi := &file_config_kayenta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kayenta_ServiceIntegrations_Aws.ProtoReflect.Descriptor instead.
func (*Kayenta_ServiceIntegrations_Aws) Descriptor() ([]byte, []int) {
	return file_config_kayenta_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *Kayenta_ServiceIntegrations_Aws) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Kayenta_ServiceIntegrations_Aws) GetAccounts() []*canary.AwsAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

var File_config_kayenta_proto protoreflect.FileDescriptor

var file_config_kayenta_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6b, 0x61, 0x79, 0x65, 0x6e, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x10, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x61, 0x77, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x61,
	0x6e, 0x61, 0x72, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79,
	0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x66, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79,
	0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x67, 0x63, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x73, 0x33, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x05, 0x0a, 0x07, 0x4b, 0x61, 0x79, 0x65, 0x6e, 0x74,
	0x61, 0x12, 0x43, 0x0a, 0x07, 0x6b, 0x61, 0x79, 0x65, 0x6e, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x4b, 0x61, 0x79, 0x65, 0x6e, 0x74, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6b,
	0x61, 0x79, 0x65, 0x6e, 0x74, 0x61, 0x1a, 0xab, 0x05, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x48,
	0x0a, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x61,
	0x79, 0x65, 0x6e, 0x74, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x52, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x03, 0x67, 0x63, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72,
	0x79, 0x2e, 0x47, 0x63, 0x73, 0x52, 0x03, 0x67, 0x63, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x50, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74,
	0x68, 0x65, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61,
	0x6e, 0x61, 0x72, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x52, 0x07, 0x64, 0x61,
	0x74, 0x61, 0x64, 0x6f, 0x67, 0x12, 0x3f, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x4b, 0x61, 0x79, 0x65, 0x6e, 0x74, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x77,
	0x73, 0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x20, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72,
	0x79, 0x2e, 0x53, 0x33, 0x52, 0x02, 0x73, 0x33, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x66, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x46, 0x78, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x66, 0x78, 0x12, 0x32, 0x0a, 0x08,
	0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x4e, 0x65,
	0x77, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69, 0x63,
	0x1a, 0x5b, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x55, 0x0a,
	0x03, 0x41, 0x77, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x34,
	0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e,
	0x41, 0x77, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65,
	0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_kayenta_proto_rawDescOnce sync.Once
	file_config_kayenta_proto_rawDescData = file_config_kayenta_proto_rawDesc
)

func file_config_kayenta_proto_rawDescGZIP() []byte {
	file_config_kayenta_proto_rawDescOnce.Do(func() {
		file_config_kayenta_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_kayenta_proto_rawDescData)
	})
	return file_config_kayenta_proto_rawDescData
}

var file_config_kayenta_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_config_kayenta_proto_goTypes = []interface{}{
	(*Kayenta)(nil),                            // 0: proto.config.Kayenta
	(*Kayenta_ServiceIntegrations)(nil),        // 1: proto.config.Kayenta.ServiceIntegrations
	(*Kayenta_ServiceIntegrations_Google)(nil), // 2: proto.config.Kayenta.ServiceIntegrations.Google
	(*Kayenta_ServiceIntegrations_Aws)(nil),    // 3: proto.config.Kayenta.ServiceIntegrations.Aws
	(*canary.Stackdriver)(nil),                 // 4: proto.canary.Stackdriver
	(*canary.Gcs)(nil),                         // 5: proto.canary.Gcs
	(*canary.Prometheus)(nil),                  // 6: proto.canary.Prometheus
	(*canary.Datadog)(nil),                     // 7: proto.canary.Datadog
	(*canary.S3)(nil),                          // 8: proto.canary.S3
	(*canary.SignalFx)(nil),                    // 9: proto.canary.SignalFx
	(*canary.NewRelic)(nil),                    // 10: proto.canary.NewRelic
	(*canary.GoogleAccount)(nil),               // 11: proto.canary.GoogleAccount
	(*canary.AwsAccount)(nil),                  // 12: proto.canary.AwsAccount
}
var file_config_kayenta_proto_depIdxs = []int32{
	1,  // 0: proto.config.Kayenta.kayenta:type_name -> proto.config.Kayenta.ServiceIntegrations
	2,  // 1: proto.config.Kayenta.ServiceIntegrations.google:type_name -> proto.config.Kayenta.ServiceIntegrations.Google
	4,  // 2: proto.config.Kayenta.ServiceIntegrations.stackdriver:type_name -> proto.canary.Stackdriver
	5,  // 3: proto.config.Kayenta.ServiceIntegrations.gcs:type_name -> proto.canary.Gcs
	6,  // 4: proto.config.Kayenta.ServiceIntegrations.prometheus:type_name -> proto.canary.Prometheus
	7,  // 5: proto.config.Kayenta.ServiceIntegrations.datadog:type_name -> proto.canary.Datadog
	3,  // 6: proto.config.Kayenta.ServiceIntegrations.aws:type_name -> proto.config.Kayenta.ServiceIntegrations.Aws
	8,  // 7: proto.config.Kayenta.ServiceIntegrations.s3:type_name -> proto.canary.S3
	9,  // 8: proto.config.Kayenta.ServiceIntegrations.signalfx:type_name -> proto.canary.SignalFx
	10, // 9: proto.config.Kayenta.ServiceIntegrations.newrelic:type_name -> proto.canary.NewRelic
	11, // 10: proto.config.Kayenta.ServiceIntegrations.Google.accounts:type_name -> proto.canary.GoogleAccount
	12, // 11: proto.config.Kayenta.ServiceIntegrations.Aws.accounts:type_name -> proto.canary.AwsAccount
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_config_kayenta_proto_init() }
func file_config_kayenta_proto_init() {
	if File_config_kayenta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_kayenta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kayenta); i {
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
		file_config_kayenta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kayenta_ServiceIntegrations); i {
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
		file_config_kayenta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kayenta_ServiceIntegrations_Google); i {
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
		file_config_kayenta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kayenta_ServiceIntegrations_Aws); i {
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
			RawDescriptor: file_config_kayenta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_kayenta_proto_goTypes,
		DependencyIndexes: file_config_kayenta_proto_depIdxs,
		MessageInfos:      file_config_kayenta_proto_msgTypes,
	}.Build()
	File_config_kayenta_proto = out.File
	file_config_kayenta_proto_rawDesc = nil
	file_config_kayenta_proto_goTypes = nil
	file_config_kayenta_proto_depIdxs = nil
}
