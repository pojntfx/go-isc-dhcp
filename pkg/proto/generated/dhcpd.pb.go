// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dhcpd.proto

package goISCDHCP

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

type DHCPD struct {
	Device               string    `protobuf:"bytes,1,opt,name=Device,proto3" json:"Device,omitempty"`
	Subnets              []*Subnet `protobuf:"bytes,2,rep,name=Subnets,proto3" json:"Subnets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DHCPD) Reset()         { *m = DHCPD{} }
func (m *DHCPD) String() string { return proto.CompactTextString(m) }
func (*DHCPD) ProtoMessage()    {}
func (*DHCPD) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f4f7f9a47f6dd7, []int{0}
}

func (m *DHCPD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DHCPD.Unmarshal(m, b)
}
func (m *DHCPD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DHCPD.Marshal(b, m, deterministic)
}
func (m *DHCPD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DHCPD.Merge(m, src)
}
func (m *DHCPD) XXX_Size() int {
	return xxx_messageInfo_DHCPD.Size(m)
}
func (m *DHCPD) XXX_DiscardUnknown() {
	xxx_messageInfo_DHCPD.DiscardUnknown(m)
}

var xxx_messageInfo_DHCPD proto.InternalMessageInfo

func (m *DHCPD) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *DHCPD) GetSubnets() []*Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

type Subnet struct {
	Network              string   `protobuf:"bytes,1,opt,name=Network,proto3" json:"Network,omitempty"`
	Netmask              string   `protobuf:"bytes,2,opt,name=Netmask,proto3" json:"Netmask,omitempty"`
	NextServer           string   `protobuf:"bytes,3,opt,name=NextServer,proto3" json:"NextServer,omitempty"`
	Filename             string   `protobuf:"bytes,4,opt,name=Filename,proto3" json:"Filename,omitempty"`
	Range                *Range   `protobuf:"bytes,5,opt,name=Range,proto3" json:"Range,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Subnet) Reset()         { *m = Subnet{} }
func (m *Subnet) String() string { return proto.CompactTextString(m) }
func (*Subnet) ProtoMessage()    {}
func (*Subnet) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f4f7f9a47f6dd7, []int{1}
}

func (m *Subnet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subnet.Unmarshal(m, b)
}
func (m *Subnet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subnet.Marshal(b, m, deterministic)
}
func (m *Subnet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subnet.Merge(m, src)
}
func (m *Subnet) XXX_Size() int {
	return xxx_messageInfo_Subnet.Size(m)
}
func (m *Subnet) XXX_DiscardUnknown() {
	xxx_messageInfo_Subnet.DiscardUnknown(m)
}

var xxx_messageInfo_Subnet proto.InternalMessageInfo

func (m *Subnet) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Subnet) GetNetmask() string {
	if m != nil {
		return m.Netmask
	}
	return ""
}

func (m *Subnet) GetNextServer() string {
	if m != nil {
		return m.NextServer
	}
	return ""
}

func (m *Subnet) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Subnet) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

