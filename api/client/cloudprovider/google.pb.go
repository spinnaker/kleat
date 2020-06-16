// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: cloudprovider/google.proto

package cloudprovider

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Configuration for the Google Compute Engine (GCE) provider.
type GoogleComputeEngine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the provider is enabled.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*GoogleComputeEngineAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	// Configuration for Spinnaker's image bakery.
	BakeryDefaults *GoogleBakeryDefaults `protobuf:"bytes,4,opt,name=bakeryDefaults,proto3" json:"bakeryDefaults,omitempty"`
}

func (x *GoogleComputeEngine) Reset() {
	*x = GoogleComputeEngine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_google_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleComputeEngine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleComputeEngine) ProtoMessage() {}

func (x *GoogleComputeEngine) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_google_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleComputeEngine.ProtoReflect.Descriptor instead.
func (*GoogleComputeEngine) Descriptor() ([]byte, []int) {
	return file_cloudprovider_google_proto_rawDescGZIP(), []int{0}
}

func (x *GoogleComputeEngine) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *GoogleComputeEngine) GetAccounts() []*GoogleComputeEngineAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *GoogleComputeEngine) GetPrimaryAccount() string {
	if x != nil {
		return x.PrimaryAccount
	}
	return ""
}

func (x *GoogleComputeEngine) GetBakeryDefaults() *GoogleBakeryDefaults {
	if x != nil {
		return x.BakeryDefaults
	}
	return nil
}

// Configuration for a Spinnaker Google account. An account maps to a
// credential that can authenticate against a GCP project.
type GoogleComputeEngineAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Deprecated): List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,3,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
	// Fiat permissions configuration.
	Permissions *client.Permissions `protobuf:"bytes,4,opt,name=permissions,proto3" json:"permissions,omitempty"`
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
	AlphaListed *wrappers.BoolValue `protobuf:"bytes,7,opt,name=alphaListed,proto3" json:"alphaListed,omitempty"`
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
	UserDataFile string `protobuf:"bytes,11,opt,name=userDataFile,proto3" json:"userDataFile,omitempty"`
}

func (x *GoogleComputeEngineAccount) Reset() {
	*x = GoogleComputeEngineAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_google_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleComputeEngineAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleComputeEngineAccount) ProtoMessage() {}

func (x *GoogleComputeEngineAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_google_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleComputeEngineAccount.ProtoReflect.Descriptor instead.
func (*GoogleComputeEngineAccount) Descriptor() ([]byte, []int) {
	return file_cloudprovider_google_proto_rawDescGZIP(), []int{1}
}

func (x *GoogleComputeEngineAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoogleComputeEngineAccount) GetRequiredGroupMemberships() []string {
	if x != nil {
		return x.RequiredGroupMemberships
	}
	return nil
}

func (x *GoogleComputeEngineAccount) GetPermissions() *client.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *GoogleComputeEngineAccount) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GoogleComputeEngineAccount) GetJsonPath() string {
	if x != nil {
		return x.JsonPath
	}
	return ""
}

func (x *GoogleComputeEngineAccount) GetAlphaListed() *wrappers.BoolValue {
	if x != nil {
		return x.AlphaListed
	}
	return nil
}

func (x *GoogleComputeEngineAccount) GetImageProjects() []string {
	if x != nil {
		return x.ImageProjects
	}
	return nil
}

func (x *GoogleComputeEngineAccount) GetConsul() *Consul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *GoogleComputeEngineAccount) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *GoogleComputeEngineAccount) GetUserDataFile() string {
	if x != nil {
		return x.UserDataFile
	}
	return ""
}

// Configuration for Spinnaker's image bakery.
type GoogleBakeryDefaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	UseInternalIp *wrappers.BoolValue `protobuf:"bytes,5,opt,name=useInternalIp,proto3" json:"useInternalIp,omitempty"`
	// The default project ID for the network and subnet to use for the VM
	// baking your image.
	NetworkProjectId string `protobuf:"bytes,6,opt,name=networkProjectId,proto3" json:"networkProjectId,omitempty"`
}

