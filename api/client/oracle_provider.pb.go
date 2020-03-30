// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oracle_provider.proto

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

// Configuration for the Oracle provider.
type OracleProvider struct {
	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*OracleAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	// Configuration for Spinnaker's image bakery.
	BakeryDefaults       *OracleBakeryDefaults `protobuf:"bytes,4,opt,name=bakeryDefaults,proto3" json:"bakeryDefaults,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OracleProvider) Reset()         { *m = OracleProvider{} }
func (m *OracleProvider) String() string { return proto.CompactTextString(m) }
func (*OracleProvider) ProtoMessage()    {}
func (*OracleProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_31bf2aa6e218d42d, []int{0}
}

func (m *OracleProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleProvider.Unmarshal(m, b)
}
func (m *OracleProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleProvider.Marshal(b, m, deterministic)
}
func (m *OracleProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleProvider.Merge(m, src)
}
func (m *OracleProvider) XXX_Size() int {
	return xxx_messageInfo_OracleProvider.Size(m)
}
func (m *OracleProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleProvider.DiscardUnknown(m)
}

var xxx_messageInfo_OracleProvider proto.InternalMessageInfo

func (m *OracleProvider) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *OracleProvider) GetAccounts() []*OracleAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *OracleProvider) GetPrimaryAccount() string {
	if m != nil {
		return m.PrimaryAccount
	}
	return ""
}

func (m *OracleProvider) GetBakeryDefaults() *OracleBakeryDefaults {
	if m != nil {
		return m.BakeryDefaults
	}
	return nil
}

// Configuration for an Oracle account. An account maps to an Oracle Cloud
// Infrastructure (OCI) user.
type OracleAccount struct {
	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Deprecated) List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,2,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
	// Fiat permissions configuration.
	Permissions *Permissions `protobuf:"bytes,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Required) The OCID of the Oracle Compartment to use.
	CompartmentId string `protobuf:"bytes,4,opt,name=compartmentId,proto3" json:"compartmentId,omitempty"`
	// The environment name for the account. Many accounts can share the
	// same environment (e.g., dev, test, prod).
	Environment string `protobuf:"bytes,5,opt,name=environment,proto3" json:"environment,omitempty"`
	// (Required) Fingerprint of the public key.
	Fingerprint string `protobuf:"bytes,6,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Passphrase used for the private key, if it is encrypted.
	PrivateKeyPassphrase string `protobuf:"bytes,7,opt,name=privateKeyPassphrase,proto3" json:"privateKeyPassphrase,omitempty"`
	// (Required) An Oracle region (e.g., `us-phoenix-1`).
	Region string `protobuf:"bytes,8,opt,name=region,proto3" json:"region,omitempty"`
	// (Required) Path to the private key in PEM format.
	SshPrivateKeyFilePath string `protobuf:"bytes,9,opt,name=sshPrivateKeyFilePath,proto3" json:"sshPrivateKeyFilePath,omitempty"`
	// (Required) The OCID of the Oracle Tenancy to use.
	TenancyId string `protobuf:"bytes,10,opt,name=tenancyId,proto3" json:"tenancyId,omitempty"`
	// (Required) The OCID of the Oracle User with which to authenticate.
	UserId               string   `protobuf:"bytes,11,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OracleAccount) Reset()         { *m = OracleAccount{} }
func (m *OracleAccount) String() string { return proto.CompactTextString(m) }
func (*OracleAccount) ProtoMessage()    {}
func (*OracleAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_31bf2aa6e218d42d, []int{1}
}

func (m *OracleAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleAccount.Unmarshal(m, b)
}
func (m *OracleAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleAccount.Marshal(b, m, deterministic)
}
func (m *OracleAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleAccount.Merge(m, src)
}
func (m *OracleAccount) XXX_Size() int {
	return xxx_messageInfo_OracleAccount.Size(m)
}
func (m *OracleAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleAccount.DiscardUnknown(m)
}

var xxx_messageInfo_OracleAccount proto.InternalMessageInfo

func (m *OracleAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OracleAccount) GetRequiredGroupMemberships() []string {
	if m != nil {
		return m.RequiredGroupMemberships
	}
	return nil
}

func (m *OracleAccount) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *OracleAccount) GetCompartmentId() string {
	if m != nil {
		return m.CompartmentId
	}
	return ""
}

