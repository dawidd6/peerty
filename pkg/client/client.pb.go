// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: pkg/client/client.proto

package client

import (
	metadata "github.com/dawidd6/p2p/pkg/metadata"
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

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata []*metadata.Metadata `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_pkg_client_client_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetMetadata() []*metadata.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_pkg_client_client_proto_rawDescGZIP(), []int{1}
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_pkg_client_client_proto_rawDescGZIP(), []int{2}
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_pkg_client_client_proto_rawDescGZIP(), []int{3}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_client_client_proto_rawDescGZIP(), []int{4}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_client_client_proto_rawDescGZIP(), []int{5}
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_client_client_proto_rawDescGZIP(), []int{6}
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_client_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_client_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_client_client_proto_rawDescGZIP(), []int{7}
}

type StopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_client_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_client_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_pkg_client_client_proto_rawDescGZIP(), []int{8}
}

type StopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopResponse) Reset() {
	*x = StopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_client_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopResponse) ProtoMessage() {}

func (x *StopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_client_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopResponse.ProtoReflect.Descriptor instead.
func (*StopResponse) Descriptor() ([]byte, []int) {
	return file_pkg_client_client_proto_rawDescGZIP(), []int{9}
}

var File_pkg_client_client_proto protoreflect.FileDescriptor

var file_pkg_client_client_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x0d, 0x0a, 0x0b, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x0a, 0x0c,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x0a,
	0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c,
	0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc7, 0x01, 0x0a,
	0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0b,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x0d, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x0c, 0x2e, 0x53, 0x74, 0x6f, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x77, 0x69, 0x64, 0x64, 0x36, 0x2f, 0x70, 0x32, 0x70,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_client_client_proto_rawDescOnce sync.Once
	file_pkg_client_client_proto_rawDescData = file_pkg_client_client_proto_rawDesc
)

func file_pkg_client_client_proto_rawDescGZIP() []byte {
	file_pkg_client_client_proto_rawDescOnce.Do(func() {
		file_pkg_client_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_client_client_proto_rawDescData)
	})
	return file_pkg_client_client_proto_rawDescData
}

var file_pkg_client_client_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_client_client_proto_goTypes = []interface{}{
	(*AddRequest)(nil),        // 0: AddRequest
	(*AddResponse)(nil),       // 1: AddResponse
	(*DeleteRequest)(nil),     // 2: DeleteRequest
	(*DeleteResponse)(nil),    // 3: DeleteResponse
	(*ListRequest)(nil),       // 4: ListRequest
	(*ListResponse)(nil),      // 5: ListResponse
	(*StartRequest)(nil),      // 6: StartRequest
	(*StartResponse)(nil),     // 7: StartResponse
	(*StopRequest)(nil),       // 8: StopRequest
	(*StopResponse)(nil),      // 9: StopResponse
	(*metadata.Metadata)(nil), // 10: Metadata
}
var file_pkg_client_client_proto_depIdxs = []int32{
	10, // 0: AddRequest.metadata:type_name -> Metadata
	0,  // 1: Client.Add:input_type -> AddRequest
	2,  // 2: Client.Delete:input_type -> DeleteRequest
	4,  // 3: Client.List:input_type -> ListRequest
	6,  // 4: Client.Start:input_type -> StartRequest
	8,  // 5: Client.Stop:input_type -> StopRequest
	1,  // 6: Client.Add:output_type -> AddResponse
	3,  // 7: Client.Delete:output_type -> DeleteResponse
	5,  // 8: Client.List:output_type -> ListResponse
	7,  // 9: Client.Start:output_type -> StartResponse
	9,  // 10: Client.Stop:output_type -> StopResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_client_client_proto_init() }
func file_pkg_client_client_proto_init() {
	if File_pkg_client_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_client_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_pkg_client_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_pkg_client_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_pkg_client_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_pkg_client_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_pkg_client_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_pkg_client_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_pkg_client_client_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
		file_pkg_client_client_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRequest); i {
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
		file_pkg_client_client_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopResponse); i {
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
			RawDescriptor: file_pkg_client_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_client_client_proto_goTypes,
		DependencyIndexes: file_pkg_client_client_proto_depIdxs,
		MessageInfos:      file_pkg_client_client_proto_msgTypes,
	}.Build()
	File_pkg_client_client_proto = out.File
	file_pkg_client_client_proto_rawDesc = nil
	file_pkg_client_client_proto_goTypes = nil
	file_pkg_client_client_proto_depIdxs = nil
}
