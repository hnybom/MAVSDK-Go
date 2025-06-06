// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: geofence.proto

package geofence

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

// Geofence types.
type FenceType int32

const (
	FenceType_FENCE_TYPE_INCLUSION FenceType = 0 // Type representing an inclusion fence
	FenceType_FENCE_TYPE_EXCLUSION FenceType = 1 // Type representing an exclusion fence
)

// Enum value maps for FenceType.
var (
	FenceType_name = map[int32]string{
		0: "FENCE_TYPE_INCLUSION",
		1: "FENCE_TYPE_EXCLUSION",
	}
	FenceType_value = map[string]int32{
		"FENCE_TYPE_INCLUSION": 0,
		"FENCE_TYPE_EXCLUSION": 1,
	}
)

func (x FenceType) Enum() *FenceType {
	p := new(FenceType)
	*p = x
	return p
}

func (x FenceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FenceType) Descriptor() protoreflect.EnumDescriptor {
	return file_geofence_proto_enumTypes[0].Descriptor()
}

func (FenceType) Type() protoreflect.EnumType {
	return &file_geofence_proto_enumTypes[0]
}

func (x FenceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FenceType.Descriptor instead.
func (FenceType) EnumDescriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{0}
}

// Possible results returned for geofence requests.
type GeofenceResult_Result int32

const (
	GeofenceResult_RESULT_UNKNOWN                 GeofenceResult_Result = 0 // Unknown result
	GeofenceResult_RESULT_SUCCESS                 GeofenceResult_Result = 1 // Request succeeded
	GeofenceResult_RESULT_ERROR                   GeofenceResult_Result = 2 // Error
	GeofenceResult_RESULT_TOO_MANY_GEOFENCE_ITEMS GeofenceResult_Result = 3 // Too many objects in the geofence
	GeofenceResult_RESULT_BUSY                    GeofenceResult_Result = 4 // Vehicle is busy
	GeofenceResult_RESULT_TIMEOUT                 GeofenceResult_Result = 5 // Request timed out
	GeofenceResult_RESULT_INVALID_ARGUMENT        GeofenceResult_Result = 6 // Invalid argument
	GeofenceResult_RESULT_NO_SYSTEM               GeofenceResult_Result = 7 // No system connected
)

// Enum value maps for GeofenceResult_Result.
var (
	GeofenceResult_Result_name = map[int32]string{
		0: "RESULT_UNKNOWN",
		1: "RESULT_SUCCESS",
		2: "RESULT_ERROR",
		3: "RESULT_TOO_MANY_GEOFENCE_ITEMS",
		4: "RESULT_BUSY",
		5: "RESULT_TIMEOUT",
		6: "RESULT_INVALID_ARGUMENT",
		7: "RESULT_NO_SYSTEM",
	}
	GeofenceResult_Result_value = map[string]int32{
		"RESULT_UNKNOWN":                 0,
		"RESULT_SUCCESS":                 1,
		"RESULT_ERROR":                   2,
		"RESULT_TOO_MANY_GEOFENCE_ITEMS": 3,
		"RESULT_BUSY":                    4,
		"RESULT_TIMEOUT":                 5,
		"RESULT_INVALID_ARGUMENT":        6,
		"RESULT_NO_SYSTEM":               7,
	}
)

func (x GeofenceResult_Result) Enum() *GeofenceResult_Result {
	p := new(GeofenceResult_Result)
	*p = x
	return p
}

func (x GeofenceResult_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GeofenceResult_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_geofence_proto_enumTypes[1].Descriptor()
}

func (GeofenceResult_Result) Type() protoreflect.EnumType {
	return &file_geofence_proto_enumTypes[1]
}

func (x GeofenceResult_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GeofenceResult_Result.Descriptor instead.
func (GeofenceResult_Result) EnumDescriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{8, 0}
}

