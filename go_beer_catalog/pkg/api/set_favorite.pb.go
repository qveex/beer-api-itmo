// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: set_favorite.proto

package __

import (
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

type SetFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeerId int64 `protobuf:"varint,2,opt,name=beerId,proto3" json:"beerId,omitempty"`
}

func (x *SetFavoriteRequest) Reset() {
	*x = SetFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_favorite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFavoriteRequest) ProtoMessage() {}

func (x *SetFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_set_favorite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFavoriteRequest.ProtoReflect.Descriptor instead.
func (*SetFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_set_favorite_proto_rawDescGZIP(), []int{0}
}

func (x *SetFavoriteRequest) GetBeerId() int64 {
	if x != nil {
		return x.BeerId
	}
	return 0
}

type SetFavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetFavoriteResponse) Reset() {
	*x = SetFavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_favorite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFavoriteResponse) ProtoMessage() {}

func (x *SetFavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_set_favorite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFavoriteResponse.ProtoReflect.Descriptor instead.
func (*SetFavoriteResponse) Descriptor() ([]byte, []int) {
	return file_set_favorite_proto_rawDescGZIP(), []int{1}
}

var File_set_favorite_proto protoreflect.FileDescriptor

var file_set_favorite_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x2c, 0x0a, 0x12, 0x53, 0x65, 0x74,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x65, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x62, 0x65, 0x65, 0x72, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03,
	0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_set_favorite_proto_rawDescOnce sync.Once
	file_set_favorite_proto_rawDescData = file_set_favorite_proto_rawDesc
)

func file_set_favorite_proto_rawDescGZIP() []byte {
	file_set_favorite_proto_rawDescOnce.Do(func() {
		file_set_favorite_proto_rawDescData = protoimpl.X.CompressGZIP(file_set_favorite_proto_rawDescData)
	})
	return file_set_favorite_proto_rawDescData
}

var file_set_favorite_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_set_favorite_proto_goTypes = []interface{}{
	(*SetFavoriteRequest)(nil),  // 0: api.SetFavoriteRequest
	(*SetFavoriteResponse)(nil), // 1: api.SetFavoriteResponse
}
var file_set_favorite_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_set_favorite_proto_init() }
func file_set_favorite_proto_init() {
	if File_set_favorite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_set_favorite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetFavoriteRequest); i {
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
		file_set_favorite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetFavoriteResponse); i {
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
			RawDescriptor: file_set_favorite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_set_favorite_proto_goTypes,
		DependencyIndexes: file_set_favorite_proto_depIdxs,
		MessageInfos:      file_set_favorite_proto_msgTypes,
	}.Build()
	File_set_favorite_proto = out.File
	file_set_favorite_proto_rawDesc = nil
	file_set_favorite_proto_goTypes = nil
	file_set_favorite_proto_depIdxs = nil
}
