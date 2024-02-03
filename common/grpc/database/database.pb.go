// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: database/database.proto

package database

import (
	strava "github.com/JanMeckelholt/running/common/grpc/strava"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ActivitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AthleteId uint64 `protobuf:"varint,1,opt,name=athleteId,proto3" json:"athleteId,omitempty"`
	Since     uint64 `protobuf:"varint,2,opt,name=since,proto3" json:"since,omitempty"`
	Until     uint64 `protobuf:"varint,3,opt,name=until,proto3" json:"until,omitempty"`
}

func (x *ActivitiesRequest) Reset() {
	*x = ActivitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivitiesRequest) ProtoMessage() {}

func (x *ActivitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivitiesRequest.ProtoReflect.Descriptor instead.
func (*ActivitiesRequest) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{0}
}

func (x *ActivitiesRequest) GetAthleteId() uint64 {
	if x != nil {
		return x.AthleteId
	}
	return 0
}

func (x *ActivitiesRequest) GetSince() uint64 {
	if x != nil {
		return x.Since
	}
	return 0
}

func (x *ActivitiesRequest) GetUntil() uint64 {
	if x != nil {
		return x.Until
	}
	return 0
}

type ActivitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activities []*strava.Activity `protobuf:"bytes,1,rep,name=Activities,proto3" json:"Activities,omitempty"`
}

func (x *ActivitiesResponse) Reset() {
	*x = ActivitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivitiesResponse) ProtoMessage() {}

func (x *ActivitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivitiesResponse.ProtoReflect.Descriptor instead.
func (*ActivitiesResponse) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{1}
}

func (x *ActivitiesResponse) GetActivities() []*strava.Activity {
	if x != nil {
		return x.Activities
	}
	return nil
}

type ClientId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *ClientId) Reset() {
	*x = ClientId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientId) ProtoMessage() {}

func (x *ClientId) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientId.ProtoReflect.Descriptor instead.
func (*ClientId) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{2}
}

func (x *ClientId) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type AthleteId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AthleteId uint64 `protobuf:"varint,1,opt,name=athleteId,proto3" json:"athleteId,omitempty"`
}

func (x *AthleteId) Reset() {
	*x = AthleteId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AthleteId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AthleteId) ProtoMessage() {}

func (x *AthleteId) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AthleteId.ProtoReflect.Descriptor instead.
func (*AthleteId) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{3}
}

func (x *AthleteId) GetAthleteId() uint64 {
	if x != nil {
		return x.AthleteId
	}
	return 0
}

type ActivityId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId string `protobuf:"bytes,1,opt,name=activityId,proto3" json:"activityId,omitempty"`
}

func (x *ActivityId) Reset() {
	*x = ActivityId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityId) ProtoMessage() {}

func (x *ActivityId) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityId.ProtoReflect.Descriptor instead.
func (*ActivityId) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{4}
}

func (x *ActivityId) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=clientSecret,proto3" json:"clientSecret,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{5}
}

func (x *Client) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Client) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

type Athlete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AthleteId    uint64             `protobuf:"varint,1,opt,name=athleteId,proto3" json:"athleteId,omitempty"`
	ClientId     string             `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Token        string             `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken string             `protobuf:"bytes,4,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	Activities   []*strava.Activity `protobuf:"bytes,5,rep,name=Activities,proto3" json:"Activities,omitempty"`
}

func (x *Athlete) Reset() {
	*x = Athlete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Athlete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Athlete) ProtoMessage() {}

func (x *Athlete) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Athlete.ProtoReflect.Descriptor instead.
func (*Athlete) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{6}
}

func (x *Athlete) GetAthleteId() uint64 {
	if x != nil {
		return x.AthleteId
	}
	return 0
}

func (x *Athlete) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Athlete) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Athlete) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Athlete) GetActivities() []*strava.Activity {
	if x != nil {
		return x.Activities
	}
	return nil
}

type UpdateClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string    `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	KvPairs  []*KvPair `protobuf:"bytes,2,rep,name=kvPairs,proto3" json:"kvPairs,omitempty"`
}

func (x *UpdateClientRequest) Reset() {
	*x = UpdateClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClientRequest) ProtoMessage() {}

func (x *UpdateClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClientRequest.ProtoReflect.Descriptor instead.
func (*UpdateClientRequest) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateClientRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *UpdateClientRequest) GetKvPairs() []*KvPair {
	if x != nil {
		return x.KvPairs
	}
	return nil
}

type UpdateAthleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AthleteId uint64    `protobuf:"varint,1,opt,name=athleteId,proto3" json:"athleteId,omitempty"`
	KvPairs   []*KvPair `protobuf:"bytes,2,rep,name=kvPairs,proto3" json:"kvPairs,omitempty"`
}

func (x *UpdateAthleteRequest) Reset() {
	*x = UpdateAthleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAthleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAthleteRequest) ProtoMessage() {}

func (x *UpdateAthleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAthleteRequest.ProtoReflect.Descriptor instead.
func (*UpdateAthleteRequest) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateAthleteRequest) GetAthleteId() uint64 {
	if x != nil {
		return x.AthleteId
	}
	return 0
}

func (x *UpdateAthleteRequest) GetKvPairs() []*KvPair {
	if x != nil {
		return x.KvPairs
	}
	return nil
}

type KvPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KvPair) Reset() {
	*x = KvPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KvPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KvPair) ProtoMessage() {}

func (x *KvPair) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KvPair.ProtoReflect.Descriptor instead.
func (*KvPair) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{9}
}

func (x *KvPair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KvPair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*TokenResponse_SuccessResponse
	//	*TokenResponse_ErrorResponse
	Response isTokenResponse_Response `protobuf_oneof:"response"`
}

func (x *TokenResponse) Reset() {
	*x = TokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{10}
}

func (m *TokenResponse) GetResponse() isTokenResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *TokenResponse) GetSuccessResponse() *SuccessResponse {
	if x, ok := x.GetResponse().(*TokenResponse_SuccessResponse); ok {
		return x.SuccessResponse
	}
	return nil
}

func (x *TokenResponse) GetErrorResponse() *ErrorResponse {
	if x, ok := x.GetResponse().(*TokenResponse_ErrorResponse); ok {
		return x.ErrorResponse
	}
	return nil
}

type isTokenResponse_Response interface {
	isTokenResponse_Response()
}

type TokenResponse_SuccessResponse struct {
	SuccessResponse *SuccessResponse `protobuf:"bytes,1,opt,name=success_response,json=successResponse,proto3,oneof"`
}

type TokenResponse_ErrorResponse struct {
	ErrorResponse *ErrorResponse `protobuf:"bytes,2,opt,name=error_response,json=errorResponse,proto3,oneof"`
}

func (*TokenResponse_SuccessResponse) isTokenResponse_Response() {}

func (*TokenResponse_ErrorResponse) isTokenResponse_Response() {}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{11}
}

func (x *SuccessResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_database_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_database_database_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_database_database_proto_rawDescGZIP(), []int{12}
}

func (x *ErrorResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_database_database_proto protoreflect.FileDescriptor

var file_database_database_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x74,
	0x68, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61,
	0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x75,
	0x6e, 0x74, 0x69, 0x6c, 0x22, 0x46, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x08,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x09, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x22,
	0x2c, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x48, 0x0a,
	0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xaf, 0x01, 0x0a, 0x07, 0x41, 0x74, 0x68, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x30, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74,
	0x72, 0x61, 0x76, 0x61, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x5d, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x07,
	0x6b, 0x76, 0x50, 0x61, 0x69, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x6b, 0x76, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x07, 0x6b, 0x76, 0x50, 0x61, 0x69, 0x72, 0x73, 0x22, 0x60, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x07, 0x6b, 0x76, 0x50, 0x61, 0x69, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x6b, 0x76, 0x50, 0x61, 0x69,
	0x72, 0x52, 0x07, 0x6b, 0x76, 0x50, 0x61, 0x69, 0x72, 0x73, 0x22, 0x30, 0x0a, 0x06, 0x6b, 0x76,
	0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa5, 0x01, 0x0a,
	0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x10, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x25, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x84, 0x05, 0x0a, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64,
	0x1a, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x74, 0x68, 0x6c,
	0x65, 0x74, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x10, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x15, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x53, 0x56, 0x12, 0x10, 0x2e, 0x73,
	0x74, 0x72, 0x61, 0x76, 0x61, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x1a, 0x10, 0x2e,
	0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61,
	0x6e, 0x4d, 0x65, 0x63, 0x6b, 0x65, 0x6c, 0x68, 0x6f, 0x6c, 0x74, 0x2f, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_database_database_proto_rawDescOnce sync.Once
	file_database_database_proto_rawDescData = file_database_database_proto_rawDesc
)

func file_database_database_proto_rawDescGZIP() []byte {
	file_database_database_proto_rawDescOnce.Do(func() {
		file_database_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_database_database_proto_rawDescData)
	})
	return file_database_database_proto_rawDescData
}

