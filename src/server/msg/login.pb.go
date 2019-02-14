// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

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
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}

// 登录
type LoginRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Uid                  string       `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Err                  *Error       `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *LoginResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *LoginResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

// 注册
type RegisteRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisteRequest) Reset()         { *m = RegisteRequest{} }
func (m *RegisteRequest) String() string { return proto.CompactTextString(m) }
func (*RegisteRequest) ProtoMessage()    {}
func (*RegisteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{2}
}

func (m *RegisteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisteRequest.Unmarshal(m, b)
}
func (m *RegisteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisteRequest.Marshal(b, m, deterministic)
}
func (m *RegisteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteRequest.Merge(m, src)
}
func (m *RegisteRequest) XXX_Size() int {
	return xxx_messageInfo_RegisteRequest.Size(m)
}
func (m *RegisteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteRequest proto.InternalMessageInfo

func (m *RegisteRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *RegisteRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisteResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Uid                  string       `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Err                  *Error       `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RegisteResponse) Reset()         { *m = RegisteResponse{} }
func (m *RegisteResponse) String() string { return proto.CompactTextString(m) }
func (*RegisteResponse) ProtoMessage()    {}
func (*RegisteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{3}
}

func (m *RegisteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisteResponse.Unmarshal(m, b)
}
func (m *RegisteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisteResponse.Marshal(b, m, deterministic)
}
func (m *RegisteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteResponse.Merge(m, src)
}
func (m *RegisteResponse) XXX_Size() int {
	return xxx_messageInfo_RegisteResponse.Size(m)
}
func (m *RegisteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteResponse proto.InternalMessageInfo

func (m *RegisteResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *RegisteResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RegisteResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterEnum("msg.ResponseCode", ResponseCode_name, ResponseCode_value)
	proto.RegisterType((*LoginRequest)(nil), "msg.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "msg.LoginResponse")
	proto.RegisterType((*RegisteRequest)(nil), "msg.RegisteRequest")
	proto.RegisterType((*RegisteResponse)(nil), "msg.RegisteResponse")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x8d, 0x29, 0xb6, 0x9d, 0xd4, 0x1a, 0xe7, 0x14, 0x8a, 0x87, 0x12, 0x28, 0x14, 0x0f,
	0x39, 0xd4, 0x27, 0x90, 0xd8, 0x82, 0xd0, 0xd3, 0x06, 0x1f, 0xa0, 0x26, 0xe3, 0x1a, 0xb0, 0x99,
	0x38, 0xb3, 0xc1, 0xd7, 0x97, 0xac, 0xa6, 0x78, 0x97, 0xde, 0xe6, 0x9f, 0x8f, 0xf9, 0xf8, 0x61,
	0x20, 0xfa, 0x60, 0x5b, 0x37, 0x59, 0x2b, 0xec, 0x18, 0xc3, 0xa3, 0xda, 0xc5, 0x94, 0x44, 0x7e,
	0x72, 0xfa, 0x04, 0xb3, 0x7d, 0x8f, 0x0d, 0x7d, 0x76, 0xa4, 0x0e, 0x13, 0x18, 0x1f, 0xca, 0x92,
	0xbb, 0xc6, 0x25, 0xc1, 0x32, 0x58, 0x4f, 0xcd, 0x10, 0x71, 0x01, 0x93, 0xf6, 0xa0, 0xfa, 0xc5,
	0x52, 0x25, 0x97, 0x1e, 0x9d, 0x72, 0xfa, 0x06, 0xd7, 0xbf, 0x16, 0x6d, 0xb9, 0x51, 0xc2, 0x15,
	0x8c, 0x4a, 0xae, 0xc8, 0x3b, 0xe6, 0x9b, 0xdb, 0xec, 0xa8, 0x36, 0x1b, 0x60, 0xce, 0x15, 0x19,
	0x8f, 0x31, 0x86, 0xb0, 0xab, 0x07, 0x5d, 0x3f, 0xe2, 0x1d, 0x84, 0x24, 0x92, 0x84, 0xcb, 0x60,
	0x1d, 0x6d, 0xc0, 0xdf, 0x6d, 0x45, 0x58, 0x4c, 0xbf, 0x4e, 0x77, 0x30, 0x37, 0x64, 0x6b, 0x75,
	0xf4, 0xbf, 0xbe, 0xef, 0x70, 0x73, 0xf2, 0x9c, 0xb5, 0xf1, 0xfd, 0x0a, 0x66, 0x7f, 0x2d, 0x38,
	0x81, 0xd1, 0xee, 0xf1, 0x79, 0x1f, 0x5f, 0x60, 0x04, 0xe3, 0xe2, 0x25, 0xcf, 0xb7, 0x45, 0x11,
	0x07, 0xaf, 0x57, 0xfe, 0x1b, 0x0f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x67, 0xc0, 0xbf, 0x94,
	0xac, 0x01, 0x00, 0x00,
}