// Point type.
type Point struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LatitudeDeg   float64                `protobuf:"fixed64,1,opt,name=latitude_deg,json=latitudeDeg,proto3" json:"latitude_deg,omitempty"`    // Latitude in degrees (range: -90 to +90)
	LongitudeDeg  float64                `protobuf:"fixed64,2,opt,name=longitude_deg,json=longitudeDeg,proto3" json:"longitude_deg,omitempty"` // Longitude in degrees (range: -180 to +180)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Point) Reset() {
	*x = Point{}
	mi := &file_geofence_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetLatitudeDeg() float64 {
	if x != nil {
		return x.LatitudeDeg
	}
	return 0
}

func (x *Point) GetLongitudeDeg() float64 {
	if x != nil {
		return x.LongitudeDeg
	}
	return 0
}

// Polygon type.
type Polygon struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Points        []*Point               `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`                                                            // Points defining the polygon
	FenceType     FenceType              `protobuf:"varint,2,opt,name=fence_type,json=fenceType,proto3,enum=mavsdk.rpc.geofence.FenceType" json:"fence_type,omitempty"` // Fence type
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Polygon) Reset() {
	*x = Polygon{}
	mi := &file_geofence_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Polygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Polygon) ProtoMessage() {}

func (x *Polygon) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Polygon.ProtoReflect.Descriptor instead.
func (*Polygon) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{1}
}

func (x *Polygon) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

func (x *Polygon) GetFenceType() FenceType {
	if x != nil {
		return x.FenceType
	}
	return FenceType_FENCE_TYPE_INCLUSION
}

// Circular type.
type Circle struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Point         *Point                 `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`                                                              // Point defining the center
	Radius        float32                `protobuf:"fixed32,2,opt,name=radius,proto3" json:"radius,omitempty"`                                                          // Radius of the circular fence
	FenceType     FenceType              `protobuf:"varint,3,opt,name=fence_type,json=fenceType,proto3,enum=mavsdk.rpc.geofence.FenceType" json:"fence_type,omitempty"` // Fence type
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Circle) Reset() {
	*x = Circle{}
	mi := &file_geofence_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Circle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Circle) ProtoMessage() {}

func (x *Circle) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Circle.ProtoReflect.Descriptor instead.
func (*Circle) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{2}
}

func (x *Circle) GetPoint() *Point {
	if x != nil {
		return x.Point
	}
	return nil
}

func (x *Circle) GetRadius() float32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *Circle) GetFenceType() FenceType {
	if x != nil {
		return x.FenceType
	}
	return FenceType_FENCE_TYPE_INCLUSION
}

// Geofence data type.
type GeofenceData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Polygons      []*Polygon             `protobuf:"bytes,1,rep,name=polygons,proto3" json:"polygons,omitempty"` // Polygon(s) representing the geofence(s)
	Circles       []*Circle              `protobuf:"bytes,2,rep,name=circles,proto3" json:"circles,omitempty"`   // Circle(s) representing the geofence(s)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeofenceData) Reset() {
	*x = GeofenceData{}
	mi := &file_geofence_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeofenceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeofenceData) ProtoMessage() {}

func (x *GeofenceData) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeofenceData.ProtoReflect.Descriptor instead.
func (*GeofenceData) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{3}
}

func (x *GeofenceData) GetPolygons() []*Polygon {
	if x != nil {
		return x.Polygons
	}
	return nil
}

func (x *GeofenceData) GetCircles() []*Circle {
	if x != nil {
		return x.Circles
	}
	return nil
}

type UploadGeofenceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GeofenceData  *GeofenceData          `protobuf:"bytes,1,opt,name=geofence_data,json=geofenceData,proto3" json:"geofence_data,omitempty"` // Circle(s) and/or Polygon(s) representing the geofence(s)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadGeofenceRequest) Reset() {
	*x = UploadGeofenceRequest{}
	mi := &file_geofence_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadGeofenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadGeofenceRequest) ProtoMessage() {}

func (x *UploadGeofenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadGeofenceRequest.ProtoReflect.Descriptor instead.
func (*UploadGeofenceRequest) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{4}
}

