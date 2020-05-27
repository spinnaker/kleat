// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: config/clouddriver.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	artifact "github.com/spinnaker/kleat/api/client/artifact"
	cloudprovider "github.com/spinnaker/kleat/api/client/cloudprovider"
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

// Configuration for the clouddriver microservice.
type Clouddriver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kubernetes     *cloudprovider.Kubernetes          `protobuf:"bytes,1,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
	Google         *cloudprovider.GoogleComputeEngine `protobuf:"bytes,2,opt,name=google,proto3" json:"google,omitempty"`
	Appengine      *cloudprovider.Appengine           `protobuf:"bytes,3,opt,name=appengine,proto3" json:"appengine,omitempty"`
	Aws            *cloudprovider.Aws                 `protobuf:"bytes,4,opt,name=aws,proto3" json:"aws,omitempty"`
	Azure          *cloudprovider.Azure               `protobuf:"bytes,5,opt,name=azure,proto3" json:"azure,omitempty"`
	Cloudfoundry   *cloudprovider.CloudFoundry        `protobuf:"bytes,6,opt,name=cloudfoundry,proto3" json:"cloudfoundry,omitempty"`
	Dcos           *cloudprovider.Dcos                `protobuf:"bytes,7,opt,name=dcos,proto3" json:"dcos,omitempty"`
	DockerRegistry *cloudprovider.DockerRegistry      `protobuf:"bytes,8,opt,name=dockerRegistry,proto3" json:"dockerRegistry,omitempty"`
	Ecs            *cloudprovider.Ecs                 `protobuf:"bytes,9,opt,name=ecs,proto3" json:"ecs,omitempty"`
	Huaweicloud    *cloudprovider.HuaweiCloud         `protobuf:"bytes,10,opt,name=huaweicloud,proto3" json:"huaweicloud,omitempty"`
	Oracle         *cloudprovider.Oracle              `protobuf:"bytes,11,opt,name=oracle,proto3" json:"oracle,omitempty"`
	Artifacts      *artifact.Artifacts                `protobuf:"bytes,12,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
}

func (x *Clouddriver) Reset() {
	*x = Clouddriver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_clouddriver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clouddriver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clouddriver) ProtoMessage() {}

func (x *Clouddriver) ProtoReflect() protoreflect.Message {
	mi := &file_config_clouddriver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clouddriver.ProtoReflect.Descriptor instead.
func (*Clouddriver) Descriptor() ([]byte, []int) {
	return file_config_clouddriver_proto_rawDescGZIP(), []int{0}
}

func (x *Clouddriver) GetKubernetes() *cloudprovider.Kubernetes {
	if x != nil {
		return x.Kubernetes
	}
	return nil
}

func (x *Clouddriver) GetGoogle() *cloudprovider.GoogleComputeEngine {
	if x != nil {
		return x.Google
	}
	return nil
}

func (x *Clouddriver) GetAppengine() *cloudprovider.Appengine {
	if x != nil {
		return x.Appengine
	}
	return nil
}

func (x *Clouddriver) GetAws() *cloudprovider.Aws {
	if x != nil {
		return x.Aws
	}
	return nil
}

func (x *Clouddriver) GetAzure() *cloudprovider.Azure {
	if x != nil {
		return x.Azure
	}
	return nil
}

func (x *Clouddriver) GetCloudfoundry() *cloudprovider.CloudFoundry {
	if x != nil {
		return x.Cloudfoundry
	}
	return nil
}

func (x *Clouddriver) GetDcos() *cloudprovider.Dcos {
	if x != nil {
		return x.Dcos
	}
	return nil
}

func (x *Clouddriver) GetDockerRegistry() *cloudprovider.DockerRegistry {
	if x != nil {
		return x.DockerRegistry
	}
	return nil
}

func (x *Clouddriver) GetEcs() *cloudprovider.Ecs {
	if x != nil {
		return x.Ecs
	}
	return nil
}

func (x *Clouddriver) GetHuaweicloud() *cloudprovider.HuaweiCloud {
	if x != nil {
		return x.Huaweicloud
	}
	return nil
}

func (x *Clouddriver) GetOracle() *cloudprovider.Oracle {
	if x != nil {
		return x.Oracle
	}
	return nil
}

func (x *Clouddriver) GetArtifacts() *artifact.Artifacts {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

var File_config_clouddriver_proto protoreflect.FileDescriptor

var file_config_clouddriver_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x18, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x64, 0x63, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x65, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x68, 0x75, 0x61, 0x77, 0x65,
	0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x05, 0x0a, 0x0b, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x52, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x09,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x77, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x77, 0x73,
	0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65,
	0x52, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79,
	0x52, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x12, 0x2d,
	0x0a, 0x04, 0x64, 0x63, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x44, 0x63, 0x6f, 0x73, 0x52, 0x04, 0x64, 0x63, 0x6f, 0x73, 0x12, 0x4b, 0x0a,
	0x0e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x03, 0x65, 0x63,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x63,
	0x73, 0x52, 0x03, 0x65, 0x63, 0x73, 0x12, 0x42, 0x0a, 0x0b, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x48, 0x75, 0x61, 0x77, 0x65, 0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x0b, 0x68,
	0x75, 0x61, 0x77, 0x65, 0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x12,
	0x37, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x09, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72,
	0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_clouddriver_proto_rawDescOnce sync.Once
	file_config_clouddriver_proto_rawDescData = file_config_clouddriver_proto_rawDesc
)

func file_config_clouddriver_proto_rawDescGZIP() []byte {
	file_config_clouddriver_proto_rawDescOnce.Do(func() {
		file_config_clouddriver_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_clouddriver_proto_rawDescData)
	})
	return file_config_clouddriver_proto_rawDescData
}

