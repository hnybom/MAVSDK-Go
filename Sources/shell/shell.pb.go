// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: shell.proto

package shell

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

// Possible results returned for shell requests
type ShellResult_Result int32

const (
	ShellResult_RESULT_UNKNOWN          ShellResult_Result = 0 // Unknown result
	ShellResult_RESULT_SUCCESS          ShellResult_Result = 1 // Request succeeded
	ShellResult_RESULT_NO_SYSTEM        ShellResult_Result = 2 // No system is connected
	ShellResult_RESULT_CONNECTION_ERROR ShellResult_Result = 3 // Connection error
	ShellResult_RESULT_NO_RESPONSE      ShellResult_Result = 4 // Response was not received
	ShellResult_RESULT_BUSY             ShellResult_Result = 5 // Shell busy (transfer in progress)
)

// Enum value maps for ShellResult_Result.
var (
	ShellResult_Result_name = map[int32]string{
		0: "RESULT_UNKNOWN",
		1: "RESULT_SUCCESS",
		2: "RESULT_NO_SYSTEM",
		3: "RESULT_CONNECTION_ERROR",
		4: "RESULT_NO_RESPONSE",
		5: "RESULT_BUSY",
	}
	ShellResult_Result_value = map[string]int32{
		"RESULT_UNKNOWN":          0,
		"RESULT_SUCCESS":          1,
		"RESULT_NO_SYSTEM":        2,
		"RESULT_CONNECTION_ERROR": 3,
		"RESULT_NO_RESPONSE":      4,
		"RESULT_BUSY":             5,
	}
)

func (x ShellResult_Result) Enum() *ShellResult_Result {
	p := new(ShellResult_Result)
	*p = x
	return p
}

func (x ShellResult_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShellResult_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_shell_proto_enumTypes[0].Descriptor()
}

func (ShellResult_Result) Type() protoreflect.EnumType {
	return &file_shell_proto_enumTypes[0]
}

func (x ShellResult_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShellResult_Result.Descriptor instead.
func (ShellResult_Result) EnumDescriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{4, 0}
}

type SendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Command       string                 `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"` // The command line to send
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	mi := &file_shell_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{0}
}

func (x *SendRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

type SendResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShellResult   *ShellResult           `protobuf:"bytes,1,opt,name=shell_result,json=shellResult,proto3" json:"shell_result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	mi := &file_shell_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{1}
}

func (x *SendResponse) GetShellResult() *ShellResult {
	if x != nil {
		return x.ShellResult
	}
	return nil
}

type SubscribeReceiveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeReceiveRequest) Reset() {
	*x = SubscribeReceiveRequest{}
	mi := &file_shell_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeReceiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReceiveRequest) ProtoMessage() {}

func (x *SubscribeReceiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeReceiveRequest.ProtoReflect.Descriptor instead.
func (*SubscribeReceiveRequest) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{2}
}

type ReceiveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          string                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` // Received data.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReceiveResponse) Reset() {
	*x = ReceiveResponse{}
	mi := &file_shell_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReceiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveResponse) ProtoMessage() {}

func (x *ReceiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveResponse.ProtoReflect.Descriptor instead.
func (*ReceiveResponse) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{3}
}

func (x *ReceiveResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// Result type.
type ShellResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        ShellResult_Result     `protobuf:"varint,1,opt,name=result,proto3,enum=mavsdk.rpc.shell.ShellResult_Result" json:"result,omitempty"` // Result enum value
	ResultStr     string                 `protobuf:"bytes,2,opt,name=result_str,json=resultStr,proto3" json:"result_str,omitempty"`                    // Human-readable English string describing the result
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShellResult) Reset() {
	*x = ShellResult{}
	mi := &file_shell_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShellResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellResult) ProtoMessage() {}

func (x *ShellResult) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellResult.ProtoReflect.Descriptor instead.
func (*ShellResult) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{4}
}

func (x *ShellResult) GetResult() ShellResult_Result {
	if x != nil {
		return x.Result
	}
	return ShellResult_RESULT_UNKNOWN
}

func (x *ShellResult) GetResultStr() string {
	if x != nil {
		return x.ResultStr
	}
	return ""
}

var File_shell_proto protoreflect.FileDescriptor

