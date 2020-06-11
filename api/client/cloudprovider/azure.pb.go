// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: cloudprovider/azure.proto

package cloudprovider

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

// Configuration for the Azure provider.
type Azure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*AzureAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	// Configuration for Spinnaker's image bakery.
	BakeryDefaults *AzureBakeryDefaults `protobuf:"bytes,4,opt,name=bakeryDefaults,proto3" json:"bakeryDefaults,omitempty"`
}

func (x *Azure) Reset() {
	*x = Azure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_azure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Azure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Azure) ProtoMessage() {}

func (x *Azure) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_azure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Azure.ProtoReflect.Descriptor instead.
func (*Azure) Descriptor() ([]byte, []int) {
	return file_cloudprovider_azure_proto_rawDescGZIP(), []int{0}
}

func (x *Azure) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Azure) GetAccounts() []*AzureAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *Azure) GetPrimaryAccount() string {
	if x != nil {
		return x.PrimaryAccount
	}
	return ""
}

func (x *Azure) GetBakeryDefaults() *AzureBakeryDefaults {
	if x != nil {
		return x.BakeryDefaults
	}
	return nil
}

// Configuration for an Azure account.
type AzureAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) The `appKey` (password) of your service principal.
	AppKey string `protobuf:"bytes,2,opt,name=appKey,proto3" json:"appKey,omitempty"`
	// (Required) The `clientId` (also called `appId`) of your service principal.
	ClientId string `protobuf:"bytes,3,opt,name=clientId,proto3" json:"clientId,omitempty"`
	// (Required) The name of a KeyVault that contains the user name, password,
	// and ssh public key used to create VMs
	DefaultKeyVault string `protobuf:"bytes,4,opt,name=defaultKeyVault,proto3" json:"defaultKeyVault,omitempty"`
	// (Required) The default resource group to contain any non-application
	// specific resources.
	DefaultResourceGroup string `protobuf:"bytes,5,opt,name=defaultResourceGroup,proto3" json:"defaultResourceGroup,omitempty"`
	// The environment name for the account. Many accounts can share the
	// same environment (e.g., dev, test, prod).
	Environment string `protobuf:"bytes,6,opt,name=environment,proto3" json:"environment,omitempty"`
	// The `objectId` of your service principal. This is only required if using
	// Packer to bake Windows images.
	ObjectId string `protobuf:"bytes,7,opt,name=objectId,proto3" json:"objectId,omitempty"`
	// The resource group to use if baking images with Packer.
	PackerResourceGroup string `protobuf:"bytes,8,opt,name=packerResourceGroup,proto3" json:"packerResourceGroup,omitempty"`
	// The storage account to use if baking images with Packer.
	PackerStorageAccount string `protobuf:"bytes,9,opt,name=packerStorageAccount,proto3" json:"packerStorageAccount,omitempty"`
	// Fiat permissions configuration.
	Permissions *client.Permissions `protobuf:"bytes,10,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Deprecated): List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,11,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
	// The Azure regions this Spinnaker account will manage.
	Regions []string `protobuf:"bytes,12,rep,name=regions,proto3" json:"regions,omitempty"`
	// (Required) The `subscriptionId` to which your service principal is
	// assigned.
	SubscriptionId string `protobuf:"bytes,13,opt,name=subscriptionId,proto3" json:"subscriptionId,omitempty"`
	// (Required) The `tenantId` to which your service principal is assigned.
	TenantId string `protobuf:"bytes,14,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// If true, the  SSH public key is used to provision the linux VM.
	// If false, the password is used instead.
	UseSshPublicKey bool `protobuf:"varint,15,opt,name=useSshPublicKey,proto3" json:"useSshPublicKey,omitempty"`
}

func (x *AzureAccount) Reset() {
	*x = AzureAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_azure_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AzureAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureAccount) ProtoMessage() {}

func (x *AzureAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_azure_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureAccount.ProtoReflect.Descriptor instead.
func (*AzureAccount) Descriptor() ([]byte, []int) {
	return file_cloudprovider_azure_proto_rawDescGZIP(), []int{1}
}

func (x *AzureAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AzureAccount) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

func (x *AzureAccount) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AzureAccount) GetDefaultKeyVault() string {
	if x != nil {
		return x.DefaultKeyVault
	}
	return ""
}

func (x *AzureAccount) GetDefaultResourceGroup() string {
	if x != nil {
		return x.DefaultResourceGroup
	}
	return ""
}

func (x *AzureAccount) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *AzureAccount) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *AzureAccount) GetPackerResourceGroup() string {
	if x != nil {
		return x.PackerResourceGroup
	}
	return ""
}

