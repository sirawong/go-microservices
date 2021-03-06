// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/holders/holers.proto

package grpc_holders

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

type ListHoldersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListHoldersRequest) Reset()         { *m = ListHoldersRequest{} }
func (m *ListHoldersRequest) String() string { return proto.CompactTextString(m) }
func (*ListHoldersRequest) ProtoMessage()    {}
func (*ListHoldersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_89804eb1b95fdd75, []int{0}
}

func (m *ListHoldersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHoldersRequest.Unmarshal(m, b)
}
func (m *ListHoldersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHoldersRequest.Marshal(b, m, deterministic)
}
func (m *ListHoldersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHoldersRequest.Merge(m, src)
}
func (m *ListHoldersRequest) XXX_Size() int {
	return xxx_messageInfo_ListHoldersRequest.Size(m)
}
func (m *ListHoldersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHoldersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListHoldersRequest proto.InternalMessageInfo

type ListHoldersResponse struct {
	Holders              []*Holder `protobuf:"bytes,1,rep,name=holders,proto3" json:"holders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListHoldersResponse) Reset()         { *m = ListHoldersResponse{} }
func (m *ListHoldersResponse) String() string { return proto.CompactTextString(m) }
func (*ListHoldersResponse) ProtoMessage()    {}
func (*ListHoldersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89804eb1b95fdd75, []int{1}
}

func (m *ListHoldersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHoldersResponse.Unmarshal(m, b)
}
func (m *ListHoldersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHoldersResponse.Marshal(b, m, deterministic)
}
func (m *ListHoldersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHoldersResponse.Merge(m, src)
}
func (m *ListHoldersResponse) XXX_Size() int {
	return xxx_messageInfo_ListHoldersResponse.Size(m)
}
func (m *ListHoldersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHoldersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListHoldersResponse proto.InternalMessageInfo

func (m *ListHoldersResponse) GetHolders() []*Holder {
	if m != nil {
		return m.Holders
	}
	return nil
}

type GetHolderByMovieIdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHolderByMovieIdRequest) Reset()         { *m = GetHolderByMovieIdRequest{} }
func (m *GetHolderByMovieIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetHolderByMovieIdRequest) ProtoMessage()    {}
func (*GetHolderByMovieIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_89804eb1b95fdd75, []int{2}
}

func (m *GetHolderByMovieIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHolderByMovieIdRequest.Unmarshal(m, b)
}
func (m *GetHolderByMovieIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHolderByMovieIdRequest.Marshal(b, m, deterministic)
}
func (m *GetHolderByMovieIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHolderByMovieIdRequest.Merge(m, src)
}
func (m *GetHolderByMovieIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetHolderByMovieIdRequest.Size(m)
}
func (m *GetHolderByMovieIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHolderByMovieIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHolderByMovieIdRequest proto.InternalMessageInfo

func (m *GetHolderByMovieIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetHolderByMovieIdResponse struct {
	Holder               *Holder  `protobuf:"bytes,1,opt,name=holder,proto3" json:"holder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHolderByMovieIdResponse) Reset()         { *m = GetHolderByMovieIdResponse{} }
func (m *GetHolderByMovieIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetHolderByMovieIdResponse) ProtoMessage()    {}
func (*GetHolderByMovieIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89804eb1b95fdd75, []int{3}
}

func (m *GetHolderByMovieIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHolderByMovieIdResponse.Unmarshal(m, b)
}
func (m *GetHolderByMovieIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHolderByMovieIdResponse.Marshal(b, m, deterministic)
}
func (m *GetHolderByMovieIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHolderByMovieIdResponse.Merge(m, src)
}
func (m *GetHolderByMovieIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetHolderByMovieIdResponse.Size(m)
}
func (m *GetHolderByMovieIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHolderByMovieIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHolderByMovieIdResponse proto.InternalMessageInfo

func (m *GetHolderByMovieIdResponse) GetHolder() *Holder {
	if m != nil {
		return m.Holder
	}
	return nil
}

type GetHolderRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHolderRequest) Reset()         { *m = GetHolderRequest{} }
func (m *GetHolderRequest) String() string { return proto.CompactTextString(m) }
func (*GetHolderRequest) ProtoMessage()    {}
func (*GetHolderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_89804eb1b95fdd75, []int{4}
}

func (m *GetHolderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHolderRequest.Unmarshal(m, b)
}
func (m *GetHolderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHolderRequest.Marshal(b, m, deterministic)
}
func (m *GetHolderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHolderRequest.Merge(m, src)
}
func (m *GetHolderRequest) XXX_Size() int {
	return xxx_messageInfo_GetHolderRequest.Size(m)
}
func (m *GetHolderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHolderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHolderRequest proto.InternalMessageInfo

func (m *GetHolderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetHolderResponse struct {
	Holder               *Holder  `protobuf:"bytes,1,opt,name=holder,proto3" json:"holder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHolderResponse) Reset()         { *m = GetHolderResponse{} }
func (m *GetHolderResponse) String() string { return proto.CompactTextString(m) }
func (*GetHolderResponse) ProtoMessage()    {}
func (*GetHolderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89804eb1b95fdd75, []int{5}
}

func (m *GetHolderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHolderResponse.Unmarshal(m, b)
}
func (m *GetHolderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHolderResponse.Marshal(b, m, deterministic)
}
func (m *GetHolderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHolderResponse.Merge(m, src)
}
func (m *GetHolderResponse) XXX_Size() int {
	return xxx_messageInfo_GetHolderResponse.Size(m)
}
func (m *GetHolderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHolderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHolderResponse proto.InternalMessageInfo

func (m *GetHolderResponse) GetHolder() *Holder {
	if m != nil {
		return m.Holder
	}
	return nil
}

type AddHolderRequest struct {
	Holder               *Holder  `protobuf:"bytes,1,opt,name=holder,proto3" json:"holder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddHolderRequest) Reset()         { *m = AddHolderRequest{} }
func (m *AddHolderRequest) String() string { return proto.CompactTextString(m) }
func (*AddHolderRequest) ProtoMessage()    {}
func (*AddHolderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_89804eb1b95fdd75, []int{6}
}

func (m *AddHolderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddHolderRequest.Unmarshal(m, b)
}
func (m *AddHolderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddHolderRequest.Marshal(b, m, deterministic)
}
func (m *AddHolderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddHolderRequest.Merge(m, src)
}
func (m *AddHolderRequest) XXX_Size() int {
	return xxx_messageInfo_AddHolderRequest.Size(m)
}
func (m *AddHolderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddHolderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddHolderRequest proto.InternalMessageInfo

func (m *AddHolderRequest) GetHolder() *Holder {
	if m != nil {
		return m.Holder
	}
	return nil
}

type AddHolderResponse struct {
	Holder               *Holder  `protobuf:"bytes,1,opt,name=holder,proto3" json:"holder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddHolderResponse) Reset()         { *m = AddHolderResponse{} }
func (m *AddHolderResponse) String() string { return proto.CompactTextString(m) }
func (*AddHolderResponse) ProtoMessage()    {}
func (*AddHolderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89804eb1b95fdd75, []int{7}
}

func (m *AddHolderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddHolderResponse.Unmarshal(m, b)
}
func (m *AddHolderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddHolderResponse.Marshal(b, m, deterministic)
}
func (m *AddHolderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddHolderResponse.Merge(m, src)
}
func (m *AddHolderResponse) XXX_Size() int {
	return xxx_messageInfo_AddHolderResponse.Size(m)
}
func (m *AddHolderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddHolderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddHolderResponse proto.InternalMessageInfo

func (m *AddHolderResponse) GetHolder() *Holder {
	if m != nil {
		return m.Holder
	}
	return nil
}

type UpdateHolderRequest struct {
	Holder               *Holder  `protobuf:"bytes,1,opt,name=holder,proto3" json:"holder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateHolderRequest) Reset()         { *m = UpdateHolderRequest{} }
func (m *UpdateHolderRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateHolderRequest) ProtoMessage()    {}
func (*UpdateHolderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_89804eb1b95fdd75, []int{8}
}

func (m *UpdateHolderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateHolderRequest.Unmarshal(m, b)
}
func (m *UpdateHolderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateHolderRequest.Marshal(b, m, deterministic)
}
func (m *UpdateHolderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateHolderRequest.Merge(m, src)
}
func (m *UpdateHolderRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateHolderRequest.Size(m)
}
func (m *UpdateHolderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateHolderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateHolderRequest proto.InternalMessageInfo

func (m *UpdateHolderRequest) GetHolder() *Holder {
	if m != nil {
		return m.Holder
	}
	return nil
}

type UpdateHolderResponse struct {
	Holder               *Holder  `protobuf:"bytes,1,opt,name=holder,proto3" json:"holder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateHolderResponse) Reset()         { *m = UpdateHolderResponse{} }
func (m *UpdateHolderResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateHolderResponse) ProtoMessage()    {}
func (*UpdateHolderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89804eb1b95fdd75, []int{9}
}

func (m *UpdateHolderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateHolderResponse.Unmarshal(m, b)
}
func (m *UpdateHolderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateHolderResponse.Marshal(b, m, deterministic)
}
func (m *UpdateHolderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateHolderResponse.Merge(m, src)
}
func (m *UpdateHolderResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateHolderResponse.Size(m)
}
func (m *UpdateHolderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateHolderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateHolderResponse proto.InternalMessageInfo

func (m *UpdateHolderResponse) GetHolder() *Holder {
	if m != nil {
		return m.Holder
	}
	return nil
}

func init() {
	proto.RegisterType((*ListHoldersRequest)(nil), "grpc.holders.ListHoldersRequest")
	proto.RegisterType((*ListHoldersResponse)(nil), "grpc.holders.ListHoldersResponse")
	proto.RegisterType((*GetHolderByMovieIdRequest)(nil), "grpc.holders.GetHolderByMovieIdRequest")
	proto.RegisterType((*GetHolderByMovieIdResponse)(nil), "grpc.holders.GetHolderByMovieIdResponse")
	proto.RegisterType((*GetHolderRequest)(nil), "grpc.holders.GetHolderRequest")
	proto.RegisterType((*GetHolderResponse)(nil), "grpc.holders.GetHolderResponse")
	proto.RegisterType((*AddHolderRequest)(nil), "grpc.holders.AddHolderRequest")
	proto.RegisterType((*AddHolderResponse)(nil), "grpc.holders.AddHolderResponse")
	proto.RegisterType((*UpdateHolderRequest)(nil), "grpc.holders.UpdateHolderRequest")
	proto.RegisterType((*UpdateHolderResponse)(nil), "grpc.holders.UpdateHolderResponse")
}

func init() { proto.RegisterFile("proto/holders/holers.proto", fileDescriptor_89804eb1b95fdd75) }

var fileDescriptor_89804eb1b95fdd75 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdd, 0x4a, 0xc3, 0x30,
	0x1c, 0xc5, 0xe9, 0x06, 0x93, 0x9d, 0x0d, 0xd9, 0xb2, 0x5e, 0xcc, 0x5c, 0x68, 0x0d, 0x82, 0x05,
	0x25, 0xc2, 0x7c, 0x01, 0xeb, 0x07, 0x3a, 0x99, 0x20, 0x85, 0x3d, 0xc0, 0x34, 0xa1, 0x16, 0xd4,
	0xd6, 0xa6, 0x0a, 0xbe, 0x9b, 0x0f, 0x27, 0x6b, 0xb3, 0xda, 0x8f, 0x75, 0x62, 0xbd, 0x0a, 0xe4,
	0x7f, 0xce, 0xef, 0xfc, 0x69, 0x0e, 0x05, 0x0d, 0xa3, 0x20, 0x0e, 0x4e, 0x9e, 0x82, 0x67, 0x21,
	0x23, 0xb5, 0x3c, 0x65, 0xa4, 0x78, 0x72, 0x49, 0xfa, 0x5e, 0x14, 0x3e, 0x72, 0x3d, 0xa2, 0x07,
	0x15, 0xe5, 0xf2, 0xe4, 0x2f, 0x52, 0xa9, 0x85, 0x27, 0xb5, 0x87, 0x99, 0x20, 0x33, 0x5f, 0xc5,
	0x37, 0xe9, 0xd4, 0x95, 0x6f, 0xef, 0x52, 0xc5, 0xec, 0x0a, 0xa3, 0xc2, 0xad, 0x0a, 0x83, 0x57,
	0x25, 0x09, 0xc7, 0x96, 0xc6, 0x8c, 0x0d, 0xab, 0x6d, 0xf7, 0x26, 0x26, 0xcf, 0x47, 0xf2, 0x54,
	0xef, 0xae, 0x44, 0xec, 0x08, 0x3b, 0xd7, 0x52, 0x53, 0xce, 0x3f, 0xef, 0x82, 0x0f, 0x5f, 0x4e,
	0x85, 0xce, 0x20, 0xdb, 0x68, 0xf9, 0x62, 0x6c, 0x58, 0x86, 0xdd, 0x75, 0x5b, 0xbe, 0x60, 0xb7,
	0xa0, 0xeb, 0xc4, 0x3a, 0xfa, 0x18, 0x9d, 0x94, 0x9a, 0x38, 0xea, 0x92, 0xb5, 0x86, 0x31, 0x0c,
	0x32, 0x56, 0x5d, 0x9e, 0x83, 0x61, 0x4e, 0xd3, 0x28, 0xe6, 0x0c, 0x03, 0x47, 0x88, 0x62, 0xcc,
	0xdf, 0x08, 0x0e, 0x86, 0x39, 0x42, 0xa3, 0x25, 0x2e, 0x30, 0x9a, 0x87, 0x62, 0x11, 0xcb, 0xff,
	0xec, 0x71, 0x09, 0xb3, 0x08, 0x69, 0xb2, 0xca, 0xe4, 0xab, 0x0d, 0xe8, 0xce, 0x38, 0xf7, 0x53,
	0xe2, 0xa2, 0x97, 0x6b, 0x11, 0xb1, 0x8a, 0xde, 0x6a, 0xed, 0xe8, 0xfe, 0x06, 0x85, 0x5e, 0xc8,
	0x03, 0xa9, 0xb6, 0x84, 0x1c, 0x16, 0x8d, 0xb5, 0xa5, 0xa3, 0xf6, 0xef, 0x42, 0x1d, 0x34, 0x43,
	0x37, 0x9b, 0x92, 0xdd, 0x1a, 0xdb, 0x0a, 0xbb, 0x57, 0x3b, 0xff, 0xa1, 0x65, 0xef, 0x5c, 0xa6,
	0x95, 0x2b, 0x54, 0xa6, 0x55, 0x0b, 0x32, 0x47, 0x3f, 0xff, 0x5a, 0xa4, 0xf4, 0xdd, 0xd6, 0xd4,
	0x81, 0xb2, 0x4d, 0x92, 0x14, 0xfb, 0xd0, 0x49, 0x7e, 0x09, 0xa7, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x6e, 0x74, 0xda, 0xe2, 0x64, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HoldersAPIClient is the client API for HoldersAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HoldersAPIClient interface {
	// Get all customers
	ListHolders(ctx context.Context, in *ListHoldersRequest, opts ...grpc.CallOption) (*ListHoldersResponse, error)
	// Get customer by Movie ID
	GetHolderByMovieId(ctx context.Context, in *GetHolderByMovieIdRequest, opts ...grpc.CallOption) (*GetHolderByMovieIdResponse, error)
	// Get customer by ID
	GetHolder(ctx context.Context, in *GetHolderRequest, opts ...grpc.CallOption) (*GetHolderResponse, error)
	// Add new customer
	AddHolder(ctx context.Context, in *AddHolderRequest, opts ...grpc.CallOption) (*AddHolderResponse, error)
	// Update customer
	UpdateHolder(ctx context.Context, in *UpdateHolderRequest, opts ...grpc.CallOption) (*UpdateHolderResponse, error)
}

type holdersAPIClient struct {
	cc *grpc.ClientConn
}

func NewHoldersAPIClient(cc *grpc.ClientConn) HoldersAPIClient {
	return &holdersAPIClient{cc}
}

func (c *holdersAPIClient) ListHolders(ctx context.Context, in *ListHoldersRequest, opts ...grpc.CallOption) (*ListHoldersResponse, error) {
	out := new(ListHoldersResponse)
	err := c.cc.Invoke(ctx, "/grpc.holders.HoldersAPI/ListHolders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdersAPIClient) GetHolderByMovieId(ctx context.Context, in *GetHolderByMovieIdRequest, opts ...grpc.CallOption) (*GetHolderByMovieIdResponse, error) {
	out := new(GetHolderByMovieIdResponse)
	err := c.cc.Invoke(ctx, "/grpc.holders.HoldersAPI/GetHolderByMovieId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdersAPIClient) GetHolder(ctx context.Context, in *GetHolderRequest, opts ...grpc.CallOption) (*GetHolderResponse, error) {
	out := new(GetHolderResponse)
	err := c.cc.Invoke(ctx, "/grpc.holders.HoldersAPI/GetHolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdersAPIClient) AddHolder(ctx context.Context, in *AddHolderRequest, opts ...grpc.CallOption) (*AddHolderResponse, error) {
	out := new(AddHolderResponse)
	err := c.cc.Invoke(ctx, "/grpc.holders.HoldersAPI/AddHolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holdersAPIClient) UpdateHolder(ctx context.Context, in *UpdateHolderRequest, opts ...grpc.CallOption) (*UpdateHolderResponse, error) {
	out := new(UpdateHolderResponse)
	err := c.cc.Invoke(ctx, "/grpc.holders.HoldersAPI/UpdateHolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HoldersAPIServer is the server API for HoldersAPI service.
type HoldersAPIServer interface {
	// Get all customers
	ListHolders(context.Context, *ListHoldersRequest) (*ListHoldersResponse, error)
	// Get customer by Movie ID
	GetHolderByMovieId(context.Context, *GetHolderByMovieIdRequest) (*GetHolderByMovieIdResponse, error)
	// Get customer by ID
	GetHolder(context.Context, *GetHolderRequest) (*GetHolderResponse, error)
	// Add new customer
	AddHolder(context.Context, *AddHolderRequest) (*AddHolderResponse, error)
	// Update customer
	UpdateHolder(context.Context, *UpdateHolderRequest) (*UpdateHolderResponse, error)
}

// UnimplementedHoldersAPIServer can be embedded to have forward compatible implementations.
type UnimplementedHoldersAPIServer struct {
}

func (*UnimplementedHoldersAPIServer) ListHolders(ctx context.Context, req *ListHoldersRequest) (*ListHoldersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHolders not implemented")
}
func (*UnimplementedHoldersAPIServer) GetHolderByMovieId(ctx context.Context, req *GetHolderByMovieIdRequest) (*GetHolderByMovieIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHolderByMovieId not implemented")
}
func (*UnimplementedHoldersAPIServer) GetHolder(ctx context.Context, req *GetHolderRequest) (*GetHolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHolder not implemented")
}
func (*UnimplementedHoldersAPIServer) AddHolder(ctx context.Context, req *AddHolderRequest) (*AddHolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHolder not implemented")
}
func (*UnimplementedHoldersAPIServer) UpdateHolder(ctx context.Context, req *UpdateHolderRequest) (*UpdateHolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHolder not implemented")
}

func RegisterHoldersAPIServer(s *grpc.Server, srv HoldersAPIServer) {
	s.RegisterService(&_HoldersAPI_serviceDesc, srv)
}

func _HoldersAPI_ListHolders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHoldersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldersAPIServer).ListHolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.holders.HoldersAPI/ListHolders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldersAPIServer).ListHolders(ctx, req.(*ListHoldersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldersAPI_GetHolderByMovieId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHolderByMovieIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldersAPIServer).GetHolderByMovieId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.holders.HoldersAPI/GetHolderByMovieId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldersAPIServer).GetHolderByMovieId(ctx, req.(*GetHolderByMovieIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldersAPI_GetHolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldersAPIServer).GetHolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.holders.HoldersAPI/GetHolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldersAPIServer).GetHolder(ctx, req.(*GetHolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldersAPI_AddHolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldersAPIServer).AddHolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.holders.HoldersAPI/AddHolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldersAPIServer).AddHolder(ctx, req.(*AddHolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoldersAPI_UpdateHolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoldersAPIServer).UpdateHolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.holders.HoldersAPI/UpdateHolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoldersAPIServer).UpdateHolder(ctx, req.(*UpdateHolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HoldersAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.holders.HoldersAPI",
	HandlerType: (*HoldersAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListHolders",
			Handler:    _HoldersAPI_ListHolders_Handler,
		},
		{
			MethodName: "GetHolderByMovieId",
			Handler:    _HoldersAPI_GetHolderByMovieId_Handler,
		},
		{
			MethodName: "GetHolder",
			Handler:    _HoldersAPI_GetHolder_Handler,
		},
		{
			MethodName: "AddHolder",
			Handler:    _HoldersAPI_AddHolder_Handler,
		},
		{
			MethodName: "UpdateHolder",
			Handler:    _HoldersAPI_UpdateHolder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/holders/holers.proto",
}
