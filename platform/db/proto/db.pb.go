// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.2
// source: platform/db/proto/db.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

type CreateCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateCustomerRequest) Reset() {
	*x = CreateCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_db_proto_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerRequest) ProtoMessage() {}

func (x *CreateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_db_proto_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_platform_db_proto_db_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCustomerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *CreateCustomerResponse) Reset() {
	*x = CreateCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_db_proto_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerResponse) ProtoMessage() {}

func (x *CreateCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_platform_db_proto_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerResponse.ProtoReflect.Descriptor instead.
func (*CreateCustomerResponse) Descriptor() ([]byte, []int) {
	return file_platform_db_proto_db_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCustomerResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type GetCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCustomerRequest) Reset() {
	*x = GetCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_db_proto_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerRequest) ProtoMessage() {}

func (x *GetCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_db_proto_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return file_platform_db_proto_db_proto_rawDescGZIP(), []int{2}
}

func (x *GetCustomerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *GetCustomerResponse) Reset() {
	*x = GetCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_db_proto_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerResponse) ProtoMessage() {}

func (x *GetCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_platform_db_proto_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerResponse) Descriptor() ([]byte, []int) {
	return file_platform_db_proto_db_proto_rawDescGZIP(), []int{3}
}

func (x *GetCustomerResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type GetCustomerByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCustomerByNameRequest) Reset() {
	*x = GetCustomerByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_db_proto_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByNameRequest) ProtoMessage() {}

func (x *GetCustomerByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_db_proto_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByNameRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerByNameRequest) Descriptor() ([]byte, []int) {
	return file_platform_db_proto_db_proto_rawDescGZIP(), []int{4}
}

func (x *GetCustomerByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCustomerByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *GetCustomerByNameResponse) Reset() {
	*x = GetCustomerByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_db_proto_db_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByNameResponse) ProtoMessage() {}

func (x *GetCustomerByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_platform_db_proto_db_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByNameResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerByNameResponse) Descriptor() ([]byte, []int) {
	return file_platform_db_proto_db_proto_rawDescGZIP(), []int{5}
}

func (x *GetCustomerByNameResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type GetItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetItemRequest) Reset() {
	*x = GetItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_db_proto_db_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemRequest) ProtoMessage() {}

func (x *GetItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_db_proto_db_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemRequest.ProtoReflect.Descriptor instead.
func (*GetItemRequest) Descriptor() ([]byte, []int) {
	return file_platform_db_proto_db_proto_rawDescGZIP(), []int{6}
}

func (x *GetItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetItemResponse) Reset() {
	*x = GetItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_db_proto_db_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemResponse) ProtoMessage() {}

func (x *GetItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_platform_db_proto_db_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemResponse.ProtoReflect.Descriptor instead.
func (*GetItemResponse) Descriptor() ([]byte, []int) {
	return file_platform_db_proto_db_proto_rawDescGZIP(), []int{7}
}

func (x *GetItemResponse) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_db_proto_db_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_platform_db_proto_db_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_platform_db_proto_db_proto_rawDescGZIP(), []int{8}
}

func (x *Customer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Title      string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Price      int64  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_db_proto_db_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_platform_db_proto_db_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_platform_db_proto_db_proto_rawDescGZIP(), []int{9}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Item) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_platform_db_proto_db_proto protoreflect.FileDescriptor

var file_platform_db_proto_db_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x64, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x6d, 0x65,
	0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x64, 0x62, 0x22, 0x2b,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x70, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72,
	0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x24, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x08, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6d,
	0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x64, 0x62, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x22, 0x2e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x73, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x56, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73,
	0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75,
	0x72, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x64, 0x62, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x2e, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0x8e, 0x05,
	0x0a, 0x09, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa3, 0x01, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x47,
	0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x64,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72,
	0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x64, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x9a, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x44, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73,
	0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75,
	0x72, 0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72,
	0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xac,
	0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67,
	0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32,
	0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f,
	0x68, 0x6f, 0x75, 0x72, 0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x4b, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72,
	0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8e, 0x01,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x40, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x6d, 0x65,
	0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x64, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4c,
	0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x72,
	0x63, 0x61, 0x72, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x2d, 0x32, 0x30, 0x32, 0x31, 0x2d, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x2d, 0x6f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x2d, 0x68, 0x6f, 0x75, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_platform_db_proto_db_proto_rawDescOnce sync.Once
	file_platform_db_proto_db_proto_rawDescData = file_platform_db_proto_db_proto_rawDesc
)

func file_platform_db_proto_db_proto_rawDescGZIP() []byte {
	file_platform_db_proto_db_proto_rawDescOnce.Do(func() {
		file_platform_db_proto_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_db_proto_db_proto_rawDescData)
	})
	return file_platform_db_proto_db_proto_rawDescData
}