func (x *AzureAccount) GetPackerStorageAccount() string {
	if x != nil {
		return x.PackerStorageAccount
	}
	return ""
}

func (x *AzureAccount) GetPermissions() *client.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *AzureAccount) GetRequiredGroupMemberships() []string {
	if x != nil {
		return x.RequiredGroupMemberships
	}
	return nil
}

func (x *AzureAccount) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *AzureAccount) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *AzureAccount) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *AzureAccount) GetUseSshPublicKey() bool {
	if x != nil {
		return x.UseSshPublicKey
	}
	return false
}

// Configuration for Spinnaker's image bakery.
type AzureBakeryDefaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of configured base images.
	BaseImages []*AzureBaseImageSettings `protobuf:"bytes,1,rep,name=baseImages,proto3" json:"baseImages,omitempty"`
}

func (x *AzureBakeryDefaults) Reset() {
	*x = AzureBakeryDefaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_azure_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AzureBakeryDefaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureBakeryDefaults) ProtoMessage() {}

func (x *AzureBakeryDefaults) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_azure_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureBakeryDefaults.ProtoReflect.Descriptor instead.
func (*AzureBakeryDefaults) Descriptor() ([]byte, []int) {
	return file_cloudprovider_azure_proto_rawDescGZIP(), []int{2}
}

func (x *AzureBakeryDefaults) GetBaseImages() []*AzureBaseImageSettings {
	if x != nil {
		return x.BaseImages
	}
	return nil
}

// Configuration for a base image for the Azure provider's bakery.
type AzureBaseImageSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base image configuration.
	BaseImage *AzureBaseImage `protobuf:"bytes,1,opt,name=baseImage,proto3" json:"baseImage,omitempty"`
}

func (x *AzureBaseImageSettings) Reset() {
	*x = AzureBaseImageSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_azure_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AzureBaseImageSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureBaseImageSettings) ProtoMessage() {}

func (x *AzureBaseImageSettings) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_azure_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureBaseImageSettings.ProtoReflect.Descriptor instead.
func (*AzureBaseImageSettings) Descriptor() ([]byte, []int) {
	return file_cloudprovider_azure_proto_rawDescGZIP(), []int{3}
}

func (x *AzureBaseImageSettings) GetBaseImage() *AzureBaseImage {
	if x != nil {
		return x.BaseImage
	}
	return nil
}

