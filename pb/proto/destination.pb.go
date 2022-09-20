// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/destination.proto

package proto

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

type OperationStatus int32

const (
	OperationStatus_SUCCESS         OperationStatus = 0
	OperationStatus_PARTIAL_FAILURE OperationStatus = 1
	OperationStatus_TOTAL_FAILURE   OperationStatus = 2
)

// Enum value maps for OperationStatus.
var (
	OperationStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "PARTIAL_FAILURE",
		2: "TOTAL_FAILURE",
	}
	OperationStatus_value = map[string]int32{
		"SUCCESS":         0,
		"PARTIAL_FAILURE": 1,
		"TOTAL_FAILURE":   2,
	}
)

func (x OperationStatus) Enum() *OperationStatus {
	p := new(OperationStatus)
	*p = x
	return p
}

func (x OperationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_destination_proto_enumTypes[0].Descriptor()
}

func (OperationStatus) Type() protoreflect.EnumType {
	return &file_proto_destination_proto_enumTypes[0]
}

func (x OperationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationStatus.Descriptor instead.
func (OperationStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_destination_proto_rawDescGZIP(), []int{0}
}

type BatchUpsertUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users    []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Mappings []byte  `protobuf:"bytes,2,opt,name=mappings,proto3" json:"mappings,omitempty"`
	Settings []byte  `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *BatchUpsertUsersRequest) Reset() {
	*x = BatchUpsertUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_destination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchUpsertUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchUpsertUsersRequest) ProtoMessage() {}

func (x *BatchUpsertUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_destination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchUpsertUsersRequest.ProtoReflect.Descriptor instead.
func (*BatchUpsertUsersRequest) Descriptor() ([]byte, []int) {
	return file_proto_destination_proto_rawDescGZIP(), []int{0}
}

func (x *BatchUpsertUsersRequest) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *BatchUpsertUsersRequest) GetMappings() []byte {
	if x != nil {
		return x.Mappings
	}
	return nil
}

func (x *BatchUpsertUsersRequest) GetSettings() []byte {
	if x != nil {
		return x.Settings
	}
	return nil
}

type BatchUpsertUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status OperationStatus  `protobuf:"varint,1,opt,name=status,proto3,enum=proto.OperationStatus" json:"status,omitempty"`
	Errors []*ResponseError `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *BatchUpsertUsersResponse) Reset() {
	*x = BatchUpsertUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_destination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchUpsertUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchUpsertUsersResponse) ProtoMessage() {}

func (x *BatchUpsertUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_destination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchUpsertUsersResponse.ProtoReflect.Descriptor instead.
func (*BatchUpsertUsersResponse) Descriptor() ([]byte, []int) {
	return file_proto_destination_proto_rawDescGZIP(), []int{1}
}

func (x *BatchUpsertUsersResponse) GetStatus() OperationStatus {
	if x != nil {
		return x.Status
	}
	return OperationStatus_SUCCESS
}

func (x *BatchUpsertUsersResponse) GetErrors() []*ResponseError {
	if x != nil {
		return x.Errors
	}
	return nil
}

type BatchUpdateUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users    []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Mappings []byte  `protobuf:"bytes,2,opt,name=mappings,proto3" json:"mappings,omitempty"`
	Settings []byte  `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *BatchUpdateUsersRequest) Reset() {
	*x = BatchUpdateUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_destination_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchUpdateUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchUpdateUsersRequest) ProtoMessage() {}

func (x *BatchUpdateUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_destination_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchUpdateUsersRequest.ProtoReflect.Descriptor instead.
func (*BatchUpdateUsersRequest) Descriptor() ([]byte, []int) {
	return file_proto_destination_proto_rawDescGZIP(), []int{2}
}

func (x *BatchUpdateUsersRequest) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *BatchUpdateUsersRequest) GetMappings() []byte {
	if x != nil {
		return x.Mappings
	}
	return nil
}

func (x *BatchUpdateUsersRequest) GetSettings() []byte {
	if x != nil {
		return x.Settings
	}
	return nil
}

type BatchUpdateUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status OperationStatus  `protobuf:"varint,1,opt,name=status,proto3,enum=proto.OperationStatus" json:"status,omitempty"`
	Errors []*ResponseError `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *BatchUpdateUsersResponse) Reset() {
	*x = BatchUpdateUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_destination_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchUpdateUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchUpdateUsersResponse) ProtoMessage() {}

