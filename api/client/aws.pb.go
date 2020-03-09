// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aws.proto

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

// Configuration for the AWS provider.
type AwsProvider struct {
	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*AwsAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	// Your AWS Access Key ID. Note that if you are baking AMIs with Rosco, you
	// may also need to set `AwsBakeryDefaults.awsAccessKey`.
	AccessKeyId string `protobuf:"bytes,4,opt,name=accessKeyId,proto3" json:"accessKeyId,omitempty"`
	// Your AWS Secret Key. Note that if you are baking AMIs with Rosco, you
	// may also need to set `AwsBakeryDefaults.awsSecretKey`.
	SecretAccessKey string `protobuf:"bytes,5,opt,name=secretAccessKey,proto3" json:"secretAccessKey,omitempty"`
	// Halyard does not expose a command to edit this field, but defaults it to
	// `{{name}}-keypair`.
	// TODO(mneterval): Can we move to Clouddriver?
	DefaultKeyPairTemplate string `protobuf:"bytes,6,opt,name=defaultKeyPairTemplate,proto3" json:"defaultKeyPairTemplate,omitempty"`
	// List of default regions.
	DefaultRegions []*AwsRegion `protobuf:"bytes,7,rep,name=defaultRegions,proto3" json:"defaultRegions,omitempty"`
	// Defaults relevant to the AWS provider.
	// TODO(mneterval): Can we move to Clouddriver?
	Defaults *AwsDefaults `protobuf:"bytes,8,opt,name=defaults,proto3" json:"defaults,omitempty"`
	// Configuration for AWS-specific features.
	Features *AwsFeatures `protobuf:"bytes,9,opt,name=features,proto3" json:"features,omitempty"`
	// Configuration for Spinnaker's image bakery.
	BakeryDefaults       *AwsBakeryDefaults `protobuf:"bytes,11,opt,name=bakeryDefaults,proto3" json:"bakeryDefaults,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AwsProvider) Reset()         { *m = AwsProvider{} }
func (m *AwsProvider) String() string { return proto.CompactTextString(m) }
func (*AwsProvider) ProtoMessage()    {}
func (*AwsProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_30665040a56c0ab2, []int{0}
}

func (m *AwsProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsProvider.Unmarshal(m, b)
}
func (m *AwsProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsProvider.Marshal(b, m, deterministic)
}
func (m *AwsProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsProvider.Merge(m, src)
}
func (m *AwsProvider) XXX_Size() int {
	return xxx_messageInfo_AwsProvider.Size(m)
}
func (m *AwsProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsProvider.DiscardUnknown(m)
}

var xxx_messageInfo_AwsProvider proto.InternalMessageInfo

func (m *AwsProvider) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *AwsProvider) GetAccounts() []*AwsAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *AwsProvider) GetPrimaryAccount() string {
	if m != nil {
		return m.PrimaryAccount
	}
	return ""
}

func (m *AwsProvider) GetAccessKeyId() string {
	if m != nil {
		return m.AccessKeyId
	}
	return ""
}

func (m *AwsProvider) GetSecretAccessKey() string {
	if m != nil {
		return m.SecretAccessKey
	}
	return ""
}

func (m *AwsProvider) GetDefaultKeyPairTemplate() string {
	if m != nil {
		return m.DefaultKeyPairTemplate
	}
	return ""
}

func (m *AwsProvider) GetDefaultRegions() []*AwsRegion {
	if m != nil {
		return m.DefaultRegions
	}
	return nil
}

func (m *AwsProvider) GetDefaults() *AwsDefaults {
	if m != nil {
		return m.Defaults
	}
	return nil
}

func (m *AwsProvider) GetFeatures() *AwsFeatures {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *AwsProvider) GetBakeryDefaults() *AwsBakeryDefaults {
	if m != nil {
		return m.BakeryDefaults
	}
	return nil
}

// Configuration for an AWS account.
type AwsAccount struct {
	// The AWS account ID to manage. See
	// http://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html
	// for more information.
	AccountId string `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	// If set, Spinnaker will configure a credentials provider that uses AWS
	// Security Token Service to assume the specified role.
	// Examples: `user/spinnaker`, `role/spinnakerManaged`.
	AssumeRole string `protobuf:"bytes,2,opt,name=assumeRole,proto3" json:"assumeRole,omitempty"`
	// The name of the AWS key-pair to use. See
	// http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html
	// for more information.
	DefaultKeyPair string `protobuf:"bytes,3,opt,name=defaultKeyPair,proto3" json:"defaultKeyPair,omitempty"`
	// The endpoint at which your Eureka discovery system is reachable. See
	// https://github.com/Netflix/eureka for more information. Example:
	// `http://.eureka.url.to.use:8080/eureka-server/v2`.
	// Using will make Spinnaker use AWS regions in the hostname to access
	// discovery so that you can have discovery for multiple regions.
	Discovery string `protobuf:"bytes,4,opt,name=discovery,proto3" json:"discovery,omitempty"`
	// The endpoint at which Edda is reachable. Edda is not a hard dependency of
	// Spinnaker, but is helpful for reducing the request volume against AWS.
	// See https://github.com/Netflix/edda for more information.
	Edda string `protobuf:"bytes,5,opt,name=edda,proto3" json:"edda,omitempty"`
	// The environment name for the account. Many accounts can share the
	// same environment (e.g., dev, test, prod).
	Environment string `protobuf:"bytes,6,opt,name=environment,proto3" json:"environment,omitempty"`
	// Fiat permissions configuration.
	Permissions *Permissions `protobuf:"bytes,7,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Deprecated): List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,8,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
	// List of configured AWS lifecycle hooks.
	LifecycleHooks []*AwsLifecycleHook `protobuf:"bytes,9,rep,name=lifecycleHooks,proto3" json:"lifecycleHooks,omitempty"`
	// List of configured AWS regions.
	Regions []*AwsRegion `protobuf:"bytes,12,rep,name=regions,proto3" json:"regions,omitempty"`
	// The name of the account.
	Name                 string   `protobuf:"bytes,13,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsAccount) Reset()         { *m = AwsAccount{} }
func (m *AwsAccount) String() string { return proto.CompactTextString(m) }
func (*AwsAccount) ProtoMessage()    {}
func (*AwsAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_30665040a56c0ab2, []int{1}
}

func (m *AwsAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsAccount.Unmarshal(m, b)
}
func (m *AwsAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsAccount.Marshal(b, m, deterministic)
}
func (m *AwsAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsAccount.Merge(m, src)
}
func (m *AwsAccount) XXX_Size() int {
	return xxx_messageInfo_AwsAccount.Size(m)
}
func (m *AwsAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsAccount.DiscardUnknown(m)
}

var xxx_messageInfo_AwsAccount proto.InternalMessageInfo

func (m *AwsAccount) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *AwsAccount) GetAssumeRole() string {
	if m != nil {
		return m.AssumeRole
	}
	return ""
}