func (m *OracleAccount) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *OracleAccount) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *OracleAccount) GetPrivateKeyPassphrase() string {
	if m != nil {
		return m.PrivateKeyPassphrase
	}
	return ""
}

func (m *OracleAccount) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *OracleAccount) GetSshPrivateKeyFilePath() string {
	if m != nil {
		return m.SshPrivateKeyFilePath
	}
	return ""
}

func (m *OracleAccount) GetTenancyId() string {
	if m != nil {
		return m.TenancyId
	}
	return ""
}

func (m *OracleAccount) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// Configuration for Spinnaker's image bakery.
type OracleBakeryDefaults struct {
	// The name of the Packer template that will be used to bake images from
	// this base image. The template file must be found in this list:
	// https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer,
	// or supplied as described here: https://spinnaker.io/setup/bakery/.
	TemplateFile string `protobuf:"bytes,1,opt,name=templateFile,proto3" json:"templateFile,omitempty"`
	// List of configured base images.
	BaseImages []*OracleBaseImageSettings `protobuf:"bytes,2,rep,name=baseImages,proto3" json:"baseImages,omitempty"`
	// (Required) The name of the Availability Domain within which a new instance
	// is launched and provisioned.
	AvailabilityDomain string `protobuf:"bytes,3,opt,name=availabilityDomain,proto3" json:"availabilityDomain,omitempty"`
	// (Required) The shape for a newly created instance.
	InstanceShape string `protobuf:"bytes,4,opt,name=instanceShape,proto3" json:"instanceShape,omitempty"`
	// (Required) The name of the subnet within which a new instance is launched
	// and provisioned.
	SubnetId             string   `protobuf:"bytes,5,opt,name=subnetId,proto3" json:"subnetId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OracleBakeryDefaults) Reset()         { *m = OracleBakeryDefaults{} }
func (m *OracleBakeryDefaults) String() string { return proto.CompactTextString(m) }
func (*OracleBakeryDefaults) ProtoMessage()    {}
func (*OracleBakeryDefaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_31bf2aa6e218d42d, []int{2}
}

func (m *OracleBakeryDefaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleBakeryDefaults.Unmarshal(m, b)
}
func (m *OracleBakeryDefaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleBakeryDefaults.Marshal(b, m, deterministic)
}
func (m *OracleBakeryDefaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleBakeryDefaults.Merge(m, src)
}
func (m *OracleBakeryDefaults) XXX_Size() int {
	return xxx_messageInfo_OracleBakeryDefaults.Size(m)
}
func (m *OracleBakeryDefaults) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleBakeryDefaults.DiscardUnknown(m)
}

var xxx_messageInfo_OracleBakeryDefaults proto.InternalMessageInfo

func (m *OracleBakeryDefaults) GetTemplateFile() string {
	if m != nil {
		return m.TemplateFile
	}
	return ""
}

func (m *OracleBakeryDefaults) GetBaseImages() []*OracleBaseImageSettings {
	if m != nil {
		return m.BaseImages
	}
	return nil
}

func (m *OracleBakeryDefaults) GetAvailabilityDomain() string {
	if m != nil {
		return m.AvailabilityDomain
	}
	return ""
}

func (m *OracleBakeryDefaults) GetInstanceShape() string {
	if m != nil {
		return m.InstanceShape
	}
	return ""
}

func (m *OracleBakeryDefaults) GetSubnetId() string {
	if m != nil {
		return m.SubnetId
	}
	return ""
}

// Configuration for a base image for the Oracle provider's bakery.
type OracleBaseImageSettings struct {
	// Oracle base image configuration.
	BaseImage *OracleBaseImage `protobuf:"bytes,1,opt,name=baseImage,proto3" json:"baseImage,omitempty"`
	// Oracle virtualization settings.
	VirtualizationSettings *OracleVirtualizationSettings `protobuf:"bytes,2,opt,name=virtualizationSettings,proto3" json:"virtualizationSettings,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                      `json:"-"`
	XXX_unrecognized       []byte                        `json:"-"`
	XXX_sizecache          int32                         `json:"-"`
}

