// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: runner/runner.proto

package runner

import (
	database "github.com/JanMeckelholt/running/common/grpc/database"
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

type HealthMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Health string `protobuf:"bytes,1,opt,name=Health,proto3" json:"Health,omitempty"`
}

func (x *HealthMessage) Reset() {
	*x = HealthMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_runner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthMessage) ProtoMessage() {}

func (x *HealthMessage) ProtoReflect() protoreflect.Message {
	mi := &file_runner_runner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthMessage.ProtoReflect.Descriptor instead.
func (*HealthMessage) Descriptor() ([]byte, []int) {
	return file_runner_runner_proto_rawDescGZIP(), []int{0}
}

func (x *HealthMessage) GetHealth() string {
	if x != nil {
		return x.Health
	}
	return ""
}

type RunnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *RunnerRequest) Reset() {
	*x = RunnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_runner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunnerRequest) ProtoMessage() {}

func (x *RunnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runner_runner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunnerRequest.ProtoReflect.Descriptor instead.
func (*RunnerRequest) Descriptor() ([]byte, []int) {
	return file_runner_runner_proto_rawDescGZIP(), []int{1}
}

func (x *RunnerRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type ActivitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Since    uint64 `protobuf:"varint,2,opt,name=since,proto3" json:"since,omitempty"`
}

func (x *ActivitiesRequest) Reset() {
	*x = ActivitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_runner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivitiesRequest) ProtoMessage() {}

func (x *ActivitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runner_runner_proto_msgTypes[2]
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
	return file_runner_runner_proto_rawDescGZIP(), []int{2}
}

func (x *ActivitiesRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ActivitiesRequest) GetSince() uint64 {
	if x != nil {
		return x.Since
	}
	return 0
}

type WeekSummariesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Weeks    uint64 `protobuf:"varint,2,opt,name=weeks,proto3" json:"weeks,omitempty"`
}

func (x *WeekSummariesRequest) Reset() {
	*x = WeekSummariesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_runner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeekSummariesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeekSummariesRequest) ProtoMessage() {}

func (x *WeekSummariesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runner_runner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeekSummariesRequest.ProtoReflect.Descriptor instead.
func (*WeekSummariesRequest) Descriptor() ([]byte, []int) {
	return file_runner_runner_proto_rawDescGZIP(), []int{3}
}

func (x *WeekSummariesRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *WeekSummariesRequest) GetWeeks() uint64 {
	if x != nil {
		return x.Weeks
	}
	return 0
}

type WeekSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Distance       uint64 `protobuf:"varint,1,opt,name=distance,proto3" json:"distance,omitempty"`
	TimeUnix       uint64 `protobuf:"varint,2,opt,name=timeUnix,proto3" json:"timeUnix,omitempty"`
	TimeStr        string `protobuf:"bytes,3,opt,name=timeStr,proto3" json:"timeStr,omitempty"`
	Climb          uint64 `protobuf:"varint,4,opt,name=climb,proto3" json:"climb,omitempty"`
	NumberOfRuns   uint64 `protobuf:"varint,5,opt,name=numberOfRuns,proto3" json:"numberOfRuns,omitempty"`
	NumberOfOthers uint64 `protobuf:"varint,6,opt,name=numberOfOthers,proto3" json:"numberOfOthers,omitempty"`
}

func (x *WeekSummary) Reset() {
	*x = WeekSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_runner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeekSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeekSummary) ProtoMessage() {}

func (x *WeekSummary) ProtoReflect() protoreflect.Message {
	mi := &file_runner_runner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeekSummary.ProtoReflect.Descriptor instead.
func (*WeekSummary) Descriptor() ([]byte, []int) {
	return file_runner_runner_proto_rawDescGZIP(), []int{4}
}

func (x *WeekSummary) GetDistance() uint64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *WeekSummary) GetTimeUnix() uint64 {
	if x != nil {
		return x.TimeUnix
	}
	return 0
}

func (x *WeekSummary) GetTimeStr() string {
	if x != nil {
		return x.TimeStr
	}
	return ""
}

func (x *WeekSummary) GetClimb() uint64 {
	if x != nil {
		return x.Climb
	}
	return 0
}

func (x *WeekSummary) GetNumberOfRuns() uint64 {
	if x != nil {
		return x.NumberOfRuns
	}
	return 0
}

func (x *WeekSummary) GetNumberOfOthers() uint64 {
	if x != nil {
		return x.NumberOfOthers
	}
	return 0
}

type WeekSummariesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weeksummaries []*WeekSummary `protobuf:"bytes,1,rep,name=weeksummaries,proto3" json:"weeksummaries,omitempty"`
}

func (x *WeekSummariesResponse) Reset() {
	*x = WeekSummariesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_runner_runner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeekSummariesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeekSummariesResponse) ProtoMessage() {}