func (x *BatchUpdateUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_destination_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchUpdateUsersResponse.ProtoReflect.Descriptor instead.
func (*BatchUpdateUsersResponse) Descriptor() ([]byte, []int) {
	return file_proto_destination_proto_rawDescGZIP(), []int{3}
}

func (x *BatchUpdateUsersResponse) GetStatus() OperationStatus {
	if x != nil {
		return x.Status
	}
	return OperationStatus_SUCCESS
}

func (x *BatchUpdateUsersResponse) GetErrors() []*ResponseError {
	if x != nil {
		return x.Errors
	}
	return nil
}

type BatchDeleteUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds  []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	Settings []byte   `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *BatchDeleteUsersRequest) Reset() {
	*x = BatchDeleteUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_destination_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchDeleteUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDeleteUsersRequest) ProtoMessage() {}

func (x *BatchDeleteUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_destination_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDeleteUsersRequest.ProtoReflect.Descriptor instead.
func (*BatchDeleteUsersRequest) Descriptor() ([]byte, []int) {
	return file_proto_destination_proto_rawDescGZIP(), []int{4}
}

func (x *BatchDeleteUsersRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *BatchDeleteUsersRequest) GetSettings() []byte {
	if x != nil {
		return x.Settings
	}
	return nil
}

type BatchDeleteUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status OperationStatus  `protobuf:"varint,1,opt,name=status,proto3,enum=proto.OperationStatus" json:"status,omitempty"`
	Errors []*ResponseError `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *BatchDeleteUsersResponse) Reset() {
	*x = BatchDeleteUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_destination_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchDeleteUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDeleteUsersResponse) ProtoMessage() {}

func (x *BatchDeleteUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_destination_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDeleteUsersResponse.ProtoReflect.Descriptor instead.
func (*BatchDeleteUsersResponse) Descriptor() ([]byte, []int) {
	return file_proto_destination_proto_rawDescGZIP(), []int{5}
}

func (x *BatchDeleteUsersResponse) GetStatus() OperationStatus {
	if x != nil {
		return x.Status
	}
	return OperationStatus_SUCCESS
}

func (x *BatchDeleteUsersResponse) GetErrors() []*ResponseError {
	if x != nil {
		return x.Errors
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName  string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"` // TODO: add additional fields
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_destination_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_destination_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_destination_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type ResponseError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseError) Reset() {
	*x = ResponseError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_destination_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseError) ProtoMessage() {}

func (x *ResponseError) ProtoReflect() protoreflect.Message {
	mi := &file_proto_destination_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseError.ProtoReflect.Descriptor instead.
func (*ResponseError) Descriptor() ([]byte, []int) {
	return file_proto_destination_proto_rawDescGZIP(), []int{7}
}

func (x *ResponseError) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ResponseError) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ResponseError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_destination_proto protoreflect.FileDescriptor

var file_proto_destination_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x74, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x78, 0x0a, 0x18, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x22, 0x74, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x78, 0x0a, 0x18, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x22, 0x50, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x22, 0x78, 0x0a, 0x18, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x9b, 0x01, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x56, 0x0a, 0x0d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2a, 0x46, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x4f, 0x54, 0x41, 0x4c,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x32, 0x92, 0x02, 0x0a, 0x11, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x12, 0x53, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_destination_proto_rawDescOnce sync.Once
	file_proto_destination_proto_rawDescData = file_proto_destination_proto_rawDesc
)

func file_proto_destination_proto_rawDescGZIP() []byte {
	file_proto_destination_proto_rawDescOnce.Do(func() {
		file_proto_destination_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_destination_proto_rawDescData)
	})
	return file_proto_destination_proto_rawDescData
}