func (m *OracleBaseImageSettings) Reset()         { *m = OracleBaseImageSettings{} }
func (m *OracleBaseImageSettings) String() string { return proto.CompactTextString(m) }
func (*OracleBaseImageSettings) ProtoMessage()    {}
func (*OracleBaseImageSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_31bf2aa6e218d42d, []int{3}
}

func (m *OracleBaseImageSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleBaseImageSettings.Unmarshal(m, b)
}
func (m *OracleBaseImageSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleBaseImageSettings.Marshal(b, m, deterministic)
}
func (m *OracleBaseImageSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleBaseImageSettings.Merge(m, src)
}
func (m *OracleBaseImageSettings) XXX_Size() int {
	return xxx_messageInfo_OracleBaseImageSettings.Size(m)
}
func (m *OracleBaseImageSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleBaseImageSettings.DiscardUnknown(m)
}

var xxx_messageInfo_OracleBaseImageSettings proto.InternalMessageInfo

func (m *OracleBaseImageSettings) GetBaseImage() *OracleBaseImage {
	if m != nil {
		return m.BaseImage
	}
	return nil
}

func (m *OracleBaseImageSettings) GetVirtualizationSettings() *OracleVirtualizationSettings {
	if m != nil {
		return m.VirtualizationSettings
	}
	return nil
}

