// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: services/catalog/proto/catalog.proto

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

type CreateItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Price int64  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CreateItemRequest) Reset() {
	*x = CreateItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_catalog_proto_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItemRequest) ProtoMessage() {}

func (x *CreateItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_catalog_proto_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItemRequest.ProtoReflect.Descriptor instead.
func (*CreateItemRequest) Descriptor() ([]byte, []int) {
	return file_services_catalog_proto_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *CreateItemRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateItemRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreateItemResponse) Reset() {
	*x = CreateItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_catalog_proto_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItemResponse) ProtoMessage() {}

func (x *CreateItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_catalog_proto_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItemResponse.ProtoReflect.Descriptor instead.
func (*CreateItemResponse) Descriptor() ([]byte, []int) {
	return file_services_catalog_proto_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *CreateItemResponse) GetItem() *Item {
	if x != nil {
		return x.Item
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
		mi := &file_services_catalog_proto_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemRequest) ProtoMessage() {}

func (x *GetItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_catalog_proto_catalog_proto_msgTypes[2]
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
	return file_services_catalog_proto_catalog_proto_rawDescGZIP(), []int{2}
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
		mi := &file_services_catalog_proto_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemResponse) ProtoMessage() {}

func (x *GetItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_catalog_proto_catalog_proto_msgTypes[3]
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
	return file_services_catalog_proto_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *GetItemResponse) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type ListItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ListItemsRequest) Reset() {
	*x = ListItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_catalog_proto_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsRequest) ProtoMessage() {}

func (x *ListItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_catalog_proto_catalog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsRequest.ProtoReflect.Descriptor instead.
func (*ListItemsRequest) Descriptor() ([]byte, []int) {
	return file_services_catalog_proto_catalog_proto_rawDescGZIP(), []int{4}
}

func (x *ListItemsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListItemsResponse) Reset() {
	*x = ListItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_catalog_proto_catalog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsResponse) ProtoMessage() {}

func (x *ListItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_catalog_proto_catalog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsResponse.ProtoReflect.Descriptor instead.
func (*ListItemsResponse) Descriptor() ([]byte, []int) {
	return file_services_catalog_proto_catalog_proto_rawDescGZIP(), []int{5}
}

func (x *ListItemsResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId   string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CustomerName string `protobuf:"bytes,3,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	Title        string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Price        int64  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_catalog_proto_catalog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_services_catalog_proto_catalog_proto_msgTypes[6]
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
	return file_services_catalog_proto_catalog_proto_rawDescGZIP(), []int{6}
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

func (x *Item) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
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

var File_services_catalog_proto_catalog_proto protoreflect.FileDescriptor

var file_services_catalog_proto_catalog_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e,
	0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30,
	0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x22, 0x3f, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x65,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f,
	0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f,
	0x75, 0x72, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61,
	0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x22, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x66, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31,
	0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68,
	0x6f, 0x75, 0x72, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x32, 0xf0, 0x03, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa1, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x48, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67,
	0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32,
	0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f,
	0x68, 0x6f, 0x75, 0x72, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x49,
	0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x45, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e,
	0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30,
	0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x6d,
	0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x47, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f,
	0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f,
	0x75, 0x72, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x6d, 0x65,
	0x72, 0x63, 0x61, 0x72, 0x69, 0x2e, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x61, 0x72, 0x69, 0x2f, 0x6d, 0x65, 0x72, 0x63,
	0x61, 0x72, 0x69, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_catalog_proto_catalog_proto_rawDescOnce sync.Once
	file_services_catalog_proto_catalog_proto_rawDescData = file_services_catalog_proto_catalog_proto_rawDesc
)

func file_services_catalog_proto_catalog_proto_rawDescGZIP() []byte {
	file_services_catalog_proto_catalog_proto_rawDescOnce.Do(func() {
		file_services_catalog_proto_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_catalog_proto_catalog_proto_rawDescData)
	})
	return file_services_catalog_proto_catalog_proto_rawDescData
}

var file_services_catalog_proto_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_services_catalog_proto_catalog_proto_goTypes = []interface{}{
	(*CreateItemRequest)(nil),  // 0: mercari.go_conference_2021_spring_office_hour.catalog.CreateItemRequest
	(*CreateItemResponse)(nil), // 1: mercari.go_conference_2021_spring_office_hour.catalog.CreateItemResponse
	(*GetItemRequest)(nil),     // 2: mercari.go_conference_2021_spring_office_hour.catalog.GetItemRequest
	(*GetItemResponse)(nil),    // 3: mercari.go_conference_2021_spring_office_hour.catalog.GetItemResponse
	(*ListItemsRequest)(nil),   // 4: mercari.go_conference_2021_spring_office_hour.catalog.ListItemsRequest
	(*ListItemsResponse)(nil),  // 5: mercari.go_conference_2021_spring_office_hour.catalog.ListItemsResponse
	(*Item)(nil),               // 6: mercari.go_conference_2021_spring_office_hour.catalog.Item
}
var file_services_catalog_proto_catalog_proto_depIdxs = []int32{
	6, // 0: mercari.go_conference_2021_spring_office_hour.catalog.CreateItemResponse.item:type_name -> mercari.go_conference_2021_spring_office_hour.catalog.Item
	6, // 1: mercari.go_conference_2021_spring_office_hour.catalog.GetItemResponse.item:type_name -> mercari.go_conference_2021_spring_office_hour.catalog.Item
	6, // 2: mercari.go_conference_2021_spring_office_hour.catalog.ListItemsResponse.items:type_name -> mercari.go_conference_2021_spring_office_hour.catalog.Item
	0, // 3: mercari.go_conference_2021_spring_office_hour.catalog.CatalogService.CreateItem:input_type -> mercari.go_conference_2021_spring_office_hour.catalog.CreateItemRequest
	2, // 4: mercari.go_conference_2021_spring_office_hour.catalog.CatalogService.GetItem:input_type -> mercari.go_conference_2021_spring_office_hour.catalog.GetItemRequest
	4, // 5: mercari.go_conference_2021_spring_office_hour.catalog.CatalogService.ListItems:input_type -> mercari.go_conference_2021_spring_office_hour.catalog.ListItemsRequest
	1, // 6: mercari.go_conference_2021_spring_office_hour.catalog.CatalogService.CreateItem:output_type -> mercari.go_conference_2021_spring_office_hour.catalog.CreateItemResponse
	3, // 7: mercari.go_conference_2021_spring_office_hour.catalog.CatalogService.GetItem:output_type -> mercari.go_conference_2021_spring_office_hour.catalog.GetItemResponse
	5, // 8: mercari.go_conference_2021_spring_office_hour.catalog.CatalogService.ListItems:output_type -> mercari.go_conference_2021_spring_office_hour.catalog.ListItemsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_services_catalog_proto_catalog_proto_init() }
func file_services_catalog_proto_catalog_proto_init() {
	if File_services_catalog_proto_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_catalog_proto_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateItemRequest); i {
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
		file_services_catalog_proto_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateItemResponse); i {
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
		file_services_catalog_proto_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_catalog_proto_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_catalog_proto_catalog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsRequest); i {
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
		file_services_catalog_proto_catalog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsResponse); i {
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
		file_services_catalog_proto_catalog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_services_catalog_proto_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_catalog_proto_catalog_proto_goTypes,
		DependencyIndexes: file_services_catalog_proto_catalog_proto_depIdxs,
		MessageInfos:      file_services_catalog_proto_catalog_proto_msgTypes,
	}.Build()
	File_services_catalog_proto_catalog_proto = out.File
	file_services_catalog_proto_catalog_proto_rawDesc = nil
	file_services_catalog_proto_catalog_proto_goTypes = nil
	file_services_catalog_proto_catalog_proto_depIdxs = nil
}
