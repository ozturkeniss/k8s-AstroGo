// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0--rc2
// source: grpc/protos/astronaut.proto

package protos

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

// Astronaut status
type AstronautStatus int32

const (
	AstronautStatus_ASTRONAUT_STATUS_AVAILABLE  AstronautStatus = 0
	AstronautStatus_ASTRONAUT_STATUS_ON_MISSION AstronautStatus = 1
	AstronautStatus_ASTRONAUT_STATUS_RESTING    AstronautStatus = 2
	AstronautStatus_ASTRONAUT_STATUS_TRAINING   AstronautStatus = 3
)

// Enum value maps for AstronautStatus.
var (
	AstronautStatus_name = map[int32]string{
		0: "ASTRONAUT_STATUS_AVAILABLE",
		1: "ASTRONAUT_STATUS_ON_MISSION",
		2: "ASTRONAUT_STATUS_RESTING",
		3: "ASTRONAUT_STATUS_TRAINING",
	}
	AstronautStatus_value = map[string]int32{
		"ASTRONAUT_STATUS_AVAILABLE":  0,
		"ASTRONAUT_STATUS_ON_MISSION": 1,
		"ASTRONAUT_STATUS_RESTING":    2,
		"ASTRONAUT_STATUS_TRAINING":   3,
	}
)

func (x AstronautStatus) Enum() *AstronautStatus {
	p := new(AstronautStatus)
	*p = x
	return p
}

func (x AstronautStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AstronautStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_protos_astronaut_proto_enumTypes[0].Descriptor()
}

func (AstronautStatus) Type() protoreflect.EnumType {
	return &file_grpc_protos_astronaut_proto_enumTypes[0]
}

func (x AstronautStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AstronautStatus.Descriptor instead.
func (AstronautStatus) EnumDescriptor() ([]byte, []int) {
	return file_grpc_protos_astronaut_proto_rawDescGZIP(), []int{0}
}

// Astronaut message
type Astronaut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status        AstronautStatus        `protobuf:"varint,3,opt,name=status,proto3,enum=astronaut.AstronautStatus" json:"status,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Astronaut) Reset() {
	*x = Astronaut{}
	mi := &file_grpc_protos_astronaut_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Astronaut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Astronaut) ProtoMessage() {}

func (x *Astronaut) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_protos_astronaut_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Astronaut.ProtoReflect.Descriptor instead.
func (*Astronaut) Descriptor() ([]byte, []int) {
	return file_grpc_protos_astronaut_proto_rawDescGZIP(), []int{0}
}

func (x *Astronaut) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Astronaut) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Astronaut) GetStatus() AstronautStatus {
	if x != nil {
		return x.Status
	}
	return AstronautStatus_ASTRONAUT_STATUS_AVAILABLE
}

func (x *Astronaut) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Astronaut) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// Create astronaut request
type CreateAstronautRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAstronautRequest) Reset() {
	*x = CreateAstronautRequest{}
	mi := &file_grpc_protos_astronaut_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAstronautRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAstronautRequest) ProtoMessage() {}

func (x *CreateAstronautRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_protos_astronaut_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAstronautRequest.ProtoReflect.Descriptor instead.
func (*CreateAstronautRequest) Descriptor() ([]byte, []int) {
	return file_grpc_protos_astronaut_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAstronautRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Create astronaut response
type CreateAstronautResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Astronaut     *Astronaut             `protobuf:"bytes,1,opt,name=astronaut,proto3" json:"astronaut,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAstronautResponse) Reset() {
	*x = CreateAstronautResponse{}
	mi := &file_grpc_protos_astronaut_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAstronautResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAstronautResponse) ProtoMessage() {}

