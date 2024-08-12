// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: example/message.proto

package example

import (
	_ "github.com/trylotus/go-lotus-proto"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	BlockNumber uint64                 `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	TxHash      []byte                 `protobuf:"bytes,3,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Index       uint64                 `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	Data        []byte                 `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_example_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_example_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Message) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *Message) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *Message) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_example_message_proto protoreflect.FileDescriptor

var file_example_message_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2f, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0c,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x04, 0xca, 0xab, 0x30, 0x00, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0xd2, 0xab, 0x30, 0x02, 0x08, 0x01, 0x52, 0x06,
	0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x3a, 0x2e, 0xca, 0xab, 0x30, 0x2a, 0x0a, 0x16, 0x10, 0x01, 0x1a, 0x02, 0x74, 0x73, 0x1a, 0x07,
	0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x1a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x0a, 0x10,
	0x1a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x1a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x72, 0x79, 0x6c, 0x6f, 0x74, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_example_message_proto_rawDescOnce sync.Once
	file_example_message_proto_rawDescData = file_example_message_proto_rawDesc
)

func file_example_message_proto_rawDescGZIP() []byte {
	file_example_message_proto_rawDescOnce.Do(func() {
		file_example_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_message_proto_rawDescData)
	})
	return file_example_message_proto_rawDescData
}

var file_example_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_example_message_proto_goTypes = []interface{}{
	(*Message)(nil),               // 0: example.Message
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_example_message_proto_depIdxs = []int32{
	1, // 0: example.Message.ts:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_message_proto_init() }
func file_example_message_proto_init() {
	if File_example_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_example_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_message_proto_goTypes,
		DependencyIndexes: file_example_message_proto_depIdxs,
		MessageInfos:      file_example_message_proto_msgTypes,
	}.Build()
	File_example_message_proto = out.File
	file_example_message_proto_rawDesc = nil
	file_example_message_proto_goTypes = nil
	file_example_message_proto_depIdxs = nil
}