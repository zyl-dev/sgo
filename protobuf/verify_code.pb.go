// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: protobuf/verify_code.proto

package protobuf

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

type VerifyCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier     string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Receiver       string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	ServiceName    string `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Action         string `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	Channel        string `protobuf:"bytes,5,opt,name=channel,proto3" json:"channel,omitempty"`
	Code           string `protobuf:"bytes,6,opt,name=code,proto3" json:"code,omitempty"`
	VerifyStatus   int32  `protobuf:"varint,7,opt,name=verify_status,json=verifyStatus,proto3" json:"verify_status,omitempty"`
	VerifyTimes    int32  `protobuf:"varint,8,opt,name=verify_times,json=verifyTimes,proto3" json:"verify_times,omitempty"`
	MaxVerifyTimes int32  `protobuf:"varint,9,opt,name=max_verify_times,json=maxVerifyTimes,proto3" json:"max_verify_times,omitempty"`
	Ttl            int32  `protobuf:"varint,10,opt,name=ttl,proto3" json:"ttl,omitempty"`
	TplId          int32  `protobuf:"varint,11,opt,name=tpl_id,json=tplId,proto3" json:"tpl_id,omitempty"`
}

func (x *VerifyCode) Reset() {
	*x = VerifyCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_verify_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCode) ProtoMessage() {}

func (x *VerifyCode) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_verify_code_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCode.ProtoReflect.Descriptor instead.
func (*VerifyCode) Descriptor() ([]byte, []int) {
	return file_protobuf_verify_code_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyCode) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *VerifyCode) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *VerifyCode) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *VerifyCode) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *VerifyCode) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *VerifyCode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *VerifyCode) GetVerifyStatus() int32 {
	if x != nil {
		return x.VerifyStatus
	}
	return 0
}

func (x *VerifyCode) GetVerifyTimes() int32 {
	if x != nil {
		return x.VerifyTimes
	}
	return 0
}

func (x *VerifyCode) GetMaxVerifyTimes() int32 {
	if x != nil {
		return x.MaxVerifyTimes
	}
	return 0
}

func (x *VerifyCode) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *VerifyCode) GetTplId() int32 {
	if x != nil {
		return x.TplId
	}
	return 0
}

type CreateCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"receiver" binding:"required"
	Receiver string `protobuf:"bytes,1,opt,name=receiver,proto3" json:"receiver" binding:"required"`
	// @inject_tag: json:"service_name" binding:"required"
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name" binding:"required"`
	Action      string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	// @inject_tag: json:"channel" binding:"required"
	Channel        string `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel" binding:"required"`
	MaxVerifyTimes int32  `protobuf:"varint,5,opt,name=max_verify_times,json=maxVerifyTimes,proto3" json:"max_verify_times,omitempty"`
	Ttl            int32  `protobuf:"varint,6,opt,name=ttl,proto3" json:"ttl,omitempty"`
	TplId          int32  `protobuf:"varint,7,opt,name=tpl_id,json=tplId,proto3" json:"tpl_id,omitempty"`
}

func (x *CreateCodeRequest) Reset() {
	*x = CreateCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_verify_code_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCodeRequest) ProtoMessage() {}

func (x *CreateCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_verify_code_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCodeRequest.ProtoReflect.Descriptor instead.
func (*CreateCodeRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_verify_code_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCodeRequest) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *CreateCodeRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *CreateCodeRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *CreateCodeRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *CreateCodeRequest) GetMaxVerifyTimes() int32 {
	if x != nil {
		return x.MaxVerifyTimes
	}
	return 0
}

func (x *CreateCodeRequest) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *CreateCodeRequest) GetTplId() int32 {
	if x != nil {
		return x.TplId
	}
	return 0
}

type CreateCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" binding:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" binding:"required"`
}

func (x *CreateCodeResponse) Reset() {
	*x = CreateCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_verify_code_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCodeResponse) ProtoMessage() {}

func (x *CreateCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_verify_code_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCodeResponse.ProtoReflect.Descriptor instead.
func (*CreateCodeResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_verify_code_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCodeResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type VerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" binding:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" binding:"required"`
	// @inject_tag: json:"code" binding:"required"
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code" binding:"required"`
}

