// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: server_utility.proto

package server_utility

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status types.
type StatusTextType int32

const (
	StatusTextType_STATUS_TEXT_TYPE_DEBUG     StatusTextType = 0 // Debug
	StatusTextType_STATUS_TEXT_TYPE_INFO      StatusTextType = 1 // Information
	StatusTextType_STATUS_TEXT_TYPE_NOTICE    StatusTextType = 2 // Notice
	StatusTextType_STATUS_TEXT_TYPE_WARNING   StatusTextType = 3 // Warning
	StatusTextType_STATUS_TEXT_TYPE_ERROR     StatusTextType = 4 // Error
	StatusTextType_STATUS_TEXT_TYPE_CRITICAL  StatusTextType = 5 // Critical
	StatusTextType_STATUS_TEXT_TYPE_ALERT     StatusTextType = 6 // Alert
	StatusTextType_STATUS_TEXT_TYPE_EMERGENCY StatusTextType = 7 // Emergency
)

// Enum value maps for StatusTextType.
var (
	StatusTextType_name = map[int32]string{
		0: "STATUS_TEXT_TYPE_DEBUG",
		1: "STATUS_TEXT_TYPE_INFO",
		2: "STATUS_TEXT_TYPE_NOTICE",
		3: "STATUS_TEXT_TYPE_WARNING",
		4: "STATUS_TEXT_TYPE_ERROR",
		5: "STATUS_TEXT_TYPE_CRITICAL",
		6: "STATUS_TEXT_TYPE_ALERT",
		7: "STATUS_TEXT_TYPE_EMERGENCY",
	}
	StatusTextType_value = map[string]int32{
		"STATUS_TEXT_TYPE_DEBUG":     0,
		"STATUS_TEXT_TYPE_INFO":      1,
		"STATUS_TEXT_TYPE_NOTICE":    2,
		"STATUS_TEXT_TYPE_WARNING":   3,
		"STATUS_TEXT_TYPE_ERROR":     4,
		"STATUS_TEXT_TYPE_CRITICAL":  5,
		"STATUS_TEXT_TYPE_ALERT":     6,
		"STATUS_TEXT_TYPE_EMERGENCY": 7,
	}
)

func (x StatusTextType) Enum() *StatusTextType {
	p := new(StatusTextType)
	*p = x
	return p
}

func (x StatusTextType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusTextType) Descriptor() protoreflect.EnumDescriptor {
	return file_server_utility_proto_enumTypes[0].Descriptor()
}

func (StatusTextType) Type() protoreflect.EnumType {
	return &file_server_utility_proto_enumTypes[0]
}

