// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: apis/crudapi.proto

package crudapi

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skip        int32  `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Limit       int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	OrderBy     string `protobuf:"bytes,4,opt,name=orderBy,proto3" json:"orderBy,omitempty"`
	OrderByDesc bool   `protobuf:"varint,5,opt,name=orderByDesc,proto3" json:"orderByDesc,omitempty"`
}

func (x *ListOptions) Reset() {
	*x = ListOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_crudapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOptions) ProtoMessage() {}

func (x *ListOptions) ProtoReflect() protoreflect.Message {
	mi := &file_apis_crudapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOptions.ProtoReflect.Descriptor instead.
func (*ListOptions) Descriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{0}
}

func (x *ListOptions) GetSkip() int32 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *ListOptions) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListOptions) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListOptions) GetOrderByDesc() bool {
	if x != nil {
		return x.OrderByDesc
	}
	return false
}

type ErrorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode   int32  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	InternalCode string `protobuf:"bytes,2,opt,name=internalCode,proto3" json:"internalCode,omitempty"`
	Message      string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Suggestion   string `protobuf:"bytes,4,opt,name=suggestion,proto3" json:"suggestion,omitempty"`
}

func (x *ErrorDetails) Reset() {
	*x = ErrorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_crudapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetails) ProtoMessage() {}

func (x *ErrorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_apis_crudapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetails.ProtoReflect.Descriptor instead.
func (*ErrorDetails) Descriptor() ([]byte, []int) {
	return file_apis_crudapi_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorDetails) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ErrorDetails) GetInternalCode() string {
	if x != nil {
		return x.InternalCode
	}
	return ""
}

func (x *ErrorDetails) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorDetails) GetSuggestion() string {
	if x != nil {
		return x.Suggestion
	}
	return ""
}

var File_apis_crudapi_proto protoreflect.FileDescriptor

var file_apis_crudapi_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x04, 0x73,
	0x6b, 0x69, 0x70, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x44, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x44, 0x65, 0x73, 0x63, 0x22, 0x8c,
	0x01, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xb1, 0x01,
	0x5a, 0x09, 0x2e, 0x2f, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x92, 0x41, 0xa2, 0x01, 0x12,
	0x2e, 0x0a, 0x10, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x43, 0x52, 0x55, 0x44, 0x20, 0x41,
	0x50, 0x49, 0x73, 0x22, 0x12, 0x1a, 0x10, 0x6b, 0x7a, 0x69, 0x72, 0x74, 0x6d, 0x40, 0x67, 0x6d,
	0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x06, 0x76, 0x30, 0x2e, 0x31, 0x2e, 0x30, 0x1a,
	0x15, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73,
	0x74, 0x3a, 0x36, 0x39, 0x39, 0x39, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x32, 0x12,
	0x30, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x64, 0x6c, 0x69, 0x6c, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x63, 0x72, 0x75,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_crudapi_proto_rawDescOnce sync.Once
	file_apis_crudapi_proto_rawDescData = file_apis_crudapi_proto_rawDesc
)

func file_apis_crudapi_proto_rawDescGZIP() []byte {
	file_apis_crudapi_proto_rawDescOnce.Do(func() {
		file_apis_crudapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_crudapi_proto_rawDescData)
	})
	return file_apis_crudapi_proto_rawDescData
}

var file_apis_crudapi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apis_crudapi_proto_goTypes = []interface{}{
	(*ListOptions)(nil),  // 0: crudapi.ListOptions
	(*ErrorDetails)(nil), // 1: crudapi.ErrorDetails
}
var file_apis_crudapi_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apis_crudapi_proto_init() }
func file_apis_crudapi_proto_init() {
	if File_apis_crudapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_crudapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOptions); i {
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
		file_apis_crudapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetails); i {
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
			RawDescriptor: file_apis_crudapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_crudapi_proto_goTypes,
		DependencyIndexes: file_apis_crudapi_proto_depIdxs,
		MessageInfos:      file_apis_crudapi_proto_msgTypes,
	}.Build()
	File_apis_crudapi_proto = out.File
	file_apis_crudapi_proto_rawDesc = nil
	file_apis_crudapi_proto_goTypes = nil
	file_apis_crudapi_proto_depIdxs = nil
}
