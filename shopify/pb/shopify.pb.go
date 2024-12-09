// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: shopify.proto

package pb

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId  string  `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ShopId     string  `protobuf:"bytes,3,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	TotalPrice float32 `protobuf:"fixed32,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	OrderId    string  `protobuf:"bytes,5,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_shopify_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_shopify_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_shopify_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Order) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

func (x *Order) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type SyncOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopName string `protobuf:"bytes,1,opt,name=shop_name,json=shopName,proto3" json:"shop_name,omitempty"`
	Limit    int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	SinceId  string `protobuf:"bytes,3,opt,name=since_id,json=sinceId,proto3" json:"since_id,omitempty"`
	Token    string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *SyncOrdersRequest) Reset() {
	*x = SyncOrdersRequest{}
	mi := &file_shopify_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SyncOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncOrdersRequest) ProtoMessage() {}

func (x *SyncOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shopify_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncOrdersRequest.ProtoReflect.Descriptor instead.
func (*SyncOrdersRequest) Descriptor() ([]byte, []int) {
	return file_shopify_proto_rawDescGZIP(), []int{1}
}

func (x *SyncOrdersRequest) GetShopName() string {
	if x != nil {
		return x.ShopName
	}
	return ""
}

func (x *SyncOrdersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SyncOrdersRequest) GetSinceId() string {
	if x != nil {
		return x.SinceId
	}
	return ""
}

func (x *SyncOrdersRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SyncOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *SyncOrdersResponse) Reset() {
	*x = SyncOrdersResponse{}
	mi := &file_shopify_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SyncOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncOrdersResponse) ProtoMessage() {}

func (x *SyncOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shopify_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncOrdersResponse.ProtoReflect.Descriptor instead.
func (*SyncOrdersResponse) Descriptor() ([]byte, []int) {
	return file_shopify_proto_rawDescGZIP(), []int{2}
}

func (x *SyncOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type StoreTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopName  string `protobuf:"bytes,1,opt,name=shop_name,json=shopName,proto3" json:"shop_name,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *StoreTokenRequest) Reset() {
	*x = StoreTokenRequest{}
	mi := &file_shopify_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreTokenRequest) ProtoMessage() {}

