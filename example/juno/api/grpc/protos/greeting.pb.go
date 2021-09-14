// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: api/grpc/protos/greeting.proto

package protos

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

type GetGreetingListParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize  int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *GetGreetingListParam) Reset() {
	*x = GetGreetingListParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreetingListParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreetingListParam) ProtoMessage() {}

func (x *GetGreetingListParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreetingListParam.ProtoReflect.Descriptor instead.
func (*GetGreetingListParam) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{0}
}

func (x *GetGreetingListParam) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *GetGreetingListParam) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetGreetingListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item  []*GetGreetingListReply_List `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
	Total int64                        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetGreetingListReply) Reset() {
	*x = GetGreetingListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreetingListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreetingListReply) ProtoMessage() {}

func (x *GetGreetingListReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreetingListReply.ProtoReflect.Descriptor instead.
func (*GetGreetingListReply) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{1}
}

func (x *GetGreetingListReply) GetItem() []*GetGreetingListReply_List {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *GetGreetingListReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetGreetingAllParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGreetingAllParam) Reset() {
	*x = GetGreetingAllParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreetingAllParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreetingAllParam) ProtoMessage() {}

func (x *GetGreetingAllParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreetingAllParam.ProtoReflect.Descriptor instead.
func (*GetGreetingAllParam) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{2}
}

type GetGreetingAllReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []*GetGreetingAllReply_List `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *GetGreetingAllReply) Reset() {
	*x = GetGreetingAllReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreetingAllReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreetingAllReply) ProtoMessage() {}

func (x *GetGreetingAllReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreetingAllReply.ProtoReflect.Descriptor instead.
func (*GetGreetingAllReply) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{3}
}

func (x *GetGreetingAllReply) GetItem() []*GetGreetingAllReply_List {
	if x != nil {
		return x.Item
	}
	return nil
}

type GetGreetingDetailParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetGreetingDetailParam) Reset() {
	*x = GetGreetingDetailParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreetingDetailParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreetingDetailParam) ProtoMessage() {}

func (x *GetGreetingDetailParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreetingDetailParam.ProtoReflect.Descriptor instead.
func (*GetGreetingDetailParam) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{4}
}

func (x *GetGreetingDetailParam) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGreetingDetailReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGreetingDetailReply) Reset() {
	*x = GetGreetingDetailReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreetingDetailReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreetingDetailReply) ProtoMessage() {}

func (x *GetGreetingDetailReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreetingDetailReply.ProtoReflect.Descriptor instead.
func (*GetGreetingDetailReply) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{5}
}

type CreateGreetingParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateGreetingParam) Reset() {
	*x = CreateGreetingParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGreetingParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGreetingParam) ProtoMessage() {}

func (x *CreateGreetingParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGreetingParam.ProtoReflect.Descriptor instead.
func (*CreateGreetingParam) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{6}
}

type CreateGreetingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateGreetingReply) Reset() {
	*x = CreateGreetingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGreetingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGreetingReply) ProtoMessage() {}

func (x *CreateGreetingReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGreetingReply.ProtoReflect.Descriptor instead.
func (*CreateGreetingReply) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{7}
}

type UpdateGreetingParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *UpdateGreetingParam) Reset() {
	*x = UpdateGreetingParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGreetingParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGreetingParam) ProtoMessage() {}

func (x *UpdateGreetingParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGreetingParam.ProtoReflect.Descriptor instead.
func (*UpdateGreetingParam) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateGreetingParam) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateGreetingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateGreetingReply) Reset() {
	*x = UpdateGreetingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGreetingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGreetingReply) ProtoMessage() {}

func (x *UpdateGreetingReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGreetingReply.ProtoReflect.Descriptor instead.
func (*UpdateGreetingReply) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{9}
}

type DeleteGreetingParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DeleteGreetingParam) Reset() {
	*x = DeleteGreetingParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGreetingParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGreetingParam) ProtoMessage() {}

func (x *DeleteGreetingParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGreetingParam.ProtoReflect.Descriptor instead.
func (*DeleteGreetingParam) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteGreetingParam) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteGreetingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGreetingReply) Reset() {
	*x = DeleteGreetingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGreetingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGreetingReply) ProtoMessage() {}

func (x *DeleteGreetingReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGreetingReply.ProtoReflect.Descriptor instead.
func (*DeleteGreetingReply) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{11}
}

type GetGreetingListReply_List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGreetingListReply_List) Reset() {
	*x = GetGreetingListReply_List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreetingListReply_List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreetingListReply_List) ProtoMessage() {}

func (x *GetGreetingListReply_List) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreetingListReply_List.ProtoReflect.Descriptor instead.
func (*GetGreetingListReply_List) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{1, 0}
}

type GetGreetingAllReply_List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGreetingAllReply_List) Reset() {
	*x = GetGreetingAllReply_List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_greeting_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGreetingAllReply_List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGreetingAllReply_List) ProtoMessage() {}

