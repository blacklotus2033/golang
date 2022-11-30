// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v2.28.1
// 	protoc        v3.20.2
// source: known.proto

package pb

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

type SqrtReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input float64 `protobuf:"fixed64,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *SqrtReq) Reset() {
	*x = SqrtReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_known_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqrtReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqrtReq) ProtoMessage() {}

func (x *SqrtReq) ProtoReflect() protoreflect.Message {
	mi := &file_known_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqrtReq.ProtoReflect.Descriptor instead.
func (*SqrtReq) Descriptor() ([]byte, []int) {
	return file_known_proto_rawDescGZIP(), []int{0}
}

func (x *SqrtReq) GetInput() float64 {
	if x != nil {
		return x.Input
	}
	return 0
}

type SqrtRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Root float64 `protobuf:"fixed64,1,opt,name=root,proto3" json:"root,omitempty"`
}

func (x *SqrtRes) Reset() {
	*x = SqrtRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_known_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SqrtRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqrtRes) ProtoMessage() {}

func (x *SqrtRes) ProtoReflect() protoreflect.Message {
	mi := &file_known_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqrtRes.ProtoReflect.Descriptor instead.
func (*SqrtRes) Descriptor() ([]byte, []int) {
	return file_known_proto_rawDescGZIP(), []int{1}
}

func (x *SqrtRes) GetRoot() float64 {
	if x != nil {
		return x.Root
	}
	return 0
}

type KnowFromReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KnowFromReq) Reset() {
	*x = KnowFromReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_known_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KnowFromReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KnowFromReq) ProtoMessage() {}

func (x *KnowFromReq) ProtoReflect() protoreflect.Message {
	mi := &file_known_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KnowFromReq.ProtoReflect.Descriptor instead.
func (*KnowFromReq) Descriptor() ([]byte, []int) {
	return file_known_proto_rawDescGZIP(), []int{2}
}

type KnowFromRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greetings string `protobuf:"bytes,1,opt,name=greetings,proto3" json:"greetings,omitempty"`
}

func (x *KnowFromRes) Reset() {
	*x = KnowFromRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_known_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KnowFromRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KnowFromRes) ProtoMessage() {}

func (x *KnowFromRes) ProtoReflect() protoreflect.Message {
	mi := &file_known_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KnowFromRes.ProtoReflect.Descriptor instead.
func (*KnowFromRes) Descriptor() ([]byte, []int) {
	return file_known_proto_rawDescGZIP(), []int{3}
}

func (x *KnowFromRes) GetGreetings() string {
	if x != nil {
		return x.Greetings
	}
	return ""
}

var File_known_proto protoreflect.FileDescriptor

var file_known_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a,
	0x07, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x1d,
	0x0a, 0x07, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x22, 0x0d, 0x0a,
	0x0b, 0x4b, 0x6e, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x22, 0x2b, 0x0a, 0x0b,
	0x4b, 0x6e, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x32, 0x4f, 0x0a, 0x05, 0x4b, 0x6e, 0x6f,
	0x77, 0x6e, 0x12, 0x1c, 0x0a, 0x04, 0x53, 0x71, 0x72, 0x74, 0x12, 0x08, 0x2e, 0x53, 0x71, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x08, 0x2e, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x28, 0x0a, 0x08, 0x4b, 0x6e, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0c, 0x2e, 0x4b,
	0x6e, 0x6f, 0x77, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x4b, 0x6e, 0x6f,
	0x77, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_known_proto_rawDescOnce sync.Once
	file_known_proto_rawDescData = file_known_proto_rawDesc
)

func file_known_proto_rawDescGZIP() []byte {
	file_known_proto_rawDescOnce.Do(func() {
		file_known_proto_rawDescData = protoimpl.X.CompressGZIP(file_known_proto_rawDescData)
	})
	return file_known_proto_rawDescData
}

var file_known_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_known_proto_goTypes = []interface{}{
	(*SqrtReq)(nil),     // 0: SqrtReq
	(*SqrtRes)(nil),     // 1: SqrtRes
	(*KnowFromReq)(nil), // 2: KnowFromReq
	(*KnowFromRes)(nil), // 3: KnowFromRes
}
var file_known_proto_depIdxs = []int32{
	0, // 0: Known.Sqrt:input_type -> SqrtReq
	2, // 1: Known.KnowFrom:input_type -> KnowFromReq
	1, // 2: Known.Sqrt:output_type -> SqrtRes
	3, // 3: Known.KnowFrom:output_type -> KnowFromRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_known_proto_init() }
func file_known_proto_init() {
	if File_known_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_known_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqrtReq); i {
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
		file_known_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SqrtRes); i {
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
		file_known_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KnowFromReq); i {
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
		file_known_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KnowFromRes); i {
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
			RawDescriptor: file_known_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_known_proto_goTypes,
		DependencyIndexes: file_known_proto_depIdxs,
		MessageInfos:      file_known_proto_msgTypes,
	}.Build()
	File_known_proto = out.File
	file_known_proto_rawDesc = nil
	file_known_proto_goTypes = nil
	file_known_proto_depIdxs = nil
}