func (m *AwsAccount) GetDefaultKeyPair() string {
	if m != nil {
		return m.DefaultKeyPair
	}
	return ""
}

func (m *AwsAccount) GetDiscovery() string {
	if m != nil {
		return m.Discovery
	}
	return ""
}

func (m *AwsAccount) GetEdda() string {
	if m != nil {
		return m.Edda
	}
	return ""
}

func (m *AwsAccount) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *AwsAccount) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *AwsAccount) GetRequiredGroupMemberships() []string {
	if m != nil {
		return m.RequiredGroupMemberships
	}
	return nil
}

func (m *AwsAccount) GetLifecycleHooks() []*AwsLifecycleHook {
	if m != nil {
		return m.LifecycleHooks
	}
	return nil
}

func (m *AwsAccount) GetRegions() []*AwsRegion {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *AwsAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// An AWS region.
type AwsRegion struct {
	// The name of the region.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsRegion) Reset()         { *m = AwsRegion{} }
func (m *AwsRegion) String() string { return proto.CompactTextString(m) }
func (*AwsRegion) ProtoMessage()    {}
func (*AwsRegion) Descriptor() ([]byte, []int) {
	return fileDescriptor_30665040a56c0ab2, []int{2}
}

func (m *AwsRegion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsRegion.Unmarshal(m, b)
}
func (m *AwsRegion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsRegion.Marshal(b, m, deterministic)
}
func (m *AwsRegion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsRegion.Merge(m, src)
}
func (m *AwsRegion) XXX_Size() int {
	return xxx_messageInfo_AwsRegion.Size(m)
}
func (m *AwsRegion) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsRegion.DiscardUnknown(m)
}

var xxx_messageInfo_AwsRegion proto.InternalMessageInfo

func (m *AwsRegion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Configuration for Spinnaker's image bakery.
type AwsBakeryDefaults struct {
	// The default access key used to communicate with AWS.
	AwsAccessKey string `protobuf:"bytes,1,opt,name=awsAccessKey,proto3" json:"awsAccessKey,omitempty"`
	// The secret key used to communicate with AWS.
	AwsSecretKey string `protobuf:"bytes,2,opt,name=awsSecretKey,proto3" json:"awsSecretKey,omitempty"`
	// If using VPC, the default ID of the subnet, such as `subnet-12345def`,
	// where Packer will launch the EC2 instance. This field is required if you
	// are using a non-default VPC.
	AwsSubnetId string `protobuf:"bytes,3,opt,name=awsSubnetId,proto3" json:"awsSubnetId,omitempty"`
	// If launching into a VPC subnet, Packer needs the VPC ID in order to
	// create a temporary security group within the VPC. Requires `subnet_id` to
	// be set. If this default value is left blank, Packer will try to get the
	// VPC ID from `awsSubnetId`.
	AwsVpcId string `protobuf:"bytes,4,opt,name=awsVpcId,proto3" json:"awsVpcId,omitempty"`
	// If using a non-default VPC, public IP addresses are not provided by
	// default. If this is enabled, your new instance will get a Public IP.
	AwsAssociatePublicIpAddress bool `protobuf:"varint,5,opt,name=awsAssociatePublicIpAddress,proto3" json:"awsAssociatePublicIpAddress,omitempty"`
	// The default type of virtualization for the AMI you are building. This
	// option must match the supported virtualization type of
	// `AwsVirtualizationSettings.sourceAmi`.
	// Acceptable values: `pv`, `hvm`.
	DefaultVirtualizationType string `protobuf:"bytes,6,opt,name=defaultVirtualizationType,proto3" json:"defaultVirtualizationType,omitempty"`
	// List of configured base images.
	BaseImages []*AwsBaseImageSettings `protobuf:"bytes,7,rep,name=baseImages,proto3" json:"baseImages,omitempty"`
	// This is the name of the packer template that will be used to bake images
	// from this base image. The template file must be found in this list
	// https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer,
	// or supplied as described here: https://spinnaker.io/setup/bakery/.
	TemplateFile         string   `protobuf:"bytes,8,opt,name=templateFile,proto3" json:"templateFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsBakeryDefaults) Reset()         { *m = AwsBakeryDefaults{} }
func (m *AwsBakeryDefaults) String() string { return proto.CompactTextString(m) }
func (*AwsBakeryDefaults) ProtoMessage()    {}
func (*AwsBakeryDefaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_30665040a56c0ab2, []int{3}
}

func (m *AwsBakeryDefaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsBakeryDefaults.Unmarshal(m, b)
}
func (m *AwsBakeryDefaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsBakeryDefaults.Marshal(b, m, deterministic)
}
func (m *AwsBakeryDefaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsBakeryDefaults.Merge(m, src)
}
func (m *AwsBakeryDefaults) XXX_Size() int {
	return xxx_messageInfo_AwsBakeryDefaults.Size(m)
}
func (m *AwsBakeryDefaults) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsBakeryDefaults.DiscardUnknown(m)
}

var xxx_messageInfo_AwsBakeryDefaults proto.InternalMessageInfo

func (m *AwsBakeryDefaults) GetAwsAccessKey() string {
	if m != nil {
		return m.AwsAccessKey
	}
	return ""
}

func (m *AwsBakeryDefaults) GetAwsSecretKey() string {
	if m != nil {
		return m.AwsSecretKey
	}
	return ""
}

func (m *AwsBakeryDefaults) GetAwsSubnetId() string {
	if m != nil {
		return m.AwsSubnetId
	}
	return ""
}

func (m *AwsBakeryDefaults) GetAwsVpcId() string {
	if m != nil {
		return m.AwsVpcId
	}
	return ""
}

func (m *AwsBakeryDefaults) GetAwsAssociatePublicIpAddress() bool {
	if m != nil {
		return m.AwsAssociatePublicIpAddress
	}
	return false
}

func (m *AwsBakeryDefaults) GetDefaultVirtualizationType() string {
	if m != nil {
		return m.DefaultVirtualizationType
	}
	return ""
}

func (m *AwsBakeryDefaults) GetBaseImages() []*AwsBaseImageSettings {
	if m != nil {
		return m.BaseImages
	}
	return nil
}

func (m *AwsBakeryDefaults) GetTemplateFile() string {
	if m != nil {
		return m.TemplateFile
	}
	return ""
}

// Configuration for a base image for the AWS provider's bakery.
type AwsBaseImageSettings struct {
	// Base image configuration.
	BaseImage *AwsBaseImage `protobuf:"bytes,1,opt,name=baseImage,proto3" json:"baseImage,omitempty"`
	// Base image virtualization settings.
	VirtualizationSettings *AwsVirtualizationSettings `protobuf:"bytes,2,opt,name=virtualizationSettings,proto3" json:"virtualizationSettings,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                   `json:"-"`
	XXX_unrecognized       []byte                     `json:"-"`
	XXX_sizecache          int32                      `json:"-"`
}

func (m *AwsBaseImageSettings) Reset()         { *m = AwsBaseImageSettings{} }
func (m *AwsBaseImageSettings) String() string { return proto.CompactTextString(m) }
func (*AwsBaseImageSettings) ProtoMessage()    {}
func (*AwsBaseImageSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_30665040a56c0ab2, []int{4}
}

func (m *AwsBaseImageSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsBaseImageSettings.Unmarshal(m, b)
}
func (m *AwsBaseImageSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsBaseImageSettings.Marshal(b, m, deterministic)
}
func (m *AwsBaseImageSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsBaseImageSettings.Merge(m, src)
}
func (m *AwsBaseImageSettings) XXX_Size() int {
	return xxx_messageInfo_AwsBaseImageSettings.Size(m)
}
func (m *AwsBaseImageSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsBaseImageSettings.DiscardUnknown(m)
}

var xxx_messageInfo_AwsBaseImageSettings proto.InternalMessageInfo

func (m *AwsBaseImageSettings) GetBaseImage() *AwsBaseImage {
	if m != nil {
		return m.BaseImage
	}
	return nil
}

func (m *AwsBaseImageSettings) GetVirtualizationSettings() *AwsVirtualizationSettings {
	if m != nil {
		return m.VirtualizationSettings
	}
	return nil
}

// Base image configuration.
type AwsBaseImage struct {
	// This is the identifier used by AWS to find this base image.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A short description to help human operators identify the
	// image.
	ShortDescription string `protobuf:"bytes,2,opt,name=shortDescription,proto3" json:"shortDescription,omitempty"`
	// A long description to help human operators identify the
	// image.
	DetailedDescription string `protobuf:"bytes,3,opt,name=detailedDescription,proto3" json:"detailedDescription,omitempty"`
	// This is used to help Spinnaker's bakery download the build
	// artifacts you supply it with. For example, specifying deb
	// indicates that your artifacts will need to be fetched from a
	// debian repository.
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

func (m *AwsBaseImage) Reset()         { *m = AwsBaseImage{} }
func (m *AwsBaseImage) String() string { return proto.CompactTextString(m) }
func (*AwsBaseImage) ProtoMessage()    {}
func (*AwsBaseImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_30665040a56c0ab2, []int{5}
}

func (m *AwsBaseImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsBaseImage.Unmarshal(m, b)
}
func (m *AwsBaseImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsBaseImage.Marshal(b, m, deterministic)
}
func (m *AwsBaseImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsBaseImage.Merge(m, src)
}
func (m *AwsBaseImage) XXX_Size() int {
	return xxx_messageInfo_AwsBaseImage.Size(m)
}
func (m *AwsBaseImage) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsBaseImage.DiscardUnknown(m)
}

var xxx_messageInfo_AwsBaseImage proto.InternalMessageInfo

func (m *AwsBaseImage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AwsBaseImage) GetShortDescription() string {
	if m != nil {
		return m.ShortDescription
	}
	return ""
}

func (m *AwsBaseImage) GetDetailedDescription() string {
	if m != nil {
		return m.DetailedDescription
	}
	return ""
}

func (m *AwsBaseImage) GetPackageType() string {
	if m != nil {
		return m.PackageType
	}
	return ""
}

func (m *AwsBaseImage) GetTemplateFile() string {
	if m != nil {
		return m.TemplateFile
	}
	return ""
}

// Base image virtualization settings.
type AwsVirtualizationSettings struct {
	// The name of the region in which to launch the EC2 instance to create the
	// AMI.
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// The type of virtualization for the AMI you are building. This
	// option must match the supported virtualization type of `sourceAmi`.
	// Acceptable values: `pv`, `hvm`.
	VirtualizationType string `protobuf:"bytes,2,opt,name=virtualizationType,proto3" json:"virtualizationType,omitempty"`
	// The EC2 instance type to use while building the AMI, such as `t2.small`.
	InstanceType string `protobuf:"bytes,3,opt,name=instanceType,proto3" json:"instanceType,omitempty"`
	// The source AMI whose root volume will be copied and provisioned on the
	// currently running instance. This must be an EBS-backed AMI with a root
	// volume snapshot that you have access to.
	SourceAmi string `protobuf:"bytes,4,opt,name=sourceAmi,proto3" json:"sourceAmi,omitempty"`
	// The username to connect to SSH with. Required if using SSH.
	SshUserName string `protobuf:"bytes,5,opt,name=sshUserName,proto3" json:"sshUserName,omitempty"`
	// The username to use to connect to WinRM.
	WinRmUserName string `protobuf:"bytes,6,opt,name=winRmUserName,proto3" json:"winRmUserName,omitempty"`
	// The maximum hourly price to pay for a spot instance to create the AMI.
	// Spot instances are a type of instance that EC2 starts when the current
	// spot price is less than the maximum price you specify. Spot price will be
	// updated based on available spot instance capacity and current spot
	// instance requests. It may save you some costs. You can set this to `auto`
	// for Packer to automatically discover the best spot price or to "0" to use
	// an on demand instance (default).
	SpotPrice string `protobuf:"bytes,7,opt,name=spotPrice,proto3" json:"spotPrice,omitempty"`
	// Required if `spotPrice` is set to `auto`. This tells Packer what sort of
	// AMI you are launching to find the best spot price. This must be one of:
	// `Linux/UNIX`, `SUSE Linux`, `Windows`, `Linux/UNIX (Amazon VPC)`,
	// `SUSE Linux (Amazon VPC)`, `Windows (Amazon VPC)`.
	SpotPriceAutoProduct string   `protobuf:"bytes,8,opt,name=spotPriceAutoProduct,proto3" json:"spotPriceAutoProduct,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsVirtualizationSettings) Reset()         { *m = AwsVirtualizationSettings{} }
func (m *AwsVirtualizationSettings) String() string { return proto.CompactTextString(m) }
func (*AwsVirtualizationSettings) ProtoMessage()    {}
func (*AwsVirtualizationSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_30665040a56c0ab2, []int{6}
}

func (m *AwsVirtualizationSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsVirtualizationSettings.Unmarshal(m, b)
}
func (m *AwsVirtualizationSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsVirtualizationSettings.Marshal(b, m, deterministic)
}
func (m *AwsVirtualizationSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsVirtualizationSettings.Merge(m, src)
}
func (m *AwsVirtualizationSettings) XXX_Size() int {
	return xxx_messageInfo_AwsVirtualizationSettings.Size(m)
}
func (m *AwsVirtualizationSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsVirtualizationSettings.DiscardUnknown(m)
}

var xxx_messageInfo_AwsVirtualizationSettings proto.InternalMessageInfo

func (m *AwsVirtualizationSettings) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *AwsVirtualizationSettings) GetVirtualizationType() string {
	if m != nil {
		return m.VirtualizationType
	}
	return ""
}

func (m *AwsVirtualizationSettings) GetInstanceType() string {
	if m != nil {
		return m.InstanceType
	}
	return ""
}

func (m *AwsVirtualizationSettings) GetSourceAmi() string {
	if m != nil {
		return m.SourceAmi
	}
	return ""
}

func (m *AwsVirtualizationSettings) GetSshUserName() string {
	if m != nil {
		return m.SshUserName
	}
	return ""
}

func (m *AwsVirtualizationSettings) GetWinRmUserName() string {
	if m != nil {
		return m.WinRmUserName
	}
	return ""
}

func (m *AwsVirtualizationSettings) GetSpotPrice() string {
	if m != nil {
		return m.SpotPrice
	}
	return ""
}

func (m *AwsVirtualizationSettings) GetSpotPriceAutoProduct() string {
	if m != nil {
		return m.SpotPriceAutoProduct
	}
	return ""
}

// Configuration for AWS Auto Scaling Lifecycle Hooks. For more information, see:
// https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html
type AwsLifecycleHook struct {
	// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an
	// unexpected failure occurs. This parameter can be either CONTINUE or ABANDON. The default value is
	// ABANDON.
	DefaultResult string `protobuf:"bytes,1,opt,name=defaultResult,proto3" json:"defaultResult,omitempty"`
	// Set the heartbeat timeout in seconds for the lifecycle hook. Instances can remain in a wait
	// state for a finite period of time. Must be greater than or equal to 30 and less than or equal to 7200.
	// The default is 3600 (one hour).
	HeartbeatTimeout int32 `protobuf:"varint,2,opt,name=heartbeatTimeout,proto3" json:"heartbeatTimeout,omitempty"`
	// Type of lifecycle transition. Acceptable values: `autoscaling:EC2_INSTANCE_LAUNCHING`,
	// `autoscaling:EC2_INSTANCE_TERMINATING`
	LifecycleTransition string `protobuf:"bytes,3,opt,name=lifecycleTransition,proto3" json:"lifecycleTransition,omitempty"`
	// The ARN of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in
	// the transition state for the lifecycle hook. This target can be either an SQS queue or an SNS topic.
	NotificationTargetARN string `protobuf:"bytes,4,opt,name=notificationTargetARN,proto3" json:"notificationTargetARN,omitempty"`
	// The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification
	// target, for example, an Amazon SNS topic or an Amazon SQS queue.
	RoleARN              string   `protobuf:"bytes,5,opt,name=roleARN,proto3" json:"roleARN,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsLifecycleHook) Reset()         { *m = AwsLifecycleHook{} }
func (m *AwsLifecycleHook) String() string { return proto.CompactTextString(m) }
func (*AwsLifecycleHook) ProtoMessage()    {}
func (*AwsLifecycleHook) Descriptor() ([]byte, []int) {
	return fileDescriptor_30665040a56c0ab2, []int{7}
}

func (m *AwsLifecycleHook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsLifecycleHook.Unmarshal(m, b)
}
func (m *AwsLifecycleHook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsLifecycleHook.Marshal(b, m, deterministic)
}
func (m *AwsLifecycleHook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsLifecycleHook.Merge(m, src)
}
func (m *AwsLifecycleHook) XXX_Size() int {
	return xxx_messageInfo_AwsLifecycleHook.Size(m)
}
func (m *AwsLifecycleHook) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsLifecycleHook.DiscardUnknown(m)
}

var xxx_messageInfo_AwsLifecycleHook proto.InternalMessageInfo

func (m *AwsLifecycleHook) GetDefaultResult() string {
	if m != nil {
		return m.DefaultResult
	}
	return ""
}

func (m *AwsLifecycleHook) GetHeartbeatTimeout() int32 {
	if m != nil {
		return m.HeartbeatTimeout
	}
	return 0
}

func (m *AwsLifecycleHook) GetLifecycleTransition() string {
	if m != nil {
		return m.LifecycleTransition
	}
	return ""
}

func (m *AwsLifecycleHook) GetNotificationTargetARN() string {
	if m != nil {
		return m.NotificationTargetARN
	}
	return ""
}

func (m *AwsLifecycleHook) GetRoleARN() string {
	if m != nil {
		return m.RoleARN
	}
	return ""
}

// Defaults relevant to the AWS provider.
type AwsDefaults struct {
	// Halyard does not expose a command to edit this field, but defaults it to
	// `BaseIAMRole`.
	// TODO(mneterval): Can we move to Clouddriver?
	IamRole              string   `protobuf:"bytes,1,opt,name=iamRole,proto3" json:"iamRole,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsDefaults) Reset()         { *m = AwsDefaults{} }
func (m *AwsDefaults) String() string { return proto.CompactTextString(m) }
func (*AwsDefaults) ProtoMessage()    {}
func (*AwsDefaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_30665040a56c0ab2, []int{8}
}

func (m *AwsDefaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsDefaults.Unmarshal(m, b)
}
func (m *AwsDefaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsDefaults.Marshal(b, m, deterministic)
}
func (m *AwsDefaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsDefaults.Merge(m, src)
}
func (m *AwsDefaults) XXX_Size() int {
	return xxx_messageInfo_AwsDefaults.Size(m)
}
func (m *AwsDefaults) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsDefaults.DiscardUnknown(m)
}

var xxx_messageInfo_AwsDefaults proto.InternalMessageInfo

func (m *AwsDefaults) GetIamRole() string {
	if m != nil {
		return m.IamRole
	}
	return ""
}

// Configuration for AWS-specific features.
type AwsFeatures struct {
	// Configuration for AWS CloudFormation.
	CloudFormation       *AwsFeatures_CloudFormation `protobuf:"bytes,1,opt,name=cloudFormation,proto3" json:"cloudFormation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AwsFeatures) Reset()         { *m = AwsFeatures{} }
func (m *AwsFeatures) String() string { return proto.CompactTextString(m) }
func (*AwsFeatures) ProtoMessage()    {}
func (*AwsFeatures) Descriptor() ([]byte, []int) {
	return fileDescriptor_30665040a56c0ab2, []int{9}
}

func (m *AwsFeatures) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsFeatures.Unmarshal(m, b)
}
func (m *AwsFeatures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsFeatures.Marshal(b, m, deterministic)
}
func (m *AwsFeatures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsFeatures.Merge(m, src)
}
func (m *AwsFeatures) XXX_Size() int {
	return xxx_messageInfo_AwsFeatures.Size(m)
}
func (m *AwsFeatures) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsFeatures.DiscardUnknown(m)
}

var xxx_messageInfo_AwsFeatures proto.InternalMessageInfo

func (m *AwsFeatures) GetCloudFormation() *AwsFeatures_CloudFormation {
	if m != nil {
		return m.CloudFormation
	}
	return nil
}

// Configuration for AWS CloudFormation.
type AwsFeatures_CloudFormation struct {
	// Whether AWS CloudFormation is enabled.
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsFeatures_CloudFormation) Reset()         { *m = AwsFeatures_CloudFormation{} }
func (m *AwsFeatures_CloudFormation) String() string { return proto.CompactTextString(m) }
func (*AwsFeatures_CloudFormation) ProtoMessage()    {}
func (*AwsFeatures_CloudFormation) Descriptor() ([]byte, []int) {
	return fileDescriptor_30665040a56c0ab2, []int{9, 0}
}

func (m *AwsFeatures_CloudFormation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsFeatures_CloudFormation.Unmarshal(m, b)
}
func (m *AwsFeatures_CloudFormation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsFeatures_CloudFormation.Marshal(b, m, deterministic)
}
func (m *AwsFeatures_CloudFormation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsFeatures_CloudFormation.Merge(m, src)
}
func (m *AwsFeatures_CloudFormation) XXX_Size() int {
	return xxx_messageInfo_AwsFeatures_CloudFormation.Size(m)
}
func (m *AwsFeatures_CloudFormation) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsFeatures_CloudFormation.DiscardUnknown(m)
}