var file_config_clouddriver_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_clouddriver_proto_goTypes = []interface{}{
	(*Clouddriver)(nil),                       // 0: proto.config.Clouddriver
	(*cloudprovider.Kubernetes)(nil),          // 1: proto.cloudprovider.Kubernetes
	(*cloudprovider.GoogleComputeEngine)(nil), // 2: proto.cloudprovider.GoogleComputeEngine
	(*cloudprovider.Appengine)(nil),           // 3: proto.cloudprovider.Appengine
	(*cloudprovider.Aws)(nil),                 // 4: proto.cloudprovider.Aws
	(*cloudprovider.Azure)(nil),               // 5: proto.cloudprovider.Azure
	(*cloudprovider.CloudFoundry)(nil),        // 6: proto.cloudprovider.CloudFoundry
	(*cloudprovider.Dcos)(nil),                // 7: proto.cloudprovider.Dcos
	(*cloudprovider.DockerRegistry)(nil),      // 8: proto.cloudprovider.DockerRegistry
	(*cloudprovider.Ecs)(nil),                 // 9: proto.cloudprovider.Ecs
	(*cloudprovider.HuaweiCloud)(nil),         // 10: proto.cloudprovider.HuaweiCloud
	(*cloudprovider.Oracle)(nil),              // 11: proto.cloudprovider.Oracle
	(*artifact.Artifacts)(nil),                // 12: proto.artifact.Artifacts
}
var file_config_clouddriver_proto_depIdxs = []int32{
	1,  // 0: proto.config.Clouddriver.kubernetes:type_name -> proto.cloudprovider.Kubernetes
	2,  // 1: proto.config.Clouddriver.google:type_name -> proto.cloudprovider.GoogleComputeEngine
	3,  // 2: proto.config.Clouddriver.appengine:type_name -> proto.cloudprovider.Appengine
	4,  // 3: proto.config.Clouddriver.aws:type_name -> proto.cloudprovider.Aws
	5,  // 4: proto.config.Clouddriver.azure:type_name -> proto.cloudprovider.Azure
	6,  // 5: proto.config.Clouddriver.cloudfoundry:type_name -> proto.cloudprovider.CloudFoundry
	7,  // 6: proto.config.Clouddriver.dcos:type_name -> proto.cloudprovider.Dcos
	8,  // 7: proto.config.Clouddriver.dockerRegistry:type_name -> proto.cloudprovider.DockerRegistry
	9,  // 8: proto.config.Clouddriver.ecs:type_name -> proto.cloudprovider.Ecs
	10, // 9: proto.config.Clouddriver.huaweicloud:type_name -> proto.cloudprovider.HuaweiCloud
	11, // 10: proto.config.Clouddriver.oracle:type_name -> proto.cloudprovider.Oracle
	12, // 11: proto.config.Clouddriver.artifacts:type_name -> proto.artifact.Artifacts
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_config_clouddriver_proto_init() }
func file_config_clouddriver_proto_init() {
	if File_config_clouddriver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_clouddriver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clouddriver); i {
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
			RawDescriptor: file_config_clouddriver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_clouddriver_proto_goTypes,
		DependencyIndexes: file_config_clouddriver_proto_depIdxs,
		MessageInfos:      file_config_clouddriver_proto_msgTypes,
	}.Build()
	File_config_clouddriver_proto = out.File
	file_config_clouddriver_proto_rawDesc = nil
	file_config_clouddriver_proto_goTypes = nil
	file_config_clouddriver_proto_depIdxs = nil
}