func (x *GoogleBakeryDefaults) Reset() {
	*x = GoogleBakeryDefaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_google_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleBakeryDefaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleBakeryDefaults) ProtoMessage() {}

func (x *GoogleBakeryDefaults) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_google_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleBakeryDefaults.ProtoReflect.Descriptor instead.
func (*GoogleBakeryDefaults) Descriptor() ([]byte, []int) {
	return file_cloudprovider_google_proto_rawDescGZIP(), []int{2}
}

func (x *GoogleBakeryDefaults) GetTemplateFile() string {
	if x != nil {
		return x.TemplateFile
	}
	return ""
}

func (x *GoogleBakeryDefaults) GetBaseImages() []*GoogleBaseImageSettings {
	if x != nil {
		return x.BaseImages
	}
	return nil
}

func (x *GoogleBakeryDefaults) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *GoogleBakeryDefaults) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *GoogleBakeryDefaults) GetUseInternalIp() *wrappers.BoolValue {
	if x != nil {
		return x.UseInternalIp
	}
	return nil
}

func (x *GoogleBakeryDefaults) GetNetworkProjectId() string {
	if x != nil {
		return x.NetworkProjectId
	}
	return ""
}

// Configuration for a base image for the Google provider's bakery.
type GoogleBaseImageSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base image configuration.
	BaseImage *GoogleBaseImage `protobuf:"bytes,1,opt,name=baseImage,proto3" json:"baseImage,omitempty"`
	// Image source configuration.
	VirtualizationSettings *VirtualizationSettings `protobuf:"bytes,2,opt,name=virtualizationSettings,proto3" json:"virtualizationSettings,omitempty"`
}

func (x *GoogleBaseImageSettings) Reset() {
	*x = GoogleBaseImageSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_google_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleBaseImageSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleBaseImageSettings) ProtoMessage() {}

func (x *GoogleBaseImageSettings) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_google_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleBaseImageSettings.ProtoReflect.Descriptor instead.
func (*GoogleBaseImageSettings) Descriptor() ([]byte, []int) {
	return file_cloudprovider_google_proto_rawDescGZIP(), []int{3}
}

func (x *GoogleBaseImageSettings) GetBaseImage() *GoogleBaseImage {
	if x != nil {
		return x.BaseImage
	}
	return nil
}

func (x *GoogleBaseImageSettings) GetVirtualizationSettings() *VirtualizationSettings {
	if x != nil {
		return x.VirtualizationSettings
	}
	return nil
}

// Base image configuration.
type GoogleBaseImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is the identifier used by GCP to find this base
	// image.
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
	// If set to true, Deck will annotate the popup tooltip to
	// indicate that the selected option represents an image family.
	ImageFamily *wrappers.BoolValue `protobuf:"bytes,5,opt,name=imageFamily,proto3" json:"imageFamily,omitempty"`
}

func (x *GoogleBaseImage) Reset() {
	*x = GoogleBaseImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_google_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleBaseImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleBaseImage) ProtoMessage() {}

func (x *GoogleBaseImage) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_google_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleBaseImage.ProtoReflect.Descriptor instead.
func (*GoogleBaseImage) Descriptor() ([]byte, []int) {
	return file_cloudprovider_google_proto_rawDescGZIP(), []int{4}
}

func (x *GoogleBaseImage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GoogleBaseImage) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

func (x *GoogleBaseImage) GetDetailedDescription() string {
	if x != nil {
		return x.DetailedDescription
	}
	return ""
}

func (x *GoogleBaseImage) GetPackageType() string {
	if x != nil {
		return x.PackageType
	}
	return ""
}

func (x *GoogleBaseImage) GetImageFamily() *wrappers.BoolValue {
	if x != nil {
		return x.ImageFamily
	}
	return nil
}