func (x *CreateAstronautResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_protos_astronaut_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAstronautResponse.ProtoReflect.Descriptor instead.
func (*CreateAstronautResponse) Descriptor() ([]byte, []int) {
	return file_grpc_protos_astronaut_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAstronautResponse) GetAstronaut() *Astronaut {
	if x != nil {
		return x.Astronaut
	}
	return nil
}

// Get astronaut request
type GetAstronautRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AstronautId   uint64                 `protobuf:"varint,1,opt,name=astronaut_id,json=astronautId,proto3" json:"astronaut_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAstronautRequest) Reset() {
	*x = GetAstronautRequest{}
	mi := &file_grpc_protos_astronaut_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAstronautRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAstronautRequest) ProtoMessage() {}

func (x *GetAstronautRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_protos_astronaut_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAstronautRequest.ProtoReflect.Descriptor instead.
func (*GetAstronautRequest) Descriptor() ([]byte, []int) {
	return file_grpc_protos_astronaut_proto_rawDescGZIP(), []int{3}
}

func (x *GetAstronautRequest) GetAstronautId() uint64 {
	if x != nil {
		return x.AstronautId
	}
	return 0
}

// Get astronaut response
type GetAstronautResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Astronaut     *Astronaut             `protobuf:"bytes,1,opt,name=astronaut,proto3" json:"astronaut,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAstronautResponse) Reset() {
	*x = GetAstronautResponse{}
	mi := &file_grpc_protos_astronaut_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAstronautResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAstronautResponse) ProtoMessage() {}

func (x *GetAstronautResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_protos_astronaut_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAstronautResponse.ProtoReflect.Descriptor instead.
func (*GetAstronautResponse) Descriptor() ([]byte, []int) {
	return file_grpc_protos_astronaut_proto_rawDescGZIP(), []int{4}
}

func (x *GetAstronautResponse) GetAstronaut() *Astronaut {
	if x != nil {
		return x.Astronaut
	}
	return nil
}

// List astronauts request
type ListAstronautsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAstronautsRequest) Reset() {
	*x = ListAstronautsRequest{}
	mi := &file_grpc_protos_astronaut_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAstronautsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAstronautsRequest) ProtoMessage() {}

func (x *ListAstronautsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_protos_astronaut_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAstronautsRequest.ProtoReflect.Descriptor instead.
func (*ListAstronautsRequest) Descriptor() ([]byte, []int) {
	return file_grpc_protos_astronaut_proto_rawDescGZIP(), []int{5}
}

// List astronauts response
type ListAstronautsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Astronauts    []*Astronaut           `protobuf:"bytes,1,rep,name=astronauts,proto3" json:"astronauts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAstronautsResponse) Reset() {
	*x = ListAstronautsResponse{}
	mi := &file_grpc_protos_astronaut_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAstronautsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAstronautsResponse) ProtoMessage() {}

func (x *ListAstronautsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_protos_astronaut_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAstronautsResponse.ProtoReflect.Descriptor instead.
func (*ListAstronautsResponse) Descriptor() ([]byte, []int) {
	return file_grpc_protos_astronaut_proto_rawDescGZIP(), []int{6}
}

func (x *ListAstronautsResponse) GetAstronauts() []*Astronaut {
	if x != nil {
		return x.Astronauts
	}
	return nil
}

// Update astronaut status request
type UpdateAstronautStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AstronautId   uint64                 `protobuf:"varint,1,opt,name=astronaut_id,json=astronautId,proto3" json:"astronaut_id,omitempty"`
	Status        AstronautStatus        `protobuf:"varint,2,opt,name=status,proto3,enum=astronaut.AstronautStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAstronautStatusRequest) Reset() {
	*x = UpdateAstronautStatusRequest{}
	mi := &file_grpc_protos_astronaut_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAstronautStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAstronautStatusRequest) ProtoMessage() {}

func (x *UpdateAstronautStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_protos_astronaut_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAstronautStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateAstronautStatusRequest) Descriptor() ([]byte, []int) {
	return file_grpc_protos_astronaut_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateAstronautStatusRequest) GetAstronautId() uint64 {
	if x != nil {
		return x.AstronautId
	}
	return 0
}

func (x *UpdateAstronautStatusRequest) GetStatus() AstronautStatus {
	if x != nil {
		return x.Status
	}
	return AstronautStatus_ASTRONAUT_STATUS_AVAILABLE
}