const file_shell_proto_rawDesc = "" +
	"\n" +
	"\vshell.proto\x12\x10mavsdk.rpc.shell\x1a\x14mavsdk_options.proto\"'\n" +
	"\vSendRequest\x12\x18\n" +
	"\acommand\x18\x01 \x01(\tR\acommand\"P\n" +
	"\fSendResponse\x12@\n" +
	"\fshell_result\x18\x01 \x01(\v2\x1d.mavsdk.rpc.shell.ShellResultR\vshellResult\"\x19\n" +
	"\x17SubscribeReceiveRequest\"%\n" +
	"\x0fReceiveResponse\x12\x12\n" +
	"\x04data\x18\x01 \x01(\tR\x04data\"\xf9\x01\n" +
	"\vShellResult\x12<\n" +
	"\x06result\x18\x01 \x01(\x0e2$.mavsdk.rpc.shell.ShellResult.ResultR\x06result\x12\x1d\n" +
	"\n" +
	"result_str\x18\x02 \x01(\tR\tresultStr\"\x8c\x01\n" +
	"\x06Result\x12\x12\n" +
	"\x0eRESULT_UNKNOWN\x10\x00\x12\x12\n" +
	"\x0eRESULT_SUCCESS\x10\x01\x12\x14\n" +
	"\x10RESULT_NO_SYSTEM\x10\x02\x12\x1b\n" +
	"\x17RESULT_CONNECTION_ERROR\x10\x03\x12\x16\n" +
	"\x12RESULT_NO_RESPONSE\x10\x04\x12\x0f\n" +
	"\vRESULT_BUSY\x10\x052\xc5\x01\n" +
	"\fShellService\x12K\n" +
	"\x04Send\x12\x1d.mavsdk.rpc.shell.SendRequest\x1a\x1e.mavsdk.rpc.shell.SendResponse\"\x04\x80\xb5\x18\x01\x12h\n" +
	"\x10SubscribeReceive\x12).mavsdk.rpc.shell.SubscribeReceiveRequest\x1a!.mavsdk.rpc.shell.ReceiveResponse\"\x04\x80\xb5\x18\x000\x01B\x15B\n" +
	"ShellProtoZ\a.;shellb\x06proto3"

var (
	file_shell_proto_rawDescOnce sync.Once
	file_shell_proto_rawDescData []byte
)

func file_shell_proto_rawDescGZIP() []byte {
	file_shell_proto_rawDescOnce.Do(func() {
		file_shell_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shell_proto_rawDesc), len(file_shell_proto_rawDesc)))
	})
	return file_shell_proto_rawDescData
}

var file_shell_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shell_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_shell_proto_goTypes = []any{
	(ShellResult_Result)(0),         // 0: mavsdk.rpc.shell.ShellResult.Result
	(*SendRequest)(nil),             // 1: mavsdk.rpc.shell.SendRequest
	(*SendResponse)(nil),            // 2: mavsdk.rpc.shell.SendResponse
	(*SubscribeReceiveRequest)(nil), // 3: mavsdk.rpc.shell.SubscribeReceiveRequest
	(*ReceiveResponse)(nil),         // 4: mavsdk.rpc.shell.ReceiveResponse
	(*ShellResult)(nil),             // 5: mavsdk.rpc.shell.ShellResult
}
var file_shell_proto_depIdxs = []int32{
	5, // 0: mavsdk.rpc.shell.SendResponse.shell_result:type_name -> mavsdk.rpc.shell.ShellResult
	0, // 1: mavsdk.rpc.shell.ShellResult.result:type_name -> mavsdk.rpc.shell.ShellResult.Result
	1, // 2: mavsdk.rpc.shell.ShellService.Send:input_type -> mavsdk.rpc.shell.SendRequest
	3, // 3: mavsdk.rpc.shell.ShellService.SubscribeReceive:input_type -> mavsdk.rpc.shell.SubscribeReceiveRequest
	2, // 4: mavsdk.rpc.shell.ShellService.Send:output_type -> mavsdk.rpc.shell.SendResponse
	4, // 5: mavsdk.rpc.shell.ShellService.SubscribeReceive:output_type -> mavsdk.rpc.shell.ReceiveResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shell_proto_init() }
func file_shell_proto_init() {
	if File_shell_proto != nil {
		return
	}
	
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shell_proto_rawDesc), len(file_shell_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shell_proto_goTypes,
		DependencyIndexes: file_shell_proto_depIdxs,
		EnumInfos:         file_shell_proto_enumTypes,
		MessageInfos:      file_shell_proto_msgTypes,
	}.Build()
	File_shell_proto = out.File
	file_shell_proto_goTypes = nil
	file_shell_proto_depIdxs = nil
}
