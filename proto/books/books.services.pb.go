// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/books/books.services.proto

package grpc_books_v1

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

type ListBooksRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBooksRequest) Reset()         { *m = ListBooksRequest{} }
func (m *ListBooksRequest) String() string { return proto.CompactTextString(m) }
func (*ListBooksRequest) ProtoMessage()    {}
func (*ListBooksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f76168824b8688, []int{0}
}

func (m *ListBooksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBooksRequest.Unmarshal(m, b)
}
func (m *ListBooksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBooksRequest.Marshal(b, m, deterministic)
}
func (m *ListBooksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBooksRequest.Merge(m, src)
}
func (m *ListBooksRequest) XXX_Size() int {
	return xxx_messageInfo_ListBooksRequest.Size(m)
}
func (m *ListBooksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBooksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBooksRequest proto.InternalMessageInfo

type ListBooksResponse struct {
	Books                []*Book  `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBooksResponse) Reset()         { *m = ListBooksResponse{} }
func (m *ListBooksResponse) String() string { return proto.CompactTextString(m) }
func (*ListBooksResponse) ProtoMessage()    {}
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f76168824b8688, []int{1}
}

func (m *ListBooksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBooksResponse.Unmarshal(m, b)
}
func (m *ListBooksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBooksResponse.Marshal(b, m, deterministic)
}
func (m *ListBooksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBooksResponse.Merge(m, src)
}
func (m *ListBooksResponse) XXX_Size() int {
	return xxx_messageInfo_ListBooksResponse.Size(m)
}
func (m *ListBooksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBooksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBooksResponse proto.InternalMessageInfo

func (m *ListBooksResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type GetBookRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookRequest) Reset()         { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()    {}
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f76168824b8688, []int{2}
}

func (m *GetBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookRequest.Unmarshal(m, b)
}
func (m *GetBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookRequest.Marshal(b, m, deterministic)
}
func (m *GetBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookRequest.Merge(m, src)
}
func (m *GetBookRequest) XXX_Size() int {
	return xxx_messageInfo_GetBookRequest.Size(m)
}
func (m *GetBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookRequest proto.InternalMessageInfo

func (m *GetBookRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetBookResponse struct {
	Book                 *Book    `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBookResponse) Reset()         { *m = GetBookResponse{} }
func (m *GetBookResponse) String() string { return proto.CompactTextString(m) }
func (*GetBookResponse) ProtoMessage()    {}
func (*GetBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f76168824b8688, []int{3}
}

func (m *GetBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBookResponse.Unmarshal(m, b)
}
func (m *GetBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBookResponse.Marshal(b, m, deterministic)
}
func (m *GetBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBookResponse.Merge(m, src)
}
func (m *GetBookResponse) XXX_Size() int {
	return xxx_messageInfo_GetBookResponse.Size(m)
}
func (m *GetBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBookResponse proto.InternalMessageInfo

func (m *GetBookResponse) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type GetBooksRequest struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBooksRequest) Reset()         { *m = GetBooksRequest{} }
func (m *GetBooksRequest) String() string { return proto.CompactTextString(m) }
func (*GetBooksRequest) ProtoMessage()    {}
func (*GetBooksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f76168824b8688, []int{4}
}

func (m *GetBooksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBooksRequest.Unmarshal(m, b)
}
func (m *GetBooksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBooksRequest.Marshal(b, m, deterministic)
}
func (m *GetBooksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBooksRequest.Merge(m, src)
}
func (m *GetBooksRequest) XXX_Size() int {
	return xxx_messageInfo_GetBooksRequest.Size(m)
}
func (m *GetBooksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBooksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBooksRequest proto.InternalMessageInfo

func (m *GetBooksRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetBooksResponse struct {
	Books                []*Book  `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBooksResponse) Reset()         { *m = GetBooksResponse{} }
func (m *GetBooksResponse) String() string { return proto.CompactTextString(m) }
func (*GetBooksResponse) ProtoMessage()    {}
func (*GetBooksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f76168824b8688, []int{5}
}

func (m *GetBooksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBooksResponse.Unmarshal(m, b)
}
func (m *GetBooksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBooksResponse.Marshal(b, m, deterministic)
}
func (m *GetBooksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBooksResponse.Merge(m, src)
}
func (m *GetBooksResponse) XXX_Size() int {
	return xxx_messageInfo_GetBooksResponse.Size(m)
}
func (m *GetBooksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBooksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBooksResponse proto.InternalMessageInfo

func (m *GetBooksResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

type AddBookRequest struct {
	Book                 *Book    `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddBookRequest) Reset()         { *m = AddBookRequest{} }
func (m *AddBookRequest) String() string { return proto.CompactTextString(m) }
func (*AddBookRequest) ProtoMessage()    {}
func (*AddBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f76168824b8688, []int{6}
}

func (m *AddBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBookRequest.Unmarshal(m, b)
}
func (m *AddBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBookRequest.Marshal(b, m, deterministic)
}
func (m *AddBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBookRequest.Merge(m, src)
}
func (m *AddBookRequest) XXX_Size() int {
	return xxx_messageInfo_AddBookRequest.Size(m)
}
func (m *AddBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddBookRequest proto.InternalMessageInfo

func (m *AddBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type AddBookResponse struct {
	Book                 *Book    `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddBookResponse) Reset()         { *m = AddBookResponse{} }
func (m *AddBookResponse) String() string { return proto.CompactTextString(m) }
func (*AddBookResponse) ProtoMessage()    {}
func (*AddBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f76168824b8688, []int{7}
}

func (m *AddBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBookResponse.Unmarshal(m, b)
}
func (m *AddBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBookResponse.Marshal(b, m, deterministic)
}
func (m *AddBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBookResponse.Merge(m, src)
}
func (m *AddBookResponse) XXX_Size() int {
	return xxx_messageInfo_AddBookResponse.Size(m)
}
func (m *AddBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddBookResponse proto.InternalMessageInfo

func (m *AddBookResponse) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

type DeleteBookRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBookRequest) Reset()         { *m = DeleteBookRequest{} }
func (m *DeleteBookRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBookRequest) ProtoMessage()    {}
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f76168824b8688, []int{8}
}

func (m *DeleteBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBookRequest.Unmarshal(m, b)
}
func (m *DeleteBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBookRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBookRequest.Merge(m, src)
}
func (m *DeleteBookRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBookRequest.Size(m)
}
func (m *DeleteBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBookRequest proto.InternalMessageInfo

func (m *DeleteBookRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteBookResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBookResponse) Reset()         { *m = DeleteBookResponse{} }
func (m *DeleteBookResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteBookResponse) ProtoMessage()    {}
func (*DeleteBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f76168824b8688, []int{9}
}

func (m *DeleteBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBookResponse.Unmarshal(m, b)
}
func (m *DeleteBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBookResponse.Marshal(b, m, deterministic)
}
func (m *DeleteBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBookResponse.Merge(m, src)
}
func (m *DeleteBookResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteBookResponse.Size(m)
}
func (m *DeleteBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBookResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListBooksRequest)(nil), "grpc.books.v1.ListBooksRequest")
	proto.RegisterType((*ListBooksResponse)(nil), "grpc.books.v1.ListBooksResponse")
	proto.RegisterType((*GetBookRequest)(nil), "grpc.books.v1.GetBookRequest")
	proto.RegisterType((*GetBookResponse)(nil), "grpc.books.v1.GetBookResponse")
	proto.RegisterType((*GetBooksRequest)(nil), "grpc.books.v1.GetBooksRequest")
	proto.RegisterType((*GetBooksResponse)(nil), "grpc.books.v1.GetBooksResponse")
	proto.RegisterType((*AddBookRequest)(nil), "grpc.books.v1.AddBookRequest")
	proto.RegisterType((*AddBookResponse)(nil), "grpc.books.v1.AddBookResponse")
	proto.RegisterType((*DeleteBookRequest)(nil), "grpc.books.v1.DeleteBookRequest")
	proto.RegisterType((*DeleteBookResponse)(nil), "grpc.books.v1.DeleteBookResponse")
}

func init() { proto.RegisterFile("proto/books/books.services.proto", fileDescriptor_78f76168824b8688) }

var fileDescriptor_78f76168824b8688 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x4a, 0xfb, 0x50,
	0x10, 0xc5, 0x49, 0xf3, 0xff, 0x6b, 0x33, 0x62, 0x9a, 0x8e, 0x2e, 0x4a, 0x40, 0x1b, 0xaf, 0x0b,
	0xeb, 0x26, 0x62, 0x5d, 0x59, 0x50, 0xa8, 0x08, 0x2a, 0x8a, 0x68, 0xde, 0xc0, 0x36, 0x43, 0x09,
	0x7e, 0x24, 0xe6, 0xc6, 0x3e, 0x86, 0xcf, 0x2c, 0xb9, 0xf7, 0xe6, 0xdb, 0x14, 0xbb, 0x09, 0x21,
	0xf3, 0x9b, 0x33, 0x67, 0xe6, 0x10, 0x70, 0xa2, 0x38, 0x4c, 0xc2, 0x93, 0x59, 0x18, 0xbe, 0x72,
	0xf9, 0x74, 0x39, 0xc5, 0xcb, 0x60, 0x4e, 0xdc, 0x15, 0x25, 0xdc, 0x5e, 0xc4, 0xd1, 0xdc, 0x95,
	0xa5, 0xe5, 0xa9, 0xfd, 0x4b, 0xc3, 0x3b, 0x71, 0xfe, 0xb2, 0xc8, 0x1a, 0x18, 0x82, 0xf5, 0x10,
	0xf0, 0xe4, 0x2a, 0xad, 0x79, 0xf4, 0xf9, 0x45, 0x3c, 0x61, 0x97, 0xd0, 0x2f, 0x7d, 0xe3, 0x51,
	0xf8, 0xc1, 0x09, 0x8f, 0xe1, 0xbf, 0x10, 0x18, 0x68, 0x8e, 0x3e, 0xda, 0x1a, 0xef, 0xb8, 0x95,
	0x49, 0x6e, 0x0a, 0x7b, 0x92, 0x60, 0x0e, 0x98, 0x37, 0x24, 0xda, 0x95, 0x22, 0x9a, 0xd0, 0x09,
	0xfc, 0x81, 0xe6, 0x68, 0x23, 0xc3, 0xeb, 0x04, 0x3e, 0x9b, 0x40, 0x2f, 0x27, 0x94, 0xfe, 0x11,
	0xfc, 0x4b, 0xbb, 0x05, 0xd4, 0x22, 0x2f, 0x00, 0x76, 0x98, 0xf7, 0x66, 0x86, 0xd1, 0x02, 0x3d,
	0xf0, 0xa5, 0x33, 0xc3, 0x4b, 0x5f, 0xd9, 0x05, 0x58, 0x05, 0xb4, 0xfe, 0x06, 0xe7, 0x60, 0x4e,
	0x7d, 0xbf, 0xbc, 0xc1, 0x9f, 0xed, 0x4d, 0xa0, 0x97, 0xb7, 0xae, 0xbf, 0x5a, 0xff, 0x9a, 0xde,
	0x28, 0xa1, 0x55, 0xb7, 0xdb, 0x05, 0x2c, 0x43, 0x72, 0xc6, 0xf8, 0x5b, 0x87, 0xae, 0x58, 0x77,
	0xfa, 0x74, 0x87, 0x8f, 0x60, 0xe4, 0x01, 0xe2, 0xb0, 0x36, 0xaf, 0x1e, 0xb7, 0xed, 0xb4, 0x03,
	0x6a, 0x81, 0x5b, 0xd8, 0x54, 0xd7, 0xc4, 0xbd, 0x1a, 0x5c, 0x0d, 0xda, 0xde, 0x6f, 0x2b, 0x2b,
	0xa5, 0x7b, 0xe8, 0x66, 0xb9, 0x60, 0x0b, 0x9b, 0xfb, 0x1a, 0xb6, 0xd6, 0x0b, 0x5b, 0xea, 0xd4,
	0x0d, 0x5b, 0xd5, 0xf4, 0x1a, 0xb6, 0xea, 0x09, 0x3d, 0x03, 0x14, 0x37, 0xc5, 0xfa, 0x41, 0x1a,
	0x99, 0xd8, 0x07, 0x2b, 0x08, 0x29, 0x39, 0xdb, 0x10, 0xff, 0xd7, 0xd9, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xdd, 0xee, 0x0a, 0x19, 0xb4, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BooksAPIClient is the client API for BooksAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BooksAPIClient interface {
	// Get all books
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	// Get book by ID
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error)
	// Get book by IDs
	GetBooks(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (*GetBooksResponse, error)
	// Add new book
	AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*AddBookResponse, error)
	// Delete book
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error)
}

type booksAPIClient struct {
	cc *grpc.ClientConn
}

func NewBooksAPIClient(cc *grpc.ClientConn) BooksAPIClient {
	return &booksAPIClient{cc}
}

func (c *booksAPIClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, "/grpc.books.v1.BooksAPI/ListBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksAPIClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error) {
	out := new(GetBookResponse)
	err := c.cc.Invoke(ctx, "/grpc.books.v1.BooksAPI/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksAPIClient) GetBooks(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (*GetBooksResponse, error) {
	out := new(GetBooksResponse)
	err := c.cc.Invoke(ctx, "/grpc.books.v1.BooksAPI/GetBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksAPIClient) AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*AddBookResponse, error) {
	out := new(AddBookResponse)
	err := c.cc.Invoke(ctx, "/grpc.books.v1.BooksAPI/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksAPIClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error) {
	out := new(DeleteBookResponse)
	err := c.cc.Invoke(ctx, "/grpc.books.v1.BooksAPI/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BooksAPIServer is the server API for BooksAPI service.
type BooksAPIServer interface {
	// Get all books
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	// Get book by ID
	GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error)
	// Get book by IDs
	GetBooks(context.Context, *GetBooksRequest) (*GetBooksResponse, error)
	// Add new book
	AddBook(context.Context, *AddBookRequest) (*AddBookResponse, error)
	// Delete book
	DeleteBook(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error)
}

// UnimplementedBooksAPIServer can be embedded to have forward compatible implementations.
type UnimplementedBooksAPIServer struct {
}

func (*UnimplementedBooksAPIServer) ListBooks(ctx context.Context, req *ListBooksRequest) (*ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (*UnimplementedBooksAPIServer) GetBook(ctx context.Context, req *GetBookRequest) (*GetBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (*UnimplementedBooksAPIServer) GetBooks(ctx context.Context, req *GetBooksRequest) (*GetBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (*UnimplementedBooksAPIServer) AddBook(ctx context.Context, req *AddBookRequest) (*AddBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (*UnimplementedBooksAPIServer) DeleteBook(ctx context.Context, req *DeleteBookRequest) (*DeleteBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}

func RegisterBooksAPIServer(s *grpc.Server, srv BooksAPIServer) {
	s.RegisterService(&_BooksAPI_serviceDesc, srv)
}

func _BooksAPI_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksAPIServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.books.v1.BooksAPI/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksAPIServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksAPI_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksAPIServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.books.v1.BooksAPI/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksAPIServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksAPI_GetBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksAPIServer).GetBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.books.v1.BooksAPI/GetBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksAPIServer).GetBooks(ctx, req.(*GetBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksAPI_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksAPIServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.books.v1.BooksAPI/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksAPIServer).AddBook(ctx, req.(*AddBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksAPI_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksAPIServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.books.v1.BooksAPI/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksAPIServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BooksAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.books.v1.BooksAPI",
	HandlerType: (*BooksAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBooks",
			Handler:    _BooksAPI_ListBooks_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BooksAPI_GetBook_Handler,
		},
		{
			MethodName: "GetBooks",
			Handler:    _BooksAPI_GetBooks_Handler,
		},
		{
			MethodName: "AddBook",
			Handler:    _BooksAPI_AddBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BooksAPI_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/books/books.services.proto",
}