func (x *VerifyCodeRequest) Reset() {
	*x = VerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_verify_code_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCodeRequest) ProtoMessage() {}

func (x *VerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_verify_code_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*VerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_verify_code_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyCodeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VerifyCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type VerifyCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"max_verify_times"
	MaxVerifyTimes int32 `protobuf:"varint,1,opt,name=max_verify_times,json=maxVerifyTimes,proto3" json:"max_verify_times"`
	// @inject_tag: json:"verified_times"
	VerifiedTimes int32 `protobuf:"varint,2,opt,name=verified_times,json=verifiedTimes,proto3" json:"verified_times"`
	// @inject_tag: json:"expired_at"
	ExpiredAt int64 `protobuf:"varint,3,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at"`
}

func (x *VerifyCodeResponse) Reset() {
	*x = VerifyCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_verify_code_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCodeResponse) ProtoMessage() {}

func (x *VerifyCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_verify_code_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCodeResponse.ProtoReflect.Descriptor instead.
func (*VerifyCodeResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_verify_code_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyCodeResponse) GetMaxVerifyTimes() int32 {
	if x != nil {
		return x.MaxVerifyTimes
	}
	return 0
}

func (x *VerifyCodeResponse) GetVerifiedTimes() int32 {
	if x != nil {
		return x.VerifiedTimes
	}
	return 0
}

func (x *VerifyCodeResponse) GetExpiredAt() int64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

var File_protobuf_verify_code_proto protoreflect.FileDescriptor

var file_protobuf_verify_code_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0xcc, 0x02, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x61,
	0x78, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x74, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x15,
	0x0a, 0x06, 0x74, 0x70, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x74, 0x70, 0x6c, 0x49, 0x64, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x28, 0x0a, 0x10,
	0x6d, 0x61, 0x78, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x70, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x70, 0x6c, 0x49, 0x64, 0x22,
	0x24, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x84,
	0x01, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x6d, 0x61, 0x78, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x41, 0x74, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x74, 0x63, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_verify_code_proto_rawDescOnce sync.Once
	file_protobuf_verify_code_proto_rawDescData = file_protobuf_verify_code_proto_rawDesc
)

func file_protobuf_verify_code_proto_rawDescGZIP() []byte {
	file_protobuf_verify_code_proto_rawDescOnce.Do(func() {
		file_protobuf_verify_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_verify_code_proto_rawDescData)
	})
	return file_protobuf_verify_code_proto_rawDescData
}

var file_protobuf_verify_code_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protobuf_verify_code_proto_goTypes = []interface{}{
	(*VerifyCode)(nil),         // 0: protobuf.VerifyCode
	(*CreateCodeRequest)(nil),  // 1: protobuf.CreateCodeRequest
	(*CreateCodeResponse)(nil), // 2: protobuf.CreateCodeResponse
	(*VerifyCodeRequest)(nil),  // 3: protobuf.VerifyCodeRequest
	(*VerifyCodeResponse)(nil), // 4: protobuf.VerifyCodeResponse
}
var file_protobuf_verify_code_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_verify_code_proto_init() }
func file_protobuf_verify_code_proto_init() {
	if File_protobuf_verify_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_verify_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCode); i {
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
		file_protobuf_verify_code_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCodeRequest); i {
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
		file_protobuf_verify_code_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCodeResponse); i {
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
		file_protobuf_verify_code_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCodeRequest); i {
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
		file_protobuf_verify_code_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCodeResponse); i {
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
			RawDescriptor: file_protobuf_verify_code_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_verify_code_proto_goTypes,
		DependencyIndexes: file_protobuf_verify_code_proto_depIdxs,
		MessageInfos:      file_protobuf_verify_code_proto_msgTypes,
	}.Build()
	File_protobuf_verify_code_proto = out.File
	file_protobuf_verify_code_proto_rawDesc = nil
	file_protobuf_verify_code_proto_goTypes = nil
	file_protobuf_verify_code_proto_depIdxs = nil
}