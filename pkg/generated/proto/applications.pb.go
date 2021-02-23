// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: applications.proto

package proto

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

// Application is the specification for an application. This specification is also used for the Kubernetes CRD as
// ApplicationSpec.
// This is also the reason why we use istio.io/tools/cmd/protoc-gen-deepcopy within the code generation, to generate the
// deepcopy function, which are required to use the Application within a CRD.
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster   string                  `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace string                  `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Links     []*ApplicationLink      `protobuf:"bytes,4,rep,name=links,proto3" json:"links,omitempty"`
	Resources []*ApplicationResources `protobuf:"bytes,5,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_applications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_applications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_applications_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *Application) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetLinks() []*ApplicationLink {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *Application) GetResources() []*ApplicationResources {
	if x != nil {
		return x.Resources
	}
	return nil
}

// ApplicationLink is the format of a link, which can be provided within an application. A link consists of a title,
// which is displayed in the frontend and the link link (title=Example, link=https://example.com).
type ApplicationLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Link  string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *ApplicationLink) Reset() {
	*x = ApplicationLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_applications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationLink) ProtoMessage() {}

func (x *ApplicationLink) ProtoReflect() protoreflect.Message {
	mi := &file_applications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationLink.ProtoReflect.Descriptor instead.
func (*ApplicationLink) Descriptor() ([]byte, []int) {
	return file_applications_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationLink) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ApplicationLink) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

// ApplicationResources is the specification to retrieve all Kubernetes resources, which can be associated with the
// application. For that, a list of kinds (deployments, pods, statefulsets) as specified in
// app/src/components/resources/helpers.tsx and a selector must be set. Currently only label selector is supported,
// e.g. app=example.
type ApplicationResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kinds    []string `protobuf:"bytes,1,rep,name=kinds,proto3" json:"kinds,omitempty"`
	Selector string   `protobuf:"bytes,2,opt,name=selector,proto3" json:"selector,omitempty"`
}

func (x *ApplicationResources) Reset() {
	*x = ApplicationResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_applications_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationResources) ProtoMessage() {}

func (x *ApplicationResources) ProtoReflect() protoreflect.Message {
	mi := &file_applications_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationResources.ProtoReflect.Descriptor instead.
func (*ApplicationResources) Descriptor() ([]byte, []int) {
	return file_applications_proto_rawDescGZIP(), []int{2}
}

func (x *ApplicationResources) GetKinds() []string {
	if x != nil {
		return x.Kinds
	}
	return nil
}

func (x *ApplicationResources) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

var File_applications_proto protoreflect.FileDescriptor

var file_applications_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0xd0, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x12, 0x40, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x22, 0x48, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x69,
	0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x62, 0x73, 0x69,
	0x6f, 0x2f, 0x6b, 0x6f, 0x62, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_applications_proto_rawDescOnce sync.Once
	file_applications_proto_rawDescData = file_applications_proto_rawDesc
)

func file_applications_proto_rawDescGZIP() []byte {
	file_applications_proto_rawDescOnce.Do(func() {
		file_applications_proto_rawDescData = protoimpl.X.CompressGZIP(file_applications_proto_rawDescData)
	})
	return file_applications_proto_rawDescData
}

var file_applications_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_applications_proto_goTypes = []interface{}{
	(*Application)(nil),          // 0: applications.Application
	(*ApplicationLink)(nil),      // 1: applications.ApplicationLink
	(*ApplicationResources)(nil), // 2: applications.ApplicationResources
}
var file_applications_proto_depIdxs = []int32{
	1, // 0: applications.Application.links:type_name -> applications.ApplicationLink
	2, // 1: applications.Application.resources:type_name -> applications.ApplicationResources
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_applications_proto_init() }
func file_applications_proto_init() {
	if File_applications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_applications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_applications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationLink); i {
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
		file_applications_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationResources); i {
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
			RawDescriptor: file_applications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_applications_proto_goTypes,
		DependencyIndexes: file_applications_proto_depIdxs,
		MessageInfos:      file_applications_proto_msgTypes,
	}.Build()
	File_applications_proto = out.File
	file_applications_proto_rawDesc = nil
	file_applications_proto_goTypes = nil
	file_applications_proto_depIdxs = nil
}