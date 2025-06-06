// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: calibration.proto

package calibration

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

// Possible results returned for calibration commands
type CalibrationResult_Result int32

const (
	CalibrationResult_RESULT_UNKNOWN          CalibrationResult_Result = 0  // Unknown result
	CalibrationResult_RESULT_SUCCESS          CalibrationResult_Result = 1  // The calibration succeeded
	CalibrationResult_RESULT_NEXT             CalibrationResult_Result = 2  // Intermediate message showing progress or instructions on the next steps
	CalibrationResult_RESULT_FAILED           CalibrationResult_Result = 3  // Calibration failed
	CalibrationResult_RESULT_NO_SYSTEM        CalibrationResult_Result = 4  // No system is connected
	CalibrationResult_RESULT_CONNECTION_ERROR CalibrationResult_Result = 5  // Connection error
	CalibrationResult_RESULT_BUSY             CalibrationResult_Result = 6  // Vehicle is busy
	CalibrationResult_RESULT_COMMAND_DENIED   CalibrationResult_Result = 7  // Command refused by vehicle
	CalibrationResult_RESULT_TIMEOUT          CalibrationResult_Result = 8  // Command timed out
	CalibrationResult_RESULT_CANCELLED        CalibrationResult_Result = 9  // Calibration process was cancelled
	CalibrationResult_RESULT_FAILED_ARMED     CalibrationResult_Result = 10 // Calibration process failed since the vehicle is armed
	CalibrationResult_RESULT_UNSUPPORTED      CalibrationResult_Result = 11 // Functionality not supported
)

// Enum value maps for CalibrationResult_Result.
var (
	CalibrationResult_Result_name = map[int32]string{
		0:  "RESULT_UNKNOWN",
		1:  "RESULT_SUCCESS",
		2:  "RESULT_NEXT",
		3:  "RESULT_FAILED",
		4:  "RESULT_NO_SYSTEM",
		5:  "RESULT_CONNECTION_ERROR",
		6:  "RESULT_BUSY",
		7:  "RESULT_COMMAND_DENIED",
		8:  "RESULT_TIMEOUT",
		9:  "RESULT_CANCELLED",
		10: "RESULT_FAILED_ARMED",
		11: "RESULT_UNSUPPORTED",
	}
	CalibrationResult_Result_value = map[string]int32{
		"RESULT_UNKNOWN":          0,
		"RESULT_SUCCESS":          1,
		"RESULT_NEXT":             2,
		"RESULT_FAILED":           3,
		"RESULT_NO_SYSTEM":        4,
		"RESULT_CONNECTION_ERROR": 5,
		"RESULT_BUSY":             6,
		"RESULT_COMMAND_DENIED":   7,
		"RESULT_TIMEOUT":          8,
		"RESULT_CANCELLED":        9,
		"RESULT_FAILED_ARMED":     10,
		"RESULT_UNSUPPORTED":      11,
	}
)

func (x CalibrationResult_Result) Enum() *CalibrationResult_Result {
	p := new(CalibrationResult_Result)
	*p = x
	return p
}

func (x CalibrationResult_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CalibrationResult_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_calibration_proto_enumTypes[0].Descriptor()
}

func (CalibrationResult_Result) Type() protoreflect.EnumType {
	return &file_calibration_proto_enumTypes[0]
}

func (x CalibrationResult_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CalibrationResult_Result.Descriptor instead.
func (CalibrationResult_Result) EnumDescriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{12, 0}
}

type SubscribeCalibrateGyroRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeCalibrateGyroRequest) Reset() {
	*x = SubscribeCalibrateGyroRequest{}
	mi := &file_calibration_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeCalibrateGyroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeCalibrateGyroRequest) ProtoMessage() {}

func (x *SubscribeCalibrateGyroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeCalibrateGyroRequest.ProtoReflect.Descriptor instead.
func (*SubscribeCalibrateGyroRequest) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{0}
}

type CalibrateGyroResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	CalibrationResult *CalibrationResult     `protobuf:"bytes,1,opt,name=calibration_result,json=calibrationResult,proto3" json:"calibration_result,omitempty"`
	ProgressData      *ProgressData          `protobuf:"bytes,2,opt,name=progress_data,json=progressData,proto3" json:"progress_data,omitempty"` // Progress data
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CalibrateGyroResponse) Reset() {
	*x = CalibrateGyroResponse{}
	mi := &file_calibration_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalibrateGyroResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalibrateGyroResponse) ProtoMessage() {}

