// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: MCD_FLOP.proto

package MCD_FLOP

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

type LogNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Sig  []byte                 `protobuf:"bytes,2,opt,name=Sig,proto3" json:"Sig,omitempty"`   //	bytes4
	Usr  []byte                 `protobuf:"bytes,3,opt,name=Usr,proto3" json:"Usr,omitempty"`   //	address
	Arg1 []byte                 `protobuf:"bytes,4,opt,name=Arg1,proto3" json:"Arg1,omitempty"` //	bytes32
	Arg2 []byte                 `protobuf:"bytes,5,opt,name=Arg2,proto3" json:"Arg2,omitempty"` //	bytes32
	Data []byte                 `protobuf:"bytes,6,opt,name=Data,proto3" json:"Data,omitempty"` //	bytes
}

func (x *LogNote) Reset() {
	*x = LogNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MCD_FLOP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogNote) ProtoMessage() {}

func (x *LogNote) ProtoReflect() protoreflect.Message {
	mi := &file_MCD_FLOP_proto_msgTypes[0]
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
	return file_MCD_FLOP_proto_rawDescGZIP(), []int{0}
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

func (x *LogNote) GetUsr() []byte {
	if x != nil {
		return x.Usr
	}
	return nil
}

func (x *LogNote) GetArg1() []byte {
	if x != nil {
		return x.Arg1
	}
	return nil
}

func (x *LogNote) GetArg2() []byte {
	if x != nil {
		return x.Arg2
	}
	return nil
}

func (x *LogNote) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Kick struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Id  []byte                 `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`   //	uint256
	Lot []byte                 `protobuf:"bytes,3,opt,name=Lot,proto3" json:"Lot,omitempty"` //	uint256
	Bid []byte                 `protobuf:"bytes,4,opt,name=Bid,proto3" json:"Bid,omitempty"` //	uint256
	Gal []byte                 `protobuf:"bytes,5,opt,name=Gal,proto3" json:"Gal,omitempty"` //	address
}

func (x *Kick) Reset() {
	*x = Kick{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MCD_FLOP_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kick) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kick) ProtoMessage() {}

func (x *Kick) ProtoReflect() protoreflect.Message {
	mi := &file_MCD_FLOP_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kick.ProtoReflect.Descriptor instead.
func (*Kick) Descriptor() ([]byte, []int) {
	return file_MCD_FLOP_proto_rawDescGZIP(), []int{1}
}

func (x *Kick) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Kick) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Kick) GetLot() []byte {
	if x != nil {
		return x.Lot
	}
	return nil
}

func (x *Kick) GetBid() []byte {
	if x != nil {
		return x.Bid
	}
	return nil
}

func (x *Kick) GetGal() []byte {
	if x != nil {
		return x.Gal
	}
	return nil
}

var File_MCD_FLOP_proto protoreflect.FileDescriptor

var file_MCD_FLOP_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x4d, 0x43, 0x44, 0x5f, 0x46, 0x4c, 0x4f, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x4d, 0x43, 0x44, 0x5f, 0x46, 0x4c, 0x4f, 0x50, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x07,
	0x4c, 0x6f, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x02, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x53, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x73, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x55, 0x73, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x67, 0x31, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x41, 0x72, 0x67, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x41,
	0x72, 0x67, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x41, 0x72, 0x67, 0x32, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x78, 0x0a, 0x04, 0x4b, 0x69, 0x63, 0x6b, 0x12, 0x2a, 0x0a, 0x02, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x4c, 0x6f, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x42, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x47,
	0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x47, 0x61, 0x6c, 0x42, 0x3b, 0x5a,
	0x39, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x64, 0x61,
	0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x4d, 0x43, 0x44, 0x5f, 0x46, 0x4c, 0x4f, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_MCD_FLOP_proto_rawDescOnce sync.Once
	file_MCD_FLOP_proto_rawDescData = file_MCD_FLOP_proto_rawDesc
)

func file_MCD_FLOP_proto_rawDescGZIP() []byte {
	file_MCD_FLOP_proto_rawDescOnce.Do(func() {
		file_MCD_FLOP_proto_rawDescData = protoimpl.X.CompressGZIP(file_MCD_FLOP_proto_rawDescData)
	})
	return file_MCD_FLOP_proto_rawDescData
}

var file_MCD_FLOP_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_MCD_FLOP_proto_goTypes = []interface{}{
	(*LogNote)(nil),               // 0: MCD_FLOP.LogNote
	(*Kick)(nil),                  // 1: MCD_FLOP.Kick
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_MCD_FLOP_proto_depIdxs = []int32{
	2, // 0: MCD_FLOP.LogNote.ts:type_name -> google.protobuf.Timestamp
	2, // 1: MCD_FLOP.Kick.ts:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_MCD_FLOP_proto_init() }
func file_MCD_FLOP_proto_init() {
	if File_MCD_FLOP_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MCD_FLOP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_MCD_FLOP_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kick); i {
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
			RawDescriptor: file_MCD_FLOP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MCD_FLOP_proto_goTypes,
		DependencyIndexes: file_MCD_FLOP_proto_depIdxs,
		MessageInfos:      file_MCD_FLOP_proto_msgTypes,
	}.Build()
	File_MCD_FLOP_proto = out.File
	file_MCD_FLOP_proto_rawDesc = nil
	file_MCD_FLOP_proto_goTypes = nil
	file_MCD_FLOP_proto_depIdxs = nil
}
