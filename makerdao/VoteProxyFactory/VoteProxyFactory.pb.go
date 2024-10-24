// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: VoteProxyFactory/VoteProxyFactory.proto

package VoteProxyFactory

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

type LinkRequested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Cold []byte                 `protobuf:"bytes,2,opt,name=Cold,proto3" json:"Cold,omitempty"` //	address
	Hot  []byte                 `protobuf:"bytes,3,opt,name=Hot,proto3" json:"Hot,omitempty"`   //	address
}

func (x *LinkRequested) Reset() {
	*x = LinkRequested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VoteProxyFactory_VoteProxyFactory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkRequested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkRequested) ProtoMessage() {}

func (x *LinkRequested) ProtoReflect() protoreflect.Message {
	mi := &file_VoteProxyFactory_VoteProxyFactory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkRequested.ProtoReflect.Descriptor instead.
func (*LinkRequested) Descriptor() ([]byte, []int) {
	return file_VoteProxyFactory_VoteProxyFactory_proto_rawDescGZIP(), []int{0}
}

func (x *LinkRequested) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *LinkRequested) GetCold() []byte {
	if x != nil {
		return x.Cold
	}
	return nil
}

func (x *LinkRequested) GetHot() []byte {
	if x != nil {
		return x.Hot
	}
	return nil
}

type LinkConfirmed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts        *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Cold      []byte                 `protobuf:"bytes,2,opt,name=Cold,proto3" json:"Cold,omitempty"`           //	address
	Hot       []byte                 `protobuf:"bytes,3,opt,name=Hot,proto3" json:"Hot,omitempty"`             //	address
	VoteProxy []byte                 `protobuf:"bytes,4,opt,name=VoteProxy,proto3" json:"VoteProxy,omitempty"` //	address
}

func (x *LinkConfirmed) Reset() {
	*x = LinkConfirmed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VoteProxyFactory_VoteProxyFactory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkConfirmed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkConfirmed) ProtoMessage() {}

func (x *LinkConfirmed) ProtoReflect() protoreflect.Message {
	mi := &file_VoteProxyFactory_VoteProxyFactory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkConfirmed.ProtoReflect.Descriptor instead.
func (*LinkConfirmed) Descriptor() ([]byte, []int) {
	return file_VoteProxyFactory_VoteProxyFactory_proto_rawDescGZIP(), []int{1}
}

func (x *LinkConfirmed) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *LinkConfirmed) GetCold() []byte {
	if x != nil {
		return x.Cold
	}
	return nil
}

func (x *LinkConfirmed) GetHot() []byte {
	if x != nil {
		return x.Hot
	}
	return nil
}

func (x *LinkConfirmed) GetVoteProxy() []byte {
	if x != nil {
		return x.VoteProxy
	}
	return nil
}

var File_VoteProxyFactory_VoteProxyFactory_proto protoreflect.FileDescriptor

var file_VoteProxyFactory_VoteProxyFactory_proto_rawDesc = []byte{
	0x0a, 0x27, 0x56, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x46, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x56, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x46, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x56, 0x6f, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0d,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a,
	0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x6c,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x43, 0x6f, 0x6c, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x48, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x48, 0x6f, 0x74, 0x22,
	0x7f, 0x0a, 0x0d, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64,
	0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x43, 0x6f, 0x6c, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x48, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x48,
	0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x56, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x56, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x42, 0x43, 0x5a, 0x41, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x6b, 0x65,
	0x72, 0x64, 0x61, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x73, 0x2f, 0x56, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x46, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_VoteProxyFactory_VoteProxyFactory_proto_rawDescOnce sync.Once
	file_VoteProxyFactory_VoteProxyFactory_proto_rawDescData = file_VoteProxyFactory_VoteProxyFactory_proto_rawDesc
)

func file_VoteProxyFactory_VoteProxyFactory_proto_rawDescGZIP() []byte {
	file_VoteProxyFactory_VoteProxyFactory_proto_rawDescOnce.Do(func() {
		file_VoteProxyFactory_VoteProxyFactory_proto_rawDescData = protoimpl.X.CompressGZIP(file_VoteProxyFactory_VoteProxyFactory_proto_rawDescData)
	})
	return file_VoteProxyFactory_VoteProxyFactory_proto_rawDescData
}

var file_VoteProxyFactory_VoteProxyFactory_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_VoteProxyFactory_VoteProxyFactory_proto_goTypes = []interface{}{
	(*LinkRequested)(nil),         // 0: VoteProxyFactory.LinkRequested
	(*LinkConfirmed)(nil),         // 1: VoteProxyFactory.LinkConfirmed
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_VoteProxyFactory_VoteProxyFactory_proto_depIdxs = []int32{
	2, // 0: VoteProxyFactory.LinkRequested.ts:type_name -> google.protobuf.Timestamp
	2, // 1: VoteProxyFactory.LinkConfirmed.ts:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_VoteProxyFactory_VoteProxyFactory_proto_init() }
func file_VoteProxyFactory_VoteProxyFactory_proto_init() {
	if File_VoteProxyFactory_VoteProxyFactory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_VoteProxyFactory_VoteProxyFactory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkRequested); i {
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
		file_VoteProxyFactory_VoteProxyFactory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkConfirmed); i {
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
			RawDescriptor: file_VoteProxyFactory_VoteProxyFactory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VoteProxyFactory_VoteProxyFactory_proto_goTypes,
		DependencyIndexes: file_VoteProxyFactory_VoteProxyFactory_proto_depIdxs,
		MessageInfos:      file_VoteProxyFactory_VoteProxyFactory_proto_msgTypes,
	}.Build()
	File_VoteProxyFactory_VoteProxyFactory_proto = out.File
	file_VoteProxyFactory_VoteProxyFactory_proto_rawDesc = nil
	file_VoteProxyFactory_VoteProxyFactory_proto_goTypes = nil
	file_VoteProxyFactory_VoteProxyFactory_proto_depIdxs = nil
}
