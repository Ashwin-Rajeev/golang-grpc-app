// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: app/app.proto

package app

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

// The request message containing the user's name.
type StudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StudentRequest) Reset() {
	*x = StudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentRequest) ProtoMessage() {}

func (x *StudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentRequest.ProtoReflect.Descriptor instead.
func (*StudentRequest) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{0}
}

func (x *StudentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// The response message containing the greetings
type StudentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code    string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Program string `protobuf:"bytes,4,opt,name=program,proto3" json:"program,omitempty"`
}

func (x *StudentResp) Reset() {
	*x = StudentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentResp) ProtoMessage() {}

func (x *StudentResp) ProtoReflect() protoreflect.Message {
	mi := &file_app_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentResp.ProtoReflect.Descriptor instead.
func (*StudentResp) Descriptor() ([]byte, []int) {
	return file_app_app_proto_rawDescGZIP(), []int{1}
}

func (x *StudentResp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StudentResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StudentResp) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *StudentResp) GetProgram() string {
	if x != nil {
		return x.Program
	}
	return ""
}

var File_app_app_proto protoreflect.FileDescriptor

var file_app_app_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x70, 0x22, 0x20, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5f, 0x0a, 0x0b, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x32, 0x42, 0x0a, 0x08, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x13, 0x42, 0x08, 0x41,
	0x70, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_app_proto_rawDescOnce sync.Once
	file_app_app_proto_rawDescData = file_app_app_proto_rawDesc
)

func file_app_app_proto_rawDescGZIP() []byte {
	file_app_app_proto_rawDescOnce.Do(func() {
		file_app_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_app_proto_rawDescData)
	})
	return file_app_app_proto_rawDescData
}

var file_app_app_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_app_proto_goTypes = []interface{}{
	(*StudentRequest)(nil), // 0: app.StudentRequest
	(*StudentResp)(nil),    // 1: app.StudentResp
}
var file_app_app_proto_depIdxs = []int32{
	0, // 0: app.Students.GetStudents:input_type -> app.StudentRequest
	1, // 1: app.Students.GetStudents:output_type -> app.StudentResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_app_proto_init() }
func file_app_app_proto_init() {
	if File_app_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentRequest); i {
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
		file_app_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentResp); i {
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
			RawDescriptor: file_app_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_app_proto_goTypes,
		DependencyIndexes: file_app_app_proto_depIdxs,
		MessageInfos:      file_app_app_proto_msgTypes,
	}.Build()
	File_app_app_proto = out.File
	file_app_app_proto_rawDesc = nil
	file_app_app_proto_goTypes = nil
	file_app_app_proto_depIdxs = nil
}