func (x *GetGreetingAllReply_List) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_greeting_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGreetingAllReply_List.ProtoReflect.Descriptor instead.
func (*GetGreetingAllReply_List) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_greeting_proto_rawDescGZIP(), []int{3, 0}
}

var File_api_grpc_protos_greeting_proto protoreflect.FileDescriptor

var file_api_grpc_protos_greeting_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x50, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x6b, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x35, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x1a,
	0x06, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x53,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x1a, 0x06, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x18, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x15,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x32, 0xba, 0x03, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x47,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x18,
	0x5a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_protos_greeting_proto_rawDescOnce sync.Once
	file_api_grpc_protos_greeting_proto_rawDescData = file_api_grpc_protos_greeting_proto_rawDesc
)

func file_api_grpc_protos_greeting_proto_rawDescGZIP() []byte {
	file_api_grpc_protos_greeting_proto_rawDescOnce.Do(func() {
		file_api_grpc_protos_greeting_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_protos_greeting_proto_rawDescData)
	})
	return file_api_grpc_protos_greeting_proto_rawDescData
}

var file_api_grpc_protos_greeting_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_grpc_protos_greeting_proto_goTypes = []interface{}{
	(*GetGreetingListParam)(nil),      // 0: protos.GetGreetingListParam
	(*GetGreetingListReply)(nil),      // 1: protos.GetGreetingListReply
	(*GetGreetingAllParam)(nil),       // 2: protos.GetGreetingAllParam
	(*GetGreetingAllReply)(nil),       // 3: protos.GetGreetingAllReply
	(*GetGreetingDetailParam)(nil),    // 4: protos.GetGreetingDetailParam
	(*GetGreetingDetailReply)(nil),    // 5: protos.GetGreetingDetailReply
	(*CreateGreetingParam)(nil),       // 6: protos.CreateGreetingParam
	(*CreateGreetingReply)(nil),       // 7: protos.CreateGreetingReply
	(*UpdateGreetingParam)(nil),       // 8: protos.UpdateGreetingParam
	(*UpdateGreetingReply)(nil),       // 9: protos.UpdateGreetingReply
	(*DeleteGreetingParam)(nil),       // 10: protos.DeleteGreetingParam
	(*DeleteGreetingReply)(nil),       // 11: protos.DeleteGreetingReply
	(*GetGreetingListReply_List)(nil), // 12: protos.GetGreetingListReply.List
	(*GetGreetingAllReply_List)(nil),  // 13: protos.GetGreetingAllReply.List
}
var file_api_grpc_protos_greeting_proto_depIdxs = []int32{
	12, // 0: protos.GetGreetingListReply.item:type_name -> protos.GetGreetingListReply.List
	13, // 1: protos.GetGreetingAllReply.item:type_name -> protos.GetGreetingAllReply.List
	0,  // 2: protos.Greeting.GetList:input_type -> protos.GetGreetingListParam
	2,  // 3: protos.Greeting.GetAll:input_type -> protos.GetGreetingAllParam
	4,  // 4: protos.Greeting.GetDetail:input_type -> protos.GetGreetingDetailParam
	6,  // 5: protos.Greeting.Create:input_type -> protos.CreateGreetingParam
	8,  // 6: protos.Greeting.Update:input_type -> protos.UpdateGreetingParam
	10, // 7: protos.Greeting.Delete:input_type -> protos.DeleteGreetingParam
	1,  // 8: protos.Greeting.GetList:output_type -> protos.GetGreetingListReply
	3,  // 9: protos.Greeting.GetAll:output_type -> protos.GetGreetingAllReply
	5,  // 10: protos.Greeting.GetDetail:output_type -> protos.GetGreetingDetailReply
	7,  // 11: protos.Greeting.Create:output_type -> protos.CreateGreetingReply
	9,  // 12: protos.Greeting.Update:output_type -> protos.UpdateGreetingReply
	11, // 13: protos.Greeting.Delete:output_type -> protos.DeleteGreetingReply
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_grpc_protos_greeting_proto_init() }
func file_api_grpc_protos_greeting_proto_init() {
	if File_api_grpc_protos_greeting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_protos_greeting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreetingListParam); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreetingListReply); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreetingAllParam); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreetingAllReply); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreetingDetailParam); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreetingDetailReply); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGreetingParam); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGreetingReply); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGreetingParam); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGreetingReply); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGreetingParam); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGreetingReply); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreetingListReply_List); i {
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
		file_api_grpc_protos_greeting_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGreetingAllReply_List); i {
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
			RawDescriptor: file_api_grpc_protos_greeting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_protos_greeting_proto_goTypes,
		DependencyIndexes: file_api_grpc_protos_greeting_proto_depIdxs,
		MessageInfos:      file_api_grpc_protos_greeting_proto_msgTypes,
	}.Build()
	File_api_grpc_protos_greeting_proto = out.File
	file_api_grpc_protos_greeting_proto_rawDesc = nil
	file_api_grpc_protos_greeting_proto_goTypes = nil
	file_api_grpc_protos_greeting_proto_depIdxs = nil
}