func (x StatusTextType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusTextType.Descriptor instead.
func (StatusTextType) EnumDescriptor() ([]byte, []int) {
	return file_server_utility_proto_rawDescGZIP(), []int{0}
}

// Possible results returned for server utility requests.
type ServerUtilityResult_Result int32

const (
	ServerUtilityResult_RESULT_UNKNOWN          ServerUtilityResult_Result = 0 // Unknown result
	ServerUtilityResult_RESULT_SUCCESS          ServerUtilityResult_Result = 1 // Request succeeded
	ServerUtilityResult_RESULT_NO_SYSTEM        ServerUtilityResult_Result = 2 // No system is connected
	ServerUtilityResult_RESULT_CONNECTION_ERROR ServerUtilityResult_Result = 3 // Connection error
	ServerUtilityResult_RESULT_INVALID_ARGUMENT ServerUtilityResult_Result = 4 // Invalid argument
)

// Enum value maps for ServerUtilityResult_Result.
var (
	ServerUtilityResult_Result_name = map[int32]string{
		0: "RESULT_UNKNOWN",
		1: "RESULT_SUCCESS",
		2: "RESULT_NO_SYSTEM",
		3: "RESULT_CONNECTION_ERROR",
		4: "RESULT_INVALID_ARGUMENT",
	}
	ServerUtilityResult_Result_value = map[string]int32{
		"RESULT_UNKNOWN":          0,
		"RESULT_SUCCESS":          1,
		"RESULT_NO_SYSTEM":        2,
		"RESULT_CONNECTION_ERROR": 3,
		"RESULT_INVALID_ARGUMENT": 4,
	}
)

func (x ServerUtilityResult_Result) Enum() *ServerUtilityResult_Result {
	p := new(ServerUtilityResult_Result)
	*p = x
	return p
}

func (x ServerUtilityResult_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerUtilityResult_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_server_utility_proto_enumTypes[1].Descriptor()
}

func (ServerUtilityResult_Result) Type() protoreflect.EnumType {
	return &file_server_utility_proto_enumTypes[1]
}

func (x ServerUtilityResult_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerUtilityResult_Result.Descriptor instead.
func (ServerUtilityResult_Result) EnumDescriptor() ([]byte, []int) {
	return file_server_utility_proto_rawDescGZIP(), []int{2, 0}
}

type SendStatusTextRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          StatusTextType         `protobuf:"varint,1,opt,name=type,proto3,enum=mavsdk.rpc.server_utility.StatusTextType" json:"type,omitempty"` // The text to send
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`                                                // Text message
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendStatusTextRequest) Reset() {
	*x = SendStatusTextRequest{}
	mi := &file_server_utility_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendStatusTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendStatusTextRequest) ProtoMessage() {}

func (x *SendStatusTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_utility_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendStatusTextRequest.ProtoReflect.Descriptor instead.
func (*SendStatusTextRequest) Descriptor() ([]byte, []int) {
	return file_server_utility_proto_rawDescGZIP(), []int{0}
}

func (x *SendStatusTextRequest) GetType() StatusTextType {
	if x != nil {
		return x.Type
	}
	return StatusTextType_STATUS_TEXT_TYPE_DEBUG
}

func (x *SendStatusTextRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SendStatusTextResponse struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	ServerUtilityResult *ServerUtilityResult   `protobuf:"bytes,1,opt,name=server_utility_result,json=serverUtilityResult,proto3" json:"server_utility_result,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *SendStatusTextResponse) Reset() {
	*x = SendStatusTextResponse{}
	mi := &file_server_utility_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendStatusTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendStatusTextResponse) ProtoMessage() {}

func (x *SendStatusTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_utility_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendStatusTextResponse.ProtoReflect.Descriptor instead.
func (*SendStatusTextResponse) Descriptor() ([]byte, []int) {
	return file_server_utility_proto_rawDescGZIP(), []int{1}
}

func (x *SendStatusTextResponse) GetServerUtilityResult() *ServerUtilityResult {
	if x != nil {
		return x.ServerUtilityResult
	}
	return nil
}

type ServerUtilityResult struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Result        ServerUtilityResult_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mavsdk.rpc.server_utility.ServerUtilityResult_Result" json:"result,omitempty"` // Result enum value
	ResultStr     string                     `protobuf:"bytes,2,opt,name=result_str,json=resultStr,proto3" json:"result_str,omitempty"`                                     // Human-readable English string describing the result
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerUtilityResult) Reset() {
	*x = ServerUtilityResult{}
	mi := &file_server_utility_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerUtilityResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerUtilityResult) ProtoMessage() {}

func (x *ServerUtilityResult) ProtoReflect() protoreflect.Message {
	mi := &file_server_utility_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerUtilityResult.ProtoReflect.Descriptor instead.
func (*ServerUtilityResult) Descriptor() ([]byte, []int) {
	return file_server_utility_proto_rawDescGZIP(), []int{2}
}

func (x *ServerUtilityResult) GetResult() ServerUtilityResult_Result {
	if x != nil {
		return x.Result
	}
	return ServerUtilityResult_RESULT_UNKNOWN
}

func (x *ServerUtilityResult) GetResultStr() string {
	if x != nil {
		return x.ResultStr
	}
	return ""
}

var File_server_utility_proto protoreflect.FileDescriptor

const file_server_utility_proto_rawDesc = "" +
	"\n" +
	"\x14server_utility.proto\x12\x19mavsdk.rpc.server_utility\x1a\x14mavsdk_options.proto\"j\n" +
	"\x15SendStatusTextRequest\x12=\n" +
	"\x04type\x18\x01 \x01(\x0e2).mavsdk.rpc.server_utility.StatusTextTypeR\x04type\x12\x12\n" +
	"\x04text\x18\x02 \x01(\tR\x04text\"|\n" +
	"\x16SendStatusTextResponse\x12b\n" +
	"\x15server_utility_result\x18\x01 \x01(\v2..mavsdk.rpc.server_utility.ServerUtilityResultR\x13serverUtilityResult\"\x86\x02\n" +
	"\x13ServerUtilityResult\x12M\n" +
	"\x06result\x18\x01 \x01(\x0e25.mavsdk.rpc.server_utility.ServerUtilityResult.ResultR\x06result\x12\x1d\n" +
	"\n" +
	"result_str\x18\x02 \x01(\tR\tresultStr\"\x80\x01\n" +
	"\x06Result\x12\x12\n" +
	"\x0eRESULT_UNKNOWN\x10\x00\x12\x12\n" +
	"\x0eRESULT_SUCCESS\x10\x01\x12\x14\n" +
	"\x10RESULT_NO_SYSTEM\x10\x02\x12\x1b\n" +
	"\x17RESULT_CONNECTION_ERROR\x10\x03\x12\x1b\n" +
	"\x17RESULT_INVALID_ARGUMENT\x10\x04*\xf9\x01\n" +
	"\x0eStatusTextType\x12\x1a\n" +
	"\x16STATUS_TEXT_TYPE_DEBUG\x10\x00\x12\x19\n" +
	"\x15STATUS_TEXT_TYPE_INFO\x10\x01\x12\x1b\n" +
	"\x17STATUS_TEXT_TYPE_NOTICE\x10\x02\x12\x1c\n" +
	"\x18STATUS_TEXT_TYPE_WARNING\x10\x03\x12\x1a\n" +
	"\x16STATUS_TEXT_TYPE_ERROR\x10\x04\x12\x1d\n" +
	"\x19STATUS_TEXT_TYPE_CRITICAL\x10\x05\x12\x1a\n" +
	"\x16STATUS_TEXT_TYPE_ALERT\x10\x06\x12\x1e\n" +
	"\x1aSTATUS_TEXT_TYPE_EMERGENCY\x10\a2\x93\x01\n" +
	"\x14ServerUtilityService\x12{\n" +
	"\x0eSendStatusText\x120.mavsdk.rpc.server_utility.SendStatusTextRequest\x1a1.mavsdk.rpc.server_utility.SendStatusTextResponse\"\x04\x80\xb5\x18\x01B&B\x12ServerUtilityProtoZ\x10.;server_utilityb\x06proto3"

var (
	file_server_utility_proto_rawDescOnce sync.Once
	file_server_utility_proto_rawDescData []byte
)

func file_server_utility_proto_rawDescGZIP() []byte {
	file_server_utility_proto_rawDescOnce.Do(func() {
		file_server_utility_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_server_utility_proto_rawDesc), len(file_server_utility_proto_rawDesc)))
	})
	return file_server_utility_proto_rawDescData
}

