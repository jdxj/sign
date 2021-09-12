// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: secret/secret.proto

package secret

import (
	crontab "github.com/jdxj/sign/internal/proto/crontab"
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

type CreateSecretReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64          `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Domain crontab.Domain `protobuf:"varint,2,opt,name=Domain,proto3,enum=Domain" json:"Domain,omitempty"`
	Key    string         `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"` // todo: 需要加密
}

func (x *CreateSecretReq) Reset() {
	*x = CreateSecretReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_secret_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSecretReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSecretReq) ProtoMessage() {}

func (x *CreateSecretReq) ProtoReflect() protoreflect.Message {
	mi := &file_secret_secret_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSecretReq.ProtoReflect.Descriptor instead.
func (*CreateSecretReq) Descriptor() ([]byte, []int) {
	return file_secret_secret_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSecretReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateSecretReq) GetDomain() crontab.Domain {
	if x != nil {
		return x.Domain
	}
	return crontab.Domain_UnknownDomain
}

func (x *CreateSecretReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type CreateSecretRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretID int64 `protobuf:"varint,1,opt,name=SecretID,proto3" json:"SecretID,omitempty"`
}

func (x *CreateSecretRsp) Reset() {
	*x = CreateSecretRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_secret_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSecretRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSecretRsp) ProtoMessage() {}

func (x *CreateSecretRsp) ProtoReflect() protoreflect.Message {
	mi := &file_secret_secret_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSecretRsp.ProtoReflect.Descriptor instead.
func (*CreateSecretRsp) Descriptor() ([]byte, []int) {
	return file_secret_secret_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSecretRsp) GetSecretID() int64 {
	if x != nil {
		return x.SecretID
	}
	return 0
}

type GetSecretReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetSecretReq) Reset() {
	*x = GetSecretReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_secret_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretReq) ProtoMessage() {}

func (x *GetSecretReq) ProtoReflect() protoreflect.Message {
	mi := &file_secret_secret_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretReq.ProtoReflect.Descriptor instead.
func (*GetSecretReq) Descriptor() ([]byte, []int) {
	return file_secret_secret_proto_rawDescGZIP(), []int{2}
}

func (x *GetSecretReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type SecretRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretID int64          `protobuf:"varint,1,opt,name=SecretID,proto3" json:"SecretID,omitempty"`
	Domain   crontab.Domain `protobuf:"varint,2,opt,name=Domain,proto3,enum=Domain" json:"Domain,omitempty"`
	Key      string         `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"` // todo: 需要解密
}

func (x *SecretRecord) Reset() {
	*x = SecretRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_secret_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretRecord) ProtoMessage() {}

func (x *SecretRecord) ProtoReflect() protoreflect.Message {
	mi := &file_secret_secret_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretRecord.ProtoReflect.Descriptor instead.
func (*SecretRecord) Descriptor() ([]byte, []int) {
	return file_secret_secret_proto_rawDescGZIP(), []int{3}
}

func (x *SecretRecord) GetSecretID() int64 {
	if x != nil {
		return x.SecretID
	}
	return 0
}

func (x *SecretRecord) GetDomain() crontab.Domain {
	if x != nil {
		return x.Domain
	}
	return crontab.Domain_UnknownDomain
}

func (x *SecretRecord) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetSecretRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*SecretRecord `protobuf:"bytes,1,rep,name=Records,proto3" json:"Records,omitempty"`
}

func (x *GetSecretRsp) Reset() {
	*x = GetSecretRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_secret_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretRsp) ProtoMessage() {}

func (x *GetSecretRsp) ProtoReflect() protoreflect.Message {
	mi := &file_secret_secret_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretRsp.ProtoReflect.Descriptor instead.
func (*GetSecretRsp) Descriptor() ([]byte, []int) {
	return file_secret_secret_proto_rawDescGZIP(), []int{4}
}

func (x *GetSecretRsp) GetRecords() []*SecretRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

type UpdateSecretReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretID int64          `protobuf:"varint,1,opt,name=SecretID,proto3" json:"SecretID,omitempty"`
	Domain   crontab.Domain `protobuf:"varint,2,opt,name=Domain,proto3,enum=Domain" json:"Domain,omitempty"`
	Key      string         `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *UpdateSecretReq) Reset() {
	*x = UpdateSecretReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_secret_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSecretReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSecretReq) ProtoMessage() {}

func (x *UpdateSecretReq) ProtoReflect() protoreflect.Message {
	mi := &file_secret_secret_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSecretReq.ProtoReflect.Descriptor instead.
func (*UpdateSecretReq) Descriptor() ([]byte, []int) {
	return file_secret_secret_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSecretReq) GetSecretID() int64 {
	if x != nil {
		return x.SecretID
	}
	return 0
}

func (x *UpdateSecretReq) GetDomain() crontab.Domain {
	if x != nil {
		return x.Domain
	}
	return crontab.Domain_UnknownDomain
}

func (x *UpdateSecretReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteSecretReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretID int64 `protobuf:"varint,1,opt,name=SecretID,proto3" json:"SecretID,omitempty"`
}

func (x *DeleteSecretReq) Reset() {
	*x = DeleteSecretReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secret_secret_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSecretReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSecretReq) ProtoMessage() {}

func (x *DeleteSecretReq) ProtoReflect() protoreflect.Message {
	mi := &file_secret_secret_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSecretReq.ProtoReflect.Descriptor instead.
func (*DeleteSecretReq) Descriptor() ([]byte, []int) {
	return file_secret_secret_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSecretReq) GetSecretID() int64 {
	if x != nil {
		return x.SecretID
	}
	return 0
}

var File_secret_secret_proto protoreflect.FileDescriptor

var file_secret_secret_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x72, 0x6f, 0x6e, 0x74, 0x61, 0x62, 0x2f, 0x63,
	0x72, 0x6f, 0x6e, 0x74, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x06, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x22, 0x2d, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x49, 0x44, 0x22, 0x26, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x5d,
	0x0a, 0x0c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x06, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x52, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x22, 0x37, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x73, 0x70, 0x12, 0x27, 0x0a,
	0x07, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x60, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x06,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x22, 0x2d, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x44, 0x32, 0xe2, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x73, 0x70, 0x12, 0x29, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x38, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x10, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x64, 0x78, 0x6a, 0x2f,
	0x73, 0x69, 0x67, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_secret_secret_proto_rawDescOnce sync.Once
	file_secret_secret_proto_rawDescData = file_secret_secret_proto_rawDesc
)

func file_secret_secret_proto_rawDescGZIP() []byte {
	file_secret_secret_proto_rawDescOnce.Do(func() {
		file_secret_secret_proto_rawDescData = protoimpl.X.CompressGZIP(file_secret_secret_proto_rawDescData)
	})
	return file_secret_secret_proto_rawDescData
}

var file_secret_secret_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_secret_secret_proto_goTypes = []interface{}{
	(*CreateSecretReq)(nil), // 0: CreateSecretReq
	(*CreateSecretRsp)(nil), // 1: CreateSecretRsp
	(*GetSecretReq)(nil),    // 2: GetSecretReq
	(*SecretRecord)(nil),    // 3: SecretRecord
	(*GetSecretRsp)(nil),    // 4: GetSecretRsp
	(*UpdateSecretReq)(nil), // 5: UpdateSecretReq
	(*DeleteSecretReq)(nil), // 6: DeleteSecretReq
	(crontab.Domain)(0),     // 7: Domain
	(*emptypb.Empty)(nil),   // 8: google.protobuf.Empty
}
var file_secret_secret_proto_depIdxs = []int32{
	7, // 0: CreateSecretReq.Domain:type_name -> Domain
	7, // 1: SecretRecord.Domain:type_name -> Domain
	3, // 2: GetSecretRsp.Records:type_name -> SecretRecord
	7, // 3: UpdateSecretReq.Domain:type_name -> Domain
	0, // 4: SecretService.CreateSecret:input_type -> CreateSecretReq
	2, // 5: SecretService.GetSecret:input_type -> GetSecretReq
	5, // 6: SecretService.UpdateSecret:input_type -> UpdateSecretReq
	6, // 7: SecretService.DeleteSecret:input_type -> DeleteSecretReq
	1, // 8: SecretService.CreateSecret:output_type -> CreateSecretRsp
	4, // 9: SecretService.GetSecret:output_type -> GetSecretRsp
	8, // 10: SecretService.UpdateSecret:output_type -> google.protobuf.Empty
	8, // 11: SecretService.DeleteSecret:output_type -> google.protobuf.Empty
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_secret_secret_proto_init() }
func file_secret_secret_proto_init() {
	if File_secret_secret_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_secret_secret_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSecretReq); i {
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
		file_secret_secret_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSecretRsp); i {
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
		file_secret_secret_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretReq); i {
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
		file_secret_secret_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretRecord); i {
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
		file_secret_secret_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretRsp); i {
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
		file_secret_secret_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSecretReq); i {
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
		file_secret_secret_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSecretReq); i {
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
			RawDescriptor: file_secret_secret_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_secret_secret_proto_goTypes,
		DependencyIndexes: file_secret_secret_proto_depIdxs,
		MessageInfos:      file_secret_secret_proto_msgTypes,
	}.Build()
	File_secret_secret_proto = out.File
	file_secret_secret_proto_rawDesc = nil
	file_secret_secret_proto_goTypes = nil
	file_secret_secret_proto_depIdxs = nil
}