func (x *CalibrateGyroResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalibrateGyroResponse.ProtoReflect.Descriptor instead.
func (*CalibrateGyroResponse) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{1}
}

func (x *CalibrateGyroResponse) GetCalibrationResult() *CalibrationResult {
	if x != nil {
		return x.CalibrationResult
	}
	return nil
}

func (x *CalibrateGyroResponse) GetProgressData() *ProgressData {
	if x != nil {
		return x.ProgressData
	}
	return nil
}

type SubscribeCalibrateAccelerometerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeCalibrateAccelerometerRequest) Reset() {
	*x = SubscribeCalibrateAccelerometerRequest{}
	mi := &file_calibration_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeCalibrateAccelerometerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeCalibrateAccelerometerRequest) ProtoMessage() {}

func (x *SubscribeCalibrateAccelerometerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeCalibrateAccelerometerRequest.ProtoReflect.Descriptor instead.
func (*SubscribeCalibrateAccelerometerRequest) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{2}
}

type CalibrateAccelerometerResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	CalibrationResult *CalibrationResult     `protobuf:"bytes,1,opt,name=calibration_result,json=calibrationResult,proto3" json:"calibration_result,omitempty"`
	ProgressData      *ProgressData          `protobuf:"bytes,2,opt,name=progress_data,json=progressData,proto3" json:"progress_data,omitempty"` // Progress data
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CalibrateAccelerometerResponse) Reset() {
	*x = CalibrateAccelerometerResponse{}
	mi := &file_calibration_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalibrateAccelerometerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalibrateAccelerometerResponse) ProtoMessage() {}

func (x *CalibrateAccelerometerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalibrateAccelerometerResponse.ProtoReflect.Descriptor instead.
func (*CalibrateAccelerometerResponse) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{3}
}

func (x *CalibrateAccelerometerResponse) GetCalibrationResult() *CalibrationResult {
	if x != nil {
		return x.CalibrationResult
	}
	return nil
}

func (x *CalibrateAccelerometerResponse) GetProgressData() *ProgressData {
	if x != nil {
		return x.ProgressData
	}
	return nil
}

type SubscribeCalibrateMagnetometerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeCalibrateMagnetometerRequest) Reset() {
	*x = SubscribeCalibrateMagnetometerRequest{}
	mi := &file_calibration_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeCalibrateMagnetometerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeCalibrateMagnetometerRequest) ProtoMessage() {}

func (x *SubscribeCalibrateMagnetometerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeCalibrateMagnetometerRequest.ProtoReflect.Descriptor instead.
func (*SubscribeCalibrateMagnetometerRequest) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{4}
}

type CalibrateMagnetometerResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	CalibrationResult *CalibrationResult     `protobuf:"bytes,1,opt,name=calibration_result,json=calibrationResult,proto3" json:"calibration_result,omitempty"`
	ProgressData      *ProgressData          `protobuf:"bytes,2,opt,name=progress_data,json=progressData,proto3" json:"progress_data,omitempty"` // Progress data
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CalibrateMagnetometerResponse) Reset() {
	*x = CalibrateMagnetometerResponse{}
	mi := &file_calibration_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalibrateMagnetometerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalibrateMagnetometerResponse) ProtoMessage() {}

func (x *CalibrateMagnetometerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalibrateMagnetometerResponse.ProtoReflect.Descriptor instead.
func (*CalibrateMagnetometerResponse) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{5}
}

func (x *CalibrateMagnetometerResponse) GetCalibrationResult() *CalibrationResult {
	if x != nil {
		return x.CalibrationResult
	}
	return nil
}

func (x *CalibrateMagnetometerResponse) GetProgressData() *ProgressData {
	if x != nil {
		return x.ProgressData
	}
	return nil
}

type SubscribeCalibrateLevelHorizonRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeCalibrateLevelHorizonRequest) Reset() {
	*x = SubscribeCalibrateLevelHorizonRequest{}
	mi := &file_calibration_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeCalibrateLevelHorizonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeCalibrateLevelHorizonRequest) ProtoMessage() {}

func (x *SubscribeCalibrateLevelHorizonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeCalibrateLevelHorizonRequest.ProtoReflect.Descriptor instead.
func (*SubscribeCalibrateLevelHorizonRequest) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{6}
}

type CalibrateLevelHorizonResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	CalibrationResult *CalibrationResult     `protobuf:"bytes,1,opt,name=calibration_result,json=calibrationResult,proto3" json:"calibration_result,omitempty"`
	ProgressData      *ProgressData          `protobuf:"bytes,2,opt,name=progress_data,json=progressData,proto3" json:"progress_data,omitempty"` // Progress data
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CalibrateLevelHorizonResponse) Reset() {
	*x = CalibrateLevelHorizonResponse{}
	mi := &file_calibration_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalibrateLevelHorizonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalibrateLevelHorizonResponse) ProtoMessage() {}

func (x *CalibrateLevelHorizonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalibrateLevelHorizonResponse.ProtoReflect.Descriptor instead.
func (*CalibrateLevelHorizonResponse) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{7}
}

func (x *CalibrateLevelHorizonResponse) GetCalibrationResult() *CalibrationResult {
	if x != nil {
		return x.CalibrationResult
	}
	return nil
}

func (x *CalibrateLevelHorizonResponse) GetProgressData() *ProgressData {
	if x != nil {
		return x.ProgressData
	}
	return nil
}

type SubscribeCalibrateGimbalAccelerometerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeCalibrateGimbalAccelerometerRequest) Reset() {
	*x = SubscribeCalibrateGimbalAccelerometerRequest{}
	mi := &file_calibration_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeCalibrateGimbalAccelerometerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeCalibrateGimbalAccelerometerRequest) ProtoMessage() {}

func (x *SubscribeCalibrateGimbalAccelerometerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeCalibrateGimbalAccelerometerRequest.ProtoReflect.Descriptor instead.
func (*SubscribeCalibrateGimbalAccelerometerRequest) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{8}
}

type CalibrateGimbalAccelerometerResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	CalibrationResult *CalibrationResult     `protobuf:"bytes,1,opt,name=calibration_result,json=calibrationResult,proto3" json:"calibration_result,omitempty"`
	ProgressData      *ProgressData          `protobuf:"bytes,2,opt,name=progress_data,json=progressData,proto3" json:"progress_data,omitempty"` // Progress data
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CalibrateGimbalAccelerometerResponse) Reset() {
	*x = CalibrateGimbalAccelerometerResponse{}
	mi := &file_calibration_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalibrateGimbalAccelerometerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalibrateGimbalAccelerometerResponse) ProtoMessage() {}

func (x *CalibrateGimbalAccelerometerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalibrateGimbalAccelerometerResponse.ProtoReflect.Descriptor instead.
func (*CalibrateGimbalAccelerometerResponse) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{9}
}

func (x *CalibrateGimbalAccelerometerResponse) GetCalibrationResult() *CalibrationResult {
	if x != nil {
		return x.CalibrationResult
	}
	return nil
}

func (x *CalibrateGimbalAccelerometerResponse) GetProgressData() *ProgressData {
	if x != nil {
		return x.ProgressData
	}
	return nil
}

type CancelRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelRequest) Reset() {
	*x = CancelRequest{}
	mi := &file_calibration_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelRequest) ProtoMessage() {}

func (x *CancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelRequest.ProtoReflect.Descriptor instead.
func (*CancelRequest) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{10}
}

type CancelResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	CalibrationResult *CalibrationResult     `protobuf:"bytes,1,opt,name=calibration_result,json=calibrationResult,proto3" json:"calibration_result,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CancelResponse) Reset() {
	*x = CancelResponse{}
	mi := &file_calibration_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelResponse) ProtoMessage() {}

func (x *CancelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelResponse.ProtoReflect.Descriptor instead.
func (*CancelResponse) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{11}
}

func (x *CancelResponse) GetCalibrationResult() *CalibrationResult {
	if x != nil {
		return x.CalibrationResult
	}
	return nil
}

// Result type.
type CalibrationResult struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Result        CalibrationResult_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mavsdk.rpc.calibration.CalibrationResult_Result" json:"result,omitempty"` // Result enum value
	ResultStr     string                   `protobuf:"bytes,2,opt,name=result_str,json=resultStr,proto3" json:"result_str,omitempty"`                                // Human-readable English string describing the result
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalibrationResult) Reset() {
	*x = CalibrationResult{}
	mi := &file_calibration_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalibrationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalibrationResult) ProtoMessage() {}