var file_server_utility_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_server_utility_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_server_utility_proto_goTypes = []any{
	(StatusTextType)(0),             // 0: mavsdk.rpc.server_utility.StatusTextType
	(ServerUtilityResult_Result)(0), // 1: mavsdk.rpc.server_utility.ServerUtilityResult.Result
	(*SendStatusTextRequest)(nil),   // 2: mavsdk.rpc.server_utility.SendStatusTextRequest
	(*SendStatusTextResponse)(nil),  // 3: mavsdk.rpc.server_utility.SendStatusTextResponse
	(*ServerUtilityResult)(nil),     // 4: mavsdk.rpc.server_utility.ServerUtilityResult
}
var file_server_utility_proto_depIdxs = []int32{
	0, // 0: mavsdk.rpc.server_utility.SendStatusTextRequest.type:type_name -> mavsdk.rpc.server_utility.StatusTextType
	4, // 1: mavsdk.rpc.server_utility.SendStatusTextResponse.server_utility_result:type_name -> mavsdk.rpc.server_utility.ServerUtilityResult
	1, // 2: mavsdk.rpc.server_utility.ServerUtilityResult.result:type_name -> mavsdk.rpc.server_utility.ServerUtilityResult.Result
	2, // 3: mavsdk.rpc.server_utility.ServerUtilityService.SendStatusText:input_type -> mavsdk.rpc.server_utility.SendStatusTextRequest
	3, // 4: mavsdk.rpc.server_utility.ServerUtilityService.SendStatusText:output_type -> mavsdk.rpc.server_utility.SendStatusTextResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_server_utility_proto_init() }
func file_server_utility_proto_init() {
	if File_server_utility_proto != nil {
		return
	}
	
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_server_utility_proto_rawDesc), len(file_server_utility_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_utility_proto_goTypes,
		DependencyIndexes: file_server_utility_proto_depIdxs,
		EnumInfos:         file_server_utility_proto_enumTypes,
		MessageInfos:      file_server_utility_proto_msgTypes,
	}.Build()
	File_server_utility_proto = out.File
	file_server_utility_proto_goTypes = nil
	file_server_utility_proto_depIdxs = nil
}
