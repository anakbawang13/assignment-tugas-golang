// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: aggregator_service/v1/wallet.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Wallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId      string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Balance     float32                `protobuf:"fixed32,5,opt,name=balance,proto3" json:"balance,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Wallet) Reset() {
	*x = Wallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregator_service_v1_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallet) ProtoMessage() {}

func (x *Wallet) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_service_v1_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallet.ProtoReflect.Descriptor instead.
func (*Wallet) Descriptor() ([]byte, []int) {
	return file_aggregator_service_v1_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *Wallet) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Wallet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Wallet) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Wallet) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Wallet) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Wallet) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Wallet) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetWalletResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallets []*Wallet `protobuf:"bytes,1,rep,name=wallets,proto3" json:"wallets,omitempty"`
}

func (x *GetWalletResponse) Reset() {
	*x = GetWalletResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregator_service_v1_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWalletResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWalletResponse) ProtoMessage() {}

func (x *GetWalletResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_service_v1_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWalletResponse.ProtoReflect.Descriptor instead.
func (*GetWalletResponse) Descriptor() ([]byte, []int) {
	return file_aggregator_service_v1_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *GetWalletResponse) GetWallets() []*Wallet {
	if x != nil {
		return x.Wallets
	}
	return nil
}

type GetWalletByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWalletByIDRequest) Reset() {
	*x = GetWalletByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregator_service_v1_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWalletByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWalletByIDRequest) ProtoMessage() {}

