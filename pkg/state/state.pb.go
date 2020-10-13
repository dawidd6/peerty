// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: pkg/state/state.proto

package state

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

type State int32

const (
	State_UNKNOWN     State = 0
	State_DOWNLOADING State = 1
	State_SEEDING     State = 2
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "UNKNOWN",
		1: "DOWNLOADING",
		2: "SEEDING",
	}
	State_value = map[string]int32{
		"UNKNOWN":     0,
		"DOWNLOADING": 1,
		"SEEDING":     2,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_state_state_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_pkg_state_state_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_pkg_state_state_proto_rawDescGZIP(), []int{0}
}

var File_pkg_state_state_proto protoreflect.FileDescriptor

var file_pkg_state_state_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x32, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x45, 0x45, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x42, 0x22, 0x5a, 0x20, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x77, 0x69, 0x64, 0x64,
	0x36, 0x2f, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_state_state_proto_rawDescOnce sync.Once
	file_pkg_state_state_proto_rawDescData = file_pkg_state_state_proto_rawDesc
)

func file_pkg_state_state_proto_rawDescGZIP() []byte {
	file_pkg_state_state_proto_rawDescOnce.Do(func() {
		file_pkg_state_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_state_state_proto_rawDescData)
	})
	return file_pkg_state_state_proto_rawDescData
}

var file_pkg_state_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_state_state_proto_goTypes = []interface{}{
	(State)(0), // 0: State
}
var file_pkg_state_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_state_state_proto_init() }
func file_pkg_state_state_proto_init() {
	if File_pkg_state_state_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_state_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_state_state_proto_goTypes,
		DependencyIndexes: file_pkg_state_state_proto_depIdxs,
		EnumInfos:         file_pkg_state_state_proto_enumTypes,
	}.Build()
	File_pkg_state_state_proto = out.File
	file_pkg_state_state_proto_rawDesc = nil
	file_pkg_state_state_proto_goTypes = nil
	file_pkg_state_state_proto_depIdxs = nil
}