func (x *CalibrationResult) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalibrationResult.ProtoReflect.Descriptor instead.
func (*CalibrationResult) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{12}
}

func (x *CalibrationResult) GetResult() CalibrationResult_Result {
	if x != nil {
		return x.Result
	}
	return CalibrationResult_RESULT_UNKNOWN
}

func (x *CalibrationResult) GetResultStr() string {
	if x != nil {
		return x.ResultStr
	}
	return ""
}

// Progress data coming from calibration.
//
// Can be a progress percentage, or an instruction text.
type ProgressData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HasProgress   bool                   `protobuf:"varint,1,opt,name=has_progress,json=hasProgress,proto3" json:"has_progress,omitempty"`         // Whether this ProgressData contains a 'progress' status or not
	Progress      float32                `protobuf:"fixed32,2,opt,name=progress,proto3" json:"progress,omitempty"`                                 // Progress (percentage)
	HasStatusText bool                   `protobuf:"varint,3,opt,name=has_status_text,json=hasStatusText,proto3" json:"has_status_text,omitempty"` // Whether this ProgressData contains a 'status_text' or not
	StatusText    string                 `protobuf:"bytes,4,opt,name=status_text,json=statusText,proto3" json:"status_text,omitempty"`             // Instruction text
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProgressData) Reset() {
	*x = ProgressData{}
	mi := &file_calibration_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProgressData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressData) ProtoMessage() {}

func (x *ProgressData) ProtoReflect() protoreflect.Message {
	mi := &file_calibration_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressData.ProtoReflect.Descriptor instead.
func (*ProgressData) Descriptor() ([]byte, []int) {
	return file_calibration_proto_rawDescGZIP(), []int{13}
}

func (x *ProgressData) GetHasProgress() bool {
	if x != nil {
		return x.HasProgress
	}
	return false
}

func (x *ProgressData) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *ProgressData) GetHasStatusText() bool {
	if x != nil {
		return x.HasStatusText
	}
	return false
}

func (x *ProgressData) GetStatusText() string {
	if x != nil {
		return x.StatusText
	}
	return ""
}

var File_calibration_proto protoreflect.FileDescriptor