var file_database_database_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_database_database_proto_goTypes = []interface{}{
	(*ActivitiesRequest)(nil),    // 0: database.ActivitiesRequest
	(*ActivitiesResponse)(nil),   // 1: database.ActivitiesResponse
	(*ClientId)(nil),             // 2: database.ClientId
	(*AthleteId)(nil),            // 3: database.AthleteId
	(*ActivityId)(nil),           // 4: database.ActivityId
	(*Client)(nil),               // 5: database.Client
	(*Athlete)(nil),              // 6: database.Athlete
	(*UpdateClientRequest)(nil),  // 7: database.UpdateClientRequest
	(*UpdateAthleteRequest)(nil), // 8: database.UpdateAthleteRequest
	(*KvPair)(nil),               // 9: database.kvPair
	(*TokenResponse)(nil),        // 10: database.TokenResponse
	(*SuccessResponse)(nil),      // 11: database.SuccessResponse
	(*ErrorResponse)(nil),        // 12: database.ErrorResponse
	(*strava.Activity)(nil),      // 13: strava.Activity
	(*empty.Empty)(nil),          // 14: google.protobuf.Empty
}
var file_database_database_proto_depIdxs = []int32{
	13, // 0: database.ActivitiesResponse.Activities:type_name -> strava.Activity
	13, // 1: database.Athlete.Activities:type_name -> strava.Activity
	9,  // 2: database.UpdateClientRequest.kvPairs:type_name -> database.kvPair
	9,  // 3: database.UpdateAthleteRequest.kvPairs:type_name -> database.kvPair
	11, // 4: database.TokenResponse.success_response:type_name -> database.SuccessResponse
	12, // 5: database.TokenResponse.error_response:type_name -> database.ErrorResponse
	5,  // 6: database.Database.UpsertClient:input_type -> database.Client
	7,  // 7: database.Database.UpdateClient:input_type -> database.UpdateClientRequest
	2,  // 8: database.Database.GetClient:input_type -> database.ClientId
	6,  // 9: database.Database.UpsertAthlete:input_type -> database.Athlete
	8,  // 10: database.Database.UpdateAthlete:input_type -> database.UpdateAthleteRequest
	3,  // 11: database.Database.GetAthlete:input_type -> database.AthleteId
	13, // 12: database.Database.UpsertActivity:input_type -> strava.Activity
	13, // 13: database.Database.UpsertActivityFromCSV:input_type -> strava.Activity
	4,  // 14: database.Database.GetActivity:input_type -> database.ActivityId
	0,  // 15: database.Database.GetActivities:input_type -> database.ActivitiesRequest
	14, // 16: database.Database.UpsertClient:output_type -> google.protobuf.Empty
	5,  // 17: database.Database.UpdateClient:output_type -> database.Client
	5,  // 18: database.Database.GetClient:output_type -> database.Client
	14, // 19: database.Database.UpsertAthlete:output_type -> google.protobuf.Empty
	6,  // 20: database.Database.UpdateAthlete:output_type -> database.Athlete
	6,  // 21: database.Database.GetAthlete:output_type -> database.Athlete
	14, // 22: database.Database.UpsertActivity:output_type -> google.protobuf.Empty
	14, // 23: database.Database.UpsertActivityFromCSV:output_type -> google.protobuf.Empty
	13, // 24: database.Database.GetActivity:output_type -> strava.Activity
	1,  // 25: database.Database.GetActivities:output_type -> database.ActivitiesResponse
	16, // [16:26] is the sub-list for method output_type
	6,  // [6:16] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_database_database_proto_init() }
func file_database_database_proto_init() {
	if File_database_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_database_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivitiesRequest); i {
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
		file_database_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivitiesResponse); i {
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
		file_database_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientId); i {
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
		file_database_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AthleteId); i {
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
		file_database_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityId); i {
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
		file_database_database_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_database_database_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Athlete); i {
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
		file_database_database_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClientRequest); i {
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
		file_database_database_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAthleteRequest); i {
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
		file_database_database_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KvPair); i {
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
		file_database_database_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenResponse); i {
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
		file_database_database_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResponse); i {
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
		file_database_database_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
	file_database_database_proto_msgTypes[10].OneofWrappers = []interface{}{
		(*TokenResponse_SuccessResponse)(nil),
		(*TokenResponse_ErrorResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_database_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_database_database_proto_goTypes,
		DependencyIndexes: file_database_database_proto_depIdxs,
		MessageInfos:      file_database_database_proto_msgTypes,
	}.Build()
	File_database_database_proto = out.File
	file_database_database_proto_rawDesc = nil
	file_database_database_proto_goTypes = nil
	file_database_database_proto_depIdxs = nil
}