func (x *WeekSummariesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runner_runner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeekSummariesResponse.ProtoReflect.Descriptor instead.
func (*WeekSummariesResponse) Descriptor() ([]byte, []int) {
	return file_runner_runner_proto_rawDescGZIP(), []int{5}
}

func (x *WeekSummariesResponse) GetWeeksummaries() []*WeekSummary {
	if x != nil {
		return x.Weeksummaries
	}
	return nil
}

var File_runner_runner_proto protoreflect.FileDescriptor

var file_runner_runner_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2f, 0x73, 0x74, 0x72, 0x61,
	0x76, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x22, 0x2b, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x45,
	0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x73, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x48, 0x0a, 0x14, 0x57, 0x65, 0x65, 0x6b, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x65, 0x65,
	0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x22,
	0xc1, 0x01, 0x0a, 0x0b, 0x57, 0x65, 0x65, 0x6b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x53,
	0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x69, 0x6d, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x63, 0x6c, 0x69, 0x6d, 0x62, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x4f, 0x66, 0x52, 0x75, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x52, 0x75, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x73, 0x22, 0x52, 0x0a, 0x15, 0x57, 0x65, 0x65, 0x6b, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0d,
	0x77, 0x65, 0x65, 0x6b, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x65,
	0x6b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0d, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x32, 0xa4, 0x03, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x12, 0x3c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12,
	0x15, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2e,
	0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3a, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x10, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x54, 0x6f, 0x44, 0x42, 0x12, 0x19, 0x2e, 0x72,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x57,
	0x65, 0x65, 0x6b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x65,
	0x6b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x15,
	0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x6e,
	0x4d, 0x65, 0x63, 0x6b, 0x65, 0x6c, 0x68, 0x6f, 0x6c, 0x74, 0x2f, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_runner_runner_proto_rawDescOnce sync.Once
	file_runner_runner_proto_rawDescData = file_runner_runner_proto_rawDesc
)

func file_runner_runner_proto_rawDescGZIP() []byte {
	file_runner_runner_proto_rawDescOnce.Do(func() {
		file_runner_runner_proto_rawDescData = protoimpl.X.CompressGZIP(file_runner_runner_proto_rawDescData)
	})
	return file_runner_runner_proto_rawDescData
}

var file_runner_runner_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_runner_runner_proto_goTypes = []interface{}{
	(*HealthMessage)(nil),               // 0: runner.HealthMessage
	(*RunnerRequest)(nil),               // 1: runner.RunnerRequest
	(*ActivitiesRequest)(nil),           // 2: runner.ActivitiesRequest
	(*WeekSummariesRequest)(nil),        // 3: runner.WeekSummariesRequest
	(*WeekSummary)(nil),                 // 4: runner.WeekSummary
	(*WeekSummariesResponse)(nil),       // 5: runner.WeekSummariesResponse
	(*strava.RunnerRequest)(nil),        // 6: strava.RunnerRequest
	(*database.Client)(nil),             // 7: database.Client
	(*database.ActivitiesRequest)(nil),  // 8: database.ActivitiesRequest
	(*strava.RunnerResponse)(nil),       // 9: strava.RunnerResponse
	(*empty.Empty)(nil),                 // 10: google.protobuf.Empty
	(*database.ActivitiesResponse)(nil), // 11: database.ActivitiesResponse
}
var file_runner_runner_proto_depIdxs = []int32{
	4,  // 0: runner.WeekSummariesResponse.weeksummaries:type_name -> runner.WeekSummary
	6,  // 1: runner.Runner.GetRunner:input_type -> strava.RunnerRequest
	7,  // 2: runner.Runner.CreateClient:input_type -> database.Client
	8,  // 3: runner.Runner.GetActivities:input_type -> database.ActivitiesRequest
	2,  // 4: runner.Runner.ActivitiesToDB:input_type -> runner.ActivitiesRequest
	3,  // 5: runner.Runner.GetWeekSummaries:input_type -> runner.WeekSummariesRequest
	0,  // 6: runner.Runner.Health:input_type -> runner.HealthMessage
	9,  // 7: runner.Runner.GetRunner:output_type -> strava.RunnerResponse
	10, // 8: runner.Runner.CreateClient:output_type -> google.protobuf.Empty
	11, // 9: runner.Runner.GetActivities:output_type -> database.ActivitiesResponse
	10, // 10: runner.Runner.ActivitiesToDB:output_type -> google.protobuf.Empty
	5,  // 11: runner.Runner.GetWeekSummaries:output_type -> runner.WeekSummariesResponse
	0,  // 12: runner.Runner.Health:output_type -> runner.HealthMessage
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_runner_runner_proto_init() }
func file_runner_runner_proto_init() {
	if File_runner_runner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_runner_runner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthMessage); i {
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
		file_runner_runner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunnerRequest); i {
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
		file_runner_runner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_runner_runner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeekSummariesRequest); i {
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
		file_runner_runner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeekSummary); i {
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
		file_runner_runner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeekSummariesResponse); i {
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
			RawDescriptor: file_runner_runner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_runner_runner_proto_goTypes,
		DependencyIndexes: file_runner_runner_proto_depIdxs,
		MessageInfos:      file_runner_runner_proto_msgTypes,
	}.Build()
	File_runner_runner_proto = out.File
	file_runner_runner_proto_rawDesc = nil
	file_runner_runner_proto_goTypes = nil
	file_runner_runner_proto_depIdxs = nil
}