// Base image configuration.
type AzureBaseImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A short description to help human operators identify the
	// image.
	ShortDescription string `protobuf:"bytes,1,opt,name=shortDescription,proto3" json:"shortDescription,omitempty"`
	// A long description to help human operators identify the
	// image.
	DetailedDescription string `protobuf:"bytes,2,opt,name=detailedDescription,proto3" json:"detailedDescription,omitempty"`
	// (Required) The Publisher name for your base image. See
	// https://aka.ms/azspinimage to get a list of images.
	Publisher string `protobuf:"bytes,3,opt,name=publisher,proto3" json:"publisher,omitempty"`
	// (Required) The offer for your base image. See https://aka.ms/azspinimage
	// to get a list of images.
	Offer string `protobuf:"bytes,4,opt,name=offer,proto3" json:"offer,omitempty"`
	// (Required) The SKU for your base image. See https://aka.ms/azspinimage to
	// get a list of images.
	Sku string `protobuf:"bytes,5,opt,name=sku,proto3" json:"sku,omitempty"`
	// The version of your base image. This defaults to `latest` if not
	// specified.
	Version string `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	// This is used to help Spinnaker's bakery download the build artifacts you
	// supply it with. For example, specifying `deb` indicates that your
	// artifacts will need to be fetched from a debian repository.
	PackageType string `protobuf:"bytes,7,opt,name=packageType,proto3" json:"packageType,omitempty"`
	// This is the name of the packer template that will be used to bake images
	// from this base image. The template file must be found in this list:
	// https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer,
	// or supplied as described here: https://spinnaker.io/setup/bakery/.
	TemplateFile string `protobuf:"bytes,8,opt,name=templateFile,proto3" json:"templateFile,omitempty"`
}

func (x *AzureBaseImage) Reset() {
	*x = AzureBaseImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_azure_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AzureBaseImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureBaseImage) ProtoMessage() {}

func (x *AzureBaseImage) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_azure_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureBaseImage.ProtoReflect.Descriptor instead.
func (*AzureBaseImage) Descriptor() ([]byte, []int) {
	return file_cloudprovider_azure_proto_rawDescGZIP(), []int{4}
}

func (x *AzureBaseImage) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

func (x *AzureBaseImage) GetDetailedDescription() string {
	if x != nil {
		return x.DetailedDescription
	}
	return ""
}

func (x *AzureBaseImage) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *AzureBaseImage) GetOffer() string {
	if x != nil {
		return x.Offer
	}
	return ""
}

func (x *AzureBaseImage) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *AzureBaseImage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AzureBaseImage) GetPackageType() string {
	if x != nil {
		return x.PackageType
	}
	return ""
}

func (x *AzureBaseImage) GetTemplateFile() string {
	if x != nil {
		return x.TemplateFile
	}
	return ""
}

var File_cloudprovider_azure_proto protoreflect.FileDescriptor

var file_cloudprovider_azure_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x1a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x05, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3d, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x41, 0x7a, 0x75, 0x72, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x50,
	0x0a, 0x0e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x7a, 0x75,
	0x72, 0x65, 0x42, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73,
	0x52, 0x0e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73,
	0x22, 0xd2, 0x04, 0x0a, 0x0c, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x32, 0x0a, 0x14, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x3a, 0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x75,
	0x73, 0x65, 0x53, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x53, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x62, 0x0a, 0x13, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x42, 0x61,
	0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x4b, 0x0a, 0x0a,
	0x62, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0a, 0x62,
	0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x5b, 0x0a, 0x16, 0x41, 0x7a, 0x75,
	0x72, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x41, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x7a, 0x75,
	0x72, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x09, 0x62, 0x61, 0x73,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x94, 0x02, 0x0a, 0x0e, 0x41, 0x7a, 0x75, 0x72, 0x65,
	0x42, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x6b, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e,
	0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudprovider_azure_proto_rawDescOnce sync.Once
	file_cloudprovider_azure_proto_rawDescData = file_cloudprovider_azure_proto_rawDesc
)

func file_cloudprovider_azure_proto_rawDescGZIP() []byte {
	file_cloudprovider_azure_proto_rawDescOnce.Do(func() {
		file_cloudprovider_azure_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudprovider_azure_proto_rawDescData)
	})
	return file_cloudprovider_azure_proto_rawDescData
}

var file_cloudprovider_azure_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloudprovider_azure_proto_goTypes = []interface{}{
	(*Azure)(nil),                  // 0: proto.cloudprovider.Azure
	(*AzureAccount)(nil),           // 1: proto.cloudprovider.AzureAccount
	(*AzureBakeryDefaults)(nil),    // 2: proto.cloudprovider.AzureBakeryDefaults
	(*AzureBaseImageSettings)(nil), // 3: proto.cloudprovider.AzureBaseImageSettings
	(*AzureBaseImage)(nil),         // 4: proto.cloudprovider.AzureBaseImage
	(*client.Permissions)(nil),     // 5: proto.Permissions
}
var file_cloudprovider_azure_proto_depIdxs = []int32{
	1, // 0: proto.cloudprovider.Azure.accounts:type_name -> proto.cloudprovider.AzureAccount
	2, // 1: proto.cloudprovider.Azure.bakeryDefaults:type_name -> proto.cloudprovider.AzureBakeryDefaults
	5, // 2: proto.cloudprovider.AzureAccount.permissions:type_name -> proto.Permissions
	3, // 3: proto.cloudprovider.AzureBakeryDefaults.baseImages:type_name -> proto.cloudprovider.AzureBaseImageSettings
	4, // 4: proto.cloudprovider.AzureBaseImageSettings.baseImage:type_name -> proto.cloudprovider.AzureBaseImage
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cloudprovider_azure_proto_init() }
func file_cloudprovider_azure_proto_init() {
	if File_cloudprovider_azure_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudprovider_azure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Azure); i {
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
		file_cloudprovider_azure_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AzureAccount); i {
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
		file_cloudprovider_azure_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AzureBakeryDefaults); i {
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
		file_cloudprovider_azure_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AzureBaseImageSettings); i {
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
		file_cloudprovider_azure_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AzureBaseImage); i {
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
			RawDescriptor: file_cloudprovider_azure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudprovider_azure_proto_goTypes,
		DependencyIndexes: file_cloudprovider_azure_proto_depIdxs,
		MessageInfos:      file_cloudprovider_azure_proto_msgTypes,
	}.Build()
	File_cloudprovider_azure_proto = out.File
	file_cloudprovider_azure_proto_rawDesc = nil
	file_cloudprovider_azure_proto_goTypes = nil
	file_cloudprovider_azure_proto_depIdxs = nil
}