// Update astronaut status response
type UpdateAstronautStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Astronaut     *Astronaut             `protobuf:"bytes,1,opt,name=astronaut,proto3" json:"astronaut,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAstronautStatusResponse) Reset() {
	*x = UpdateAstronautStatusResponse{}
	mi := &file_grpc_protos_astronaut_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAstronautStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAstronautStatusResponse) ProtoMessage() {}

func (x *UpdateAstronautStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_protos_astronaut_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAstronautStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateAstronautStatusResponse) Descriptor() ([]byte, []int) {
	return file_grpc_protos_astronaut_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateAstronautStatusResponse) GetAstronaut() *Astronaut {
	if x != nil {
		return x.Astronaut
	}
	return nil
}

var File_grpc_protos_astronaut_proto protoreflect.FileDescriptor

const file_grpc_protos_astronaut_proto_rawDesc = "" +
	"\n" +
	"\x1bgrpc/protos/astronaut.proto\x12\tastronaut\"\xa1\x01\n" +
	"\tAstronaut\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x122\n" +
	"\x06status\x18\x03 \x01(\x0e2\x1a.astronaut.AstronautStatusR\x06status\x12\x1d\n" +
	"\n" +
	"created_at\x18\x04 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x05 \x01(\tR\tupdatedAt\",\n" +
	"\x16CreateAstronautRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"M\n" +
	"\x17CreateAstronautResponse\x122\n" +
	"\tastronaut\x18\x01 \x01(\v2\x14.astronaut.AstronautR\tastronaut\"8\n" +
	"\x13GetAstronautRequest\x12!\n" +
	"\fastronaut_id\x18\x01 \x01(\x04R\vastronautId\"J\n" +
	"\x14GetAstronautResponse\x122\n" +
	"\tastronaut\x18\x01 \x01(\v2\x14.astronaut.AstronautR\tastronaut\"\x17\n" +
	"\x15ListAstronautsRequest\"N\n" +
	"\x16ListAstronautsResponse\x124\n" +
	"\n" +
	"astronauts\x18\x01 \x03(\v2\x14.astronaut.AstronautR\n" +
	"astronauts\"u\n" +
	"\x1cUpdateAstronautStatusRequest\x12!\n" +
	"\fastronaut_id\x18\x01 \x01(\x04R\vastronautId\x122\n" +
	"\x06status\x18\x02 \x01(\x0e2\x1a.astronaut.AstronautStatusR\x06status\"S\n" +
	"\x1dUpdateAstronautStatusResponse\x122\n" +
	"\tastronaut\x18\x01 \x01(\v2\x14.astronaut.AstronautR\tastronaut*\x8f\x01\n" +
	"\x0fAstronautStatus\x12\x1e\n" +
	"\x1aASTRONAUT_STATUS_AVAILABLE\x10\x00\x12\x1f\n" +
	"\x1bASTRONAUT_STATUS_ON_MISSION\x10\x01\x12\x1c\n" +
	"\x18ASTRONAUT_STATUS_RESTING\x10\x02\x12\x1d\n" +
	"\x19ASTRONAUT_STATUS_TRAINING\x10\x032\x80\x03\n" +
	"\x10AstronautService\x12X\n" +
	"\x0fCreateAstronaut\x12!.astronaut.CreateAstronautRequest\x1a\".astronaut.CreateAstronautResponse\x12O\n" +
	"\fGetAstronaut\x12\x1e.astronaut.GetAstronautRequest\x1a\x1f.astronaut.GetAstronautResponse\x12U\n" +
	"\x0eListAstronauts\x12 .astronaut.ListAstronautsRequest\x1a!.astronaut.ListAstronautsResponse\x12j\n" +
	"\x15UpdateAstronautStatus\x12'.astronaut.UpdateAstronautStatusRequest\x1a(.astronaut.UpdateAstronautStatusResponseB$Z\"astrogo/grpc/astronaut/grpc/protosb\x06proto3"

var (
	file_grpc_protos_astronaut_proto_rawDescOnce sync.Once
	file_grpc_protos_astronaut_proto_rawDescData []byte
)

func file_grpc_protos_astronaut_proto_rawDescGZIP() []byte {
	file_grpc_protos_astronaut_proto_rawDescOnce.Do(func() {
		file_grpc_protos_astronaut_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_grpc_protos_astronaut_proto_rawDesc), len(file_grpc_protos_astronaut_proto_rawDesc)))
	})
	return file_grpc_protos_astronaut_proto_rawDescData
}

var file_grpc_protos_astronaut_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_protos_astronaut_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_grpc_protos_astronaut_proto_goTypes = []any{
	(AstronautStatus)(0),                  // 0: astronaut.AstronautStatus
	(*Astronaut)(nil),                     // 1: astronaut.Astronaut
	(*CreateAstronautRequest)(nil),        // 2: astronaut.CreateAstronautRequest
	(*CreateAstronautResponse)(nil),       // 3: astronaut.CreateAstronautResponse
	(*GetAstronautRequest)(nil),           // 4: astronaut.GetAstronautRequest
	(*GetAstronautResponse)(nil),          // 5: astronaut.GetAstronautResponse
	(*ListAstronautsRequest)(nil),         // 6: astronaut.ListAstronautsRequest
	(*ListAstronautsResponse)(nil),        // 7: astronaut.ListAstronautsResponse
	(*UpdateAstronautStatusRequest)(nil),  // 8: astronaut.UpdateAstronautStatusRequest
	(*UpdateAstronautStatusResponse)(nil), // 9: astronaut.UpdateAstronautStatusResponse
}
var file_grpc_protos_astronaut_proto_depIdxs = []int32{
	0,  // 0: astronaut.Astronaut.status:type_name -> astronaut.AstronautStatus
	1,  // 1: astronaut.CreateAstronautResponse.astronaut:type_name -> astronaut.Astronaut
	1,  // 2: astronaut.GetAstronautResponse.astronaut:type_name -> astronaut.Astronaut
	1,  // 3: astronaut.ListAstronautsResponse.astronauts:type_name -> astronaut.Astronaut
	0,  // 4: astronaut.UpdateAstronautStatusRequest.status:type_name -> astronaut.AstronautStatus
	1,  // 5: astronaut.UpdateAstronautStatusResponse.astronaut:type_name -> astronaut.Astronaut
	2,  // 6: astronaut.AstronautService.CreateAstronaut:input_type -> astronaut.CreateAstronautRequest
	4,  // 7: astronaut.AstronautService.GetAstronaut:input_type -> astronaut.GetAstronautRequest
	6,  // 8: astronaut.AstronautService.ListAstronauts:input_type -> astronaut.ListAstronautsRequest
	8,  // 9: astronaut.AstronautService.UpdateAstronautStatus:input_type -> astronaut.UpdateAstronautStatusRequest
	3,  // 10: astronaut.AstronautService.CreateAstronaut:output_type -> astronaut.CreateAstronautResponse
	5,  // 11: astronaut.AstronautService.GetAstronaut:output_type -> astronaut.GetAstronautResponse
	7,  // 12: astronaut.AstronautService.ListAstronauts:output_type -> astronaut.ListAstronautsResponse
	9,  // 13: astronaut.AstronautService.UpdateAstronautStatus:output_type -> astronaut.UpdateAstronautStatusResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_grpc_protos_astronaut_proto_init() }
func file_grpc_protos_astronaut_proto_init() {
	if File_grpc_protos_astronaut_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_grpc_protos_astronaut_proto_rawDesc), len(file_grpc_protos_astronaut_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_protos_astronaut_proto_goTypes,
		DependencyIndexes: file_grpc_protos_astronaut_proto_depIdxs,
		EnumInfos:         file_grpc_protos_astronaut_proto_enumTypes,
		MessageInfos:      file_grpc_protos_astronaut_proto_msgTypes,
	}.Build()
	File_grpc_protos_astronaut_proto = out.File
	file_grpc_protos_astronaut_proto_goTypes = nil
	file_grpc_protos_astronaut_proto_depIdxs = nil
}