type Range struct {
	Start                string   `protobuf:"bytes,1,opt,name=Start,proto3" json:"Start,omitempty"`
	End                  string   `protobuf:"bytes,2,opt,name=End,proto3" json:"End,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f4f7f9a47f6dd7, []int{2}
}

func (m *Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Range.Unmarshal(m, b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Range.Marshal(b, m, deterministic)
}
func (m *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(m, src)
}
func (m *Range) XXX_Size() int {
	return xxx_messageInfo_Range.Size(m)
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Range) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type DHCPDManaged struct {
	Id                   string    `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Device               string    `protobuf:"bytes,2,opt,name=Device,proto3" json:"Device,omitempty"`
	Subnets              []*Subnet `protobuf:"bytes,3,rep,name=Subnets,proto3" json:"Subnets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DHCPDManaged) Reset()         { *m = DHCPDManaged{} }
func (m *DHCPDManaged) String() string { return proto.CompactTextString(m) }
func (*DHCPDManaged) ProtoMessage()    {}
func (*DHCPDManaged) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f4f7f9a47f6dd7, []int{3}
}

func (m *DHCPDManaged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DHCPDManaged.Unmarshal(m, b)
}
func (m *DHCPDManaged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DHCPDManaged.Marshal(b, m, deterministic)
}
func (m *DHCPDManaged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DHCPDManaged.Merge(m, src)
}
func (m *DHCPDManaged) XXX_Size() int {
	return xxx_messageInfo_DHCPDManaged.Size(m)
}
func (m *DHCPDManaged) XXX_DiscardUnknown() {
	xxx_messageInfo_DHCPDManaged.DiscardUnknown(m)
}

var xxx_messageInfo_DHCPDManaged proto.InternalMessageInfo

func (m *DHCPDManaged) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DHCPDManaged) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *DHCPDManaged) GetSubnets() []*Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

type DHCPDManagerListArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DHCPDManagerListArgs) Reset()         { *m = DHCPDManagerListArgs{} }
func (m *DHCPDManagerListArgs) String() string { return proto.CompactTextString(m) }
func (*DHCPDManagerListArgs) ProtoMessage()    {}
func (*DHCPDManagerListArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f4f7f9a47f6dd7, []int{4}
}

func (m *DHCPDManagerListArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DHCPDManagerListArgs.Unmarshal(m, b)
}
func (m *DHCPDManagerListArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DHCPDManagerListArgs.Marshal(b, m, deterministic)
}
func (m *DHCPDManagerListArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DHCPDManagerListArgs.Merge(m, src)
}
func (m *DHCPDManagerListArgs) XXX_Size() int {
	return xxx_messageInfo_DHCPDManagerListArgs.Size(m)
}
func (m *DHCPDManagerListArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_DHCPDManagerListArgs.DiscardUnknown(m)
}

var xxx_messageInfo_DHCPDManagerListArgs proto.InternalMessageInfo

type DHCPDManagedId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DHCPDManagedId) Reset()         { *m = DHCPDManagedId{} }
func (m *DHCPDManagedId) String() string { return proto.CompactTextString(m) }
func (*DHCPDManagedId) ProtoMessage()    {}
func (*DHCPDManagedId) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f4f7f9a47f6dd7, []int{5}
}

func (m *DHCPDManagedId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DHCPDManagedId.Unmarshal(m, b)
}
func (m *DHCPDManagedId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DHCPDManagedId.Marshal(b, m, deterministic)
}
func (m *DHCPDManagedId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DHCPDManagedId.Merge(m, src)
}
func (m *DHCPDManagedId) XXX_Size() int {
	return xxx_messageInfo_DHCPDManagedId.Size(m)
}
func (m *DHCPDManagedId) XXX_DiscardUnknown() {
	xxx_messageInfo_DHCPDManagedId.DiscardUnknown(m)
}

var xxx_messageInfo_DHCPDManagedId proto.InternalMessageInfo

func (m *DHCPDManagedId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DHCPDManagerListReply struct {
	DHCPDsManaged        []*DHCPDManaged `protobuf:"bytes,1,rep,name=DHCPDsManaged,proto3" json:"DHCPDsManaged,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DHCPDManagerListReply) Reset()         { *m = DHCPDManagerListReply{} }
func (m *DHCPDManagerListReply) String() string { return proto.CompactTextString(m) }
func (*DHCPDManagerListReply) ProtoMessage()    {}
func (*DHCPDManagerListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f4f7f9a47f6dd7, []int{6}
}

func (m *DHCPDManagerListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DHCPDManagerListReply.Unmarshal(m, b)
}
func (m *DHCPDManagerListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DHCPDManagerListReply.Marshal(b, m, deterministic)
}
func (m *DHCPDManagerListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DHCPDManagerListReply.Merge(m, src)
}
func (m *DHCPDManagerListReply) XXX_Size() int {
	return xxx_messageInfo_DHCPDManagerListReply.Size(m)
}
func (m *DHCPDManagerListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DHCPDManagerListReply.DiscardUnknown(m)
}

var xxx_messageInfo_DHCPDManagerListReply proto.InternalMessageInfo

func (m *DHCPDManagerListReply) GetDHCPDsManaged() []*DHCPDManaged {
	if m != nil {
		return m.DHCPDsManaged
	}
	return nil
}

func init() {
	proto.RegisterType((*DHCPD)(nil), "goISCDHCP.DHCPD")
	proto.RegisterType((*Subnet)(nil), "goISCDHCP.Subnet")
	proto.RegisterType((*Range)(nil), "goISCDHCP.Range")
	proto.RegisterType((*DHCPDManaged)(nil), "goISCDHCP.DHCPDManaged")
	proto.RegisterType((*DHCPDManagerListArgs)(nil), "goISCDHCP.DHCPDManagerListArgs")
	proto.RegisterType((*DHCPDManagedId)(nil), "goISCDHCP.DHCPDManagedId")
	proto.RegisterType((*DHCPDManagerListReply)(nil), "goISCDHCP.DHCPDManagerListReply")
}

func init() {
	proto.RegisterFile("dhcpd.proto", fileDescriptor_87f4f7f9a47f6dd7)
}

var fileDescriptor_87f4f7f9a47f6dd7 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdd, 0xce, 0xd2, 0x40,
	0x10, 0xa5, 0x2d, 0x2d, 0x32, 0x28, 0xc1, 0x09, 0xe2, 0xda, 0x0b, 0x6d, 0xf6, 0xc2, 0x90, 0x98,
	0x60, 0x82, 0x17, 0x5e, 0x18, 0x13, 0x0d, 0xf8, 0xd3, 0x88, 0xc4, 0xb4, 0x89, 0xf7, 0x85, 0x4e,
	0x2a, 0x01, 0x5a, 0xb2, 0x5d, 0x51, 0x9f, 0xc3, 0x17, 0xf0, 0x51, 0x4d, 0xb7, 0x4b, 0x53, 0xf8,
	0xbe, 0xf2, 0x5d, 0x75, 0x67, 0xcf, 0x39, 0x33, 0xe7, 0x4c, 0xb3, 0xd0, 0x8b, 0x7f, 0xac, 0x0f,
	0xf1, 0xe4, 0x20, 0x32, 0x99, 0x61, 0x37, 0xc9, 0xfc, 0x70, 0x36, 0xff, 0x3c, 0xfb, 0xc6, 0x17,
	0x60, 0x17, 0xdf, 0x39, 0x8e, 0xc0, 0x99, 0xd3, 0x71, 0xb3, 0x26, 0x66, 0x78, 0xc6, 0xb8, 0x1b,
	0xe8, 0x0a, 0x5f, 0x40, 0x27, 0xfc, 0xb9, 0x4a, 0x49, 0xe6, 0xcc, 0xf4, 0xac, 0x71, 0x6f, 0xfa,
	0x70, 0x52, 0xa9, 0x27, 0x25, 0x12, 0x9c, 0x18, 0xfc, 0x9f, 0x01, 0x4e, 0x79, 0x46, 0x06, 0x9d,
	0x25, 0xc9, 0x5f, 0x99, 0xd8, 0xea, 0x86, 0xa7, 0x52, 0x23, 0xfb, 0x28, 0xdf, 0x32, 0xb3, 0x42,
	0x8a, 0x12, 0x9f, 0x02, 0x2c, 0xe9, 0xb7, 0x0c, 0x49, 0x1c, 0x49, 0x30, 0x4b, 0x81, 0xb5, 0x1b,
	0x74, 0xe1, 0xde, 0xc7, 0xcd, 0x8e, 0xd2, 0x68, 0x4f, 0xac, 0xad, 0xd0, 0xaa, 0xc6, 0xe7, 0x60,
	0x07, 0x51, 0x9a, 0x10, 0xb3, 0x3d, 0x63, 0xdc, 0x9b, 0x0e, 0x6a, 0x2e, 0xd5, 0x7d, 0x50, 0xc2,
	0xfc, 0xa5, 0xe6, 0xe1, 0x10, 0xec, 0x50, 0x46, 0x42, 0x6a, 0x7b, 0x65, 0x81, 0x03, 0xb0, 0x3e,
	0xa4, 0xb1, 0x36, 0x56, 0x1c, 0xf9, 0x1a, 0xee, 0xab, 0x0d, 0x7d, 0x8d, 0xd2, 0x28, 0xa1, 0x18,
	0xfb, 0x60, 0xfa, 0xb1, 0x16, 0x99, 0x7e, 0x5c, 0x5b, 0x9c, 0xd9, 0xb4, 0x38, 0xeb, 0xce, 0xc5,
	0x8d, 0x60, 0x58, 0x1b, 0x22, 0x16, 0x9b, 0x5c, 0xbe, 0x17, 0x49, 0xce, 0x3d, 0xe8, 0xd7, 0x87,
	0xfb, 0x37, 0xc6, 0xf3, 0xef, 0xf0, 0xe8, 0x52, 0x19, 0xd0, 0x61, 0xf7, 0x07, 0xdf, 0xc2, 0x03,
	0x05, 0xe4, 0x5a, 0xcb, 0x0c, 0xe5, 0xe2, 0x71, 0xcd, 0x45, 0xbd, 0x75, 0x70, 0xce, 0x9e, 0xfe,
	0x35, 0xcf, 0x72, 0x0b, 0x7c, 0x0d, 0xce, 0x4c, 0x50, 0x24, 0x09, 0x07, 0x97, 0x2d, 0xdc, 0x27,
	0x0d, 0x4d, 0xfd, 0x98, 0xb7, 0xf0, 0x0b, 0xb4, 0x0b, 0x57, 0xf8, 0xec, 0x76, 0x52, 0x15, 0xd6,
	0xf5, 0xae, 0x10, 0x54, 0x26, 0xde, 0xc2, 0x37, 0x60, 0x7d, 0x22, 0x89, 0xcd, 0x03, 0xdd, 0xa6,
	0x80, 0xbc, 0x85, 0xef, 0x8a, 0x5f, 0xb5, 0x23, 0x49, 0xd7, 0xf4, 0xd7, 0xb2, 0xac, 0x1c, 0xf5,
	0x80, 0x5e, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x67, 0x7f, 0xa2, 0x53, 0x4f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DHCPDManagerClient is the client API for DHCPDManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DHCPDManagerClient interface {
	Create(ctx context.Context, in *DHCPD, opts ...grpc.CallOption) (*DHCPDManagedId, error)
	List(ctx context.Context, in *DHCPDManagerListArgs, opts ...grpc.CallOption) (*DHCPDManagerListReply, error)
	Get(ctx context.Context, in *DHCPDManagedId, opts ...grpc.CallOption) (*DHCPDManaged, error)
	Delete(ctx context.Context, in *DHCPDManagedId, opts ...grpc.CallOption) (*DHCPDManagedId, error)
}

type dHCPDManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewDHCPDManagerClient(cc grpc.ClientConnInterface) DHCPDManagerClient {
	return &dHCPDManagerClient{cc}
}

func (c *dHCPDManagerClient) Create(ctx context.Context, in *DHCPD, opts ...grpc.CallOption) (*DHCPDManagedId, error) {
	out := new(DHCPDManagedId)
	err := c.cc.Invoke(ctx, "/goISCDHCP.DHCPDManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dHCPDManagerClient) List(ctx context.Context, in *DHCPDManagerListArgs, opts ...grpc.CallOption) (*DHCPDManagerListReply, error) {
	out := new(DHCPDManagerListReply)
	err := c.cc.Invoke(ctx, "/goISCDHCP.DHCPDManager/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dHCPDManagerClient) Get(ctx context.Context, in *DHCPDManagedId, opts ...grpc.CallOption) (*DHCPDManaged, error) {
	out := new(DHCPDManaged)
	err := c.cc.Invoke(ctx, "/goISCDHCP.DHCPDManager/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dHCPDManagerClient) Delete(ctx context.Context, in *DHCPDManagedId, opts ...grpc.CallOption) (*DHCPDManagedId, error) {
	out := new(DHCPDManagedId)
	err := c.cc.Invoke(ctx, "/goISCDHCP.DHCPDManager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DHCPDManagerServer is the server API for DHCPDManager service.
type DHCPDManagerServer interface {
	Create(context.Context, *DHCPD) (*DHCPDManagedId, error)
	List(context.Context, *DHCPDManagerListArgs) (*DHCPDManagerListReply, error)
	Get(context.Context, *DHCPDManagedId) (*DHCPDManaged, error)
	Delete(context.Context, *DHCPDManagedId) (*DHCPDManagedId, error)
}

// UnimplementedDHCPDManagerServer can be embedded to have forward compatible implementations.
type UnimplementedDHCPDManagerServer struct {
}

func (*UnimplementedDHCPDManagerServer) Create(ctx context.Context, req *DHCPD) (*DHCPDManagedId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedDHCPDManagerServer) List(ctx context.Context, req *DHCPDManagerListArgs) (*DHCPDManagerListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedDHCPDManagerServer) Get(ctx context.Context, req *DHCPDManagedId) (*DHCPDManaged, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDHCPDManagerServer) Delete(ctx context.Context, req *DHCPDManagedId) (*DHCPDManagedId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterDHCPDManagerServer(s *grpc.Server, srv DHCPDManagerServer) {
	s.RegisterService(&_DHCPDManager_serviceDesc, srv)
}

func _DHCPDManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DHCPD)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHCPDManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goISCDHCP.DHCPDManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHCPDManagerServer).Create(ctx, req.(*DHCPD))
	}
	return interceptor(ctx, in, info, handler)
}

func _DHCPDManager_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DHCPDManagerListArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHCPDManagerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goISCDHCP.DHCPDManager/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHCPDManagerServer).List(ctx, req.(*DHCPDManagerListArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _DHCPDManager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DHCPDManagedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHCPDManagerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goISCDHCP.DHCPDManager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHCPDManagerServer).Get(ctx, req.(*DHCPDManagedId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DHCPDManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DHCPDManagedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHCPDManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goISCDHCP.DHCPDManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHCPDManagerServer).Delete(ctx, req.(*DHCPDManagedId))
	}
	return interceptor(ctx, in, info, handler)
}

var _DHCPDManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goISCDHCP.DHCPDManager",
	HandlerType: (*DHCPDManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DHCPDManager_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DHCPDManager_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DHCPDManager_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DHCPDManager_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dhcpd.proto",
}
