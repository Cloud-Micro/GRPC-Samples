// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order_service.proto

package ecommerce

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Order struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price                float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_93a121d2d2ec3d32, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Order) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Order) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type CombinedShipment struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	OrderIDList          []string `protobuf:"bytes,3,rep,name=orderIDList,proto3" json:"orderIDList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CombinedShipment) Reset()         { *m = CombinedShipment{} }
func (m *CombinedShipment) String() string { return proto.CompactTextString(m) }
func (*CombinedShipment) ProtoMessage()    {}
func (*CombinedShipment) Descriptor() ([]byte, []int) {
	return fileDescriptor_93a121d2d2ec3d32, []int{1}
}

func (m *CombinedShipment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CombinedShipment.Unmarshal(m, b)
}
func (m *CombinedShipment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CombinedShipment.Marshal(b, m, deterministic)
}
func (m *CombinedShipment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CombinedShipment.Merge(m, src)
}
func (m *CombinedShipment) XXX_Size() int {
	return xxx_messageInfo_CombinedShipment.Size(m)
}
func (m *CombinedShipment) XXX_DiscardUnknown() {
	xxx_messageInfo_CombinedShipment.DiscardUnknown(m)
}

var xxx_messageInfo_CombinedShipment proto.InternalMessageInfo

func (m *CombinedShipment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CombinedShipment) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CombinedShipment) GetOrderIDList() []string {
	if m != nil {
		return m.OrderIDList
	}
	return nil
}

func init() {
	proto.RegisterType((*Order)(nil), "ecommerce.Order")
	proto.RegisterType((*CombinedShipment)(nil), "ecommerce.CombinedShipment")
}

func init() { proto.RegisterFile("order_service.proto", fileDescriptor_93a121d2d2ec3d32) }

var fileDescriptor_93a121d2d2ec3d32 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0xc9, 0xf4, 0x0f, 0xf4, 0x7e, 0x9f, 0xb5, 0x44, 0x91, 0xa1, 0x8a, 0x0c, 0x5d, 0xcd,
	0x6a, 0x5a, 0xf4, 0x05, 0x8a, 0xba, 0x11, 0x14, 0xa1, 0x05, 0x57, 0x82, 0xa4, 0x99, 0xeb, 0x34,
	0xd0, 0x49, 0xc2, 0x4d, 0x46, 0x17, 0x3e, 0x8a, 0x2f, 0x2b, 0x93, 0xa9, 0x3a, 0x58, 0x14, 0xdc,
	0xe5, 0x9e, 0xdc, 0xf3, 0xe3, 0xdc, 0x03, 0x07, 0x86, 0x72, 0xa4, 0x47, 0x87, 0xf4, 0xac, 0x24,
	0x66, 0x96, 0x8c, 0x37, 0x7c, 0x80, 0xd2, 0x94, 0x25, 0x92, 0xc4, 0xf1, 0x69, 0x61, 0x4c, 0xb1,
	0xc1, 0x69, 0xf8, 0x58, 0x55, 0x4f, 0xd3, 0x17, 0x12, 0xd6, 0x22, 0xb9, 0x66, 0x75, 0xf2, 0x0a,
	0xbd, 0xbb, 0x9a, 0xc0, 0x87, 0x10, 0xa9, 0x3c, 0x66, 0x09, 0x4b, 0x07, 0x8b, 0x48, 0xe5, 0x9c,
	0x43, 0x57, 0x8b, 0x12, 0xe3, 0x28, 0x28, 0xe1, 0xcd, 0x13, 0xf8, 0x97, 0xa3, 0x93, 0xa4, 0xac,
	0x57, 0x46, 0xc7, 0x9d, 0xf0, 0xd5, 0x96, 0xf8, 0x21, 0xf4, 0x2c, 0x29, 0x89, 0x71, 0x37, 0x61,
	0x69, 0xb4, 0x68, 0x06, 0x7e, 0x04, 0x7d, 0xe7, 0x85, 0xaf, 0x5c, 0xdc, 0x0b, 0x96, 0xed, 0x34,
	0x79, 0x80, 0xd1, 0xa5, 0x29, 0x57, 0x4a, 0x63, 0xbe, 0x5c, 0x2b, 0x5b, 0xa2, 0xf6, 0x3b, 0x39,
	0xbe, 0xbc, 0x51, 0xdb, 0x5b, 0x67, 0x09, 0xa7, 0x5f, 0x5f, 0xdd, 0x28, 0xe7, 0xe3, 0x4e, 0xd2,
	0xa9, 0xb3, 0xb4, 0xa4, 0xb3, 0xb7, 0x08, 0xf6, 0xc3, 0x6d, 0xb7, 0x42, 0x8b, 0x02, 0x03, 0x7d,
	0x0e, 0xc3, 0x02, 0x7d, 0x50, 0x97, 0x0d, 0xe7, 0x24, 0x6b, 0x1a, 0xca, 0x3e, 0x1a, 0xca, 0x96,
	0x9e, 0x94, 0x2e, 0xee, 0xc5, 0xa6, 0xc2, 0xf1, 0x28, 0xfb, 0xac, 0x32, 0x6b, 0x7a, 0x9a, 0xc3,
	0x7f, 0x87, 0x82, 0xe4, 0x3a, 0x8c, 0x7f, 0xf6, 0xcf, 0x58, 0x4d, 0xa8, 0x6c, 0x2e, 0x3c, 0x6e,
	0x09, 0x3b, 0x3b, 0xe3, 0x5f, 0x99, 0x29, 0xe3, 0x17, 0xb0, 0x67, 0xc9, 0x48, 0x74, 0xee, 0x47,
	0xc4, 0x71, 0x4b, 0xf9, 0xde, 0x71, 0xca, 0x66, 0x6c, 0xd5, 0x0f, 0xec, 0xf3, 0xf7, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x31, 0xa9, 0xe5, 0x9e, 0x41, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderManagementClient is the client API for OrderManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderManagementClient interface {
	GetOrderStatus(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Order, error)
	SearchOrders(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (OrderManagement_SearchOrdersClient, error)
	UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_UpdateOrdersClient, error)
	ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_ProcessOrdersClient, error)
}

type orderManagementClient struct {
	cc *grpc.ClientConn
}

func NewOrderManagementClient(cc *grpc.ClientConn) OrderManagementClient {
	return &orderManagementClient{cc}
}

func (c *orderManagementClient) GetOrderStatus(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/ecommerce.OrderManagement/getOrderStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderManagementClient) SearchOrders(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (OrderManagement_SearchOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderManagement_serviceDesc.Streams[0], "/ecommerce.OrderManagement/searchOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementSearchOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderManagement_SearchOrdersClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderManagementSearchOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementSearchOrdersClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementClient) UpdateOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_UpdateOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderManagement_serviceDesc.Streams[1], "/ecommerce.OrderManagement/updateOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementUpdateOrdersClient{stream}
	return x, nil
}

type OrderManagement_UpdateOrdersClient interface {
	Send(*Order) error
	CloseAndRecv() (*wrappers.StringValue, error)
	grpc.ClientStream
}

type orderManagementUpdateOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementUpdateOrdersClient) Send(m *Order) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementUpdateOrdersClient) CloseAndRecv() (*wrappers.StringValue, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(wrappers.StringValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderManagementClient) ProcessOrders(ctx context.Context, opts ...grpc.CallOption) (OrderManagement_ProcessOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderManagement_serviceDesc.Streams[2], "/ecommerce.OrderManagement/processOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementProcessOrdersClient{stream}
	return x, nil
}

type OrderManagement_ProcessOrdersClient interface {
	Send(*Order) error
	Recv() (*CombinedShipment, error)
	grpc.ClientStream
}

type orderManagementProcessOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementProcessOrdersClient) Send(m *Order) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderManagementProcessOrdersClient) Recv() (*CombinedShipment, error) {
	m := new(CombinedShipment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderManagementServer is the server API for OrderManagement service.
type OrderManagementServer interface {
	GetOrderStatus(context.Context, *wrappers.StringValue) (*Order, error)
	SearchOrders(*wrappers.StringValue, OrderManagement_SearchOrdersServer) error
	UpdateOrders(OrderManagement_UpdateOrdersServer) error
	ProcessOrders(OrderManagement_ProcessOrdersServer) error
}

// UnimplementedOrderManagementServer can be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServer struct {
}

func (*UnimplementedOrderManagementServer) GetOrderStatus(ctx context.Context, req *wrappers.StringValue) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderStatus not implemented")
}
func (*UnimplementedOrderManagementServer) SearchOrders(req *wrappers.StringValue, srv OrderManagement_SearchOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchOrders not implemented")
}
func (*UnimplementedOrderManagementServer) UpdateOrders(srv OrderManagement_UpdateOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateOrders not implemented")
}
func (*UnimplementedOrderManagementServer) ProcessOrders(srv OrderManagement_ProcessOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessOrders not implemented")
}

func RegisterOrderManagementServer(s *grpc.Server, srv OrderManagementServer) {
	s.RegisterService(&_OrderManagement_serviceDesc, srv)
}

func _OrderManagement_GetOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).GetOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecommerce.OrderManagement/GetOrderStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).GetOrderStatus(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderManagement_SearchOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrappers.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderManagementServer).SearchOrders(m, &orderManagementSearchOrdersServer{stream})
}

type OrderManagement_SearchOrdersServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderManagementSearchOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementSearchOrdersServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderManagement_UpdateOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).UpdateOrders(&orderManagementUpdateOrdersServer{stream})
}

type OrderManagement_UpdateOrdersServer interface {
	SendAndClose(*wrappers.StringValue) error
	Recv() (*Order, error)
	grpc.ServerStream
}

type orderManagementUpdateOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementUpdateOrdersServer) SendAndClose(m *wrappers.StringValue) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementUpdateOrdersServer) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OrderManagement_ProcessOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderManagementServer).ProcessOrders(&orderManagementProcessOrdersServer{stream})
}

type OrderManagement_ProcessOrdersServer interface {
	Send(*CombinedShipment) error
	Recv() (*Order, error)
	grpc.ServerStream
}

type orderManagementProcessOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementProcessOrdersServer) Send(m *CombinedShipment) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderManagementProcessOrdersServer) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _OrderManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.OrderManagement",
	HandlerType: (*OrderManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getOrderStatus",
			Handler:    _OrderManagement_GetOrderStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "searchOrders",
			Handler:       _OrderManagement_SearchOrders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "updateOrders",
			Handler:       _OrderManagement_UpdateOrders_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "processOrders",
			Handler:       _OrderManagement_ProcessOrders_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "order_service.proto",
}