const file_calibration_proto_rawDesc = "" +
	"\n" +
	"\x11calibration.proto\x12\x16mavsdk.rpc.calibration\x1a\x14mavsdk_options.proto\"\x1f\n" +
	"\x1dSubscribeCalibrateGyroRequest\"\xbc\x01\n" +
	"\x15CalibrateGyroResponse\x12X\n" +
	"\x12calibration_result\x18\x01 \x01(\v2).mavsdk.rpc.calibration.CalibrationResultR\x11calibrationResult\x12I\n" +
	"\rprogress_data\x18\x02 \x01(\v2$.mavsdk.rpc.calibration.ProgressDataR\fprogressData\"(\n" +
	"&SubscribeCalibrateAccelerometerRequest\"\xc5\x01\n" +
	"\x1eCalibrateAccelerometerResponse\x12X\n" +
	"\x12calibration_result\x18\x01 \x01(\v2).mavsdk.rpc.calibration.CalibrationResultR\x11calibrationResult\x12I\n" +
	"\rprogress_data\x18\x02 \x01(\v2$.mavsdk.rpc.calibration.ProgressDataR\fprogressData\"'\n" +
	"%SubscribeCalibrateMagnetometerRequest\"\xc4\x01\n" +
	"\x1dCalibrateMagnetometerResponse\x12X\n" +
	"\x12calibration_result\x18\x01 \x01(\v2).mavsdk.rpc.calibration.CalibrationResultR\x11calibrationResult\x12I\n" +
	"\rprogress_data\x18\x02 \x01(\v2$.mavsdk.rpc.calibration.ProgressDataR\fprogressData\"'\n" +
	"%SubscribeCalibrateLevelHorizonRequest\"\xc4\x01\n" +
	"\x1dCalibrateLevelHorizonResponse\x12X\n" +
	"\x12calibration_result\x18\x01 \x01(\v2).mavsdk.rpc.calibration.CalibrationResultR\x11calibrationResult\x12I\n" +
	"\rprogress_data\x18\x02 \x01(\v2$.mavsdk.rpc.calibration.ProgressDataR\fprogressData\".\n" +
	",SubscribeCalibrateGimbalAccelerometerRequest\"\xcb\x01\n" +
	"$CalibrateGimbalAccelerometerResponse\x12X\n" +
	"\x12calibration_result\x18\x01 \x01(\v2).mavsdk.rpc.calibration.CalibrationResultR\x11calibrationResult\x12I\n" +
	"\rprogress_data\x18\x02 \x01(\v2$.mavsdk.rpc.calibration.ProgressDataR\fprogressData\"\x0f\n" +
	"\rCancelRequest\"j\n" +
	"\x0eCancelResponse\x12X\n" +
	"\x12calibration_result\x18\x01 \x01(\v2).mavsdk.rpc.calibration.CalibrationResultR\x11calibrationResult\"\x8d\x03\n" +
	"\x11CalibrationResult\x12H\n" +
	"\x06result\x18\x01 \x01(\x0e20.mavsdk.rpc.calibration.CalibrationResult.ResultR\x06result\x12\x1d\n" +
	"\n" +
	"result_str\x18\x02 \x01(\tR\tresultStr\"\x8e\x02\n" +
	"\x06Result\x12\x12\n" +
	"\x0eRESULT_UNKNOWN\x10\x00\x12\x12\n" +
	"\x0eRESULT_SUCCESS\x10\x01\x12\x0f\n" +
	"\vRESULT_NEXT\x10\x02\x12\x11\n" +
	"\rRESULT_FAILED\x10\x03\x12\x14\n" +
	"\x10RESULT_NO_SYSTEM\x10\x04\x12\x1b\n" +
	"\x17RESULT_CONNECTION_ERROR\x10\x05\x12\x0f\n" +
	"\vRESULT_BUSY\x10\x06\x12\x19\n" +
	"\x15RESULT_COMMAND_DENIED\x10\a\x12\x12\n" +
	"\x0eRESULT_TIMEOUT\x10\b\x12\x14\n" +
	"\x10RESULT_CANCELLED\x10\t\x12\x17\n" +
	"\x13RESULT_FAILED_ARMED\x10\n" +
	"\x12\x16\n" +
	"\x12RESULT_UNSUPPORTED\x10\v\"\xb5\x01\n" +
	"\fProgressData\x12,\n" +
	"\fhas_progress\x18\x01 \x01(\bB\t\x82\xb5\x18\x05falseR\vhasProgress\x12#\n" +
	"\bprogress\x18\x02 \x01(\x02B\a\x82\xb5\x18\x03NaNR\bprogress\x121\n" +
	"\x0fhas_status_text\x18\x03 \x01(\bB\t\x82\xb5\x18\x05falseR\rhasStatusText\x12\x1f\n" +
	"\vstatus_text\x18\x04 \x01(\tR\n" +
	"statusText2\xac\a\n" +
	"\x12CalibrationService\x12\x8a\x01\n" +
	"\x16SubscribeCalibrateGyro\x125.mavsdk.rpc.calibration.SubscribeCalibrateGyroRequest\x1a-.mavsdk.rpc.calibration.CalibrateGyroResponse\"\b\x80\xb5\x18\x00\x88\xb5\x18\x010\x01\x12\xa5\x01\n" +
	"\x1fSubscribeCalibrateAccelerometer\x12>.mavsdk.rpc.calibration.SubscribeCalibrateAccelerometerRequest\x1a6.mavsdk.rpc.calibration.CalibrateAccelerometerResponse\"\b\x80\xb5\x18\x00\x88\xb5\x18\x010\x01\x12\xa2\x01\n" +
	"\x1eSubscribeCalibrateMagnetometer\x12=.mavsdk.rpc.calibration.SubscribeCalibrateMagnetometerRequest\x1a5.mavsdk.rpc.calibration.CalibrateMagnetometerResponse\"\b\x80\xb5\x18\x00\x88\xb5\x18\x010\x01\x12\xa2\x01\n" +
	"\x1eSubscribeCalibrateLevelHorizon\x12=.mavsdk.rpc.calibration.SubscribeCalibrateLevelHorizonRequest\x1a5.mavsdk.rpc.calibration.CalibrateLevelHorizonResponse\"\b\x80\xb5\x18\x00\x88\xb5\x18\x010\x01\x12\xb7\x01\n" +
	"%SubscribeCalibrateGimbalAccelerometer\x12D.mavsdk.rpc.calibration.SubscribeCalibrateGimbalAccelerometerRequest\x1a<.mavsdk.rpc.calibration.CalibrateGimbalAccelerometerResponse\"\b\x80\xb5\x18\x00\x88\xb5\x18\x010\x01\x12]\n" +
	"\x06Cancel\x12%.mavsdk.rpc.calibration.CancelRequest\x1a&.mavsdk.rpc.calibration.CancelResponse\"\x04\x80\xb5\x18\x01B!B\x10CalibrationProtoZ\r.;calibrationb\x06proto3"