// Image source configuration.
type VirtualizationSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  The source image. If both sourceImage and sourceImageFamily
	//  are set, sourceImage will take precedence.
	SourceImage string `protobuf:"bytes,1,opt,name=sourceImage,proto3" json:"sourceImage,omitempty"`
	// The source image family to create the image from. The newest,
	// non-deprecated image is used. If both sourceImage and
	// sourceImageFamily are set, sourceImage will take precedence.
	SourceImageFamily string `protobuf:"bytes,2,opt,name=sourceImageFamily,proto3" json:"sourceImageFamily,omitempty"`
}

func (x *VirtualizationSettings) Reset() {
	*x = VirtualizationSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_google_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualizationSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualizationSettings) ProtoMessage() {}

func (x *VirtualizationSettings) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_google_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualizationSettings.ProtoReflect.Descriptor instead.
func (*VirtualizationSettings) Descriptor() ([]byte, []int) {
	return file_cloudprovider_google_proto_rawDescGZIP(), []int{5}
}

func (x *VirtualizationSettings) GetSourceImage() string {
	if x != nil {
		return x.SourceImage
	}
	return ""
}

func (x *VirtualizationSettings) GetSourceImageFamily() string {
	if x != nil {
		return x.SourceImageFamily
	}
	return ""
}

// Configuration for Consul.
type Consul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether Consul is enabled.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Reachable Consul node endpoint connected to the Consul cluster.
	// Defaults to localhost.
	AgentEndpoint string `protobuf:"bytes,2,opt,name=agentEndpoint,proto3" json:"agentEndpoint,omitempty"`
	// Port consul is running on for every agent. Defaults to 8500.
	AgentPort int32 `protobuf:"varint,3,opt,name=agentPort,proto3" json:"agentPort,omitempty"`
	// List of data centers to cache and keep updated. Defaults to all.
	Datacenters []string `protobuf:"bytes,4,rep,name=datacenters,proto3" json:"datacenters,omitempty"`
}

func (x *Consul) Reset() {
	*x = Consul{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_google_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consul) ProtoMessage() {}

func (x *Consul) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_google_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consul.ProtoReflect.Descriptor instead.
func (*Consul) Descriptor() ([]byte, []int) {
	return file_cloudprovider_google_proto_rawDescGZIP(), []int{6}
}

func (x *Consul) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *Consul) GetAgentEndpoint() string {
	if x != nil {
		return x.AgentEndpoint
	}
	return ""
}

func (x *Consul) GetAgentPort() int32 {
	if x != nil {
		return x.AgentPort
	}
	return 0
}

func (x *Consul) GetDatacenters() []string {
	if x != nil {
		return x.Datacenters
	}
	return nil
}

var File_cloudprovider_google_proto protoreflect.FileDescriptor

var file_cloudprovider_google_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x4b, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12,
	0x26, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x51, 0x0a, 0x0e, 0x62, 0x61, 0x6b, 0x65, 0x72,
	0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x42, 0x61, 0x6b, 0x65,
	0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x0e, 0x62, 0x61, 0x6b, 0x65,
	0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xaf, 0x03, 0x0a, 0x1a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a,
	0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f,
	0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f,
	0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3c, 0x0a, 0x0b, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x22, 0xa4, 0x02, 0x0a,
	0x14, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x42, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x62, 0x61, 0x73,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0a, 0x62, 0x61, 0x73,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x40, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x22, 0xc2, 0x01, 0x0a, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x42, 0x61,
	0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x42, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x42,
	0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x63, 0x0a, 0x16, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x16, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xdf, 0x01, 0x0a, 0x0f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x10,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x0b,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22, 0x68, 0x0a, 0x16, 0x56, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x22, 0xa4, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12,
	0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61,
	0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudprovider_google_proto_rawDescOnce sync.Once
	file_cloudprovider_google_proto_rawDescData = file_cloudprovider_google_proto_rawDesc
)

func file_cloudprovider_google_proto_rawDescGZIP() []byte {
	file_cloudprovider_google_proto_rawDescOnce.Do(func() {
		file_cloudprovider_google_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudprovider_google_proto_rawDescData)
	})
	return file_cloudprovider_google_proto_rawDescData
}

