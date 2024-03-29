// Code generated - DO NOT EDIT.
// This file is a generated protocol buffer definition and any manual changes will be lost.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: DirectMom/DirectMom.proto

package DirectMom

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

type Disable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Who []byte                 `protobuf:"bytes,2,opt,name=Who,proto3" json:"Who,omitempty"` //	address
}

func (x *Disable) Reset() {
	*x = Disable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DirectMom_DirectMom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disable) ProtoMessage() {}

func (x *Disable) ProtoReflect() protoreflect.Message {
	mi := &file_DirectMom_DirectMom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disable.ProtoReflect.Descriptor instead.
func (*Disable) Descriptor() ([]byte, []int) {
	return file_DirectMom_DirectMom_proto_rawDescGZIP(), []int{0}
}

func (x *Disable) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Disable) GetWho() []byte {
	if x != nil {
		return x.Who
	}
	return nil
}

type SetAuthority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts           *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	OldAuthority []byte                 `protobuf:"bytes,2,opt,name=OldAuthority,proto3" json:"OldAuthority,omitempty"` //	address
	NewAuthority []byte                 `protobuf:"bytes,3,opt,name=NewAuthority,proto3" json:"NewAuthority,omitempty"` //	address
}

func (x *SetAuthority) Reset() {
	*x = SetAuthority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DirectMom_DirectMom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAuthority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAuthority) ProtoMessage() {}

func (x *SetAuthority) ProtoReflect() protoreflect.Message {
	mi := &file_DirectMom_DirectMom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAuthority.ProtoReflect.Descriptor instead.
func (*SetAuthority) Descriptor() ([]byte, []int) {
	return file_DirectMom_DirectMom_proto_rawDescGZIP(), []int{1}
}

func (x *SetAuthority) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *SetAuthority) GetOldAuthority() []byte {
	if x != nil {
		return x.OldAuthority
	}
	return nil
}

func (x *SetAuthority) GetNewAuthority() []byte {
	if x != nil {
		return x.NewAuthority
	}
	return nil
}

type SetOwner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts       *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	OldOwner []byte                 `protobuf:"bytes,2,opt,name=OldOwner,proto3" json:"OldOwner,omitempty"` //	address
	NewOwner []byte                 `protobuf:"bytes,3,opt,name=NewOwner,proto3" json:"NewOwner,omitempty"` //	address
}

func (x *SetOwner) Reset() {
	*x = SetOwner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DirectMom_DirectMom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOwner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOwner) ProtoMessage() {}

func (x *SetOwner) ProtoReflect() protoreflect.Message {
	mi := &file_DirectMom_DirectMom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOwner.ProtoReflect.Descriptor instead.
func (*SetOwner) Descriptor() ([]byte, []int) {
	return file_DirectMom_DirectMom_proto_rawDescGZIP(), []int{2}
}

func (x *SetOwner) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *SetOwner) GetOldOwner() []byte {
	if x != nil {
		return x.OldOwner
	}
	return nil
}

func (x *SetOwner) GetNewOwner() []byte {
	if x != nil {
		return x.NewOwner
	}
	return nil
}

var File_DirectMom_DirectMom_proto protoreflect.FileDescriptor

var file_DirectMom_DirectMom_proto_rawDesc = []byte{
	0x0a, 0x19, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4d, 0x6f, 0x6d, 0x2f, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x4d, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x4d, 0x6f, 0x6d, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x57, 0x68, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x57, 0x68, 0x6f,
	0x22, 0x82, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x4f, 0x6c, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x4f, 0x6c, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x6e, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x4f, 0x6c, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x4f, 0x6c, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x65, 0x77,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x4e, 0x65, 0x77,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x3c, 0x5a, 0x3a, 0x62, 0x6c, 0x65, 0x70, 0x2e, 0x61, 0x69,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x4d, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DirectMom_DirectMom_proto_rawDescOnce sync.Once
	file_DirectMom_DirectMom_proto_rawDescData = file_DirectMom_DirectMom_proto_rawDesc
)

func file_DirectMom_DirectMom_proto_rawDescGZIP() []byte {
	file_DirectMom_DirectMom_proto_rawDescOnce.Do(func() {
		file_DirectMom_DirectMom_proto_rawDescData = protoimpl.X.CompressGZIP(file_DirectMom_DirectMom_proto_rawDescData)
	})
	return file_DirectMom_DirectMom_proto_rawDescData
}

var file_DirectMom_DirectMom_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_DirectMom_DirectMom_proto_goTypes = []interface{}{
	(*Disable)(nil),               // 0: DirectMom.Disable
	(*SetAuthority)(nil),          // 1: DirectMom.SetAuthority
	(*SetOwner)(nil),              // 2: DirectMom.SetOwner
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_DirectMom_DirectMom_proto_depIdxs = []int32{
	3, // 0: DirectMom.Disable.ts:type_name -> google.protobuf.Timestamp
	3, // 1: DirectMom.SetAuthority.ts:type_name -> google.protobuf.Timestamp
	3, // 2: DirectMom.SetOwner.ts:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_DirectMom_DirectMom_proto_init() }
func file_DirectMom_DirectMom_proto_init() {
	if File_DirectMom_DirectMom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DirectMom_DirectMom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Disable); i {
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
		file_DirectMom_DirectMom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAuthority); i {
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
		file_DirectMom_DirectMom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOwner); i {
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
			RawDescriptor: file_DirectMom_DirectMom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DirectMom_DirectMom_proto_goTypes,
		DependencyIndexes: file_DirectMom_DirectMom_proto_depIdxs,
		MessageInfos:      file_DirectMom_DirectMom_proto_msgTypes,
	}.Build()
	File_DirectMom_DirectMom_proto = out.File
	file_DirectMom_DirectMom_proto_rawDesc = nil
	file_DirectMom_DirectMom_proto_goTypes = nil
	file_DirectMom_DirectMom_proto_depIdxs = nil
}