var file_platform_db_proto_db_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_platform_db_proto_db_proto_goTypes = []interface{}{
	(*CreateCustomerRequest)(nil),     // 0: mercari.go_conference_2021_spring_office_hour.db.CreateCustomerRequest
	(*CreateCustomerResponse)(nil),    // 1: mercari.go_conference_2021_spring_office_hour.db.CreateCustomerResponse
	(*GetCustomerRequest)(nil),        // 2: mercari.go_conference_2021_spring_office_hour.db.GetCustomerRequest
	(*GetCustomerResponse)(nil),       // 3: mercari.go_conference_2021_spring_office_hour.db.GetCustomerResponse
	(*GetCustomerByNameRequest)(nil),  // 4: mercari.go_conference_2021_spring_office_hour.db.GetCustomerByNameRequest
	(*GetCustomerByNameResponse)(nil), // 5: mercari.go_conference_2021_spring_office_hour.db.GetCustomerByNameResponse
	(*GetItemRequest)(nil),            // 6: mercari.go_conference_2021_spring_office_hour.db.GetItemRequest
	(*GetItemResponse)(nil),           // 7: mercari.go_conference_2021_spring_office_hour.db.GetItemResponse
	(*Customer)(nil),                  // 8: mercari.go_conference_2021_spring_office_hour.db.Customer
	(*Item)(nil),                      // 9: mercari.go_conference_2021_spring_office_hour.db.Item
}
var file_platform_db_proto_db_proto_depIdxs = []int32{
	8, // 0: mercari.go_conference_2021_spring_office_hour.db.CreateCustomerResponse.customer:type_name -> mercari.go_conference_2021_spring_office_hour.db.Customer
	8, // 1: mercari.go_conference_2021_spring_office_hour.db.GetCustomerResponse.customer:type_name -> mercari.go_conference_2021_spring_office_hour.db.Customer
	8, // 2: mercari.go_conference_2021_spring_office_hour.db.GetCustomerByNameResponse.customer:type_name -> mercari.go_conference_2021_spring_office_hour.db.Customer
	9, // 3: mercari.go_conference_2021_spring_office_hour.db.GetItemResponse.item:type_name -> mercari.go_conference_2021_spring_office_hour.db.Item
	0, // 4: mercari.go_conference_2021_spring_office_hour.db.DBService.CreateCustomer:input_type -> mercari.go_conference_2021_spring_office_hour.db.CreateCustomerRequest
	2, // 5: mercari.go_conference_2021_spring_office_hour.db.DBService.GetCustomer:input_type -> mercari.go_conference_2021_spring_office_hour.db.GetCustomerRequest
	4, // 6: mercari.go_conference_2021_spring_office_hour.db.DBService.GetCustomerByName:input_type -> mercari.go_conference_2021_spring_office_hour.db.GetCustomerByNameRequest
	6, // 7: mercari.go_conference_2021_spring_office_hour.db.DBService.GetItem:input_type -> mercari.go_conference_2021_spring_office_hour.db.GetItemRequest
	1, // 8: mercari.go_conference_2021_spring_office_hour.db.DBService.CreateCustomer:output_type -> mercari.go_conference_2021_spring_office_hour.db.CreateCustomerResponse
	3, // 9: mercari.go_conference_2021_spring_office_hour.db.DBService.GetCustomer:output_type -> mercari.go_conference_2021_spring_office_hour.db.GetCustomerResponse
	5, // 10: mercari.go_conference_2021_spring_office_hour.db.DBService.GetCustomerByName:output_type -> mercari.go_conference_2021_spring_office_hour.db.GetCustomerByNameResponse
	7, // 11: mercari.go_conference_2021_spring_office_hour.db.DBService.GetItem:output_type -> mercari.go_conference_2021_spring_office_hour.db.GetItemResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_platform_db_proto_db_proto_init() }
func file_platform_db_proto_db_proto_init() {
	if File_platform_db_proto_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_platform_db_proto_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerRequest); i {
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
		file_platform_db_proto_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerResponse); i {
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
		file_platform_db_proto_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerRequest); i {
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
		file_platform_db_proto_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerResponse); i {
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
		file_platform_db_proto_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerByNameRequest); i {
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
		file_platform_db_proto_db_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerByNameResponse); i {
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
		file_platform_db_proto_db_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemRequest); i {
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
		file_platform_db_proto_db_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemResponse); i {
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
		file_platform_db_proto_db_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_platform_db_proto_db_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
			RawDescriptor: file_platform_db_proto_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_platform_db_proto_db_proto_goTypes,
		DependencyIndexes: file_platform_db_proto_db_proto_depIdxs,
		MessageInfos:      file_platform_db_proto_db_proto_msgTypes,
	}.Build()
	File_platform_db_proto_db_proto = out.File
	file_platform_db_proto_db_proto_rawDesc = nil
	file_platform_db_proto_db_proto_goTypes = nil
	file_platform_db_proto_db_proto_depIdxs = nil
}