var file_cloudprovider_google_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cloudprovider_google_proto_goTypes = []interface{}{
	(*GoogleComputeEngine)(nil),        // 0: proto.cloudprovider.GoogleComputeEngine
	(*GoogleComputeEngineAccount)(nil), // 1: proto.cloudprovider.GoogleComputeEngineAccount
	(*GoogleBakeryDefaults)(nil),       // 2: proto.cloudprovider.GoogleBakeryDefaults
	(*GoogleBaseImageSettings)(nil),    // 3: proto.cloudprovider.GoogleBaseImageSettings
	(*GoogleBaseImage)(nil),            // 4: proto.cloudprovider.GoogleBaseImage
	(*VirtualizationSettings)(nil),     // 5: proto.cloudprovider.VirtualizationSettings
	(*Consul)(nil),                     // 6: proto.cloudprovider.Consul
	(*wrappers.BoolValue)(nil),         // 7: google.protobuf.BoolValue
	(*client.Permissions)(nil),         // 8: proto.Permissions
}
var file_cloudprovider_google_proto_depIdxs = []int32{
	7,  // 0: proto.cloudprovider.GoogleComputeEngine.enabled:type_name -> google.protobuf.BoolValue
	1,  // 1: proto.cloudprovider.GoogleComputeEngine.accounts:type_name -> proto.cloudprovider.GoogleComputeEngineAccount
	2,  // 2: proto.cloudprovider.GoogleComputeEngine.bakeryDefaults:type_name -> proto.cloudprovider.GoogleBakeryDefaults
	8,  // 3: proto.cloudprovider.GoogleComputeEngineAccount.permissions:type_name -> proto.Permissions
	7,  // 4: proto.cloudprovider.GoogleComputeEngineAccount.alphaListed:type_name -> google.protobuf.BoolValue
	6,  // 5: proto.cloudprovider.GoogleComputeEngineAccount.consul:type_name -> proto.cloudprovider.Consul
	3,  // 6: proto.cloudprovider.GoogleBakeryDefaults.baseImages:type_name -> proto.cloudprovider.GoogleBaseImageSettings
	7,  // 7: proto.cloudprovider.GoogleBakeryDefaults.useInternalIp:type_name -> google.protobuf.BoolValue
	4,  // 8: proto.cloudprovider.GoogleBaseImageSettings.baseImage:type_name -> proto.cloudprovider.GoogleBaseImage
	5,  // 9: proto.cloudprovider.GoogleBaseImageSettings.virtualizationSettings:type_name -> proto.cloudprovider.VirtualizationSettings
	7,  // 10: proto.cloudprovider.GoogleBaseImage.imageFamily:type_name -> google.protobuf.BoolValue
	7,  // 11: proto.cloudprovider.Consul.enabled:type_name -> google.protobuf.BoolValue
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_cloudprovider_google_proto_init() }
func file_cloudprovider_google_proto_init() {
	if File_cloudprovider_google_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudprovider_google_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleComputeEngine); i {
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
		file_cloudprovider_google_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleComputeEngineAccount); i {
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
		file_cloudprovider_google_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleBakeryDefaults); i {
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
		file_cloudprovider_google_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleBaseImageSettings); i {
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
		file_cloudprovider_google_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleBaseImage); i {
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
		file_cloudprovider_google_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualizationSettings); i {
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
		file_cloudprovider_google_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consul); i {
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
			RawDescriptor: file_cloudprovider_google_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudprovider_google_proto_goTypes,
		DependencyIndexes: file_cloudprovider_google_proto_depIdxs,
		MessageInfos:      file_cloudprovider_google_proto_msgTypes,
	}.Build()
	File_cloudprovider_google_proto = out.File
	file_cloudprovider_google_proto_rawDesc = nil
	file_cloudprovider_google_proto_goTypes = nil
	file_cloudprovider_google_proto_depIdxs = nil
}
