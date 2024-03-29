// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: JoinFab/JoinFab.proto

package JoinFab

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NewJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Join []byte                 `protobuf:"bytes,2,opt,name=Join,proto3" json:"Join,omitempty"` //	address
	Data []byte                 `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"` //	bytes
}

func (x *NewJoin) Reset() {
	*x = NewJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_JoinFab_JoinFab_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewJoin) ProtoMessage() {}

func (x *NewJoin) ProtoReflect() protoreflect.Message {
	mi := &file_JoinFab_JoinFab_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewJoin.ProtoReflect.Descriptor instead.
func (*NewJoin) Descriptor() ([]byte, []int) {
	return file_JoinFab_JoinFab_proto_rawDescGZIP(), []int{0}
}

func (x *NewJoin) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *NewJoin) GetJoin() []byte {
	if x != nil {
		return x.Join
	}
	return nil
}

func (x *NewJoin) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_JoinFab_JoinFab_proto protoreflect.FileDescriptor

var file_JoinFab_JoinFab_proto_rawDesc = []byte{
	0x0a, 0x15, 0x4a, 0x6f, 0x69, 0x6e, 0x46, 0x61, 0x62, 0x2f, 0x4a, 0x6f, 0x69, 0x6e, 0x46, 0x61,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x4a, 0x6f, 0x69, 0x6e, 0x46, 0x61, 0x62,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5d, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x2a, 0x0a, 0x02,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x3a, 0x5a, 0x38, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x6b, 0x65,
	0x72, 0x64, 0x61, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x73, 0x2f, 0x4a, 0x6f, 0x69, 0x6e, 0x46, 0x61, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_JoinFab_JoinFab_proto_rawDescOnce sync.Once
	file_JoinFab_JoinFab_proto_rawDescData = file_JoinFab_JoinFab_proto_rawDesc
)

func file_JoinFab_JoinFab_proto_rawDescGZIP() []byte {
	file_JoinFab_JoinFab_proto_rawDescOnce.Do(func() {
		file_JoinFab_JoinFab_proto_rawDescData = protoimpl.X.CompressGZIP(file_JoinFab_JoinFab_proto_rawDescData)
	})
	return file_JoinFab_JoinFab_proto_rawDescData
}

var file_JoinFab_JoinFab_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_JoinFab_JoinFab_proto_goTypes = []interface{}{
	(*NewJoin)(nil),               // 0: JoinFab.NewJoin
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_JoinFab_JoinFab_proto_depIdxs = []int32{
	1, // 0: JoinFab.NewJoin.ts:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_JoinFab_JoinFab_proto_init() }
func file_JoinFab_JoinFab_proto_init() {
	if File_JoinFab_JoinFab_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_JoinFab_JoinFab_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewJoin); i {
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
			RawDescriptor: file_JoinFab_JoinFab_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JoinFab_JoinFab_proto_goTypes,
		DependencyIndexes: file_JoinFab_JoinFab_proto_depIdxs,
		MessageInfos:      file_JoinFab_JoinFab_proto_msgTypes,
	}.Build()
	File_JoinFab_JoinFab_proto = out.File
	file_JoinFab_JoinFab_proto_rawDesc = nil
	file_JoinFab_JoinFab_proto_goTypes = nil
	file_JoinFab_JoinFab_proto_depIdxs = nil
}
