// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: account/manager.proto

package managerpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerTelegramId  int64   `protobuf:"varint,1,opt,name=manager_telegram_id,json=managerTelegramId,proto3" json:"manager_telegram_id,omitempty"`
	ManagerFullName    string  `protobuf:"bytes,2,opt,name=manager_full_name,json=managerFullName,proto3" json:"manager_full_name,omitempty"`
	ManagerEmail       string  `protobuf:"bytes,3,opt,name=manager_email,json=managerEmail,proto3" json:"manager_email,omitempty"`
	ManagerPhoneNumber *string `protobuf:"bytes,4,opt,name=manager_phone_number,json=managerPhoneNumber,proto3,oneof" json:"manager_phone_number,omitempty"`
	ManagerCompany     *string `protobuf:"bytes,5,opt,name=manager_company,json=managerCompany,proto3,oneof" json:"manager_company,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_account_manager_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_manager_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_account_manager_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetManagerTelegramId() int64 {
	if x != nil {
		return x.ManagerTelegramId
	}
	return 0
}

func (x *RegisterRequest) GetManagerFullName() string {
	if x != nil {
		return x.ManagerFullName
	}
	return ""
}

func (x *RegisterRequest) GetManagerEmail() string {
	if x != nil {
		return x.ManagerEmail
	}
	return ""
}

func (x *RegisterRequest) GetManagerPhoneNumber() string {
	if x != nil && x.ManagerPhoneNumber != nil {
		return *x.ManagerPhoneNumber
	}
	return ""
}

func (x *RegisterRequest) GetManagerCompany() string {
	if x != nil && x.ManagerCompany != nil {
		return *x.ManagerCompany
	}
	return ""
}

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerTelegramId int64 `protobuf:"varint,1,opt,name=manager_telegram_id,json=managerTelegramId,proto3" json:"manager_telegram_id,omitempty"`
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	mi := &file_account_manager_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_manager_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_account_manager_proto_rawDescGZIP(), []int{1}
}

func (x *GetInfoRequest) GetManagerTelegramId() int64 {
	if x != nil {
		return x.ManagerTelegramId
	}
	return 0
}

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerFullName    string  `protobuf:"bytes,1,opt,name=manager_full_name,json=managerFullName,proto3" json:"manager_full_name,omitempty"`
	ManagerEmail       string  `protobuf:"bytes,2,opt,name=manager_email,json=managerEmail,proto3" json:"manager_email,omitempty"`
	ManagerPhoneNumber *string `protobuf:"bytes,3,opt,name=manager_phone_number,json=managerPhoneNumber,proto3,oneof" json:"manager_phone_number,omitempty"`
	ManagerCompany     *string `protobuf:"bytes,4,opt,name=manager_company,json=managerCompany,proto3,oneof" json:"manager_company,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	mi := &file_account_manager_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_manager_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_account_manager_proto_rawDescGZIP(), []int{2}
}

func (x *GetInfoResponse) GetManagerFullName() string {
	if x != nil {
		return x.ManagerFullName
	}
	return ""
}

func (x *GetInfoResponse) GetManagerEmail() string {
	if x != nil {
		return x.ManagerEmail
	}
	return ""
}

func (x *GetInfoResponse) GetManagerPhoneNumber() string {
	if x != nil && x.ManagerPhoneNumber != nil {
		return *x.ManagerPhoneNumber
	}
	return ""
}

func (x *GetInfoResponse) GetManagerCompany() string {
	if x != nil && x.ManagerCompany != nil {
		return *x.ManagerCompany
	}
	return ""
}

type GetApiTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerTelegramId int64 `protobuf:"varint,1,opt,name=manager_telegram_id,json=managerTelegramId,proto3" json:"manager_telegram_id,omitempty"`
}

func (x *GetApiTokenRequest) Reset() {
	*x = GetApiTokenRequest{}
	mi := &file_account_manager_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetApiTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiTokenRequest) ProtoMessage() {}

func (x *GetApiTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_manager_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiTokenRequest.ProtoReflect.Descriptor instead.
func (*GetApiTokenRequest) Descriptor() ([]byte, []int) {
	return file_account_manager_proto_rawDescGZIP(), []int{3}
}

func (x *GetApiTokenRequest) GetManagerTelegramId() int64 {
	if x != nil {
		return x.ManagerTelegramId
	}
	return 0
}

type GetApiTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerApiToken string `protobuf:"bytes,1,opt,name=manager_api_token,json=managerApiToken,proto3" json:"manager_api_token,omitempty"`
}

func (x *GetApiTokenResponse) Reset() {
	*x = GetApiTokenResponse{}
	mi := &file_account_manager_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetApiTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiTokenResponse) ProtoMessage() {}

