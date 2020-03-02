// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google.proto

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

// Configuration for the Google Compute Engine (GCE) provider.
type Google struct {
	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*GoogleAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	// Configuration for Spinnaker's image bakery.
	BakeryDefaults       *GoogleBakeryDefaults `protobuf:"bytes,4,opt,name=bakeryDefaults,proto3" json:"bakeryDefaults,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Google) Reset()         { *m = Google{} }
func (m *Google) String() string { return proto.CompactTextString(m) }
func (*Google) ProtoMessage()    {}
func (*Google) Descriptor() ([]byte, []int) {
	return fileDescriptor_e86eda198d7dac6f, []int{0}
}

func (m *Google) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Google.Unmarshal(m, b)
}
func (m *Google) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Google.Marshal(b, m, deterministic)
}
func (m *Google) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Google.Merge(m, src)
}
func (m *Google) XXX_Size() int {
	return xxx_messageInfo_Google.Size(m)
}
func (m *Google) XXX_DiscardUnknown() {
	xxx_messageInfo_Google.DiscardUnknown(m)
}

var xxx_messageInfo_Google proto.InternalMessageInfo

func (m *Google) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Google) GetAccounts() []*GoogleAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *Google) GetPrimaryAccount() string {
	if m != nil {
		return m.PrimaryAccount
	}
	return ""
}

func (m *Google) GetBakeryDefaults() *GoogleBakeryDefaults {
	if m != nil {
		return m.BakeryDefaults
	}
	return nil
}

// Configuration for a Spinnaker Google account. An account maps to a
// credential that can authenticate against a GCP project.
type GoogleAccount struct {
	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Deprecated): List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,3,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
	// Fiat permissions configuration.
	Permissions *Permissions `protobuf:"bytes,4,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// The GCP project this Spinnaker account will manage.
	Project string `protobuf:"bytes,5,opt,name=project,proto3" json:"project,omitempty"`
	// The path to a JSON service account that Spinnaker will use as
	// credentials. This is only needed if Spinnaker is not deployed on a
	// Google Compute Engine VM, or needs permissions not afforded to the VM
	// it is running on. See
	// https://cloud.google.com/compute/docs/access/service-accounts for
	// more information.
	JsonPath string `protobuf:"bytes,6,opt,name=jsonPath,proto3" json:"jsonPath,omitempty"`
	// Enable this flag if your GCP project has access to alpha features and
	// you want Spinnaker to take advantage of them.
	AlphaListed bool `protobuf:"varint,7,opt,name=alphaListed,proto3" json:"alphaListed,omitempty"`
	// A list of GCP projects from which Spinnaker will be able to cache and
	// deploy images. When this is omitted, it defaults to the current
	// project. Each project must have granted the IAM role
	// compute.imageUser to the service account associated with the JSON key
	// used by this account, as well as to the Google APIs service account
	// automatically created for the project being managed (should look
	// similar to 12345678912@cloudservices.gserviceaccount.com). See
	// https://cloud.google.com/compute/docs/images/sharing-images-across-projects
	// for more information about sharing images across GCP projects.
	ImageProjects []string `protobuf:"bytes,8,rep,name=imageProjects,proto3" json:"imageProjects,omitempty"`
	// Configuration for Consul.
	Consul *Consul `protobuf:"bytes,9,opt,name=consul,proto3" json:"consul,omitempty"`
	// A list of regions for caching and mutating calls. This overwrites
	// any default regions set on the provider.
	Regions []string `protobuf:"bytes,10,rep,name=regions,proto3" json:"regions,omitempty"`
	//  The path to user data template file. Spinnaker has the ability to
	//  inject userdata into generated instance templates. The mechanism is
	//  via a template file that is token replaced to provide some specifics
	//  about the deployment. See
	//  https://github.com/spinnaker/clouddriver/blob/master/clouddriver-aws/UserData.md
	//  for more information.
	UserDataFile         string   `protobuf:"bytes,11,opt,name=userDataFile,proto3" json:"userDataFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleAccount) Reset()         { *m = GoogleAccount{} }
func (m *GoogleAccount) String() string { return proto.CompactTextString(m) }
func (*GoogleAccount) ProtoMessage()    {}
func (*GoogleAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_e86eda198d7dac6f, []int{1}
}

func (m *GoogleAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleAccount.Unmarshal(m, b)
}
func (m *GoogleAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleAccount.Marshal(b, m, deterministic)
}
func (m *GoogleAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleAccount.Merge(m, src)
}
func (m *GoogleAccount) XXX_Size() int {
	return xxx_messageInfo_GoogleAccount.Size(m)
}
func (m *GoogleAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleAccount.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleAccount proto.InternalMessageInfo

func (m *GoogleAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GoogleAccount) GetRequiredGroupMemberships() []string {
	if m != nil {
		return m.RequiredGroupMemberships
	}
	return nil
}

func (m *GoogleAccount) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *GoogleAccount) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GoogleAccount) GetJsonPath() string {
	if m != nil {
		return m.JsonPath
	}
	return ""
}

func (m *GoogleAccount) GetAlphaListed() bool {
	if m != nil {
		return m.AlphaListed
	}
	return false
}

func (m *GoogleAccount) GetImageProjects() []string {
	if m != nil {
		return m.ImageProjects
	}
	return nil
}

func (m *GoogleAccount) GetConsul() *Consul {
	if m != nil {
		return m.Consul
	}
	return nil
}

func (m *GoogleAccount) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *GoogleAccount) GetUserDataFile() string {
	if m != nil {
		return m.UserDataFile
	}
	return ""
}

// Configuration for Spinnaker's image bakery.
type GoogleBakeryDefaults struct {
	// The name of the Packer template that will be used to bake images from
	// this base image. The template file must be found in this list:
	// https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer,
	// or supplied as described here: https://spinnaker.io/setup/bakery/.
	TemplateFile string `protobuf:"bytes,1,opt,name=templateFile,proto3" json:"templateFile,omitempty"`
	// List of configured base images.
	BaseImages []*GoogleBaseImageSettings `protobuf:"bytes,2,rep,name=baseImages,proto3" json:"baseImages,omitempty"`
	// The default zone in which to bake an image.
	Zone string `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"`
	// The Google Compute network ID or URL to use for the launched
	// instance. Defaults to default.
	Network string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	// If true, use the instance's internal IP instead of its external IP
	// during baking.
	UseInternalIp bool `protobuf:"varint,5,opt,name=useInternalIp,proto3" json:"useInternalIp,omitempty"`
	// The default project ID for the network and subnet to use for the VM
	// baking your image.
	NetworkProjectId     string   `protobuf:"bytes,6,opt,name=networkProjectId,proto3" json:"networkProjectId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleBakeryDefaults) Reset()         { *m = GoogleBakeryDefaults{} }
func (m *GoogleBakeryDefaults) String() string { return proto.CompactTextString(m) }
func (*GoogleBakeryDefaults) ProtoMessage()    {}
func (*GoogleBakeryDefaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_e86eda198d7dac6f, []int{2}
}

func (m *GoogleBakeryDefaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleBakeryDefaults.Unmarshal(m, b)
}
func (m *GoogleBakeryDefaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleBakeryDefaults.Marshal(b, m, deterministic)
}
func (m *GoogleBakeryDefaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleBakeryDefaults.Merge(m, src)
}
func (m *GoogleBakeryDefaults) XXX_Size() int {
	return xxx_messageInfo_GoogleBakeryDefaults.Size(m)
}
func (m *GoogleBakeryDefaults) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleBakeryDefaults.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleBakeryDefaults proto.InternalMessageInfo

func (m *GoogleBakeryDefaults) GetTemplateFile() string {
	if m != nil {
		return m.TemplateFile
	}
	return ""
}

func (m *GoogleBakeryDefaults) GetBaseImages() []*GoogleBaseImageSettings {
	if m != nil {
		return m.BaseImages
	}
	return nil
}

func (m *GoogleBakeryDefaults) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *GoogleBakeryDefaults) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *GoogleBakeryDefaults) GetUseInternalIp() bool {
	if m != nil {
		return m.UseInternalIp
	}
	return false
}

func (m *GoogleBakeryDefaults) GetNetworkProjectId() string {
	if m != nil {
		return m.NetworkProjectId
	}
	return ""
}

// Configuration for a base image for the Google provider's bakery.
type GoogleBaseImageSettings struct {
	// Base image configuration.
	BaseImage *GoogleBaseImage `protobuf:"bytes,1,opt,name=baseImage,proto3" json:"baseImage,omitempty"`
	// Image source configuration.
	VirtualizationSettings *VirtualizationSettings `protobuf:"bytes,2,opt,name=virtualizationSettings,proto3" json:"virtualizationSettings,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *GoogleBaseImageSettings) Reset()         { *m = GoogleBaseImageSettings{} }
func (m *GoogleBaseImageSettings) String() string { return proto.CompactTextString(m) }
func (*GoogleBaseImageSettings) ProtoMessage()    {}
func (*GoogleBaseImageSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_e86eda198d7dac6f, []int{3}
}

func (m *GoogleBaseImageSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleBaseImageSettings.Unmarshal(m, b)
}
func (m *GoogleBaseImageSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleBaseImageSettings.Marshal(b, m, deterministic)
}
func (m *GoogleBaseImageSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleBaseImageSettings.Merge(m, src)
}
func (m *GoogleBaseImageSettings) XXX_Size() int {
	return xxx_messageInfo_GoogleBaseImageSettings.Size(m)
}
func (m *GoogleBaseImageSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleBaseImageSettings.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleBaseImageSettings proto.InternalMessageInfo

func (m *GoogleBaseImageSettings) GetBaseImage() *GoogleBaseImage {
	if m != nil {
		return m.BaseImage
	}
	return nil
}

func (m *GoogleBaseImageSettings) GetVirtualizationSettings() *VirtualizationSettings {
	if m != nil {
		return m.VirtualizationSettings
	}
	return nil
}

// Base image configuration.
type GoogleBaseImage struct {
	// This is the identifier used by GCP to find this base
	// image.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A short description to help human operators identify the
	// image.
	ShortDescription string `protobuf:"bytes,2,opt,name=shortDescription,proto3" json:"shortDescription,omitempty"`
	// A long description to help human operators identify the
	// image.
	DetailedDescription string `protobuf:"bytes,3,opt,name=detailedDescription,proto3" json:"detailedDescription,omitempty"`
	//  This is used to help Spinnaker’s bakery download the build
	//  artifacts you supply it with. For example, specifying deb
	//  indicates that your artifacts will need to be fetched from a
	//  debian repository.
	PackageType string `protobuf:"bytes,4,opt,name=packageType,proto3" json:"packageType,omitempty"`
	// If set to true, Deck will annotate the popup tooltip to
	// indicate that the selected option represents an image family.
	ImageFamily          bool     `protobuf:"varint,5,opt,name=imageFamily,proto3" json:"imageFamily,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleBaseImage) Reset()         { *m = GoogleBaseImage{} }
func (m *GoogleBaseImage) String() string { return proto.CompactTextString(m) }
func (*GoogleBaseImage) ProtoMessage()    {}
func (*GoogleBaseImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e86eda198d7dac6f, []int{4}
}

func (m *GoogleBaseImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleBaseImage.Unmarshal(m, b)
}
func (m *GoogleBaseImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleBaseImage.Marshal(b, m, deterministic)
}
func (m *GoogleBaseImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleBaseImage.Merge(m, src)
}
func (m *GoogleBaseImage) XXX_Size() int {
	return xxx_messageInfo_GoogleBaseImage.Size(m)
}
func (m *GoogleBaseImage) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleBaseImage.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleBaseImage proto.InternalMessageInfo

func (m *GoogleBaseImage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GoogleBaseImage) GetShortDescription() string {
	if m != nil {
		return m.ShortDescription
	}
	return ""
}

func (m *GoogleBaseImage) GetDetailedDescription() string {
	if m != nil {
		return m.DetailedDescription
	}
	return ""
}

func (m *GoogleBaseImage) GetPackageType() string {
	if m != nil {
		return m.PackageType
	}
	return ""
}

func (m *GoogleBaseImage) GetImageFamily() bool {
	if m != nil {
		return m.ImageFamily
	}
	return false
}

// Image source configuration.
type VirtualizationSettings struct {
	//  The source image. If both sourceImage and sourceImageFamily
	//  are set, sourceImage will take precedence.
	SourceImage string `protobuf:"bytes,1,opt,name=sourceImage,proto3" json:"sourceImage,omitempty"`
	// The source image family to create the image from. The newest,
	// non-deprecated image is used. If both sourceImage and
	// sourceImageFamily are set, sourceImage will take precedence.
	SourceImageFamily    string   `protobuf:"bytes,2,opt,name=sourceImageFamily,proto3" json:"sourceImageFamily,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualizationSettings) Reset()         { *m = VirtualizationSettings{} }
func (m *VirtualizationSettings) String() string { return proto.CompactTextString(m) }
func (*VirtualizationSettings) ProtoMessage()    {}
func (*VirtualizationSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_e86eda198d7dac6f, []int{5}
}

func (m *VirtualizationSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualizationSettings.Unmarshal(m, b)
}
func (m *VirtualizationSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualizationSettings.Marshal(b, m, deterministic)
}
func (m *VirtualizationSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualizationSettings.Merge(m, src)
}
func (m *VirtualizationSettings) XXX_Size() int {
	return xxx_messageInfo_VirtualizationSettings.Size(m)
}
func (m *VirtualizationSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualizationSettings.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualizationSettings proto.InternalMessageInfo

func (m *VirtualizationSettings) GetSourceImage() string {
	if m != nil {
		return m.SourceImage
	}
	return ""
}

func (m *VirtualizationSettings) GetSourceImageFamily() string {
	if m != nil {
		return m.SourceImageFamily
	}
	return ""
}

// Configuration for Consul.
type Consul struct {
	// Whether Consul is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Reachable Consul node endpoint connected to the Consul cluster.
	// Defaults to localhost.
	AgentEndpoint string `protobuf:"bytes,2,opt,name=agentEndpoint,proto3" json:"agentEndpoint,omitempty"`
	// Port consul is running on for every agent. Defaults to 8500.
	AgentPort int32 `protobuf:"varint,3,opt,name=agentPort,proto3" json:"agentPort,omitempty"`
	// List of data centers to cache and keep updated. Defaults to all.
	Datacenters          []string `protobuf:"bytes,4,rep,name=datacenters,proto3" json:"datacenters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consul) Reset()         { *m = Consul{} }
func (m *Consul) String() string { return proto.CompactTextString(m) }
func (*Consul) ProtoMessage()    {}
func (*Consul) Descriptor() ([]byte, []int) {
	return fileDescriptor_e86eda198d7dac6f, []int{6}
}

func (m *Consul) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consul.Unmarshal(m, b)
}
func (m *Consul) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consul.Marshal(b, m, deterministic)
}
func (m *Consul) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consul.Merge(m, src)
}
func (m *Consul) XXX_Size() int {
	return xxx_messageInfo_Consul.Size(m)
}
func (m *Consul) XXX_DiscardUnknown() {
	xxx_messageInfo_Consul.DiscardUnknown(m)
}

var xxx_messageInfo_Consul proto.InternalMessageInfo

func (m *Consul) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Consul) GetAgentEndpoint() string {
	if m != nil {
		return m.AgentEndpoint
	}
	return ""
}

func (m *Consul) GetAgentPort() int32 {
	if m != nil {
		return m.AgentPort
	}
	return 0
}

func (m *Consul) GetDatacenters() []string {
	if m != nil {
		return m.Datacenters
	}
	return nil
}

func init() {
	proto.RegisterType((*Google)(nil), "proto.Google")
	proto.RegisterType((*GoogleAccount)(nil), "proto.GoogleAccount")
	proto.RegisterType((*GoogleBakeryDefaults)(nil), "proto.GoogleBakeryDefaults")
	proto.RegisterType((*GoogleBaseImageSettings)(nil), "proto.GoogleBaseImageSettings")
	proto.RegisterType((*GoogleBaseImage)(nil), "proto.GoogleBaseImage")
	proto.RegisterType((*VirtualizationSettings)(nil), "proto.VirtualizationSettings")
	proto.RegisterType((*Consul)(nil), "proto.Consul")
}

func init() { proto.RegisterFile("google.proto", fileDescriptor_e86eda198d7dac6f) }

var fileDescriptor_e86eda198d7dac6f = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x6e, 0x13, 0x3d,
	0x10, 0xd5, 0x26, 0x6d, 0x9a, 0x9d, 0x34, 0xfd, 0xbe, 0x9a, 0xaa, 0xac, 0xca, 0x8f, 0xa2, 0x55,
	0x41, 0x11, 0x42, 0x55, 0x55, 0x7a, 0xc5, 0x05, 0x12, 0xb4, 0xb4, 0x8a, 0x04, 0x52, 0x64, 0x7e,
	0xee, 0x9d, 0xdd, 0x21, 0x71, 0xbb, 0xb1, 0x17, 0xdb, 0x0b, 0x4a, 0x9f, 0x80, 0x27, 0xe1, 0x3d,
	0x10, 0x2f, 0x85, 0xb8, 0x42, 0xf6, 0x3a, 0xe9, 0x6e, 0x9a, 0x5c, 0x25, 0x73, 0x66, 0xce, 0xf8,
	0x8c, 0xcf, 0xac, 0x61, 0x7b, 0x2c, 0xe5, 0x38, 0xc3, 0xa3, 0x5c, 0x49, 0x23, 0xc9, 0xa6, 0xfb,
	0x39, 0xd8, 0xcd, 0x51, 0x4d, 0xb9, 0xd6, 0x5c, 0x0a, 0x5d, 0x66, 0xe2, 0x5f, 0x01, 0xb4, 0x2e,
	0x5d, 0x29, 0x89, 0x60, 0x0b, 0x05, 0x1b, 0x65, 0x98, 0x46, 0x41, 0x2f, 0xe8, 0xb7, 0xe9, 0x3c,
	0x24, 0xc7, 0xd0, 0x66, 0x49, 0x22, 0x0b, 0x61, 0x74, 0xd4, 0xe8, 0x35, 0xfb, 0x9d, 0x93, 0xbd,
	0x92, 0x7e, 0x54, 0x52, 0x5f, 0x97, 0x49, 0xba, 0xa8, 0x22, 0x4f, 0x61, 0x27, 0x57, 0x7c, 0xca,
	0xd4, 0xcc, 0xe7, 0xa2, 0x66, 0x2f, 0xe8, 0x87, 0x74, 0x09, 0x25, 0x67, 0xb0, 0x33, 0x62, 0xd7,
	0xa8, 0x66, 0xe7, 0xf8, 0x85, 0x15, 0x99, 0xd1, 0xd1, 0x46, 0x2f, 0xe8, 0x77, 0x4e, 0x1e, 0xd4,
	0xfa, 0xbf, 0xa9, 0x95, 0xd0, 0x25, 0x4a, 0xfc, 0xa7, 0x01, 0xdd, 0x9a, 0x10, 0x42, 0x60, 0x43,
	0xb0, 0x29, 0xba, 0x39, 0x42, 0xea, 0xfe, 0x93, 0x97, 0x10, 0x29, 0xfc, 0x5a, 0x70, 0x85, 0xe9,
	0xa5, 0x92, 0x45, 0xfe, 0x1e, 0xa7, 0x23, 0x54, 0x7a, 0xc2, 0x73, 0x1d, 0x35, 0x7b, 0xcd, 0x7e,
	0x48, 0xd7, 0xe6, 0xc9, 0x29, 0x74, 0x2a, 0x57, 0xe7, 0x35, 0x12, 0xaf, 0x71, 0x78, 0x9b, 0xa1,
	0xd5, 0x32, 0x7b, 0xa1, 0xb9, 0x92, 0x57, 0x98, 0x98, 0x68, 0xd3, 0x09, 0x99, 0x87, 0xe4, 0x00,
	0xda, 0x57, 0x5a, 0x8a, 0x21, 0x33, 0x93, 0xa8, 0xe5, 0x52, 0x8b, 0x98, 0xf4, 0xa0, 0xc3, 0xb2,
	0x7c, 0xc2, 0xde, 0x71, 0x6d, 0x30, 0x8d, 0xb6, 0x9c, 0x15, 0x55, 0x88, 0x1c, 0x42, 0x97, 0x4f,
	0xd9, 0x18, 0x87, 0x65, 0x37, 0x1d, 0xb5, 0x9d, 0xfc, 0x3a, 0x48, 0x9e, 0x40, 0x2b, 0x91, 0x42,
	0x17, 0x59, 0x14, 0x3a, 0xb9, 0x5d, 0x2f, 0xf7, 0xcc, 0x81, 0xd4, 0x27, 0xad, 0x48, 0x85, 0x63,
	0x37, 0x16, 0xb8, 0x36, 0xf3, 0x90, 0xc4, 0xb0, 0x5d, 0x68, 0x54, 0xe7, 0xcc, 0xb0, 0x0b, 0x9e,
	0x61, 0xd4, 0x71, 0x42, 0x6b, 0x58, 0xfc, 0x37, 0x80, 0xbd, 0x55, 0x1e, 0x59, 0xb2, 0xc1, 0x69,
	0x9e, 0x31, 0x83, 0x8e, 0x5c, 0x3a, 0x51, 0xc3, 0xc8, 0x2b, 0x80, 0x11, 0xd3, 0x38, 0xb0, 0xb2,
	0xe7, 0x8b, 0xf5, 0x78, 0xc9, 0x78, 0x9f, 0xfe, 0x80, 0xc6, 0x70, 0x31, 0xd6, 0xb4, 0xc2, 0xb0,
	0x2e, 0xdf, 0x48, 0x81, 0x7e, 0xb5, 0xdc, 0x7f, 0x3b, 0x8e, 0x40, 0xf3, 0x5d, 0xaa, 0x6b, 0xe7,
	0x52, 0x48, 0xe7, 0xa1, 0xbd, 0xb5, 0x42, 0xe3, 0x40, 0x18, 0x54, 0x82, 0x65, 0x83, 0xdc, 0x79,
	0xd2, 0xa6, 0x75, 0x90, 0x3c, 0x83, 0xff, 0x3d, 0xc1, 0x5f, 0xe4, 0x20, 0xf5, 0x0e, 0xdd, 0xc1,
	0xe3, 0x9f, 0x01, 0xdc, 0x5f, 0xa3, 0x93, 0x9c, 0x42, 0xb8, 0x50, 0xea, 0x86, 0xef, 0x9c, 0xec,
	0xaf, 0x1e, 0x8d, 0xde, 0x16, 0x92, 0x4f, 0xb0, 0xff, 0x8d, 0x2b, 0x53, 0xb0, 0x8c, 0xdf, 0x30,
	0xc3, 0xa5, 0x98, 0xf7, 0x8b, 0x1a, 0xae, 0xc5, 0x23, 0xdf, 0xe2, 0xf3, 0xca, 0x22, 0xba, 0x86,
	0x1c, 0xff, 0x0e, 0xe0, 0xbf, 0xa5, 0x53, 0xc9, 0x0e, 0x34, 0x78, 0xea, 0x6d, 0x69, 0xf0, 0xd4,
	0x0e, 0xae, 0x27, 0x52, 0x99, 0x73, 0xd4, 0x89, 0xe2, 0xb9, 0xe5, 0xbb, 0x43, 0x43, 0x7a, 0x07,
	0x27, 0xc7, 0x70, 0x2f, 0x45, 0xc3, 0x78, 0x86, 0x69, 0xb5, 0xbc, 0xf4, 0x61, 0x55, 0xca, 0x2e,
	0x75, 0xce, 0x92, 0x6b, 0x36, 0xc6, 0x8f, 0xb3, 0x1c, 0xbd, 0x35, 0x55, 0xc8, 0x56, 0xb8, 0xfd,
	0xbd, 0x60, 0x53, 0x9e, 0xcd, 0xbc, 0x39, 0x55, 0x28, 0x9e, 0xc0, 0xfe, 0xea, 0xb9, 0x2d, 0x57,
	0xcb, 0x42, 0x25, 0x95, 0xeb, 0x0e, 0x69, 0x15, 0x22, 0xcf, 0x61, 0xb7, 0x12, 0xfa, 0x33, 0xca,
	0xf1, 0xee, 0x26, 0xe2, 0x1f, 0x01, 0xb4, 0xce, 0x16, 0x9f, 0xc7, 0x9a, 0x47, 0xf1, 0x10, 0xba,
	0x6c, 0x8c, 0xc2, 0xbc, 0x15, 0x69, 0x2e, 0xb9, 0x30, 0xbe, 0x5d, 0x1d, 0x24, 0x0f, 0x21, 0x74,
	0xc0, 0x50, 0xaa, 0xf2, 0x0d, 0xdc, 0xa4, 0xb7, 0x80, 0x15, 0x9e, 0x32, 0xc3, 0x12, 0xb4, 0x0b,
	0x68, 0xdf, 0x15, 0xfb, 0x01, 0x56, 0xa1, 0x51, 0xcb, 0x19, 0xfe, 0xe2, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x59, 0xcb, 0x90, 0x96, 0xd0, 0x05, 0x00, 0x00,
}