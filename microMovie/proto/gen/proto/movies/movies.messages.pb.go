// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/movies/movies.messages.proto

package grpc_movies

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//  Movie definition
type Movie struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Diractor             string   `protobuf:"bytes,2,opt,name=diractor,proto3" json:"diractor,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Genres               string   `protobuf:"bytes,4,opt,name=genres,proto3" json:"genres,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Movie) Reset()         { *m = Movie{} }
func (m *Movie) String() string { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()    {}
func (*Movie) Descriptor() ([]byte, []int) {
	return fileDescriptor_31e3310ebfd18983, []int{0}
}

func (m *Movie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Movie.Unmarshal(m, b)
}
func (m *Movie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Movie.Marshal(b, m, deterministic)
}
func (m *Movie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Movie.Merge(m, src)
}
func (m *Movie) XXX_Size() int {
	return xxx_messageInfo_Movie.Size(m)
}
func (m *Movie) XXX_DiscardUnknown() {
	xxx_messageInfo_Movie.DiscardUnknown(m)
}

var xxx_messageInfo_Movie proto.InternalMessageInfo

func (m *Movie) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Movie) GetDiractor() string {
	if m != nil {
		return m.Diractor
	}
	return ""
}

func (m *Movie) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Movie) GetGenres() string {
	if m != nil {
		return m.Genres
	}
	return ""
}

func init() {
	proto.RegisterType((*Movie)(nil), "grpc.movies.Movie")
}

func init() { proto.RegisterFile("proto/movies/movies.messages.proto", fileDescriptor_31e3310ebfd18983) }

var fileDescriptor_31e3310ebfd18983 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xcd, 0x2f, 0xcb, 0x4c, 0x2d, 0x86, 0x52, 0x7a, 0xb9, 0xa9, 0xc5, 0xc5, 0x89,
	0xe9, 0xa9, 0xc5, 0x7a, 0x60, 0x49, 0x21, 0xee, 0xf4, 0xa2, 0x82, 0x64, 0x3d, 0x88, 0x9c, 0x52,
	0x22, 0x17, 0xab, 0x2f, 0x88, 0x25, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0xc4, 0x94, 0x99, 0x22, 0x24, 0xc5, 0xc5, 0x91, 0x92, 0x59, 0x94, 0x98, 0x5c, 0x92,
	0x5f, 0x24, 0xc1, 0x04, 0x16, 0x85, 0xf3, 0x85, 0x44, 0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52,
	0x25, 0x98, 0xc1, 0x12, 0x10, 0x8e, 0x90, 0x18, 0x17, 0x5b, 0x7a, 0x6a, 0x5e, 0x51, 0x6a, 0xb1,
	0x04, 0x0b, 0x58, 0x18, 0xca, 0x4b, 0x62, 0x03, 0x5b, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x61, 0xdb, 0x83, 0x84, 0x9c, 0x00, 0x00, 0x00,
}