// Oracle base image configuration.
type OracleBaseImage struct {
	// The name of the base image.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A short description to help human operators identify the
	// image.
	ShortDescription string `protobuf:"bytes,2,opt,name=shortDescription,proto3" json:"shortDescription,omitempty"`
	// A long description to help human operators identify the
	// image.
	DetailedDescription string `protobuf:"bytes,3,opt,name=detailedDescription,proto3" json:"detailedDescription,omitempty"`
	//  This is used to help Spinnaker's bakery download the build
	//  artifacts you supply it with. For example, specifying deb
	//  indicates that your artifacts will need to be fetched from a
	//  debian repository.
	PackageType string `protobuf:"bytes,4,opt,name=packageType,proto3" json:"packageType,omitempty"`
	// The name of the Packer template that will be used to bake images from
	// this base image. The template file must be found in this list:
	// https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer,
	// or supplied as described here: https://spinnaker.io/setup/bakery/.
	TemplateFile         string   `protobuf:"bytes,5,opt,name=templateFile,proto3" json:"templateFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OracleBaseImage) Reset()         { *m = OracleBaseImage{} }
func (m *OracleBaseImage) String() string { return proto.CompactTextString(m) }
func (*OracleBaseImage) ProtoMessage()    {}
func (*OracleBaseImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_31bf2aa6e218d42d, []int{4}
}

func (m *OracleBaseImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleBaseImage.Unmarshal(m, b)
}
func (m *OracleBaseImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleBaseImage.Marshal(b, m, deterministic)
}
func (m *OracleBaseImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleBaseImage.Merge(m, src)
}
func (m *OracleBaseImage) XXX_Size() int {
	return xxx_messageInfo_OracleBaseImage.Size(m)
}
func (m *OracleBaseImage) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleBaseImage.DiscardUnknown(m)
}

var xxx_messageInfo_OracleBaseImage proto.InternalMessageInfo

func (m *OracleBaseImage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OracleBaseImage) GetShortDescription() string {
	if m != nil {
		return m.ShortDescription
	}
	return ""
}

func (m *OracleBaseImage) GetDetailedDescription() string {
	if m != nil {
		return m.DetailedDescription
	}
	return ""
}

func (m *OracleBaseImage) GetPackageType() string {
	if m != nil {
		return m.PackageType
	}
	return ""
}

func (m *OracleBaseImage) GetTemplateFile() string {
	if m != nil {
		return m.TemplateFile
	}
	return ""
}

// Oracle virtualization settings.
type OracleVirtualizationSettings struct {
	// (Required) The OCID of the base image ID for the baking configuration.
	BaseImageId string `protobuf:"bytes,1,opt,name=baseImageId,proto3" json:"baseImageId,omitempty"`
	// (Required) The ssh username for the baking configuration.
	SshUserName          string   `protobuf:"bytes,2,opt,name=sshUserName,proto3" json:"sshUserName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OracleVirtualizationSettings) Reset()         { *m = OracleVirtualizationSettings{} }
func (m *OracleVirtualizationSettings) String() string { return proto.CompactTextString(m) }
func (*OracleVirtualizationSettings) ProtoMessage()    {}
func (*OracleVirtualizationSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_31bf2aa6e218d42d, []int{5}
}

func (m *OracleVirtualizationSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleVirtualizationSettings.Unmarshal(m, b)
}
func (m *OracleVirtualizationSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleVirtualizationSettings.Marshal(b, m, deterministic)
}
func (m *OracleVirtualizationSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleVirtualizationSettings.Merge(m, src)
}
func (m *OracleVirtualizationSettings) XXX_Size() int {
	return xxx_messageInfo_OracleVirtualizationSettings.Size(m)
}
func (m *OracleVirtualizationSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleVirtualizationSettings.DiscardUnknown(m)
}

var xxx_messageInfo_OracleVirtualizationSettings proto.InternalMessageInfo

func (m *OracleVirtualizationSettings) GetBaseImageId() string {
	if m != nil {
		return m.BaseImageId
	}
	return ""
}

func (m *OracleVirtualizationSettings) GetSshUserName() string {
	if m != nil {
		return m.SshUserName
	}
	return ""
}

func init() {
	proto.RegisterType((*OracleProvider)(nil), "proto.OracleProvider")
	proto.RegisterType((*OracleAccount)(nil), "proto.OracleAccount")
	proto.RegisterType((*OracleBakeryDefaults)(nil), "proto.OracleBakeryDefaults")
	proto.RegisterType((*OracleBaseImageSettings)(nil), "proto.OracleBaseImageSettings")
	proto.RegisterType((*OracleBaseImage)(nil), "proto.OracleBaseImage")
	proto.RegisterType((*OracleVirtualizationSettings)(nil), "proto.OracleVirtualizationSettings")
}

func init() { proto.RegisterFile("oracle_provider.proto", fileDescriptor_31bf2aa6e218d42d) }

var fileDescriptor_31bf2aa6e218d42d = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0xd3, 0xbf, 0xf8, 0xe6, 0x6b, 0xbe, 0xef, 0x1b, 0xda, 0x62, 0x95, 0x0a, 0x45, 0xe1,
	0x2f, 0x62, 0x91, 0x54, 0xa1, 0x2b, 0x16, 0x48, 0x94, 0x0a, 0x14, 0x21, 0x20, 0x72, 0x81, 0x05,
	0x2c, 0xd0, 0xd8, 0xbe, 0x8d, 0xaf, 0x6a, 0x8f, 0x87, 0x99, 0x71, 0xa4, 0xf0, 0x4c, 0xec, 0x79,
	0x02, 0x16, 0xbc, 0x0b, 0x0f, 0x81, 0x3c, 0x76, 0x52, 0x27, 0x4d, 0x59, 0x25, 0x73, 0xce, 0xb9,
	0xf7, 0xce, 0x99, 0x73, 0x65, 0xd8, 0xcf, 0x14, 0x0f, 0x13, 0xfc, 0x22, 0x55, 0x36, 0xa5, 0x08,
	0x55, 0x5f, 0xaa, 0xcc, 0x64, 0x6c, 0xcb, 0xfe, 0x1c, 0xfe, 0x2f, 0x51, 0xa5, 0xa4, 0x35, 0x65,
	0x42, 0x97, 0x4c, 0xf7, 0x97, 0x03, 0xed, 0x77, 0xb6, 0x66, 0x5c, 0x95, 0x30, 0x0f, 0x76, 0x50,
	0xf0, 0x20, 0xc1, 0xc8, 0x73, 0x3a, 0x4e, 0xaf, 0xe9, 0xcf, 0x8f, 0xec, 0x18, 0x9a, 0x3c, 0x0c,
	0xb3, 0x5c, 0x18, 0xed, 0x35, 0x3a, 0x1b, 0xbd, 0xd6, 0x70, 0xaf, 0x6c, 0xd3, 0x2f, 0x5b, 0x3c,
	0x2f, 0x49, 0x7f, 0xa1, 0x62, 0x0f, 0xa1, 0x2d, 0x15, 0xa5, 0x5c, 0xcd, 0x2a, 0xce, 0xdb, 0xe8,
	0x38, 0x3d, 0xd7, 0x5f, 0x41, 0xd9, 0x0b, 0x68, 0x07, 0xfc, 0x12, 0xd5, 0xec, 0x0c, 0x2f, 0x78,
	0x9e, 0x18, 0xed, 0x6d, 0x76, 0x9c, 0x5e, 0x6b, 0x78, 0x67, 0xa9, 0xff, 0xe9, 0x92, 0xc4, 0x5f,
	0x29, 0xe9, 0xfe, 0xd8, 0x80, 0xdd, 0xa5, 0x8b, 0x30, 0x06, 0x9b, 0x82, 0xa7, 0x68, 0x7d, 0xb8,
	0xbe, 0xfd, 0xcf, 0x9e, 0x82, 0xa7, 0xf0, 0x6b, 0x4e, 0x0a, 0xa3, 0x57, 0x2a, 0xcb, 0xe5, 0x1b,
	0x4c, 0x03, 0x54, 0x3a, 0x26, 0x59, 0x9a, 0x72, 0xfd, 0x1b, 0x79, 0x76, 0x02, 0xad, 0xda, 0x13,
	0x5a, 0x2f, 0xad, 0x21, 0xab, 0xee, 0x38, 0xbe, 0x62, 0xfc, 0xba, 0x8c, 0xdd, 0x87, 0xdd, 0x30,
	0x4b, 0x25, 0x57, 0x26, 0x45, 0x61, 0x46, 0x91, 0xf5, 0xe6, 0xfa, 0xcb, 0x20, 0xeb, 0x40, 0x0b,
	0xc5, 0x94, 0x54, 0x26, 0x0a, 0xc0, 0xdb, 0xb2, 0x9a, 0x3a, 0x54, 0x28, 0x2e, 0x48, 0x4c, 0x50,
	0x49, 0x45, 0xc2, 0x78, 0xdb, 0xa5, 0xa2, 0x06, 0xb1, 0x21, 0xec, 0x49, 0x45, 0x53, 0x6e, 0xf0,
	0x35, 0xce, 0xc6, 0x5c, 0x6b, 0x19, 0x2b, 0xae, 0xd1, 0xdb, 0xb1, 0xd2, 0xb5, 0x1c, 0x3b, 0x80,
	0x6d, 0x85, 0x13, 0xca, 0x84, 0xd7, 0xb4, 0xaa, 0xea, 0xc4, 0x4e, 0x60, 0x5f, 0xeb, 0x78, 0xbc,
	0x28, 0x79, 0x49, 0x09, 0x8e, 0xb9, 0x89, 0x3d, 0xd7, 0xca, 0xd6, 0x93, 0xec, 0x08, 0x5c, 0x83,
	0x82, 0x8b, 0x70, 0x36, 0x8a, 0x3c, 0xb0, 0xca, 0x2b, 0xa0, 0x98, 0x95, 0x6b, 0x54, 0xa3, 0xc8,
	0x6b, 0x95, 0xb3, 0xca, 0x53, 0xf7, 0xb7, 0x03, 0x7b, 0xeb, 0x22, 0x66, 0x5d, 0xf8, 0xc7, 0x60,
	0x2a, 0x13, 0x6e, 0xb0, 0x18, 0x51, 0x05, 0xb9, 0x84, 0xb1, 0x67, 0x00, 0x01, 0xd7, 0x38, 0x4a,
	0xf9, 0x04, 0xe7, 0x7b, 0x79, 0x77, 0x65, 0x6f, 0x2a, 0xfa, 0x1c, 0x8d, 0x21, 0x31, 0xd1, 0x7e,
	0xad, 0x82, 0xf5, 0x81, 0xf1, 0x29, 0xa7, 0x84, 0x07, 0x94, 0x90, 0x99, 0x9d, 0x65, 0x29, 0x27,
	0x51, 0xed, 0xe9, 0x1a, 0xa6, 0x88, 0x93, 0x84, 0x36, 0x5c, 0x84, 0x78, 0x1e, 0x73, 0x89, 0xf3,
	0x38, 0x97, 0x40, 0x76, 0x08, 0x4d, 0x9d, 0x07, 0x02, 0x8b, 0xbc, 0xcb, 0x2c, 0x17, 0xe7, 0xee,
	0x77, 0x07, 0x6e, 0xdf, 0x70, 0x33, 0x76, 0x02, 0xee, 0xe2, 0x6e, 0xd6, 0x6e, 0x6b, 0x78, 0xb0,
	0xde, 0x8c, 0x7f, 0x25, 0x64, 0x9f, 0xe1, 0x60, 0x4a, 0xca, 0xe4, 0x3c, 0xa1, 0x6f, 0xdc, 0x50,
	0x26, 0xe6, 0xfd, 0xbc, 0x86, 0x6d, 0x71, 0x6f, 0xa9, 0xc5, 0xc7, 0xb5, 0x52, 0xff, 0x86, 0x16,
	0xdd, 0x9f, 0x0e, 0xfc, 0xbb, 0x32, 0x9b, 0xb5, 0xa1, 0x41, 0x51, 0x15, 0x47, 0x83, 0x22, 0xf6,
	0x18, 0xfe, 0xd3, 0x71, 0xa6, 0xcc, 0x19, 0xea, 0x50, 0x91, 0x2c, 0xea, 0xed, 0x68, 0xd7, 0xbf,
	0x86, 0xb3, 0x63, 0xb8, 0x15, 0xa1, 0xe1, 0x94, 0x60, 0x54, 0x97, 0x97, 0x2f, 0xbe, 0x8e, 0x2a,
	0x36, 0x5f, 0xf2, 0xf0, 0x92, 0x4f, 0xf0, 0xfd, 0x6c, 0xf1, 0xe0, 0x75, 0xe8, 0xda, 0xa2, 0x6c,
	0x5d, 0x5f, 0x94, 0x6e, 0x00, 0x47, 0x7f, 0xf3, 0x5f, 0x4c, 0x59, 0xbc, 0xe8, 0x68, 0x6e, 0xae,
	0x0e, 0x15, 0x0a, 0xad, 0xe3, 0x0f, 0x1a, 0xd5, 0xdb, 0xe2, 0xb3, 0x52, 0x1a, 0xac, 0x43, 0xa7,
	0x8f, 0x3e, 0x3d, 0x98, 0x90, 0x89, 0xf3, 0xa0, 0x1f, 0x66, 0xe9, 0x40, 0x4b, 0x12, 0xa2, 0xd8,
	0xe8, 0xc1, 0x65, 0x82, 0xdc, 0x0c, 0xb8, 0xa4, 0x41, 0x98, 0x10, 0x0a, 0x13, 0x6c, 0xdb, 0x40,
	0x9e, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x59, 0x4a, 0xfb, 0xbb, 0xb2, 0x05, 0x00, 0x00,
}