func (x *GetApiTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_manager_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiTokenResponse.ProtoReflect.Descriptor instead.
func (*GetApiTokenResponse) Descriptor() ([]byte, []int) {
	return file_account_manager_proto_rawDescGZIP(), []int{4}
}

func (x *GetApiTokenResponse) GetManagerApiToken() string {
	if x != nil {
		return x.ManagerApiToken
	}
	return ""
}

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerTelegramId int64 `protobuf:"varint,1,opt,name=manager_telegram_id,json=managerTelegramId,proto3" json:"manager_telegram_id,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	mi := &file_account_manager_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_manager_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_account_manager_proto_rawDescGZIP(), []int{5}
}

func (x *AuthRequest) GetManagerTelegramId() int64 {
	if x != nil {
		return x.ManagerTelegramId
	}
	return 0
}

type AuthApiTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerApiToken string `protobuf:"bytes,1,opt,name=manager_api_token,json=managerApiToken,proto3" json:"manager_api_token,omitempty"`
}

func (x *AuthApiTokenRequest) Reset() {
	*x = AuthApiTokenRequest{}
	mi := &file_account_manager_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthApiTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthApiTokenRequest) ProtoMessage() {}

func (x *AuthApiTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_manager_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthApiTokenRequest.ProtoReflect.Descriptor instead.
func (*AuthApiTokenRequest) Descriptor() ([]byte, []int) {
	return file_account_manager_proto_rawDescGZIP(), []int{6}
}

func (x *AuthApiTokenRequest) GetManagerApiToken() string {
	if x != nil {
		return x.ManagerApiToken
	}
	return ""
}

type AuthApiTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagerTelegramId int64 `protobuf:"varint,1,opt,name=manager_telegram_id,json=managerTelegramId,proto3" json:"manager_telegram_id,omitempty"`
}

func (x *AuthApiTokenResponse) Reset() {
	*x = AuthApiTokenResponse{}
	mi := &file_account_manager_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthApiTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthApiTokenResponse) ProtoMessage() {}

func (x *AuthApiTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_manager_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthApiTokenResponse.ProtoReflect.Descriptor instead.
func (*AuthApiTokenResponse) Descriptor() ([]byte, []int) {
	return file_account_manager_proto_rawDescGZIP(), []int{7}
}

func (x *AuthApiTokenResponse) GetManagerTelegramId() int64 {
	if x != nil {
		return x.ManagerTelegramId
	}
	return 0
}

var File_account_manager_proto protoreflect.FileDescriptor

var file_account_manager_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa4, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x74,
	0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61,
	0x6d, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x66,
	0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x35, 0x0a, 0x14, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x12, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x88, 0x01, 0x01, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x40, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x54, 0x65,
	0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x22, 0xf4, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x35, 0x0a,
	0x14, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x12, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x88,
	0x01, 0x01, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x12, 0x0a, 0x10, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22,
	0x44, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x54, 0x65, 0x6c, 0x65, 0x67,
	0x72, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3d, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x54, 0x65, 0x6c,
	0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x41,
	0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x14, 0x41, 0x75,
	0x74, 0x68, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x74, 0x65,
	0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d,
	0x49, 0x64, 0x32, 0xe2, 0x02, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x3e,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x40,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4f, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x41, 0x70,
	0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x41, 0x70, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x6b, 0x65, 0x72, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x2d, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_manager_proto_rawDescOnce sync.Once
	file_account_manager_proto_rawDescData = file_account_manager_proto_rawDesc
)

func file_account_manager_proto_rawDescGZIP() []byte {
	file_account_manager_proto_rawDescOnce.Do(func() {
		file_account_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_manager_proto_rawDescData)
	})
	return file_account_manager_proto_rawDescData
}

var file_account_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_account_manager_proto_goTypes = []any{
	(*RegisterRequest)(nil),      // 0: managerpb.RegisterRequest
	(*GetInfoRequest)(nil),       // 1: managerpb.GetInfoRequest
	(*GetInfoResponse)(nil),      // 2: managerpb.GetInfoResponse
	(*GetApiTokenRequest)(nil),   // 3: managerpb.GetApiTokenRequest
	(*GetApiTokenResponse)(nil),  // 4: managerpb.GetApiTokenResponse
	(*AuthRequest)(nil),          // 5: managerpb.AuthRequest
	(*AuthApiTokenRequest)(nil),  // 6: managerpb.AuthApiTokenRequest
	(*AuthApiTokenResponse)(nil), // 7: managerpb.AuthApiTokenResponse
	(*emptypb.Empty)(nil),        // 8: google.protobuf.Empty
}
var file_account_manager_proto_depIdxs = []int32{
	0, // 0: managerpb.Manager.Register:input_type -> managerpb.RegisterRequest
	1, // 1: managerpb.Manager.GetInfo:input_type -> managerpb.GetInfoRequest
	3, // 2: managerpb.Manager.GetApiToken:input_type -> managerpb.GetApiTokenRequest
	5, // 3: managerpb.Manager.Auth:input_type -> managerpb.AuthRequest
	6, // 4: managerpb.Manager.AuthApiToken:input_type -> managerpb.AuthApiTokenRequest
	8, // 5: managerpb.Manager.Register:output_type -> google.protobuf.Empty
	2, // 6: managerpb.Manager.GetInfo:output_type -> managerpb.GetInfoResponse
	4, // 7: managerpb.Manager.GetApiToken:output_type -> managerpb.GetApiTokenResponse
	8, // 8: managerpb.Manager.Auth:output_type -> google.protobuf.Empty
	7, // 9: managerpb.Manager.AuthApiToken:output_type -> managerpb.AuthApiTokenResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_account_manager_proto_init() }
func file_account_manager_proto_init() {
	if File_account_manager_proto != nil {
		return
	}
	file_account_manager_proto_msgTypes[0].OneofWrappers = []any{}
	file_account_manager_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_account_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_manager_proto_goTypes,
		DependencyIndexes: file_account_manager_proto_depIdxs,
		MessageInfos:      file_account_manager_proto_msgTypes,
	}.Build()
	File_account_manager_proto = out.File
	file_account_manager_proto_rawDesc = nil
	file_account_manager_proto_goTypes = nil
	file_account_manager_proto_depIdxs = nil
}