var xxx_messageInfo_AwsFeatures_CloudFormation proto.InternalMessageInfo

func (m *AwsFeatures_CloudFormation) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func init() {
	proto.RegisterType((*AwsProvider)(nil), "proto.AwsProvider")
	proto.RegisterType((*AwsAccount)(nil), "proto.AwsAccount")
	proto.RegisterType((*AwsRegion)(nil), "proto.AwsRegion")
	proto.RegisterType((*AwsBakeryDefaults)(nil), "proto.AwsBakeryDefaults")
	proto.RegisterType((*AwsBaseImageSettings)(nil), "proto.AwsBaseImageSettings")
	proto.RegisterType((*AwsBaseImage)(nil), "proto.AwsBaseImage")
	proto.RegisterType((*AwsVirtualizationSettings)(nil), "proto.AwsVirtualizationSettings")
	proto.RegisterType((*AwsLifecycleHook)(nil), "proto.AwsLifecycleHook")
	proto.RegisterType((*AwsDefaults)(nil), "proto.AwsDefaults")
	proto.RegisterType((*AwsFeatures)(nil), "proto.AwsFeatures")
	proto.RegisterType((*AwsFeatures_CloudFormation)(nil), "proto.AwsFeatures.CloudFormation")
}

func init() { proto.RegisterFile("aws.proto", fileDescriptor_30665040a56c0ab2) }

