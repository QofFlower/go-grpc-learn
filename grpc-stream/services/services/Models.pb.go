// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: Models.proto

package services

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderMain *OrderMain `protobuf:"bytes,1,opt,name=order_main,json=orderMain,proto3" json:"order_main,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_Models_proto_rawDescGZIP(), []int{0}
}

func (x *OrderRequest) GetOrderMain() *OrderMain {
	if x != nil {
		return x.OrderMain
	}
	return nil
}

type ProdModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdId    int32   `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	ProdName  string  `protobuf:"bytes,2,opt,name=prod_name,json=prodName,proto3" json:"prod_name,omitempty"`
	ProdPrice float32 `protobuf:"fixed32,3,opt,name=prod_price,json=prodPrice,proto3" json:"prod_price,omitempty"`
}

func (x *ProdModel) Reset() {
	*x = ProdModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdModel) ProtoMessage() {}

func (x *ProdModel) ProtoReflect() protoreflect.Message {
	mi := &file_Models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdModel.ProtoReflect.Descriptor instead.
func (*ProdModel) Descriptor() ([]byte, []int) {
	return file_Models_proto_rawDescGZIP(), []int{1}
}

func (x *ProdModel) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

func (x *ProdModel) GetProdName() string {
	if x != nil {
		return x.ProdName
	}
	return ""
}

func (x *ProdModel) GetProdPrice() float32 {
	if x != nil {
		return x.ProdPrice
	}
	return 0
}

type OrderMain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId      int32                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderNo      string                 `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	UserId       int32                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderMoney   float32                `protobuf:"fixed32,4,opt,name=order_money,json=orderMoney,proto3" json:"order_money,omitempty"`
	OrderTime    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=order_time,json=orderTime,proto3" json:"order_time,omitempty"`          // 日期类型
	OrderDetails []*OrderDetail         `protobuf:"bytes,6,rep,name=order_details,json=orderDetails,proto3" json:"order_details,omitempty"` // 套娃大法好
}

func (x *OrderMain) Reset() {
	*x = OrderMain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderMain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderMain) ProtoMessage() {}

func (x *OrderMain) ProtoReflect() protoreflect.Message {
	mi := &file_Models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderMain.ProtoReflect.Descriptor instead.
func (*OrderMain) Descriptor() ([]byte, []int) {
	return file_Models_proto_rawDescGZIP(), []int{2}
}

func (x *OrderMain) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderMain) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *OrderMain) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderMain) GetOrderMoney() float32 {
	if x != nil {
		return x.OrderMoney
	}
	return 0
}

func (x *OrderMain) GetOrderTime() *timestamppb.Timestamp {
	if x != nil {
		return x.OrderTime
	}
	return nil
}

func (x *OrderMain) GetOrderDetails() []*OrderDetail {
	if x != nil {
		return x.OrderDetails
	}
	return nil
}

// 子订单模型
type OrderDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DetailId  int32   `protobuf:"varint,1,opt,name=detail_id,json=detailId,proto3" json:"detail_id,omitempty"`
	OrderNo   string  `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	ProdId    int32   `protobuf:"varint,3,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	ProdPrice float32 `protobuf:"fixed32,4,opt,name=prod_price,json=prodPrice,proto3" json:"prod_price,omitempty"`
	ProdNum   int32   `protobuf:"varint,5,opt,name=prod_num,json=prodNum,proto3" json:"prod_num,omitempty"`
}

func (x *OrderDetail) Reset() {
	*x = OrderDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetail) ProtoMessage() {}

func (x *OrderDetail) ProtoReflect() protoreflect.Message {
	mi := &file_Models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetail.ProtoReflect.Descriptor instead.
func (*OrderDetail) Descriptor() ([]byte, []int) {
	return file_Models_proto_rawDescGZIP(), []int{3}
}

func (x *OrderDetail) GetDetailId() int32 {
	if x != nil {
		return x.DetailId
	}
	return 0
}

func (x *OrderDetail) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *OrderDetail) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

func (x *OrderDetail) GetProdPrice() float32 {
	if x != nil {
		return x.ProdPrice
	}
	return 0
}

func (x *OrderDetail) GetProdNum() int32 {
	if x != nil {
		return x.ProdNum
	}
	return 0
}

// 用户模型
type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserScore int32 `protobuf:"varint,2,opt,name=user_score,json=userScore,proto3" json:"user_score,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_Models_proto_rawDescGZIP(), []int{4}
}

func (x *UserInfo) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserInfo) GetUserScore() int32 {
	if x != nil {
		return x.UserScore
	}
	return 0
}

var File_Models_proto protoreflect.FileDescriptor

var file_Models_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x61,
	0x69, 0x6e, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6e, 0x22, 0x60, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x72, 0x6f,
	0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22,
	0xfe, 0x01, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4e, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x0a, 0x05, 0x25, 0x00, 0x00, 0x80, 0x3f, 0x52, 0x0a, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x98, 0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x4e, 0x75, 0x6d, 0x22, 0x42, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Models_proto_rawDescOnce sync.Once
	file_Models_proto_rawDescData = file_Models_proto_rawDesc
)

func file_Models_proto_rawDescGZIP() []byte {
	file_Models_proto_rawDescOnce.Do(func() {
		file_Models_proto_rawDescData = protoimpl.X.CompressGZIP(file_Models_proto_rawDescData)
	})
	return file_Models_proto_rawDescData
}

var file_Models_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Models_proto_goTypes = []interface{}{
	(*OrderRequest)(nil),          // 0: services.OrderRequest
	(*ProdModel)(nil),             // 1: services.ProdModel
	(*OrderMain)(nil),             // 2: services.OrderMain
	(*OrderDetail)(nil),           // 3: services.OrderDetail
	(*UserInfo)(nil),              // 4: services.UserInfo
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_Models_proto_depIdxs = []int32{
	2, // 0: services.OrderRequest.order_main:type_name -> services.OrderMain
	5, // 1: services.OrderMain.order_time:type_name -> google.protobuf.Timestamp
	3, // 2: services.OrderMain.order_details:type_name -> services.OrderDetail
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Models_proto_init() }
func file_Models_proto_init() {
	if File_Models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
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
		file_Models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdModel); i {
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
		file_Models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderMain); i {
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
		file_Models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetail); i {
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
		file_Models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
			RawDescriptor: file_Models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Models_proto_goTypes,
		DependencyIndexes: file_Models_proto_depIdxs,
		MessageInfos:      file_Models_proto_msgTypes,
	}.Build()
	File_Models_proto = out.File
	file_Models_proto_rawDesc = nil
	file_Models_proto_goTypes = nil
	file_Models_proto_depIdxs = nil
}
