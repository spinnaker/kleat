// Code generated by protoc-gen-go. DO NOT EDIT.
// source: halconfig.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HalConfig struct {
	PersistentStorage    *HalConfig_PersistentStorage `protobuf:"bytes,1,opt,name=persistentStorage,proto3" json:"persistentStorage,omitempty"`
	Providers            *HalConfig_Providers         `protobuf:"bytes,2,opt,name=providers,proto3" json:"providers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *HalConfig) Reset()         { *m = HalConfig{} }
func (m *HalConfig) String() string { return proto.CompactTextString(m) }
func (*HalConfig) ProtoMessage()    {}
func (*HalConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cda1892e07435feb, []int{0}
}

func (m *HalConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HalConfig.Unmarshal(m, b)
}
func (m *HalConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HalConfig.Marshal(b, m, deterministic)
}
func (m *HalConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HalConfig.Merge(m, src)
}
func (m *HalConfig) XXX_Size() int {
	return xxx_messageInfo_HalConfig.Size(m)
}
func (m *HalConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HalConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HalConfig proto.InternalMessageInfo

func (m *HalConfig) GetPersistentStorage() *HalConfig_PersistentStorage {
	if m != nil {
		return m.PersistentStorage
	}
	return nil
}

func (m *HalConfig) GetProviders() *HalConfig_Providers {
	if m != nil {
		return m.Providers
	}
	return nil
}

type HalConfig_PersistentStorage struct {
	PersistentStoreType  string   `protobuf:"bytes,1,opt,name=persistentStoreType,proto3" json:"persistentStoreType,omitempty"`
	Gcs                  *GCS     `protobuf:"bytes,2,opt,name=gcs,proto3" json:"gcs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HalConfig_PersistentStorage) Reset()         { *m = HalConfig_PersistentStorage{} }
func (m *HalConfig_PersistentStorage) String() string { return proto.CompactTextString(m) }
func (*HalConfig_PersistentStorage) ProtoMessage()    {}
func (*HalConfig_PersistentStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cda1892e07435feb, []int{0, 0}
}

func (m *HalConfig_PersistentStorage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HalConfig_PersistentStorage.Unmarshal(m, b)
}
func (m *HalConfig_PersistentStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HalConfig_PersistentStorage.Marshal(b, m, deterministic)
}
func (m *HalConfig_PersistentStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HalConfig_PersistentStorage.Merge(m, src)
}
func (m *HalConfig_PersistentStorage) XXX_Size() int {
	return xxx_messageInfo_HalConfig_PersistentStorage.Size(m)
}
func (m *HalConfig_PersistentStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_HalConfig_PersistentStorage.DiscardUnknown(m)
}

var xxx_messageInfo_HalConfig_PersistentStorage proto.InternalMessageInfo

func (m *HalConfig_PersistentStorage) GetPersistentStoreType() string {
	if m != nil {
		return m.PersistentStoreType
	}
	return ""
}

func (m *HalConfig_PersistentStorage) GetGcs() *GCS {
	if m != nil {
		return m.Gcs
	}
	return nil
}

