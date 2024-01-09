// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: interface/calculator.proto

package calculator

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

type CalculatorRequest_Operation int32

const (
	CalculatorRequest_UNKNOWN CalculatorRequest_Operation = 0
	CalculatorRequest_SUM     CalculatorRequest_Operation = 1
)

// Enum value maps for CalculatorRequest_Operation.
var (
	CalculatorRequest_Operation_name = map[int32]string{
		0: "UNKNOWN",
		1: "SUM",
	}
	CalculatorRequest_Operation_value = map[string]int32{
		"UNKNOWN": 0,
		"SUM":     1,
	}
)

func (x CalculatorRequest_Operation) Enum() *CalculatorRequest_Operation {
	p := new(CalculatorRequest_Operation)
	*p = x
	return p
}

func (x CalculatorRequest_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CalculatorRequest_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_interface_calculator_proto_enumTypes[0].Descriptor()
}

func (CalculatorRequest_Operation) Type() protoreflect.EnumType {
	return &file_interface_calculator_proto_enumTypes[0]
}

func (x CalculatorRequest_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CalculatorRequest_Operation.Descriptor instead.
func (CalculatorRequest_Operation) EnumDescriptor() ([]byte, []int) {
	return file_interface_calculator_proto_rawDescGZIP(), []int{0, 0}
}

type CalculatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation    CalculatorRequest_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=protocols.CalculatorRequest_Operation" json:"operation,omitempty"`
	FirstNumber  int64                       `protobuf:"varint,2,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber int64                       `protobuf:"varint,3,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
}

func (x *CalculatorRequest) Reset() {
	*x = CalculatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest) ProtoMessage() {}

func (x *CalculatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interface_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest.ProtoReflect.Descriptor instead.
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return file_interface_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *CalculatorRequest) GetOperation() CalculatorRequest_Operation {
	if x != nil {
		return x.Operation
	}
	return CalculatorRequest_UNKNOWN
}

func (x *CalculatorRequest) GetFirstNumber() int64 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *CalculatorRequest) GetSecondNumber() int64 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type CalculatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalculatorResponse) Reset() {
	*x = CalculatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorResponse) ProtoMessage() {}

func (x *CalculatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interface_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorResponse.ProtoReflect.Descriptor instead.
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return file_interface_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculatorResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_interface_calculator_proto protoreflect.FileDescriptor

var file_interface_calculator_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x09, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x55, 0x4d, 0x10, 0x01, 0x22, 0x2c,
	0x0a, 0x12, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xa2, 0x01, 0x0a,
	0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x49, 0x0a, 0x0a, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x31, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x32, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x16, 0x5a, 0x14, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_interface_calculator_proto_rawDescOnce sync.Once
	file_interface_calculator_proto_rawDescData = file_interface_calculator_proto_rawDesc
)

func file_interface_calculator_proto_rawDescGZIP() []byte {
	file_interface_calculator_proto_rawDescOnce.Do(func() {
		file_interface_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_interface_calculator_proto_rawDescData)
	})
	return file_interface_calculator_proto_rawDescData
}

var file_interface_calculator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_interface_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_interface_calculator_proto_goTypes = []interface{}{
	(CalculatorRequest_Operation)(0), // 0: protocols.CalculatorRequest.Operation
	(*CalculatorRequest)(nil),        // 1: protocols.CalculatorRequest
	(*CalculatorResponse)(nil),       // 2: protocols.CalculatorResponse
}
var file_interface_calculator_proto_depIdxs = []int32{
	0, // 0: protocols.CalculatorRequest.operation:type_name -> protocols.CalculatorRequest.Operation
	1, // 1: protocols.Calculator.Calculate1:input_type -> protocols.CalculatorRequest
	1, // 2: protocols.Calculator.Calculate2:input_type -> protocols.CalculatorRequest
	2, // 3: protocols.Calculator.Calculate1:output_type -> protocols.CalculatorResponse
	2, // 4: protocols.Calculator.Calculate2:output_type -> protocols.CalculatorResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_interface_calculator_proto_init() }
func file_interface_calculator_proto_init() {
	if File_interface_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interface_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorRequest); i {
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
		file_interface_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorResponse); i {
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
			RawDescriptor: file_interface_calculator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interface_calculator_proto_goTypes,
		DependencyIndexes: file_interface_calculator_proto_depIdxs,
		EnumInfos:         file_interface_calculator_proto_enumTypes,
		MessageInfos:      file_interface_calculator_proto_msgTypes,
	}.Build()
	File_interface_calculator_proto = out.File
	file_interface_calculator_proto_rawDesc = nil
	file_interface_calculator_proto_goTypes = nil
	file_interface_calculator_proto_depIdxs = nil
}