var (
	file_calibration_proto_rawDescOnce sync.Once
	file_calibration_proto_rawDescData []byte
)

func file_calibration_proto_rawDescGZIP() []byte {
	file_calibration_proto_rawDescOnce.Do(func() {
		file_calibration_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_calibration_proto_rawDesc), len(file_calibration_proto_rawDesc)))
	})
	return file_calibration_proto_rawDescData
}

var file_calibration_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_calibration_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_calibration_proto_goTypes = []any{
	(CalibrationResult_Result)(0),                        // 0: mavsdk.rpc.calibration.CalibrationResult.Result
	(*SubscribeCalibrateGyroRequest)(nil),                // 1: mavsdk.rpc.calibration.SubscribeCalibrateGyroRequest
	(*CalibrateGyroResponse)(nil),                        // 2: mavsdk.rpc.calibration.CalibrateGyroResponse
	(*SubscribeCalibrateAccelerometerRequest)(nil),       // 3: mavsdk.rpc.calibration.SubscribeCalibrateAccelerometerRequest
	(*CalibrateAccelerometerResponse)(nil),               // 4: mavsdk.rpc.calibration.CalibrateAccelerometerResponse
	(*SubscribeCalibrateMagnetometerRequest)(nil),        // 5: mavsdk.rpc.calibration.SubscribeCalibrateMagnetometerRequest
	(*CalibrateMagnetometerResponse)(nil),                // 6: mavsdk.rpc.calibration.CalibrateMagnetometerResponse
	(*SubscribeCalibrateLevelHorizonRequest)(nil),        // 7: mavsdk.rpc.calibration.SubscribeCalibrateLevelHorizonRequest
	(*CalibrateLevelHorizonResponse)(nil),                // 8: mavsdk.rpc.calibration.CalibrateLevelHorizonResponse
	(*SubscribeCalibrateGimbalAccelerometerRequest)(nil), // 9: mavsdk.rpc.calibration.SubscribeCalibrateGimbalAccelerometerRequest
	(*CalibrateGimbalAccelerometerResponse)(nil),         // 10: mavsdk.rpc.calibration.CalibrateGimbalAccelerometerResponse
	(*CancelRequest)(nil),                                // 11: mavsdk.rpc.calibration.CancelRequest
	(*CancelResponse)(nil),                               // 12: mavsdk.rpc.calibration.CancelResponse
	(*CalibrationResult)(nil),                            // 13: mavsdk.rpc.calibration.CalibrationResult
	(*ProgressData)(nil),                                 // 14: mavsdk.rpc.calibration.ProgressData
}
var file_calibration_proto_depIdxs = []int32{
	13, // 0: mavsdk.rpc.calibration.CalibrateGyroResponse.calibration_result:type_name -> mavsdk.rpc.calibration.CalibrationResult
	14, // 1: mavsdk.rpc.calibration.CalibrateGyroResponse.progress_data:type_name -> mavsdk.rpc.calibration.ProgressData
	13, // 2: mavsdk.rpc.calibration.CalibrateAccelerometerResponse.calibration_result:type_name -> mavsdk.rpc.calibration.CalibrationResult
	14, // 3: mavsdk.rpc.calibration.CalibrateAccelerometerResponse.progress_data:type_name -> mavsdk.rpc.calibration.ProgressData
	13, // 4: mavsdk.rpc.calibration.CalibrateMagnetometerResponse.calibration_result:type_name -> mavsdk.rpc.calibration.CalibrationResult
	14, // 5: mavsdk.rpc.calibration.CalibrateMagnetometerResponse.progress_data:type_name -> mavsdk.rpc.calibration.ProgressData
	13, // 6: mavsdk.rpc.calibration.CalibrateLevelHorizonResponse.calibration_result:type_name -> mavsdk.rpc.calibration.CalibrationResult
	14, // 7: mavsdk.rpc.calibration.CalibrateLevelHorizonResponse.progress_data:type_name -> mavsdk.rpc.calibration.ProgressData
	13, // 8: mavsdk.rpc.calibration.CalibrateGimbalAccelerometerResponse.calibration_result:type_name -> mavsdk.rpc.calibration.CalibrationResult
	14, // 9: mavsdk.rpc.calibration.CalibrateGimbalAccelerometerResponse.progress_data:type_name -> mavsdk.rpc.calibration.ProgressData
	13, // 10: mavsdk.rpc.calibration.CancelResponse.calibration_result:type_name -> mavsdk.rpc.calibration.CalibrationResult
	0,  // 11: mavsdk.rpc.calibration.CalibrationResult.result:type_name -> mavsdk.rpc.calibration.CalibrationResult.Result
	1,  // 12: mavsdk.rpc.calibration.CalibrationService.SubscribeCalibrateGyro:input_type -> mavsdk.rpc.calibration.SubscribeCalibrateGyroRequest
	3,  // 13: mavsdk.rpc.calibration.CalibrationService.SubscribeCalibrateAccelerometer:input_type -> mavsdk.rpc.calibration.SubscribeCalibrateAccelerometerRequest
	5,  // 14: mavsdk.rpc.calibration.CalibrationService.SubscribeCalibrateMagnetometer:input_type -> mavsdk.rpc.calibration.SubscribeCalibrateMagnetometerRequest
	7,  // 15: mavsdk.rpc.calibration.CalibrationService.SubscribeCalibrateLevelHorizon:input_type -> mavsdk.rpc.calibration.SubscribeCalibrateLevelHorizonRequest
	9,  // 16: mavsdk.rpc.calibration.CalibrationService.SubscribeCalibrateGimbalAccelerometer:input_type -> mavsdk.rpc.calibration.SubscribeCalibrateGimbalAccelerometerRequest
	11, // 17: mavsdk.rpc.calibration.CalibrationService.Cancel:input_type -> mavsdk.rpc.calibration.CancelRequest
	2,  // 18: mavsdk.rpc.calibration.CalibrationService.SubscribeCalibrateGyro:output_type -> mavsdk.rpc.calibration.CalibrateGyroResponse
	4,  // 19: mavsdk.rpc.calibration.CalibrationService.SubscribeCalibrateAccelerometer:output_type -> mavsdk.rpc.calibration.CalibrateAccelerometerResponse
	6,  // 20: mavsdk.rpc.calibration.CalibrationService.SubscribeCalibrateMagnetometer:output_type -> mavsdk.rpc.calibration.CalibrateMagnetometerResponse
	8,  // 21: mavsdk.rpc.calibration.CalibrationService.SubscribeCalibrateLevelHorizon:output_type -> mavsdk.rpc.calibration.CalibrateLevelHorizonResponse
	10, // 22: mavsdk.rpc.calibration.CalibrationService.SubscribeCalibrateGimbalAccelerometer:output_type -> mavsdk.rpc.calibration.CalibrateGimbalAccelerometerResponse
	12, // 23: mavsdk.rpc.calibration.CalibrationService.Cancel:output_type -> mavsdk.rpc.calibration.CancelResponse
	18, // [18:24] is the sub-list for method output_type
	12, // [12:18] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_calibration_proto_init() }
func file_calibration_proto_init() {
	if File_calibration_proto != nil {
		return
	}
	
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_calibration_proto_rawDesc), len(file_calibration_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calibration_proto_goTypes,
		DependencyIndexes: file_calibration_proto_depIdxs,
		EnumInfos:         file_calibration_proto_enumTypes,
		MessageInfos:      file_calibration_proto_msgTypes,
	}.Build()
	File_calibration_proto = out.File
	file_calibration_proto_goTypes = nil
	file_calibration_proto_depIdxs = nil
}
