// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SmartcycleService.proto

package smartcycle

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SensorRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SensorRequest) Reset()         { *m = SensorRequest{} }
func (m *SensorRequest) String() string { return proto.CompactTextString(m) }
func (*SensorRequest) ProtoMessage()    {}
func (*SensorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d170b0bc86c15d60, []int{0}
}

func (m *SensorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SensorRequest.Unmarshal(m, b)
}
func (m *SensorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SensorRequest.Marshal(b, m, deterministic)
}
func (m *SensorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensorRequest.Merge(m, src)
}
func (m *SensorRequest) XXX_Size() int {
	return xxx_messageInfo_SensorRequest.Size(m)
}
func (m *SensorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SensorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SensorRequest proto.InternalMessageInfo

type SensorData struct {
	Speed                int32    `protobuf:"varint,1,opt,name=speed,proto3" json:"speed,omitempty"`
	Heartrate            int32    `protobuf:"varint,2,opt,name=heartrate,proto3" json:"heartrate,omitempty"`
	Cadence              int32    `protobuf:"varint,3,opt,name=cadence,proto3" json:"cadence,omitempty"`
	Temperature          int32    `protobuf:"varint,4,opt,name=temperature,proto3" json:"temperature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SensorData) Reset()         { *m = SensorData{} }
func (m *SensorData) String() string { return proto.CompactTextString(m) }
func (*SensorData) ProtoMessage()    {}
func (*SensorData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d170b0bc86c15d60, []int{1}
}

func (m *SensorData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SensorData.Unmarshal(m, b)
}
func (m *SensorData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SensorData.Marshal(b, m, deterministic)
}
func (m *SensorData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensorData.Merge(m, src)
}
func (m *SensorData) XXX_Size() int {
	return xxx_messageInfo_SensorData.Size(m)
}
func (m *SensorData) XXX_DiscardUnknown() {
	xxx_messageInfo_SensorData.DiscardUnknown(m)
}

var xxx_messageInfo_SensorData proto.InternalMessageInfo

func (m *SensorData) GetSpeed() int32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *SensorData) GetHeartrate() int32 {
	if m != nil {
		return m.Heartrate
	}
	return 0
}

func (m *SensorData) GetCadence() int32 {
	if m != nil {
		return m.Cadence
	}
	return 0
}

func (m *SensorData) GetTemperature() int32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

type HeadlightRequest struct {
	On                   bool     `protobuf:"varint,1,opt,name=on,proto3" json:"on,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeadlightRequest) Reset()         { *m = HeadlightRequest{} }
func (m *HeadlightRequest) String() string { return proto.CompactTextString(m) }
func (*HeadlightRequest) ProtoMessage()    {}
func (*HeadlightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d170b0bc86c15d60, []int{2}
}

func (m *HeadlightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeadlightRequest.Unmarshal(m, b)
}
func (m *HeadlightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeadlightRequest.Marshal(b, m, deterministic)
}
func (m *HeadlightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadlightRequest.Merge(m, src)
}
func (m *HeadlightRequest) XXX_Size() int {
	return xxx_messageInfo_HeadlightRequest.Size(m)
}
func (m *HeadlightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadlightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeadlightRequest proto.InternalMessageInfo

func (m *HeadlightRequest) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

type HeadlightData struct {
	On                   bool     `protobuf:"varint,1,opt,name=on,proto3" json:"on,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeadlightData) Reset()         { *m = HeadlightData{} }
func (m *HeadlightData) String() string { return proto.CompactTextString(m) }
func (*HeadlightData) ProtoMessage()    {}
func (*HeadlightData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d170b0bc86c15d60, []int{3}
}

func (m *HeadlightData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeadlightData.Unmarshal(m, b)
}
func (m *HeadlightData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeadlightData.Marshal(b, m, deterministic)
}
func (m *HeadlightData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadlightData.Merge(m, src)
}
func (m *HeadlightData) XXX_Size() int {
	return xxx_messageInfo_HeadlightData.Size(m)
}
func (m *HeadlightData) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadlightData.DiscardUnknown(m)
}

var xxx_messageInfo_HeadlightData proto.InternalMessageInfo

func (m *HeadlightData) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func init() {
	proto.RegisterType((*SensorRequest)(nil), "smartcycle.SensorRequest")
	proto.RegisterType((*SensorData)(nil), "smartcycle.SensorData")
	proto.RegisterType((*HeadlightRequest)(nil), "smartcycle.HeadlightRequest")
	proto.RegisterType((*HeadlightData)(nil), "smartcycle.HeadlightData")
}

func init() { proto.RegisterFile("SmartcycleService.proto", fileDescriptor_d170b0bc86c15d60) }

var fileDescriptor_d170b0bc86c15d60 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0xbf, 0x96, 0x0f, 0xff, 0x5c, 0x83, 0xe8, 0xc4, 0x68, 0x21, 0x24, 0x92, 0x59, 0x99,
	0x98, 0xcc, 0x42, 0xdf, 0x80, 0x10, 0x65, 0xc1, 0x82, 0xd0, 0x85, 0xeb, 0xeb, 0x70, 0x03, 0x4d,
	0xe8, 0x4c, 0x9d, 0x99, 0x82, 0x12, 0x9f, 0xc6, 0x8d, 0xaf, 0x69, 0x98, 0x6a, 0x5b, 0xc1, 0xe5,
	0x3d, 0xbf, 0x33, 0x93, 0x73, 0xee, 0x85, 0xab, 0x38, 0x45, 0xe3, 0xe4, 0x9b, 0x5c, 0x52, 0x4c,
	0x66, 0x95, 0x48, 0x12, 0x99, 0xd1, 0x4e, 0x33, 0xb0, 0x25, 0xe0, 0x6d, 0x68, 0xc5, 0xa4, 0xac,
	0x36, 0x53, 0x7a, 0xc9, 0xc9, 0x3a, 0xfe, 0x0e, 0x50, 0x08, 0x43, 0x74, 0xc8, 0x2e, 0xa0, 0x69,
	0x33, 0xa2, 0x59, 0x14, 0xf4, 0x83, 0x9b, 0xe6, 0xb4, 0x18, 0x58, 0x0f, 0x8e, 0x17, 0x84, 0xc6,
	0x19, 0x74, 0x14, 0x85, 0x9e, 0x54, 0x02, 0x8b, 0xe0, 0x50, 0xe2, 0x8c, 0x94, 0xa4, 0xa8, 0xe1,
	0xd9, 0xcf, 0xc8, 0xfa, 0x70, 0xe2, 0x28, 0xcd, 0xc8, 0xa0, 0xcb, 0x0d, 0x45, 0xff, 0x3d, 0xad,
	0x4b, 0x9c, 0xc3, 0xd9, 0x88, 0x70, 0xb6, 0x4c, 0xe6, 0x0b, 0xf7, 0x9d, 0x88, 0x9d, 0x42, 0xa8,
	0x95, 0x0f, 0x70, 0x34, 0x0d, 0xb5, 0xe2, 0xd7, 0xd0, 0x2a, 0x3d, 0x3e, 0xe4, 0x8e, 0xe1, 0xee,
	0x33, 0x80, 0xf3, 0xbd, 0xee, 0x6c, 0x08, 0xad, 0x47, 0x72, 0xb5, 0x6e, 0x1d, 0x51, 0xed, 0x41,
	0xfc, 0x5a, 0x42, 0xf7, 0x72, 0x1f, 0x6d, 0x9f, 0xf0, 0x7f, 0x6c, 0x0c, 0xed, 0x78, 0x9d, 0x38,
	0xb9, 0x28, 0x23, 0xb0, 0x5e, 0xdd, 0xbc, 0x9b, 0xbe, 0xdb, 0xf9, 0x93, 0x16, 0xbf, 0x0d, 0x1e,
	0xe0, 0x56, 0xea, 0x54, 0x60, 0x8a, 0x1b, 0xad, 0x04, 0xae, 0x6d, 0xdd, 0xbc, 0xca, 0x37, 0xc9,
	0xeb, 0x92, 0x94, 0x15, 0x29, 0x59, 0x8b, 0x73, 0x1a, 0xb4, 0xab, 0x56, 0x93, 0xed, 0x25, 0x27,
	0xc1, 0x47, 0xd8, 0x18, 0x8d, 0x9f, 0x9e, 0x0f, 0xfc, 0x61, 0xef, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x83, 0x3a, 0xa7, 0x20, 0xf3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SmartcycleServiceClient is the client API for SmartcycleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmartcycleServiceClient interface {
	GetSensorData(ctx context.Context, in *SensorRequest, opts ...grpc.CallOption) (*SensorData, error)
	SwitchHeadlight(ctx context.Context, in *HeadlightRequest, opts ...grpc.CallOption) (*HeadlightData, error)
}

type smartcycleServiceClient struct {
	cc *grpc.ClientConn
}

func NewSmartcycleServiceClient(cc *grpc.ClientConn) SmartcycleServiceClient {
	return &smartcycleServiceClient{cc}
}

func (c *smartcycleServiceClient) GetSensorData(ctx context.Context, in *SensorRequest, opts ...grpc.CallOption) (*SensorData, error) {
	out := new(SensorData)
	err := c.cc.Invoke(ctx, "/smartcycle.SmartcycleService/GetSensorData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartcycleServiceClient) SwitchHeadlight(ctx context.Context, in *HeadlightRequest, opts ...grpc.CallOption) (*HeadlightData, error) {
	out := new(HeadlightData)
	err := c.cc.Invoke(ctx, "/smartcycle.SmartcycleService/SwitchHeadlight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmartcycleServiceServer is the server API for SmartcycleService service.
type SmartcycleServiceServer interface {
	GetSensorData(context.Context, *SensorRequest) (*SensorData, error)
	SwitchHeadlight(context.Context, *HeadlightRequest) (*HeadlightData, error)
}

// UnimplementedSmartcycleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSmartcycleServiceServer struct {
}

func (*UnimplementedSmartcycleServiceServer) GetSensorData(ctx context.Context, req *SensorRequest) (*SensorData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSensorData not implemented")
}
func (*UnimplementedSmartcycleServiceServer) SwitchHeadlight(ctx context.Context, req *HeadlightRequest) (*HeadlightData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchHeadlight not implemented")
}

func RegisterSmartcycleServiceServer(s *grpc.Server, srv SmartcycleServiceServer) {
	s.RegisterService(&_SmartcycleService_serviceDesc, srv)
}

func _SmartcycleService_GetSensorData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartcycleServiceServer).GetSensorData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcycle.SmartcycleService/GetSensorData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartcycleServiceServer).GetSensorData(ctx, req.(*SensorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartcycleService_SwitchHeadlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeadlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartcycleServiceServer).SwitchHeadlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcycle.SmartcycleService/SwitchHeadlight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartcycleServiceServer).SwitchHeadlight(ctx, req.(*HeadlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SmartcycleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smartcycle.SmartcycleService",
	HandlerType: (*SmartcycleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSensorData",
			Handler:    _SmartcycleService_GetSensorData_Handler,
		},
		{
			MethodName: "SwitchHeadlight",
			Handler:    _SmartcycleService_SwitchHeadlight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "SmartcycleService.proto",
}
