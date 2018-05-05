// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package go_micro_srv_vessel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_b980bac217217251, []int{0}
}
func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (dst *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(dst, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_b980bac217217251, []int{1}
}
func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (dst *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(dst, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels" json:"vessels,omitempty"`
	Created              bool      `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_b980bac217217251, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func init() {
	proto.RegisterType((*Vessel)(nil), "go.micro.srv.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "go.micro.srv.vessel.Specification")
	proto.RegisterType((*Response)(nil), "go.micro.srv.vessel.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
	Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselServiceClient) Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
	Create(context.Context, *Vessel, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func (h *VesselService) Create(ctx context.Context, in *Vessel, out *Response) error {
	return h.VesselServiceHandler.Create(ctx, in, out)
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_vessel_b980bac217217251) }

var fileDescriptor_vessel_b980bac217217251 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0xdd, 0x02, 0xa5, 0x8c, 0xc1, 0xc3, 0x78, 0x59, 0x50, 0x92, 0xa6, 0x27, 0x4e, 0x35,
	0x81, 0xf8, 0x00, 0xc6, 0x84, 0x44, 0x8f, 0xc5, 0xe8, 0x91, 0x2c, 0xdb, 0x11, 0x37, 0xa1, 0xdd,
	0xa6, 0x6d, 0x0a, 0xbe, 0x85, 0x8f, 0xe0, 0x0b, 0xf8, 0x8e, 0x26, 0x53, 0xc0, 0x60, 0x1a, 0x38,
	0x75, 0xfe, 0x7c, 0x33, 0xf3, 0xf5, 0x97, 0x85, 0x41, 0x96, 0xdb, 0xd2, 0xde, 0x55, 0x54, 0x14,
	0xb4, 0xde, 0x7d, 0x42, 0xae, 0xe1, 0xf5, 0xca, 0x86, 0x89, 0xd1, 0xb9, 0x0d, 0x8b, 0xbc, 0x0a,
	0xeb, 0x56, 0xf0, 0x2d, 0xc0, 0x7d, 0xe5, 0x10, 0xaf, 0xc0, 0x31, 0xb1, 0x14, 0xbe, 0x18, 0xf7,
	0x22, 0xc7, 0xc4, 0x38, 0x04, 0x4f, 0xab, 0x4c, 0x69, 0x53, 0x7e, 0x4a, 0xc7, 0x17, 0xe3, 0x4e,
	0x74, 0xc8, 0x71, 0x04, 0x90, 0xa8, 0xed, 0x62, 0x43, 0x66, 0xf5, 0x51, 0xca, 0x16, 0x77, 0x7b,
	0x89, 0xda, 0xbe, 0x71, 0x01, 0x11, 0xda, 0xa9, 0x4a, 0x48, 0xb6, 0x79, 0x19, 0xc7, 0x78, 0x0b,
	0x3d, 0x55, 0x29, 0xb3, 0x56, 0xcb, 0x35, 0xc9, 0x8e, 0x2f, 0xc6, 0x5e, 0xf4, 0x57, 0xc0, 0x01,
	0x78, 0x76, 0x93, 0x52, 0xbe, 0x30, 0xb1, 0x74, 0x79, 0xaa, 0xcb, 0xf9, 0x53, 0x1c, 0x3c, 0x43,
	0x7f, 0x9e, 0x91, 0x36, 0xef, 0x46, 0xab, 0xd2, 0xd8, 0xf4, 0xc8, 0x98, 0x38, 0x69, 0xcc, 0xf9,
	0x67, 0x2c, 0xf8, 0x12, 0xe0, 0x45, 0x54, 0x64, 0x36, 0x2d, 0x08, 0xa7, 0xe0, 0xd6, 0x14, 0x78,
	0xcb, 0xe5, 0xe4, 0x26, 0x6c, 0x20, 0x14, 0xd6, 0x74, 0xa2, 0x9d, 0x14, 0xef, 0xa1, 0x5b, 0x47,
	0x85, 0x74, 0xfc, 0xd6, 0xb9, 0xa9, 0xbd, 0x16, 0x25, 0x74, 0x75, 0x4e, 0xaa, 0xa4, 0x98, 0x69,
	0x79, 0xd1, 0x3e, 0x9d, 0xfc, 0x08, 0xe8, 0xd7, 0xea, 0x39, 0xe5, 0x95, 0xd1, 0x84, 0x2f, 0xd0,
	0x9f, 0x99, 0x34, 0x7e, 0x38, 0xc0, 0x09, 0x1a, 0x4f, 0x1c, 0x41, 0x19, 0x8e, 0x1a, 0x35, 0xfb,
	0x7f, 0x0d, 0x2e, 0x70, 0x06, 0xee, 0x23, 0x9f, 0xc4, 0x53, 0x8e, 0xcf, 0xee, 0x59, 0xba, 0xfc,
	0x9a, 0xa6, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x15, 0x18, 0x29, 0x6a, 0x02, 0x00, 0x00,
}
