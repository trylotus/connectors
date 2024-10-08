// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: RenBTC/RenBTC.proto

package RenBTC

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

type AdminChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts            *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	PreviousAdmin []byte                 `protobuf:"bytes,2,opt,name=PreviousAdmin,proto3" json:"PreviousAdmin,omitempty"` //	address
	NewAdmin      []byte                 `protobuf:"bytes,3,opt,name=NewAdmin,proto3" json:"NewAdmin,omitempty"`           //	address
}

func (x *AdminChanged) Reset() {
	*x = AdminChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RenBTC_RenBTC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminChanged) ProtoMessage() {}

func (x *AdminChanged) ProtoReflect() protoreflect.Message {
	mi := &file_RenBTC_RenBTC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminChanged.ProtoReflect.Descriptor instead.
func (*AdminChanged) Descriptor() ([]byte, []int) {
	return file_RenBTC_RenBTC_proto_rawDescGZIP(), []int{0}
}

func (x *AdminChanged) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *AdminChanged) GetPreviousAdmin() []byte {
	if x != nil {
		return x.PreviousAdmin
	}
	return nil
}

func (x *AdminChanged) GetNewAdmin() []byte {
	if x != nil {
		return x.NewAdmin
	}
	return nil
}

type Upgraded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts             *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Implementation []byte                 `protobuf:"bytes,2,opt,name=Implementation,proto3" json:"Implementation,omitempty"` //	address
}

func (x *Upgraded) Reset() {
	*x = Upgraded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RenBTC_RenBTC_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Upgraded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Upgraded) ProtoMessage() {}

func (x *Upgraded) ProtoReflect() protoreflect.Message {
	mi := &file_RenBTC_RenBTC_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Upgraded.ProtoReflect.Descriptor instead.
func (*Upgraded) Descriptor() ([]byte, []int) {
	return file_RenBTC_RenBTC_proto_rawDescGZIP(), []int{1}
}

func (x *Upgraded) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Upgraded) GetImplementation() []byte {
	if x != nil {
		return x.Implementation
	}
	return nil
}

var File_RenBTC_RenBTC_proto protoreflect.FileDescriptor

var file_RenBTC_RenBTC_proto_rawDesc = []byte{
	0x0a, 0x13, 0x52, 0x65, 0x6e, 0x42, 0x54, 0x43, 0x2f, 0x52, 0x65, 0x6e, 0x42, 0x54, 0x43, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x52, 0x65, 0x6e, 0x42, 0x54, 0x43, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c,
	0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x2a,
	0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x5e, 0x0a, 0x08,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x02, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x49, 0x6d,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x39, 0x5a, 0x37,
	0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x6f,
	0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73,
	0x2f, 0x52, 0x65, 0x6e, 0x42, 0x54, 0x43, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RenBTC_RenBTC_proto_rawDescOnce sync.Once
	file_RenBTC_RenBTC_proto_rawDescData = file_RenBTC_RenBTC_proto_rawDesc
)

func file_RenBTC_RenBTC_proto_rawDescGZIP() []byte {
	file_RenBTC_RenBTC_proto_rawDescOnce.Do(func() {
		file_RenBTC_RenBTC_proto_rawDescData = protoimpl.X.CompressGZIP(file_RenBTC_RenBTC_proto_rawDescData)
	})
	return file_RenBTC_RenBTC_proto_rawDescData
}

var file_RenBTC_RenBTC_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_RenBTC_RenBTC_proto_goTypes = []interface{}{
	(*AdminChanged)(nil),          // 0: RenBTC.AdminChanged
	(*Upgraded)(nil),              // 1: RenBTC.Upgraded
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_RenBTC_RenBTC_proto_depIdxs = []int32{
	2, // 0: RenBTC.AdminChanged.ts:type_name -> google.protobuf.Timestamp
	2, // 1: RenBTC.Upgraded.ts:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RenBTC_RenBTC_proto_init() }
func file_RenBTC_RenBTC_proto_init() {
	if File_RenBTC_RenBTC_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RenBTC_RenBTC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminChanged); i {
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
		file_RenBTC_RenBTC_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Upgraded); i {
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
			RawDescriptor: file_RenBTC_RenBTC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RenBTC_RenBTC_proto_goTypes,
		DependencyIndexes: file_RenBTC_RenBTC_proto_depIdxs,
		MessageInfos:      file_RenBTC_RenBTC_proto_msgTypes,
	}.Build()
	File_RenBTC_RenBTC_proto = out.File
	file_RenBTC_RenBTC_proto_rawDesc = nil
	file_RenBTC_RenBTC_proto_goTypes = nil
	file_RenBTC_RenBTC_proto_depIdxs = nil
}
