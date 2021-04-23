// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.2
// source: services/authority/proto/authority.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/mercari/go-conference-2021-spring-office-hour/services/customer/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_authority_proto_authority_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_authority_proto_authority_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_services_authority_proto_authority_proto_rawDescGZIP(), []int{0}
}

func (x *SignupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *proto1.Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *SignupResponse) Reset() {
	*x = SignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_authority_proto_authority_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupResponse) ProtoMessage() {}

func (x *SignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_authority_proto_authority_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupResponse.ProtoReflect.Descriptor instead.
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return file_services_authority_proto_authority_proto_rawDescGZIP(), []int{1}
}

func (x *SignupResponse) GetCustomer() *proto1.Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type SigninRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SigninRequest) Reset() {
	*x = SigninRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_authority_proto_authority_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigninRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninRequest) ProtoMessage() {}

func (x *SigninRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_authority_proto_authority_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninRequest.ProtoReflect.Descriptor instead.
func (*SigninRequest) Descriptor() ([]byte, []int) {
	return file_services_authority_proto_authority_proto_rawDescGZIP(), []int{2}
}

func (x *SigninRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SigninResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *SigninResponse) Reset() {
	*x = SigninResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_authority_proto_authority_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigninResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigninResponse) ProtoMessage() {}

func (x *SigninResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_authority_proto_authority_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigninResponse.ProtoReflect.Descriptor instead.
func (*SigninResponse) Descriptor() ([]byte, []int) {
	return file_services_authority_proto_authority_proto_rawDescGZIP(), []int{3}
}

func (x *SigninResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type ListPublicKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPublicKeysRequest) Reset() {
	*x = ListPublicKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_authority_proto_authority_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPublicKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublicKeysRequest) ProtoMessage() {}

func (x *ListPublicKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_authority_proto_authority_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublicKeysRequest.ProtoReflect.Descriptor instead.
func (*ListPublicKeysRequest) Descriptor() ([]byte, []int) {
	return file_services_authority_proto_authority_proto_rawDescGZIP(), []int{4}
}

type ListPublicKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwks string `protobuf:"bytes,1,opt,name=jwks,proto3" json:"jwks,omitempty"`
}

func (x *ListPublicKeysResponse) Reset() {
	*x = ListPublicKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_authority_proto_authority_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPublicKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublicKeysResponse) ProtoMessage() {}

func (x *ListPublicKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_authority_proto_authority_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublicKeysResponse.ProtoReflect.Descriptor instead.
func (*ListPublicKeysResponse) Descriptor() ([]byte, []int) {
	return file_services_authority_proto_authority_proto_rawDescGZIP(), []int{5}
}

func (x *ListPublicKeysResponse) GetJwks() string {
	if x != nil {
		return x.Jwks
	}
	return ""
}

var File_services_authority_proto_authority_proto protoreflect.FileDescriptor

var file_services_authority_proto_authority_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x6d, 0x65, 0x72, 0x63,
	0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x1a, 0x26, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0d, 0x53,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x6e, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5c, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67,
	0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32,
	0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f,
	0x68, 0x6f, 0x75, 0x72, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x22, 0x23, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6a, 0x77, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x77, 0x6b,
	0x73, 0x32, 0xfe, 0x03, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x12, 0x46, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73,
	0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75,
	0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x12, 0x46, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e,
	0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30,
	0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xb1,
	0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x73, 0x12, 0x4e, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73,
	0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75,
	0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x4f, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73,
	0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75,
	0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2d, 0x32, 0x30, 0x32, 0x31, 0x2d, 0x73, 0x70, 0x72, 0x69,
	0x6e, 0x67, 0x2d, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x2d, 0x68, 0x6f, 0x75, 0x72, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_authority_proto_authority_proto_rawDescOnce sync.Once
	file_services_authority_proto_authority_proto_rawDescData = file_services_authority_proto_authority_proto_rawDesc
)

func file_services_authority_proto_authority_proto_rawDescGZIP() []byte {
	file_services_authority_proto_authority_proto_rawDescOnce.Do(func() {
		file_services_authority_proto_authority_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_authority_proto_authority_proto_rawDescData)
	})
	return file_services_authority_proto_authority_proto_rawDescData
}

var file_services_authority_proto_authority_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_authority_proto_authority_proto_goTypes = []interface{}{
	(*SignupRequest)(nil),          // 0: mercari.go_conference_2021_spring_office_hour.authority.SignupRequest
	(*SignupResponse)(nil),         // 1: mercari.go_conference_2021_spring_office_hour.authority.SignupResponse
	(*SigninRequest)(nil),          // 2: mercari.go_conference_2021_spring_office_hour.authority.SigninRequest
	(*SigninResponse)(nil),         // 3: mercari.go_conference_2021_spring_office_hour.authority.SigninResponse
	(*ListPublicKeysRequest)(nil),  // 4: mercari.go_conference_2021_spring_office_hour.authority.ListPublicKeysRequest
	(*ListPublicKeysResponse)(nil), // 5: mercari.go_conference_2021_spring_office_hour.authority.ListPublicKeysResponse
	(*proto1.Customer)(nil),        // 6: mercari.go_conference_2021_spring_office_hour.customer.Customer
}
var file_services_authority_proto_authority_proto_depIdxs = []int32{
	6, // 0: mercari.go_conference_2021_spring_office_hour.authority.SignupResponse.customer:type_name -> mercari.go_conference_2021_spring_office_hour.customer.Customer
	0, // 1: mercari.go_conference_2021_spring_office_hour.authority.AuthorityService.Signup:input_type -> mercari.go_conference_2021_spring_office_hour.authority.SignupRequest
	2, // 2: mercari.go_conference_2021_spring_office_hour.authority.AuthorityService.Signin:input_type -> mercari.go_conference_2021_spring_office_hour.authority.SigninRequest
	4, // 3: mercari.go_conference_2021_spring_office_hour.authority.AuthorityService.ListPublicKeys:input_type -> mercari.go_conference_2021_spring_office_hour.authority.ListPublicKeysRequest
	1, // 4: mercari.go_conference_2021_spring_office_hour.authority.AuthorityService.Signup:output_type -> mercari.go_conference_2021_spring_office_hour.authority.SignupResponse
	3, // 5: mercari.go_conference_2021_spring_office_hour.authority.AuthorityService.Signin:output_type -> mercari.go_conference_2021_spring_office_hour.authority.SigninResponse
	5, // 6: mercari.go_conference_2021_spring_office_hour.authority.AuthorityService.ListPublicKeys:output_type -> mercari.go_conference_2021_spring_office_hour.authority.ListPublicKeysResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_authority_proto_authority_proto_init() }
func file_services_authority_proto_authority_proto_init() {
	if File_services_authority_proto_authority_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_authority_proto_authority_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupRequest); i {
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
		file_services_authority_proto_authority_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupResponse); i {
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
		file_services_authority_proto_authority_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigninRequest); i {
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
		file_services_authority_proto_authority_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigninResponse); i {
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
		file_services_authority_proto_authority_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPublicKeysRequest); i {
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
		file_services_authority_proto_authority_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPublicKeysResponse); i {
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
			RawDescriptor: file_services_authority_proto_authority_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_authority_proto_authority_proto_goTypes,
		DependencyIndexes: file_services_authority_proto_authority_proto_depIdxs,
		MessageInfos:      file_services_authority_proto_authority_proto_msgTypes,
	}.Build()
	File_services_authority_proto_authority_proto = out.File
	file_services_authority_proto_authority_proto_rawDesc = nil
	file_services_authority_proto_authority_proto_goTypes = nil
	file_services_authority_proto_authority_proto_depIdxs = nil
}
