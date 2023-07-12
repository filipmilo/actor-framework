// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0
// source: messages/proto/messages.proto

package messages

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

type Sum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *Sum) Reset() {
	*x = Sum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sum) ProtoMessage() {}

func (x *Sum) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sum.ProtoReflect.Descriptor instead.
func (*Sum) Descriptor() ([]byte, []int) {
	return file_messages_proto_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Sum) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type AdderMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AdderMessage) Reset() {
	*x = AdderMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdderMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdderMessage) ProtoMessage() {}

func (x *AdderMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdderMessage.ProtoReflect.Descriptor instead.
func (*AdderMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_messages_proto_rawDescGZIP(), []int{1}
}

func (x *AdderMessage) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_messages_proto_messages_proto protoreflect.FileDescriptor

var file_messages_proto_messages_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x17, 0x0a, 0x03, 0x53, 0x75, 0x6d,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73,
	0x75, 0x6d, 0x22, 0x26, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x1a, 0x5a, 0x18, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x2d, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_messages_proto_rawDescOnce sync.Once
	file_messages_proto_messages_proto_rawDescData = file_messages_proto_messages_proto_rawDesc
)

func file_messages_proto_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_messages_proto_rawDescData)
	})
	return file_messages_proto_messages_proto_rawDescData
}

var file_messages_proto_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_messages_proto_messages_proto_goTypes = []interface{}{
	(*Sum)(nil),          // 0: messages.Sum
	(*AdderMessage)(nil), // 1: messages.AdderMessage
}
var file_messages_proto_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_messages_proto_messages_proto_init() }
func file_messages_proto_messages_proto_init() {
	if File_messages_proto_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sum); i {
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
		file_messages_proto_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdderMessage); i {
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
			RawDescriptor: file_messages_proto_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_messages_proto_depIdxs,
		MessageInfos:      file_messages_proto_messages_proto_msgTypes,
	}.Build()
	File_messages_proto_messages_proto = out.File
	file_messages_proto_messages_proto_rawDesc = nil
	file_messages_proto_messages_proto_goTypes = nil
	file_messages_proto_messages_proto_depIdxs = nil
}
