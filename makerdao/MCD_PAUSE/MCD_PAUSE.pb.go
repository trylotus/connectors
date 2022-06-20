// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: MCD_PAUSE.proto

package MCD_PAUSE

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

type LogSetOwner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Owner []byte                 `protobuf:"bytes,2,opt,name=Owner,proto3" json:"Owner,omitempty"` //	address
}

func (x *LogSetOwner) Reset() {
	*x = LogSetOwner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MCD_PAUSE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogSetOwner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSetOwner) ProtoMessage() {}

func (x *LogSetOwner) ProtoReflect() protoreflect.Message {
	mi := &file_MCD_PAUSE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSetOwner.ProtoReflect.Descriptor instead.
func (*LogSetOwner) Descriptor() ([]byte, []int) {
	return file_MCD_PAUSE_proto_rawDescGZIP(), []int{0}
}

func (x *LogSetOwner) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *LogSetOwner) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

type LogNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Sig []byte                 `protobuf:"bytes,2,opt,name=Sig,proto3" json:"Sig,omitempty"` //	bytes4
	Guy []byte                 `protobuf:"bytes,3,opt,name=Guy,proto3" json:"Guy,omitempty"` //	address
	Foo []byte                 `protobuf:"bytes,4,opt,name=Foo,proto3" json:"Foo,omitempty"` //	bytes32
	Bar []byte                 `protobuf:"bytes,5,opt,name=Bar,proto3" json:"Bar,omitempty"` //	bytes32
	Wad []byte                 `protobuf:"bytes,6,opt,name=Wad,proto3" json:"Wad,omitempty"` //	uint256
	Fax []byte                 `protobuf:"bytes,7,opt,name=Fax,proto3" json:"Fax,omitempty"` //	bytes
}

func (x *LogNote) Reset() {
	*x = LogNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MCD_PAUSE_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogNote) ProtoMessage() {}

func (x *LogNote) ProtoReflect() protoreflect.Message {
	mi := &file_MCD_PAUSE_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogNote.ProtoReflect.Descriptor instead.
func (*LogNote) Descriptor() ([]byte, []int) {
	return file_MCD_PAUSE_proto_rawDescGZIP(), []int{1}
}

func (x *LogNote) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *LogNote) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *LogNote) GetGuy() []byte {
	if x != nil {
		return x.Guy
	}
	return nil
}

func (x *LogNote) GetFoo() []byte {
	if x != nil {
		return x.Foo
	}
	return nil
}

func (x *LogNote) GetBar() []byte {
	if x != nil {
		return x.Bar
	}
	return nil
}

func (x *LogNote) GetWad() []byte {
	if x != nil {
		return x.Wad
	}
	return nil
}

func (x *LogNote) GetFax() []byte {
	if x != nil {
		return x.Fax
	}
	return nil
}

type LogSetAuthority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts        *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Authority []byte                 `protobuf:"bytes,2,opt,name=Authority,proto3" json:"Authority,omitempty"` //	address
}

func (x *LogSetAuthority) Reset() {
	*x = LogSetAuthority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MCD_PAUSE_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogSetAuthority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSetAuthority) ProtoMessage() {}

func (x *LogSetAuthority) ProtoReflect() protoreflect.Message {
	mi := &file_MCD_PAUSE_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSetAuthority.ProtoReflect.Descriptor instead.
func (*LogSetAuthority) Descriptor() ([]byte, []int) {
	return file_MCD_PAUSE_proto_rawDescGZIP(), []int{2}
}

func (x *LogSetAuthority) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *LogSetAuthority) GetAuthority() []byte {
	if x != nil {
		return x.Authority
	}
	return nil
}

var File_MCD_PAUSE_proto protoreflect.FileDescriptor

var file_MCD_PAUSE_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x4d, 0x43, 0x44, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x4d, 0x43, 0x44, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a,
	0x0b, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x02,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0xa1,
	0x01, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x53, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x75, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x47, 0x75, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x46, 0x6f,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03,
	0x42, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x42, 0x61, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x57, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x57, 0x61, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x46, 0x61, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x46,
	0x61, 0x78, 0x22, 0x5b, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x42,
	0x3c, 0x5a, 0x3a, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x72,
	0x64, 0x61, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x4d, 0x43, 0x44, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MCD_PAUSE_proto_rawDescOnce sync.Once
	file_MCD_PAUSE_proto_rawDescData = file_MCD_PAUSE_proto_rawDesc
)

func file_MCD_PAUSE_proto_rawDescGZIP() []byte {
	file_MCD_PAUSE_proto_rawDescOnce.Do(func() {
		file_MCD_PAUSE_proto_rawDescData = protoimpl.X.CompressGZIP(file_MCD_PAUSE_proto_rawDescData)
	})
	return file_MCD_PAUSE_proto_rawDescData
}

var file_MCD_PAUSE_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_MCD_PAUSE_proto_goTypes = []interface{}{
	(*LogSetOwner)(nil),           // 0: MCD_PAUSE.LogSetOwner
	(*LogNote)(nil),               // 1: MCD_PAUSE.LogNote
	(*LogSetAuthority)(nil),       // 2: MCD_PAUSE.LogSetAuthority
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_MCD_PAUSE_proto_depIdxs = []int32{
	3, // 0: MCD_PAUSE.LogSetOwner.ts:type_name -> google.protobuf.Timestamp
	3, // 1: MCD_PAUSE.LogNote.ts:type_name -> google.protobuf.Timestamp
	3, // 2: MCD_PAUSE.LogSetAuthority.ts:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_MCD_PAUSE_proto_init() }
func file_MCD_PAUSE_proto_init() {
	if File_MCD_PAUSE_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MCD_PAUSE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogSetOwner); i {
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
		file_MCD_PAUSE_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogNote); i {
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
		file_MCD_PAUSE_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogSetAuthority); i {
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
			RawDescriptor: file_MCD_PAUSE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MCD_PAUSE_proto_goTypes,
		DependencyIndexes: file_MCD_PAUSE_proto_depIdxs,
		MessageInfos:      file_MCD_PAUSE_proto_msgTypes,
	}.Build()
	File_MCD_PAUSE_proto = out.File
	file_MCD_PAUSE_proto_rawDesc = nil
	file_MCD_PAUSE_proto_goTypes = nil
	file_MCD_PAUSE_proto_depIdxs = nil
}