var file_proto_destination_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_destination_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_destination_proto_goTypes = []interface{}{
	(OperationStatus)(0),             // 0: proto.OperationStatus
	(*BatchUpsertUsersRequest)(nil),  // 1: proto.BatchUpsertUsersRequest
	(*BatchUpsertUsersResponse)(nil), // 2: proto.BatchUpsertUsersResponse
	(*BatchUpdateUsersRequest)(nil),  // 3: proto.BatchUpdateUsersRequest
	(*BatchUpdateUsersResponse)(nil), // 4: proto.BatchUpdateUsersResponse
	(*BatchDeleteUsersRequest)(nil),  // 5: proto.BatchDeleteUsersRequest
	(*BatchDeleteUsersResponse)(nil), // 6: proto.BatchDeleteUsersResponse
	(*User)(nil),                     // 7: proto.User
	(*ResponseError)(nil),            // 8: proto.ResponseError
}
var file_proto_destination_proto_depIdxs = []int32{
	7,  // 0: proto.BatchUpsertUsersRequest.users:type_name -> proto.User
	0,  // 1: proto.BatchUpsertUsersResponse.status:type_name -> proto.OperationStatus
	8,  // 2: proto.BatchUpsertUsersResponse.errors:type_name -> proto.ResponseError
	7,  // 3: proto.BatchUpdateUsersRequest.users:type_name -> proto.User
	0,  // 4: proto.BatchUpdateUsersResponse.status:type_name -> proto.OperationStatus
	8,  // 5: proto.BatchUpdateUsersResponse.errors:type_name -> proto.ResponseError
	0,  // 6: proto.BatchDeleteUsersResponse.status:type_name -> proto.OperationStatus
	8,  // 7: proto.BatchDeleteUsersResponse.errors:type_name -> proto.ResponseError
	1,  // 8: proto.DestinationPlugin.BatchUpsertUsers:input_type -> proto.BatchUpsertUsersRequest
	3,  // 9: proto.DestinationPlugin.BatchUpdateUsers:input_type -> proto.BatchUpdateUsersRequest
	5,  // 10: proto.DestinationPlugin.BatchDeleteUsers:input_type -> proto.BatchDeleteUsersRequest
	2,  // 11: proto.DestinationPlugin.BatchUpsertUsers:output_type -> proto.BatchUpsertUsersResponse
	4,  // 12: proto.DestinationPlugin.BatchUpdateUsers:output_type -> proto.BatchUpdateUsersResponse
	6,  // 13: proto.DestinationPlugin.BatchDeleteUsers:output_type -> proto.BatchDeleteUsersResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_destination_proto_init() }
func file_proto_destination_proto_init() {
	if File_proto_destination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_destination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchUpsertUsersRequest); i {
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
		file_proto_destination_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchUpsertUsersResponse); i {
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
		file_proto_destination_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchUpdateUsersRequest); i {
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
		file_proto_destination_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchUpdateUsersResponse); i {
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
		file_proto_destination_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchDeleteUsersRequest); i {
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
		file_proto_destination_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchDeleteUsersResponse); i {
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
		file_proto_destination_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_destination_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseError); i {
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
			RawDescriptor: file_proto_destination_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_destination_proto_goTypes,
		DependencyIndexes: file_proto_destination_proto_depIdxs,
		EnumInfos:         file_proto_destination_proto_enumTypes,
		MessageInfos:      file_proto_destination_proto_msgTypes,
	}.Build()
	File_proto_destination_proto = out.File
	file_proto_destination_proto_rawDesc = nil
	file_proto_destination_proto_goTypes = nil
	file_proto_destination_proto_depIdxs = nil
}