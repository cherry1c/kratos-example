// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: c2/v1/message.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 流式请求
type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req string `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_c2_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_c2_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *StreamRequest) GetReq() string {
	if x != nil {
		return x.Req
	}
	return ""
}

// 流式响应
type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res string `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c2_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_c2_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_c2_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *StreamResponse) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

var File_c2_v1_message_proto protoreflect.FileDescriptor

var file_c2_v1_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x32, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0d, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x22, 0x0a,
	0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65,
	0x73, 0x32, 0x4c, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x42, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x2e, 0x63,
	0x32, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x34, 0x0a, 0x14, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x32, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x63, 0x32, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x56, 0x31, 0x50, 0x01, 0x5a, 0x0f, 0x63, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x32, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_c2_v1_message_proto_rawDescOnce sync.Once
	file_c2_v1_message_proto_rawDescData = file_c2_v1_message_proto_rawDesc
)

func file_c2_v1_message_proto_rawDescGZIP() []byte {
	file_c2_v1_message_proto_rawDescOnce.Do(func() {
		file_c2_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_c2_v1_message_proto_rawDescData)
	})
	return file_c2_v1_message_proto_rawDescData
}

var file_c2_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_c2_v1_message_proto_goTypes = []any{
	(*StreamRequest)(nil),  // 0: c2.v1.StreamRequest
	(*StreamResponse)(nil), // 1: c2.v1.StreamResponse
}
var file_c2_v1_message_proto_depIdxs = []int32{
	0, // 0: c2.v1.Stream.Conversations:input_type -> c2.v1.StreamRequest
	1, // 1: c2.v1.Stream.Conversations:output_type -> c2.v1.StreamResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_c2_v1_message_proto_init() }
func file_c2_v1_message_proto_init() {
	if File_c2_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_c2_v1_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StreamRequest); i {
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
		file_c2_v1_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StreamResponse); i {
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
			RawDescriptor: file_c2_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_c2_v1_message_proto_goTypes,
		DependencyIndexes: file_c2_v1_message_proto_depIdxs,
		MessageInfos:      file_c2_v1_message_proto_msgTypes,
	}.Build()
	File_c2_v1_message_proto = out.File
	file_c2_v1_message_proto_rawDesc = nil
	file_c2_v1_message_proto_goTypes = nil
	file_c2_v1_message_proto_depIdxs = nil
}