func (x *UploadGeofenceRequest) GetGeofenceData() *GeofenceData {
	if x != nil {
		return x.GeofenceData
	}
	return nil
}

type UploadGeofenceResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	GeofenceResult *GeofenceResult        `protobuf:"bytes,1,opt,name=geofence_result,json=geofenceResult,proto3" json:"geofence_result,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UploadGeofenceResponse) Reset() {
	*x = UploadGeofenceResponse{}
	mi := &file_geofence_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadGeofenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadGeofenceResponse) ProtoMessage() {}

func (x *UploadGeofenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadGeofenceResponse.ProtoReflect.Descriptor instead.
func (*UploadGeofenceResponse) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{5}
}

func (x *UploadGeofenceResponse) GetGeofenceResult() *GeofenceResult {
	if x != nil {
		return x.GeofenceResult
	}
	return nil
}

type ClearGeofenceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClearGeofenceRequest) Reset() {
	*x = ClearGeofenceRequest{}
	mi := &file_geofence_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClearGeofenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearGeofenceRequest) ProtoMessage() {}

func (x *ClearGeofenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearGeofenceRequest.ProtoReflect.Descriptor instead.
func (*ClearGeofenceRequest) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{6}
}

type ClearGeofenceResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	GeofenceResult *GeofenceResult        `protobuf:"bytes,1,opt,name=geofence_result,json=geofenceResult,proto3" json:"geofence_result,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ClearGeofenceResponse) Reset() {
	*x = ClearGeofenceResponse{}
	mi := &file_geofence_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClearGeofenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearGeofenceResponse) ProtoMessage() {}

func (x *ClearGeofenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearGeofenceResponse.ProtoReflect.Descriptor instead.
func (*ClearGeofenceResponse) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{7}
}

func (x *ClearGeofenceResponse) GetGeofenceResult() *GeofenceResult {
	if x != nil {
		return x.GeofenceResult
	}
	return nil
}

// Result type.
type GeofenceResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        GeofenceResult_Result  `protobuf:"varint,1,opt,name=result,proto3,enum=mavsdk.rpc.geofence.GeofenceResult_Result" json:"result,omitempty"` // Result enum value
	ResultStr     string                 `protobuf:"bytes,2,opt,name=result_str,json=resultStr,proto3" json:"result_str,omitempty"`                          // Human-readable English string describing the result
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeofenceResult) Reset() {
	*x = GeofenceResult{}
	mi := &file_geofence_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeofenceResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeofenceResult) ProtoMessage() {}

func (x *GeofenceResult) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeofenceResult.ProtoReflect.Descriptor instead.
func (*GeofenceResult) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{8}
}

func (x *GeofenceResult) GetResult() GeofenceResult_Result {
	if x != nil {
		return x.Result
	}
	return GeofenceResult_RESULT_UNKNOWN
}

func (x *GeofenceResult) GetResultStr() string {
	if x != nil {
		return x.ResultStr
	}
	return ""
}

var File_geofence_proto protoreflect.FileDescriptor

