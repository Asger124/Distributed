// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: time.proto

package proto

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

type Timerequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Timerequest) Reset() {
	*x = Timerequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timerequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timerequest) ProtoMessage() {}

func (x *Timerequest) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timerequest.ProtoReflect.Descriptor instead.
func (*Timerequest) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{0}
}

type Timeresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Timeresponse) Reset() {
	*x = Timeresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timeresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timeresponse) ProtoMessage() {}

func (x *Timeresponse) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timeresponse.ProtoReflect.Descriptor instead.
func (*Timeresponse) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{1}
}

func (x *Timeresponse) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

var File_time_proto protoreflect.FileDescriptor

var file_time_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b,
	0x74, 0x69, 0x6d, 0x65, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x0c, 0x74,
	0x69, 0x6d, 0x65, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32,
	0x35, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26,
	0x0a, 0x07, 0x67, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0c, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x30, 0x35,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_time_proto_rawDescOnce sync.Once
	file_time_proto_rawDescData = file_time_proto_rawDesc
)

func file_time_proto_rawDescGZIP() []byte {
	file_time_proto_rawDescOnce.Do(func() {
		file_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_time_proto_rawDescData)
	})
	return file_time_proto_rawDescData
}

var file_time_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_time_proto_goTypes = []interface{}{
	(*Timerequest)(nil),  // 0: timerequest
	(*Timeresponse)(nil), // 1: timeresponse
}
var file_time_proto_depIdxs = []int32{
	0, // 0: timeService.getTime:input_type -> timerequest
	1, // 1: timeService.getTime:output_type -> timeresponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_time_proto_init() }
func file_time_proto_init() {
	if File_time_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timerequest); i {
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
		file_time_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timeresponse); i {
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
			RawDescriptor: file_time_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_time_proto_goTypes,
		DependencyIndexes: file_time_proto_depIdxs,
		MessageInfos:      file_time_proto_msgTypes,
	}.Build()
	File_time_proto = out.File
	file_time_proto_rawDesc = nil
	file_time_proto_goTypes = nil
	file_time_proto_depIdxs = nil
}