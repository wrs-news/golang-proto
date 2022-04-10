// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: proto/core/error.proto

package core

import (
	_ "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//	*Error_Null
	//	*Error_Error
	Kind isError_Kind `protobuf_oneof:"kind"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_core_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_core_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_core_error_proto_rawDescGZIP(), []int{0}
}

func (m *Error) GetKind() isError_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Error) GetNull() _struct.NullValue {
	if x, ok := x.GetKind().(*Error_Null); ok {
		return x.Null
	}
	return _struct.NullValue(0)
}

func (x *Error) GetError() *ErrorValue {
	if x, ok := x.GetKind().(*Error_Error); ok {
		return x.Error
	}
	return nil
}

type isError_Kind interface {
	isError_Kind()
}

type Error_Null struct {
	Null _struct.NullValue `protobuf:"varint,1,opt,name=null,proto3,enum=google.protobuf.NullValue,oneof"`
}

type Error_Error struct {
	Error *ErrorValue `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*Error_Null) isError_Kind() {}

func (*Error_Error) isError_Kind() {}

type ErrorValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InternalError *InternalError `protobuf:"bytes,1,opt,name=internal_error,json=internalError,proto3" json:"internal_error,omitempty"`
	ExternalError *ExternalError `protobuf:"bytes,2,opt,name=external_error,json=externalError,proto3" json:"external_error,omitempty"`
}

func (x *ErrorValue) Reset() {
	*x = ErrorValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_core_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorValue) ProtoMessage() {}

func (x *ErrorValue) ProtoReflect() protoreflect.Message {
	mi := &file_proto_core_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorValue.ProtoReflect.Descriptor instead.
func (*ErrorValue) Descriptor() ([]byte, []int) {
	return file_proto_core_error_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorValue) GetInternalError() *InternalError {
	if x != nil {
		return x.InternalError
	}
	return nil
}

func (x *ErrorValue) GetExternalError() *ExternalError {
	if x != nil {
		return x.ExternalError
	}
	return nil
}

type InternalError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Backtrace string `protobuf:"bytes,3,opt,name=backtrace,proto3" json:"backtrace,omitempty"`
}

func (x *InternalError) Reset() {
	*x = InternalError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_core_error_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalError) ProtoMessage() {}

func (x *InternalError) ProtoReflect() protoreflect.Message {
	mi := &file_proto_core_error_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalError.ProtoReflect.Descriptor instead.
func (*InternalError) Descriptor() ([]byte, []int) {
	return file_proto_core_error_proto_rawDescGZIP(), []int{2}
}

func (x *InternalError) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *InternalError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *InternalError) GetBacktrace() string {
	if x != nil {
		return x.Backtrace
	}
	return ""
}

type ExternalError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ExternalError) Reset() {
	*x = ExternalError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_core_error_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalError) ProtoMessage() {}

func (x *ExternalError) ProtoReflect() protoreflect.Message {
	mi := &file_proto_core_error_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalError.ProtoReflect.Descriptor instead.
func (*ExternalError) Descriptor() ([]byte, []int) {
	return file_proto_core_error_proto_rawDescGZIP(), []int{3}
}

func (x *ExternalError) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ExternalError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_core_error_proto protoreflect.FileDescriptor

var file_proto_core_error_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x04,
	0x6e, 0x75, 0x6c, 0x6c, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x06,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x3a, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x0d,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5b, 0x0a,
	0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x62, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x22, 0x3d, 0x0a, 0x0d, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x72, 0x73, 0x2d, 0x6e, 0x65, 0x77, 0x73,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_core_error_proto_rawDescOnce sync.Once
	file_proto_core_error_proto_rawDescData = file_proto_core_error_proto_rawDesc
)

func file_proto_core_error_proto_rawDescGZIP() []byte {
	file_proto_core_error_proto_rawDescOnce.Do(func() {
		file_proto_core_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_core_error_proto_rawDescData)
	})
	return file_proto_core_error_proto_rawDescData
}

var file_proto_core_error_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_core_error_proto_goTypes = []interface{}{
	(*Error)(nil),          // 0: core.Error
	(*ErrorValue)(nil),     // 1: core.ErrorValue
	(*InternalError)(nil),  // 2: core.InternalError
	(*ExternalError)(nil),  // 3: core.ExternalError
	(_struct.NullValue)(0), // 4: google.protobuf.NullValue
}
var file_proto_core_error_proto_depIdxs = []int32{
	4, // 0: core.Error.null:type_name -> google.protobuf.NullValue
	1, // 1: core.Error.error:type_name -> core.ErrorValue
	2, // 2: core.ErrorValue.internal_error:type_name -> core.InternalError
	3, // 3: core.ErrorValue.external_error:type_name -> core.ExternalError
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_core_error_proto_init() }
func file_proto_core_error_proto_init() {
	if File_proto_core_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_core_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_proto_core_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorValue); i {
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
		file_proto_core_error_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalError); i {
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
		file_proto_core_error_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalError); i {
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
	file_proto_core_error_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Error_Null)(nil),
		(*Error_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_core_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_core_error_proto_goTypes,
		DependencyIndexes: file_proto_core_error_proto_depIdxs,
		MessageInfos:      file_proto_core_error_proto_msgTypes,
	}.Build()
	File_proto_core_error_proto = out.File
	file_proto_core_error_proto_rawDesc = nil
	file_proto_core_error_proto_goTypes = nil
	file_proto_core_error_proto_depIdxs = nil
}