const file_geofence_proto_rawDesc = "" +
	"\n" +
	"\x0egeofence.proto\x12\x13mavsdk.rpc.geofence\x1a\x14mavsdk_options.proto\"O\n" +
	"\x05Point\x12!\n" +
	"\flatitude_deg\x18\x01 \x01(\x01R\vlatitudeDeg\x12#\n" +
	"\rlongitude_deg\x18\x02 \x01(\x01R\flongitudeDeg\"|\n" +
	"\aPolygon\x122\n" +
	"\x06points\x18\x01 \x03(\v2\x1a.mavsdk.rpc.geofence.PointR\x06points\x12=\n" +
	"\n" +
	"fence_type\x18\x02 \x01(\x0e2\x1e.mavsdk.rpc.geofence.FenceTypeR\tfenceType\"\x9a\x01\n" +
	"\x06Circle\x120\n" +
	"\x05point\x18\x01 \x01(\v2\x1a.mavsdk.rpc.geofence.PointR\x05point\x12\x1f\n" +
	"\x06radius\x18\x02 \x01(\x02B\a\x82\xb5\x18\x03NaNR\x06radius\x12=\n" +
	"\n" +
	"fence_type\x18\x03 \x01(\x0e2\x1e.mavsdk.rpc.geofence.FenceTypeR\tfenceType\"\x7f\n" +
	"\fGeofenceData\x128\n" +
	"\bpolygons\x18\x01 \x03(\v2\x1c.mavsdk.rpc.geofence.PolygonR\bpolygons\x125\n" +
	"\acircles\x18\x02 \x03(\v2\x1b.mavsdk.rpc.geofence.CircleR\acircles\"_\n" +
	"\x15UploadGeofenceRequest\x12F\n" +
	"\rgeofence_data\x18\x01 \x01(\v2!.mavsdk.rpc.geofence.GeofenceDataR\fgeofenceData\"f\n" +
	"\x16UploadGeofenceResponse\x12L\n" +
	"\x0fgeofence_result\x18\x01 \x01(\v2#.mavsdk.rpc.geofence.GeofenceResultR\x0egeofenceResult\"\x16\n" +
	"\x14ClearGeofenceRequest\"e\n" +
	"\x15ClearGeofenceResponse\x12L\n" +
	"\x0fgeofence_result\x18\x01 \x01(\v2#.mavsdk.rpc.geofence.GeofenceResultR\x0egeofenceResult\"\xb4\x02\n" +
	"\x0eGeofenceResult\x12B\n" +
	"\x06result\x18\x01 \x01(\x0e2*.mavsdk.rpc.geofence.GeofenceResult.ResultR\x06result\x12\x1d\n" +
	"\n" +
	"result_str\x18\x02 \x01(\tR\tresultStr\"\xbe\x01\n" +
	"\x06Result\x12\x12\n" +
	"\x0eRESULT_UNKNOWN\x10\x00\x12\x12\n" +
	"\x0eRESULT_SUCCESS\x10\x01\x12\x10\n" +
	"\fRESULT_ERROR\x10\x02\x12\"\n" +
	"\x1eRESULT_TOO_MANY_GEOFENCE_ITEMS\x10\x03\x12\x0f\n" +
	"\vRESULT_BUSY\x10\x04\x12\x12\n" +
	"\x0eRESULT_TIMEOUT\x10\x05\x12\x1b\n" +
	"\x17RESULT_INVALID_ARGUMENT\x10\x06\x12\x14\n" +
	"\x10RESULT_NO_SYSTEM\x10\a*?\n" +
	"\tFenceType\x12\x18\n" +
	"\x14FENCE_TYPE_INCLUSION\x10\x00\x12\x18\n" +
	"\x14FENCE_TYPE_EXCLUSION\x10\x012\xe8\x01\n" +
	"\x0fGeofenceService\x12k\n" +
	"\x0eUploadGeofence\x12*.mavsdk.rpc.geofence.UploadGeofenceRequest\x1a+.mavsdk.rpc.geofence.UploadGeofenceResponse\"\x00\x12h\n" +
	"\rClearGeofence\x12).mavsdk.rpc.geofence.ClearGeofenceRequest\x1a*.mavsdk.rpc.geofence.ClearGeofenceResponse\"\x00B\x1bB\rGeofenceProtoZ\n" +
	".;geofenceb\x06proto3"

var (
	file_geofence_proto_rawDescOnce sync.Once
	file_geofence_proto_rawDescData []byte
)

func file_geofence_proto_rawDescGZIP() []byte {
	file_geofence_proto_rawDescOnce.Do(func() {
		file_geofence_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_geofence_proto_rawDesc), len(file_geofence_proto_rawDesc)))
	})
	return file_geofence_proto_rawDescData
}

