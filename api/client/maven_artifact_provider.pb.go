// Code generated by protoc-gen-go. DO NOT EDIT.
// source: maven_artifact_provider.proto

package client

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

// Configuration for the Maven artifact provider.
type MavenArtifactProvider struct {
	// Whether the Maven artifact provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured Maven accounts.
	Accounts             []*MavenArtifactAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MavenArtifactProvider) Reset()         { *m = MavenArtifactProvider{} }
func (m *MavenArtifactProvider) String() string { return proto.CompactTextString(m) }
func (*MavenArtifactProvider) ProtoMessage()    {}
func (*MavenArtifactProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_b468ccf3cf656006, []int{0}
}

func (m *MavenArtifactProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MavenArtifactProvider.Unmarshal(m, b)
}
func (m *MavenArtifactProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MavenArtifactProvider.Marshal(b, m, deterministic)
}
func (m *MavenArtifactProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MavenArtifactProvider.Merge(m, src)
}
func (m *MavenArtifactProvider) XXX_Size() int {
	return xxx_messageInfo_MavenArtifactProvider.Size(m)
}
func (m *MavenArtifactProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_MavenArtifactProvider.DiscardUnknown(m)
}

var xxx_messageInfo_MavenArtifactProvider proto.InternalMessageInfo

func (m *MavenArtifactProvider) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *MavenArtifactProvider) GetAccounts() []*MavenArtifactAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// Configuration for a Maven artifact account.
type MavenArtifactAccount struct {
	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) Full URI for the Maven repository (e.g.,
	// `http://some.host.com/repository/path`).
	RepositoryUrl        string   `protobuf:"bytes,2,opt,name=repositoryUrl,proto3" json:"repositoryUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MavenArtifactAccount) Reset()         { *m = MavenArtifactAccount{} }
func (m *MavenArtifactAccount) String() string { return proto.CompactTextString(m) }
func (*MavenArtifactAccount) ProtoMessage()    {}
func (*MavenArtifactAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_b468ccf3cf656006, []int{1}
}

func (m *MavenArtifactAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MavenArtifactAccount.Unmarshal(m, b)
}
func (m *MavenArtifactAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MavenArtifactAccount.Marshal(b, m, deterministic)
}
func (m *MavenArtifactAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MavenArtifactAccount.Merge(m, src)
}
func (m *MavenArtifactAccount) XXX_Size() int {
	return xxx_messageInfo_MavenArtifactAccount.Size(m)
}
func (m *MavenArtifactAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MavenArtifactAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MavenArtifactAccount proto.InternalMessageInfo

func (m *MavenArtifactAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MavenArtifactAccount) GetRepositoryUrl() string {
	if m != nil {
		return m.RepositoryUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*MavenArtifactProvider)(nil), "proto.MavenArtifactProvider")
	proto.RegisterType((*MavenArtifactAccount)(nil), "proto.MavenArtifactAccount")
}

func init() { proto.RegisterFile("maven_artifact_provider.proto", fileDescriptor_b468ccf3cf656006) }

var fileDescriptor_b468ccf3cf656006 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xce, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x06, 0x60, 0x76, 0xfd, 0x5a, 0x23, 0x5e, 0x82, 0x42, 0x40, 0x84, 0x65, 0x51, 0xdc, 0x53,
	0x03, 0x7a, 0xf0, 0x5c, 0xef, 0x42, 0x29, 0x78, 0xf1, 0x52, 0xa6, 0xe9, 0xa8, 0xb1, 0x69, 0x12,
	0xa6, 0xd3, 0x82, 0xff, 0x5e, 0x4c, 0xab, 0x50, 0xf0, 0x94, 0xcc, 0xcc, 0xc3, 0x3b, 0x23, 0xae,
	0x3b, 0x18, 0xd1, 0x57, 0x40, 0x6c, 0xdf, 0xc0, 0x70, 0x15, 0x29, 0x8c, 0xb6, 0x41, 0xca, 0x22,
	0x05, 0x0e, 0xf2, 0x28, 0x3d, 0xbb, 0x4f, 0x71, 0xf9, 0xfc, 0xe3, 0xf2, 0x99, 0x15, 0xb3, 0x92,
	0x4a, 0x9c, 0xa0, 0x87, 0xda, 0x61, 0xa3, 0x56, 0xdb, 0xd5, 0x7e, 0x53, 0xfe, 0x96, 0xf2, 0x51,
	0x6c, 0xc0, 0x98, 0x30, 0x78, 0xee, 0xd5, 0x7a, 0x7b, 0xb0, 0x3f, 0xbb, 0xbf, 0x9a, 0x32, 0xb3,
	0x45, 0x52, 0x3e, 0x99, 0xf2, 0x0f, 0xef, 0x0a, 0x71, 0xf1, 0x9f, 0x90, 0x52, 0x1c, 0x7a, 0xe8,
	0x30, 0xed, 0x39, 0x2d, 0xd3, 0x5f, 0xde, 0x88, 0x73, 0xc2, 0x18, 0x7a, 0xcb, 0x81, 0xbe, 0x5e,
	0xc8, 0xa9, 0x75, 0x1a, 0x2e, 0x9b, 0x4f, 0x77, 0xaf, 0xb7, 0xef, 0x96, 0x3f, 0x86, 0x3a, 0x33,
	0xa1, 0xd3, 0x7d, 0xb4, 0xde, 0x43, 0x8b, 0xa4, 0x5b, 0x87, 0xc0, 0x1a, 0xa2, 0xd5, 0xc6, 0x59,
	0xf4, 0x5c, 0x1f, 0xa7, 0x03, 0x1f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xbc, 0xc9, 0x6b,
	0x15, 0x01, 0x00, 0x00,
}