func (x *GetWalletByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_service_v1_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWalletByIDRequest.ProtoReflect.Descriptor instead.
func (*GetWalletByIDRequest) Descriptor() ([]byte, []int) {
	return file_aggregator_service_v1_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *GetWalletByIDRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetWalletByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallet *Wallet `protobuf:"bytes,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
}

func (x *GetWalletByIDResponse) Reset() {
	*x = GetWalletByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregator_service_v1_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWalletByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWalletByIDResponse) ProtoMessage() {}

func (x *GetWalletByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_service_v1_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWalletByIDResponse.ProtoReflect.Descriptor instead.
func (*GetWalletByIDResponse) Descriptor() ([]byte, []int) {
	return file_aggregator_service_v1_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *GetWalletByIDResponse) GetWallet() *Wallet {
	if x != nil {
		return x.Wallet
	}
	return nil
}

type UpdateWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId      string  `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Balance     float32 `protobuf:"fixed32,5,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *UpdateWalletRequest) Reset() {
	*x = UpdateWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregator_service_v1_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWalletRequest) ProtoMessage() {}

func (x *UpdateWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_service_v1_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWalletRequest.ProtoReflect.Descriptor instead.
func (*UpdateWalletRequest) Descriptor() ([]byte, []int) {
	return file_aggregator_service_v1_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateWalletRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateWalletRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateWalletRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateWalletRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateWalletRequest) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type DeleteWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteWalletRequest) Reset() {
	*x = DeleteWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregator_service_v1_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWalletRequest) ProtoMessage() {}

func (x *DeleteWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_service_v1_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWalletRequest.ProtoReflect.Descriptor instead.
func (*DeleteWalletRequest) Descriptor() ([]byte, []int) {
	return file_aggregator_service_v1_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteWalletRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserId      string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Balance     float32 `protobuf:"fixed32,4,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *CreateWalletRequest) Reset() {
	*x = CreateWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregator_service_v1_wallet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletRequest) ProtoMessage() {}

func (x *CreateWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_service_v1_wallet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletRequest.ProtoReflect.Descriptor instead.
func (*CreateWalletRequest) Descriptor() ([]byte, []int) {
	return file_aggregator_service_v1_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *CreateWalletRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateWalletRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateWalletRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateWalletRequest) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type MutationWalletResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MutationWalletResponse) Reset() {
	*x = MutationWalletResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregator_service_v1_wallet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutationWalletResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutationWalletResponse) ProtoMessage() {}

func (x *MutationWalletResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aggregator_service_v1_wallet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutationWalletResponse.ProtoReflect.Descriptor instead.
func (*MutationWalletResponse) Descriptor() ([]byte, []int) {
	return file_aggregator_service_v1_wallet_proto_rawDescGZIP(), []int{7}
}

func (x *MutationWalletResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_aggregator_service_v1_wallet_proto protoreflect.FileDescriptor

var file_aggregator_service_v1_wallet_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x06,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x07, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x73, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22,
	0x8e, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7e, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x32, 0x0a, 0x16, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xee, 0x03, 0x0a, 0x0d,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12,
	0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x5d,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x2c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x5d, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x2c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x31, 0x5a, 0x2f,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x74, 0x75, 0x67, 0x61, 0x73,
	0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aggregator_service_v1_wallet_proto_rawDescOnce sync.Once
	file_aggregator_service_v1_wallet_proto_rawDescData = file_aggregator_service_v1_wallet_proto_rawDesc
)

func file_aggregator_service_v1_wallet_proto_rawDescGZIP() []byte {
	file_aggregator_service_v1_wallet_proto_rawDescOnce.Do(func() {
		file_aggregator_service_v1_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_aggregator_service_v1_wallet_proto_rawDescData)
	})
	return file_aggregator_service_v1_wallet_proto_rawDescData
}

var file_aggregator_service_v1_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_aggregator_service_v1_wallet_proto_goTypes = []any{
	(*Wallet)(nil),                 // 0: proto.wallet_service.v1.Wallet
	(*GetWalletResponse)(nil),      // 1: proto.wallet_service.v1.GetWalletResponse
	(*GetWalletByIDRequest)(nil),   // 2: proto.wallet_service.v1.GetWalletByIDRequest
	(*GetWalletByIDResponse)(nil),  // 3: proto.wallet_service.v1.GetWalletByIDResponse
	(*UpdateWalletRequest)(nil),    // 4: proto.wallet_service.v1.UpdateWalletRequest
	(*DeleteWalletRequest)(nil),    // 5: proto.wallet_service.v1.DeleteWalletRequest
	(*CreateWalletRequest)(nil),    // 6: proto.wallet_service.v1.CreateWalletRequest
	(*MutationWalletResponse)(nil), // 7: proto.wallet_service.v1.MutationWalletResponse
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_aggregator_service_v1_wallet_proto_depIdxs = []int32{
	8, // 0: proto.wallet_service.v1.Wallet.created_at:type_name -> google.protobuf.Timestamp
	8, // 1: proto.wallet_service.v1.Wallet.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: proto.wallet_service.v1.GetWalletResponse.wallets:type_name -> proto.wallet_service.v1.Wallet
	0, // 3: proto.wallet_service.v1.GetWalletByIDResponse.wallet:type_name -> proto.wallet_service.v1.Wallet
	9, // 4: proto.wallet_service.v1.WalletService.GetWallets:input_type -> google.protobuf.Empty
	2, // 5: proto.wallet_service.v1.WalletService.GetWalletByID:input_type -> proto.wallet_service.v1.GetWalletByIDRequest
	6, // 6: proto.wallet_service.v1.WalletService.CreateWallet:input_type -> proto.wallet_service.v1.CreateWalletRequest
	4, // 7: proto.wallet_service.v1.WalletService.UpdateWallet:input_type -> proto.wallet_service.v1.UpdateWalletRequest
	5, // 8: proto.wallet_service.v1.WalletService.DeleteWallet:input_type -> proto.wallet_service.v1.DeleteWalletRequest
	1, // 9: proto.wallet_service.v1.WalletService.GetWallets:output_type -> proto.wallet_service.v1.GetWalletResponse
	3, // 10: proto.wallet_service.v1.WalletService.GetWalletByID:output_type -> proto.wallet_service.v1.GetWalletByIDResponse
	0, // 11: proto.wallet_service.v1.WalletService.CreateWallet:output_type -> proto.wallet_service.v1.Wallet
	0, // 12: proto.wallet_service.v1.WalletService.UpdateWallet:output_type -> proto.wallet_service.v1.Wallet
	0, // 13: proto.wallet_service.v1.WalletService.DeleteWallet:output_type -> proto.wallet_service.v1.Wallet
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_aggregator_service_v1_wallet_proto_init() }
func file_aggregator_service_v1_wallet_proto_init() {
	if File_aggregator_service_v1_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aggregator_service_v1_wallet_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Wallet); i {
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
		file_aggregator_service_v1_wallet_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetWalletResponse); i {
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
		file_aggregator_service_v1_wallet_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetWalletByIDRequest); i {
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
		file_aggregator_service_v1_wallet_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetWalletByIDResponse); i {
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
		file_aggregator_service_v1_wallet_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateWalletRequest); i {
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
		file_aggregator_service_v1_wallet_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteWalletRequest); i {
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
		file_aggregator_service_v1_wallet_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CreateWalletRequest); i {
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
		file_aggregator_service_v1_wallet_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*MutationWalletResponse); i {
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
			RawDescriptor: file_aggregator_service_v1_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aggregator_service_v1_wallet_proto_goTypes,
		DependencyIndexes: file_aggregator_service_v1_wallet_proto_depIdxs,
		MessageInfos:      file_aggregator_service_v1_wallet_proto_msgTypes,
	}.Build()
	File_aggregator_service_v1_wallet_proto = out.File
	file_aggregator_service_v1_wallet_proto_rawDesc = nil
	file_aggregator_service_v1_wallet_proto_goTypes = nil
	file_aggregator_service_v1_wallet_proto_depIdxs = nil
}