var fileDescriptor_30665040a56c0ab2 = []byte{
	// 988 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0xdd, 0x6e, 0x23, 0x35,
	0x14, 0x56, 0x92, 0xed, 0x26, 0x39, 0xe9, 0x86, 0xd6, 0xbb, 0x94, 0xd9, 0x2e, 0x82, 0x10, 0x21,
	0x88, 0x2a, 0x11, 0x41, 0x58, 0x21, 0x04, 0x48, 0x6c, 0x60, 0x55, 0x88, 0x16, 0x56, 0x91, 0x5b,
	0x2a, 0x6e, 0x9d, 0x99, 0xd3, 0xd4, 0xea, 0xcc, 0x78, 0xb0, 0x3d, 0x89, 0xc2, 0x35, 0x6f, 0xc2,
	0x53, 0x70, 0x85, 0xc4, 0x83, 0x70, 0xc9, 0x43, 0x70, 0x85, 0xec, 0xf1, 0xfc, 0xe5, 0x87, 0xab,
	0xc4, 0xdf, 0xf7, 0x1d, 0x7b, 0xce, 0x39, 0xdf, 0xf1, 0x0c, 0x74, 0xd9, 0x5a, 0x8d, 0x13, 0x29,
	0xb4, 0x20, 0x47, 0xf6, 0xe7, 0xfc, 0x34, 0x41, 0x19, 0x71, 0xa5, 0xb8, 0x88, 0x1d, 0x33, 0xfc,
	0xbb, 0x05, 0xbd, 0xe9, 0x5a, 0xcd, 0xa5, 0x58, 0xf1, 0x00, 0x25, 0xf1, 0xa0, 0x8d, 0x31, 0x5b,
	0x84, 0x18, 0x78, 0x8d, 0x41, 0x63, 0xd4, 0xa1, 0xf9, 0x92, 0x7c, 0x04, 0x1d, 0xe6, 0xfb, 0x22,
	0x8d, 0xb5, 0xf2, 0x9a, 0x83, 0xd6, 0xa8, 0x37, 0x39, 0xcd, 0xf6, 0x18, 0x4f, 0xd7, 0x6a, 0x9a,
	0x31, 0xb4, 0x90, 0x90, 0x0f, 0xa0, 0x9f, 0x48, 0x1e, 0x31, 0xb9, 0x71, 0x9c, 0xd7, 0x1a, 0x34,
	0x46, 0x5d, 0xba, 0x85, 0x92, 0x01, 0xf4, 0x98, 0xef, 0xa3, 0x52, 0xaf, 0x70, 0x33, 0x0b, 0xbc,
	0x07, 0x56, 0x54, 0x85, 0xc8, 0x08, 0xde, 0x50, 0xe8, 0x4b, 0xd4, 0xd3, 0x1c, 0xf4, 0x8e, 0xac,
	0x6a, 0x1b, 0x26, 0x9f, 0xc1, 0x59, 0x80, 0xb7, 0x2c, 0x0d, 0xf5, 0x2b, 0xdc, 0xcc, 0x19, 0x97,
	0xd7, 0x18, 0x25, 0x21, 0xd3, 0xe8, 0x3d, 0xb4, 0x01, 0x07, 0x58, 0xf2, 0x39, 0xf4, 0x1d, 0x43,
	0x71, 0x69, 0x8a, 0xe3, 0xb5, 0x6d, 0x82, 0x27, 0x65, 0x82, 0x19, 0x41, 0xb7, 0x74, 0x64, 0x0c,
	0x1d, 0x87, 0x28, 0xaf, 0x33, 0x68, 0x8c, 0x7a, 0x13, 0x52, 0xc6, 0xbc, 0x74, 0x0c, 0x2d, 0x34,
	0x46, 0x7f, 0x8b, 0x4c, 0xa7, 0x12, 0x95, 0xd7, 0xdd, 0xd6, 0x5f, 0x3a, 0x86, 0x16, 0x1a, 0xf2,
	0x02, 0xfa, 0x0b, 0x76, 0x8f, 0x72, 0x93, 0xef, 0xe5, 0xf5, 0x6c, 0x94, 0x57, 0x46, 0x7d, 0x53,
	0xe3, 0xe9, 0x96, 0x7e, 0xf8, 0x47, 0x0b, 0xa0, 0x6c, 0x10, 0x79, 0x1b, 0xba, 0xae, 0x45, 0xb3,
	0xac, 0xc3, 0x5d, 0x5a, 0x02, 0xe4, 0x1d, 0x00, 0xa6, 0x54, 0x1a, 0x21, 0x15, 0x21, 0x7a, 0x4d,
	0x4b, 0x57, 0x10, 0xd3, 0xd4, 0x7a, 0x09, 0xf3, 0xa6, 0xd6, 0x51, 0x73, 0x4a, 0xc0, 0x95, 0x2f,
	0x56, 0x28, 0x37, 0xae, 0xa5, 0x25, 0x40, 0x08, 0x3c, 0xc0, 0x20, 0x60, 0xae, 0x8b, 0xf6, 0xbf,
	0xb1, 0x01, 0xc6, 0x2b, 0x2e, 0x45, 0x1c, 0x61, 0xac, 0x5d, 0xbf, 0xaa, 0x10, 0x79, 0x0e, 0xbd,
	0x8a, 0x7d, 0xbd, 0x76, 0xad, 0x7a, 0xf3, 0x92, 0xa1, 0x55, 0x19, 0xf9, 0x02, 0x3c, 0x89, 0xbf,
	0xa4, 0x5c, 0x62, 0xf0, 0x9d, 0x14, 0x69, 0xf2, 0x23, 0x46, 0x0b, 0x94, 0xea, 0x8e, 0x27, 0xa6,
	0x61, 0xad, 0x51, 0x97, 0x1e, 0xe4, 0xc9, 0xd7, 0xd0, 0x0f, 0xf9, 0x2d, 0xfa, 0x1b, 0x3f, 0xc4,
	0xef, 0x85, 0xb8, 0x37, 0x2d, 0x33, 0xb6, 0x78, 0xab, 0x2c, 0xfe, 0x0f, 0x55, 0x9e, 0x6e, 0xc9,
	0xc9, 0x05, 0xb4, 0xa5, 0x33, 0xd4, 0xf1, 0x01, 0x43, 0xe5, 0x02, 0x53, 0x94, 0x98, 0x45, 0xe8,
	0x3d, 0xca, 0x8a, 0x62, 0xfe, 0x0f, 0xdf, 0x85, 0x6e, 0xa1, 0x2c, 0x04, 0x8d, 0x8a, 0xe0, 0xdf,
	0x26, 0x9c, 0xee, 0x58, 0x80, 0x0c, 0xe1, 0x98, 0xd9, 0x8e, 0xbb, 0x69, 0xc9, 0x22, 0x6a, 0x98,
	0xd3, 0x5c, 0xd9, 0x01, 0x32, 0x9a, 0x66, 0xa1, 0x29, 0x30, 0x3b, 0x9a, 0x6b, 0x75, 0x95, 0x2e,
	0x62, 0x34, 0x6e, 0x69, 0xb9, 0xd1, 0x2c, 0x21, 0x72, 0x0e, 0x1d, 0xb6, 0x56, 0x37, 0x89, 0x5f,
	0x4c, 0x6e, 0xb1, 0x26, 0x2f, 0xe0, 0x99, 0x39, 0x51, 0x29, 0xe1, 0x73, 0xa6, 0x71, 0x9e, 0x2e,
	0x42, 0xee, 0xcf, 0x92, 0x69, 0x10, 0x48, 0x54, 0xca, 0x36, 0xbf, 0x43, 0xff, 0x4f, 0x42, 0xbe,
	0x82, 0xa7, 0xce, 0x57, 0x37, 0x5c, 0xea, 0x94, 0x85, 0xfc, 0x57, 0xa6, 0xb9, 0x88, 0xaf, 0x37,
	0x49, 0x3e, 0xd1, 0x87, 0x05, 0xe4, 0x4b, 0x80, 0x05, 0x53, 0x38, 0x8b, 0xd8, 0x12, 0xf3, 0x81,
	0x7e, 0x56, 0x1d, 0x1b, 0xc7, 0x5d, 0xa1, 0xd6, 0x3c, 0x5e, 0x2a, 0x5a, 0x91, 0x9b, 0xf2, 0x68,
	0x77, 0x3b, 0x5c, 0xf2, 0x10, 0xed, 0x6c, 0x77, 0x69, 0x0d, 0x1b, 0xfe, 0xde, 0x80, 0x27, 0xfb,
	0x36, 0x22, 0x9f, 0x40, 0xb7, 0xd8, 0xca, 0x16, 0xbf, 0x37, 0x79, 0xbc, 0xe7, 0x60, 0x5a, 0xaa,
	0xc8, 0xcf, 0x70, 0xb6, 0xaa, 0xa5, 0x90, 0x6f, 0x66, 0x1b, 0xd3, 0x9b, 0x0c, 0xca, 0xf8, 0x9b,
	0xbd, 0x3a, 0x7a, 0x20, 0x7e, 0xf8, 0x57, 0x03, 0x8e, 0xab, 0xa7, 0x92, 0x3e, 0x34, 0x79, 0x3e,
	0xfa, 0x4d, 0x1e, 0x90, 0x0b, 0x38, 0x51, 0x77, 0x42, 0xea, 0x97, 0xa8, 0x7c, 0xc9, 0x13, 0x13,
	0xec, 0xdc, 0xb0, 0x83, 0x93, 0x8f, 0xe1, 0x71, 0x80, 0x9a, 0xf1, 0x10, 0x83, 0xaa, 0x3c, 0x73,
	0xc6, 0x3e, 0xca, 0x78, 0x28, 0x61, 0xfe, 0x3d, 0x5b, 0xa2, 0xed, 0x9a, 0xbb, 0xde, 0x2b, 0xd0,
	0x4e, 0xa9, 0x8f, 0xf6, 0x94, 0xfa, 0xcf, 0x26, 0x3c, 0x3d, 0x98, 0x3a, 0x39, 0x83, 0x87, 0xd9,
	0x14, 0xb9, 0xac, 0xdc, 0x8a, 0x8c, 0x81, 0xac, 0x76, 0x8d, 0x93, 0xe5, 0xb6, 0x87, 0x31, 0x4f,
	0xc2, 0x63, 0xa5, 0x59, 0xec, 0x67, 0x0f, 0x9b, 0xa5, 0x55, 0xc3, 0xcc, 0xcd, 0xa6, 0x44, 0x2a,
	0x7d, 0x9c, 0x46, 0x3c, 0xbf, 0xd9, 0x0a, 0xc0, 0x64, 0xab, 0xd4, 0xdd, 0x4f, 0x0a, 0xe5, 0x6b,
	0x33, 0xaa, 0x59, 0x2a, 0x55, 0x88, 0xbc, 0x0f, 0x8f, 0xd6, 0x3c, 0xa6, 0x51, 0xa1, 0xc9, 0x7c,
	0x5c, 0x07, 0xed, 0x29, 0x89, 0xd0, 0x73, 0xc9, 0x7d, 0xb4, 0x37, 0x9d, 0x39, 0x25, 0x07, 0xc8,
	0x04, 0x9e, 0x14, 0x8b, 0x69, 0xaa, 0xc5, 0x5c, 0x8a, 0x20, 0xf5, 0xb5, 0x33, 0xe9, 0x5e, 0x6e,
	0xf8, 0x4f, 0x03, 0x4e, 0xb6, 0xef, 0x2b, 0xf3, 0x30, 0xc5, 0xfb, 0x4c, 0xa5, 0xa1, 0x76, 0xf5,
	0xab, 0x83, 0xc6, 0x20, 0x77, 0xc8, 0xa4, 0x5e, 0x20, 0xd3, 0xd7, 0x3c, 0x42, 0x91, 0x6a, 0x5b,
	0xc4, 0x23, 0xba, 0x83, 0x1b, 0x83, 0x14, 0x77, 0xe0, 0xb5, 0x64, 0xb1, 0xe2, 0x55, 0x83, 0xec,
	0xa1, 0xc8, 0x73, 0x78, 0x33, 0x16, 0x9a, 0xdf, 0x72, 0x3f, 0x6b, 0x04, 0x93, 0x4b, 0xd4, 0x53,
	0xfa, 0xda, 0x15, 0x77, 0x3f, 0x69, 0x3e, 0x53, 0xa4, 0x08, 0xd1, 0xe8, 0xb2, 0x22, 0xe7, 0xcb,
	0xe1, 0x87, 0xf6, 0x7b, 0xa6, 0xb8, 0x0b, 0x3d, 0x68, 0x73, 0x16, 0xd9, 0xd7, 0x59, 0x96, 0x5c,
	0xbe, 0x1c, 0xfe, 0xd6, 0xb0, 0xca, 0xfc, 0xa5, 0x4b, 0x66, 0xd0, 0xf7, 0x43, 0x91, 0x06, 0x97,
	0x42, 0x46, 0xf6, 0x34, 0x37, 0xba, 0xef, 0xed, 0xbe, 0xa0, 0xc7, 0xdf, 0xd6, 0x84, 0x74, 0x2b,
	0xf0, 0xfc, 0x02, 0xfa, 0x75, 0xc5, 0xe1, 0xcf, 0xaa, 0xc5, 0x43, 0xbb, 0xfb, 0xa7, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xfb, 0xd3, 0x74, 0x96, 0xae, 0x09, 0x00, 0x00,
}