func (x *StoreTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shopify_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreTokenRequest.ProtoReflect.Descriptor instead.
func (*StoreTokenRequest) Descriptor() ([]byte, []int) {
	return file_shopify_proto_rawDescGZIP(), []int{3}
}

func (x *StoreTokenRequest) GetShopName() string {
	if x != nil {
		return x.ShopName
	}
	return ""
}

func (x *StoreTokenRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *StoreTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type StoreTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoreTokenResponse) Reset() {
	*x = StoreTokenResponse{}
	mi := &file_shopify_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StoreTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreTokenResponse) ProtoMessage() {}

func (x *StoreTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shopify_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreTokenResponse.ProtoReflect.Descriptor instead.
func (*StoreTokenResponse) Descriptor() ([]byte, []int) {
	return file_shopify_proto_rawDescGZIP(), []int{4}
}

type GetOrdersForShopAndAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopName  string `protobuf:"bytes,1,opt,name=shop_name,json=shopName,proto3" json:"shop_name,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *GetOrdersForShopAndAccountRequest) Reset() {
	*x = GetOrdersForShopAndAccountRequest{}
	mi := &file_shopify_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrdersForShopAndAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersForShopAndAccountRequest) ProtoMessage() {}

func (x *GetOrdersForShopAndAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shopify_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersForShopAndAccountRequest.ProtoReflect.Descriptor instead.
func (*GetOrdersForShopAndAccountRequest) Descriptor() ([]byte, []int) {
	return file_shopify_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrdersForShopAndAccountRequest) GetShopName() string {
	if x != nil {
		return x.ShopName
	}
	return ""
}

func (x *GetOrdersForShopAndAccountRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type GetOrdersForShopAndAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetOrdersForShopAndAccountResponse) Reset() {
	*x = GetOrdersForShopAndAccountResponse{}
	mi := &file_shopify_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrdersForShopAndAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersForShopAndAccountResponse) ProtoMessage() {}

func (x *GetOrdersForShopAndAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shopify_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersForShopAndAccountResponse.ProtoReflect.Descriptor instead.
func (*GetOrdersForShopAndAccountResponse) Descriptor() ([]byte, []int) {
	return file_shopify_proto_rawDescGZIP(), []int{6}
}

func (x *GetOrdersForShopAndAccountResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type UpdateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order     *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ShopName  string `protobuf:"bytes,3,opt,name=shop_name,json=shopName,proto3" json:"shop_name,omitempty"`
}

func (x *UpdateOrderRequest) Reset() {
	*x = UpdateOrderRequest{}
	mi := &file_shopify_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderRequest) ProtoMessage() {}

func (x *UpdateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shopify_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderRequest) Descriptor() ([]byte, []int) {
	return file_shopify_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateOrderRequest) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *UpdateOrderRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *UpdateOrderRequest) GetShopName() string {
	if x != nil {
		return x.ShopName
	}
	return ""
}

type UpdateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *UpdateOrderResponse) Reset() {
	*x = UpdateOrderResponse{}
	mi := &file_shopify_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderResponse) ProtoMessage() {}

func (x *UpdateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shopify_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderResponse) Descriptor() ([]byte, []int) {
	return file_shopify_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_shopify_proto protoreflect.FileDescriptor

var file_shopify_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x8b, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x77, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x69, 0x6e,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x37, 0x0a, 0x12, 0x53, 0x79,
	0x6e, 0x63, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x22, 0x65, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x5f, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72,
	0x53, 0x68, 0x6f, 0x70, 0x41, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x47, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46, 0x6f,
	0x72, 0x53, 0x68, 0x6f, 0x70, 0x41, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x71, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0xb7, 0x02, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x70, 0x69, 0x66,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6b, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46,
	0x6f, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x41, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46,
	0x6f, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x41, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x41, 0x6e, 0x64,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shopify_proto_rawDescOnce sync.Once
	file_shopify_proto_rawDescData = file_shopify_proto_rawDesc
)

func file_shopify_proto_rawDescGZIP() []byte {
	file_shopify_proto_rawDescOnce.Do(func() {
		file_shopify_proto_rawDescData = protoimpl.X.CompressGZIP(file_shopify_proto_rawDescData)
	})
	return file_shopify_proto_rawDescData
}

var file_shopify_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_shopify_proto_goTypes = []any{
	(*Order)(nil),                              // 0: pb.Order
	(*SyncOrdersRequest)(nil),                  // 1: pb.SyncOrdersRequest
	(*SyncOrdersResponse)(nil),                 // 2: pb.SyncOrdersResponse
	(*StoreTokenRequest)(nil),                  // 3: pb.StoreTokenRequest
	(*StoreTokenResponse)(nil),                 // 4: pb.StoreTokenResponse
	(*GetOrdersForShopAndAccountRequest)(nil),  // 5: pb.GetOrdersForShopAndAccountRequest
	(*GetOrdersForShopAndAccountResponse)(nil), // 6: pb.GetOrdersForShopAndAccountResponse
	(*UpdateOrderRequest)(nil),                 // 7: pb.UpdateOrderRequest
	(*UpdateOrderResponse)(nil),                // 8: pb.UpdateOrderResponse
}
var file_shopify_proto_depIdxs = []int32{
	0, // 0: pb.SyncOrdersResponse.orders:type_name -> pb.Order
	0, // 1: pb.GetOrdersForShopAndAccountResponse.orders:type_name -> pb.Order
	0, // 2: pb.UpdateOrderRequest.order:type_name -> pb.Order
	0, // 3: pb.UpdateOrderResponse.order:type_name -> pb.Order
	1, // 4: pb.ShopifyService.SyncOrders:input_type -> pb.SyncOrdersRequest
	3, // 5: pb.ShopifyService.StoreToken:input_type -> pb.StoreTokenRequest
	5, // 6: pb.ShopifyService.GetOrdersForShopAndAccount:input_type -> pb.GetOrdersForShopAndAccountRequest
	7, // 7: pb.ShopifyService.UpdateOrder:input_type -> pb.UpdateOrderRequest
	2, // 8: pb.ShopifyService.SyncOrders:output_type -> pb.SyncOrdersResponse
	4, // 9: pb.ShopifyService.StoreToken:output_type -> pb.StoreTokenResponse
	6, // 10: pb.ShopifyService.GetOrdersForShopAndAccount:output_type -> pb.GetOrdersForShopAndAccountResponse
	8, // 11: pb.ShopifyService.UpdateOrder:output_type -> pb.UpdateOrderResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_shopify_proto_init() }
func file_shopify_proto_init() {
	if File_shopify_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shopify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shopify_proto_goTypes,
		DependencyIndexes: file_shopify_proto_depIdxs,
		MessageInfos:      file_shopify_proto_msgTypes,
	}.Build()
	File_shopify_proto = out.File
	file_shopify_proto_rawDesc = nil
	file_shopify_proto_goTypes = nil
	file_shopify_proto_depIdxs = nil
}