var file_geofence_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_geofence_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_geofence_proto_goTypes = []any{
	(FenceType)(0),                 // 0: mavsdk.rpc.geofence.FenceType
	(GeofenceResult_Result)(0),     // 1: mavsdk.rpc.geofence.GeofenceResult.Result
	(*Point)(nil),                  // 2: mavsdk.rpc.geofence.Point
	(*Polygon)(nil),                // 3: mavsdk.rpc.geofence.Polygon
	(*Circle)(nil),                 // 4: mavsdk.rpc.geofence.Circle
	(*GeofenceData)(nil),           // 5: mavsdk.rpc.geofence.GeofenceData
	(*UploadGeofenceRequest)(nil),  // 6: mavsdk.rpc.geofence.UploadGeofenceRequest
	(*UploadGeofenceResponse)(nil), // 7: mavsdk.rpc.geofence.UploadGeofenceResponse
	(*ClearGeofenceRequest)(nil),   // 8: mavsdk.rpc.geofence.ClearGeofenceRequest
	(*ClearGeofenceResponse)(nil),  // 9: mavsdk.rpc.geofence.ClearGeofenceResponse
	(*GeofenceResult)(nil),         // 10: mavsdk.rpc.geofence.GeofenceResult
}
var file_geofence_proto_depIdxs = []int32{
	2,  // 0: mavsdk.rpc.geofence.Polygon.points:type_name -> mavsdk.rpc.geofence.Point
	0,  // 1: mavsdk.rpc.geofence.Polygon.fence_type:type_name -> mavsdk.rpc.geofence.FenceType
	2,  // 2: mavsdk.rpc.geofence.Circle.point:type_name -> mavsdk.rpc.geofence.Point
	0,  // 3: mavsdk.rpc.geofence.Circle.fence_type:type_name -> mavsdk.rpc.geofence.FenceType
	3,  // 4: mavsdk.rpc.geofence.GeofenceData.polygons:type_name -> mavsdk.rpc.geofence.Polygon
	4,  // 5: mavsdk.rpc.geofence.GeofenceData.circles:type_name -> mavsdk.rpc.geofence.Circle
	5,  // 6: mavsdk.rpc.geofence.UploadGeofenceRequest.geofence_data:type_name -> mavsdk.rpc.geofence.GeofenceData
	10, // 7: mavsdk.rpc.geofence.UploadGeofenceResponse.geofence_result:type_name -> mavsdk.rpc.geofence.GeofenceResult
	10, // 8: mavsdk.rpc.geofence.ClearGeofenceResponse.geofence_result:type_name -> mavsdk.rpc.geofence.GeofenceResult
	1,  // 9: mavsdk.rpc.geofence.GeofenceResult.result:type_name -> mavsdk.rpc.geofence.GeofenceResult.Result
	6,  // 10: mavsdk.rpc.geofence.GeofenceService.UploadGeofence:input_type -> mavsdk.rpc.geofence.UploadGeofenceRequest
	8,  // 11: mavsdk.rpc.geofence.GeofenceService.ClearGeofence:input_type -> mavsdk.rpc.geofence.ClearGeofenceRequest
	7,  // 12: mavsdk.rpc.geofence.GeofenceService.UploadGeofence:output_type -> mavsdk.rpc.geofence.UploadGeofenceResponse
	9,  // 13: mavsdk.rpc.geofence.GeofenceService.ClearGeofence:output_type -> mavsdk.rpc.geofence.ClearGeofenceResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_geofence_proto_init() }
func file_geofence_proto_init() {
	if File_geofence_proto != nil {
		return
	}
	
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_geofence_proto_rawDesc), len(file_geofence_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_geofence_proto_goTypes,
		DependencyIndexes: file_geofence_proto_depIdxs,
		EnumInfos:         file_geofence_proto_enumTypes,
		MessageInfos:      file_geofence_proto_msgTypes,
	}.Build()
	File_geofence_proto = out.File
	file_geofence_proto_goTypes = nil
	file_geofence_proto_depIdxs = nil
}