type HalConfig_Providers struct {
	Kubernetes           *Kubernetes     `protobuf:"bytes,1,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
	Google               *Google         `protobuf:"bytes,2,opt,name=google,proto3" json:"google,omitempty"`
	Appengine            *Appengine      `protobuf:"bytes,3,opt,name=appengine,proto3" json:"appengine,omitempty"`
	Aws                  *Aws            `protobuf:"bytes,4,opt,name=aws,proto3" json:"aws,omitempty"`
	Azure                *Azure          `protobuf:"bytes,5,opt,name=azure,proto3" json:"azure,omitempty"`
	Cloudfoundry         *CloudFoundry   `protobuf:"bytes,6,opt,name=cloudfoundry,proto3" json:"cloudfoundry,omitempty"`
	Dcos                 *Dcos           `protobuf:"bytes,7,opt,name=dcos,proto3" json:"dcos,omitempty"`
	DockerRegistry       *DockerRegistry `protobuf:"bytes,8,opt,name=dockerRegistry,proto3" json:"dockerRegistry,omitempty"`
	Ecs                  *Ecs            `protobuf:"bytes,9,opt,name=ecs,proto3" json:"ecs,omitempty"`
	Huaweicloud          *HuaweiCloud    `protobuf:"bytes,10,opt,name=huaweicloud,proto3" json:"huaweicloud,omitempty"`
	Oracle               *Oracle         `protobuf:"bytes,11,opt,name=oracle,proto3" json:"oracle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *HalConfig_Providers) Reset()         { *m = HalConfig_Providers{} }
func (m *HalConfig_Providers) String() string { return proto.CompactTextString(m) }
func (*HalConfig_Providers) ProtoMessage()    {}
func (*HalConfig_Providers) Descriptor() ([]byte, []int) {
	return fileDescriptor_cda1892e07435feb, []int{0, 1}
}

func (m *HalConfig_Providers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HalConfig_Providers.Unmarshal(m, b)
}
func (m *HalConfig_Providers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HalConfig_Providers.Marshal(b, m, deterministic)
}
func (m *HalConfig_Providers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HalConfig_Providers.Merge(m, src)
}
func (m *HalConfig_Providers) XXX_Size() int {
	return xxx_messageInfo_HalConfig_Providers.Size(m)
}
func (m *HalConfig_Providers) XXX_DiscardUnknown() {
	xxx_messageInfo_HalConfig_Providers.DiscardUnknown(m)
}

var xxx_messageInfo_HalConfig_Providers proto.InternalMessageInfo

func (m *HalConfig_Providers) GetKubernetes() *Kubernetes {
	if m != nil {
		return m.Kubernetes
	}
	return nil
}

func (m *HalConfig_Providers) GetGoogle() *Google {
	if m != nil {
		return m.Google
	}
	return nil
}

func (m *HalConfig_Providers) GetAppengine() *Appengine {
	if m != nil {
		return m.Appengine
	}
	return nil
}

func (m *HalConfig_Providers) GetAws() *Aws {
	if m != nil {
		return m.Aws
	}
	return nil
}

func (m *HalConfig_Providers) GetAzure() *Azure {
	if m != nil {
		return m.Azure
	}
	return nil
}

func (m *HalConfig_Providers) GetCloudfoundry() *CloudFoundry {
	if m != nil {
		return m.Cloudfoundry
	}
	return nil
}

func (m *HalConfig_Providers) GetDcos() *Dcos {
	if m != nil {
		return m.Dcos
	}
	return nil
}

func (m *HalConfig_Providers) GetDockerRegistry() *DockerRegistry {
	if m != nil {
		return m.DockerRegistry
	}
	return nil
}

func (m *HalConfig_Providers) GetEcs() *Ecs {
	if m != nil {
		return m.Ecs
	}
	return nil
}

func (m *HalConfig_Providers) GetHuaweicloud() *HuaweiCloud {
	if m != nil {
		return m.Huaweicloud
	}
	return nil
}

func (m *HalConfig_Providers) GetOracle() *Oracle {
	if m != nil {
		return m.Oracle
	}
	return nil
}

func init() {
	proto.RegisterType((*HalConfig)(nil), "proto.HalConfig")
	proto.RegisterType((*HalConfig_PersistentStorage)(nil), "proto.HalConfig.PersistentStorage")
	proto.RegisterType((*HalConfig_Providers)(nil), "proto.HalConfig.Providers")
}

func init() { proto.RegisterFile("halconfig.proto", fileDescriptor_cda1892e07435feb) }

var fileDescriptor_cda1892e07435feb = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6b, 0x14, 0x31,
	0x10, 0xc7, 0xa9, 0xd7, 0x3b, 0xcd, 0xec, 0x69, 0x7b, 0x29, 0xc5, 0xb0, 0x08, 0xca, 0x81, 0xe0,
	0xd3, 0xe1, 0x2f, 0xd0, 0x17, 0x1f, 0xca, 0xf9, 0x0b, 0x7c, 0xb0, 0x4c, 0x7d, 0x97, 0x34, 0x9b,
	0x6e, 0x97, 0x1e, 0x9b, 0x25, 0xd9, 0xf5, 0xa8, 0x7f, 0xb2, 0xfe, 0x13, 0x92, 0x49, 0xb2, 0xb7,
	0xdb, 0xf3, 0x29, 0xf0, 0xfd, 0x7e, 0x26, 0x33, 0xc3, 0x7c, 0xe1, 0xe8, 0x5a, 0x6e, 0x94, 0xa9,
	0xaf, 0xaa, 0x72, 0xd5, 0x58, 0xd3, 0x1a, 0x3e, 0xa5, 0x27, 0x7f, 0xdc, 0x68, 0xeb, 0x2a, 0xd7,
	0xea, 0xba, 0x75, 0xad, 0xb1, 0xb2, 0xd4, 0xc1, 0xcf, 0x8f, 0x6f, 0xba, 0x4b, 0x6d, 0x6b, 0xdd,
	0x6a, 0x17, 0x95, 0x79, 0x69, 0x4c, 0xb9, 0x49, 0xfe, 0x91, 0x6c, 0x1a, 0x5d, 0x97, 0x55, 0x9d,
	0x04, 0x26, 0xb7, 0x89, 0xcc, 0xe4, 0xef, 0xce, 0x26, 0x9d, 0xab, 0x8d, 0xe9, 0x8a, 0x2b, 0xd3,
	0xd5, 0x85, 0xbd, 0x8d, 0x1a, 0x14, 0xca, 0x24, 0xf8, 0xb4, 0x30, 0xea, 0x46, 0xdb, 0x9f, 0x56,
	0x97, 0x95, 0x6b, 0x7b, 0x84, 0x69, 0x95, 0x88, 0xc5, 0x75, 0x27, 0xb7, 0xba, 0xa2, 0x7f, 0xd2,
	0x2c, 0xc6, 0x4a, 0x95, 0x66, 0x59, 0xfe, 0x99, 0x02, 0xfb, 0x2a, 0x37, 0x6b, 0xda, 0x8f, 0x9f,
	0xc3, 0x62, 0xb7, 0xd4, 0x45, 0x58, 0x4a, 0x1c, 0x3c, 0x3b, 0x78, 0x91, 0xbd, 0x5e, 0x86, 0x82,
	0x55, 0x0f, 0xaf, 0xce, 0xef, 0x92, 0xb8, 0x5f, 0xcc, 0xdf, 0x03, 0x6b, 0xac, 0xf9, 0x55, 0x15,
	0xda, 0x3a, 0x71, 0x8f, 0x7e, 0xca, 0xf7, 0x7f, 0x4a, 0x04, 0xee, 0xe0, 0x5c, 0xc1, 0x62, 0xaf,
	0x03, 0x7f, 0x09, 0x27, 0xe3, 0x1e, 0xfa, 0xc7, 0x6d, 0x13, 0x46, 0x64, 0xf8, 0x3f, 0x8b, 0x3f,
	0x81, 0x49, 0xa9, 0x52, 0x6b, 0x88, 0xad, 0xbf, 0xac, 0x2f, 0xd0, 0xcb, 0xf9, 0xdf, 0x09, 0xb0,
	0xbe, 0x3b, 0x7f, 0x05, 0xb0, 0x3b, 0x5d, 0xdc, 0x7b, 0x11, 0x4b, 0xbe, 0xf5, 0x06, 0x0e, 0x20,
	0xfe, 0x1c, 0x66, 0xe1, 0xb6, 0xb1, 0xc3, 0xc3, 0xd4, 0x81, 0x44, 0x8c, 0x26, 0x5f, 0x01, 0xeb,
	0x8f, 0x2e, 0x26, 0x44, 0x1e, 0x47, 0xf2, 0x2c, 0xe9, 0xb8, 0x43, 0xfc, 0xd4, 0x72, 0xeb, 0xc4,
	0xe1, 0x68, 0xea, 0xb3, 0xad, 0x43, 0x2f, 0xf3, 0x25, 0x4c, 0x29, 0x26, 0x62, 0x4a, 0xfe, 0x3c,
	0xf9, 0x5e, 0xc3, 0x60, 0xf1, 0x77, 0x30, 0x1f, 0xa6, 0x47, 0xcc, 0x08, 0x3d, 0x89, 0xe8, 0xda,
	0x5b, 0x9f, 0x83, 0x85, 0x23, 0x90, 0x3f, 0x85, 0x43, 0x1f, 0x31, 0x71, 0x9f, 0x0a, 0xb2, 0x58,
	0xf0, 0x51, 0x19, 0x87, 0x64, 0xf0, 0x0f, 0xf0, 0x28, 0xe4, 0x0e, 0x63, 0xec, 0xc4, 0x03, 0x42,
	0x4f, 0x13, 0x3a, 0x32, 0xf1, 0x0e, 0xec, 0x57, 0xd3, 0xca, 0x09, 0x36, 0x5a, 0xed, 0x93, 0x72,
	0xe8, 0x65, 0xfe, 0x16, 0xb2, 0x41, 0x64, 0x05, 0x10, 0xc5, 0x53, 0x62, 0xc8, 0xa1, 0xd9, 0x71,
	0x88, 0xf9, 0x2b, 0x84, 0x54, 0x8b, 0x6c, 0x74, 0x85, 0xef, 0x24, 0x62, 0x34, 0x2f, 0x67, 0xa4,
	0xbe, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x9b, 0xcc, 0xce, 0xd2, 0x03, 0x00, 0x00,
}
