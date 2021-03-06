// Code generated by protoc-gen-go. DO NOT EDIT.
// source: err.proto

package msg

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

type ResponseCode int32

const (
	ResponseCode_FAIL    ResponseCode = 0
	ResponseCode_SUCCESS ResponseCode = 1
)

var ResponseCode_name = map[int32]string{
	0: "FAIL",
	1: "SUCCESS",
}

var ResponseCode_value = map[string]int32{
	"FAIL":    0,
	"SUCCESS": 1,
}

func (x ResponseCode) String() string {
	return proto.EnumName(ResponseCode_name, int32(x))
}

func (ResponseCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4a1db73bc95ee8c, []int{0}
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4a1db73bc95ee8c, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("msg.ResponseCode", ResponseCode_name, ResponseCode_value)
	proto.RegisterType((*Error)(nil), "msg.Error")
}

func init() { proto.RegisterFile("err.proto", fileDescriptor_b4a1db73bc95ee8c) }

var fileDescriptor_b4a1db73bc95ee8c = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2d, 0x2a, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0xd2, 0xe5, 0x62, 0x75, 0x2d,
	0x2a, 0xca, 0x2f, 0x12, 0x12, 0xe2, 0x62, 0x49, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x0d, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x40, 0x6a, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83,
	0x40, 0x4c, 0x2d, 0x55, 0x2e, 0x9e, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x67, 0x90,
	0x0a, 0x0e, 0x2e, 0x16, 0x37, 0x47, 0x4f, 0x1f, 0x01, 0x06, 0x21, 0x6e, 0x2e, 0xf6, 0xe0, 0x50,
	0x67, 0x67, 0xd7, 0xe0, 0x60, 0x01, 0xc6, 0x24, 0x36, 0xb0, 0x0d, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2c, 0x88, 0xf9, 0xcf, 0x6e, 0x00, 0x00, 0x00,
}
