// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item.proto

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

type ItemRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemRequest) Reset()         { *m = ItemRequest{} }
func (m *ItemRequest) String() string { return proto.CompactTextString(m) }
func (*ItemRequest) ProtoMessage()    {}
func (*ItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{0}
}

func (m *ItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemRequest.Unmarshal(m, b)
}
func (m *ItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemRequest.Marshal(b, m, deterministic)
}
func (m *ItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemRequest.Merge(m, src)
}
func (m *ItemRequest) XXX_Size() int {
	return xxx_messageInfo_ItemRequest.Size(m)
}
func (m *ItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ItemRequest proto.InternalMessageInfo

type ItemResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Items                []*Item      `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ItemResponse) Reset()         { *m = ItemResponse{} }
func (m *ItemResponse) String() string { return proto.CompactTextString(m) }
func (*ItemResponse) ProtoMessage()    {}
func (*ItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{1}
}

func (m *ItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemResponse.Unmarshal(m, b)
}
func (m *ItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemResponse.Marshal(b, m, deterministic)
}
func (m *ItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemResponse.Merge(m, src)
}
func (m *ItemResponse) XXX_Size() int {
	return xxx_messageInfo_ItemResponse.Size(m)
}
func (m *ItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ItemResponse proto.InternalMessageInfo

func (m *ItemResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *ItemResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *ItemResponse) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type EquipRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EquipRequest) Reset()         { *m = EquipRequest{} }
func (m *EquipRequest) String() string { return proto.CompactTextString(m) }
func (*EquipRequest) ProtoMessage()    {}
func (*EquipRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{2}
}

func (m *EquipRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EquipRequest.Unmarshal(m, b)
}
func (m *EquipRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EquipRequest.Marshal(b, m, deterministic)
}
func (m *EquipRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquipRequest.Merge(m, src)
}
func (m *EquipRequest) XXX_Size() int {
	return xxx_messageInfo_EquipRequest.Size(m)
}
func (m *EquipRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EquipRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EquipRequest proto.InternalMessageInfo

type EquipResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Equips               []*Equip     `protobuf:"bytes,3,rep,name=equips,proto3" json:"equips,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EquipResponse) Reset()         { *m = EquipResponse{} }
func (m *EquipResponse) String() string { return proto.CompactTextString(m) }
func (*EquipResponse) ProtoMessage()    {}
func (*EquipResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{3}
}

func (m *EquipResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EquipResponse.Unmarshal(m, b)
}
func (m *EquipResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EquipResponse.Marshal(b, m, deterministic)
}
func (m *EquipResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquipResponse.Merge(m, src)
}
func (m *EquipResponse) XXX_Size() int {
	return xxx_messageInfo_EquipResponse.Size(m)
}
func (m *EquipResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EquipResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EquipResponse proto.InternalMessageInfo

func (m *EquipResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *EquipResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *EquipResponse) GetEquips() []*Equip {
	if m != nil {
		return m.Equips
	}
	return nil
}

type ConsumeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeRequest) Reset()         { *m = ConsumeRequest{} }
func (m *ConsumeRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumeRequest) ProtoMessage()    {}
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{4}
}

func (m *ConsumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeRequest.Unmarshal(m, b)
}
func (m *ConsumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeRequest.Marshal(b, m, deterministic)
}
func (m *ConsumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeRequest.Merge(m, src)
}
func (m *ConsumeRequest) XXX_Size() int {
	return xxx_messageInfo_ConsumeRequest.Size(m)
}
func (m *ConsumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeRequest proto.InternalMessageInfo

type ConsumeResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Consumes             []*Consume   `protobuf:"bytes,3,rep,name=consumes,proto3" json:"consumes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConsumeResponse) Reset()         { *m = ConsumeResponse{} }
func (m *ConsumeResponse) String() string { return proto.CompactTextString(m) }
func (*ConsumeResponse) ProtoMessage()    {}
func (*ConsumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{5}
}

func (m *ConsumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeResponse.Unmarshal(m, b)
}
func (m *ConsumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeResponse.Marshal(b, m, deterministic)
}
func (m *ConsumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeResponse.Merge(m, src)
}
func (m *ConsumeResponse) XXX_Size() int {
	return xxx_messageInfo_ConsumeResponse.Size(m)
}
func (m *ConsumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeResponse proto.InternalMessageInfo

func (m *ConsumeResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *ConsumeResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *ConsumeResponse) GetConsumes() []*Consume {
	if m != nil {
		return m.Consumes
	}
	return nil
}

type HeroChipRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeroChipRequest) Reset()         { *m = HeroChipRequest{} }
func (m *HeroChipRequest) String() string { return proto.CompactTextString(m) }
func (*HeroChipRequest) ProtoMessage()    {}
func (*HeroChipRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{6}
}

func (m *HeroChipRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroChipRequest.Unmarshal(m, b)
}
func (m *HeroChipRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroChipRequest.Marshal(b, m, deterministic)
}
func (m *HeroChipRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroChipRequest.Merge(m, src)
}
func (m *HeroChipRequest) XXX_Size() int {
	return xxx_messageInfo_HeroChipRequest.Size(m)
}
func (m *HeroChipRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroChipRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeroChipRequest proto.InternalMessageInfo

type HeroChipResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	HeroChips            []*HeroChip  `protobuf:"bytes,3,rep,name=heroChips,proto3" json:"heroChips,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HeroChipResponse) Reset()         { *m = HeroChipResponse{} }
func (m *HeroChipResponse) String() string { return proto.CompactTextString(m) }
func (*HeroChipResponse) ProtoMessage()    {}
func (*HeroChipResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{7}
}

func (m *HeroChipResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroChipResponse.Unmarshal(m, b)
}
func (m *HeroChipResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroChipResponse.Marshal(b, m, deterministic)
}
func (m *HeroChipResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroChipResponse.Merge(m, src)
}
func (m *HeroChipResponse) XXX_Size() int {
	return xxx_messageInfo_HeroChipResponse.Size(m)
}
func (m *HeroChipResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroChipResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeroChipResponse proto.InternalMessageInfo

func (m *HeroChipResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *HeroChipResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *HeroChipResponse) GetHeroChips() []*HeroChip {
	if m != nil {
		return m.HeroChips
	}
	return nil
}

type Item struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Price                int32    `protobuf:"varint,3,opt,name=Price,proto3" json:"Price,omitempty"`
	Desc                 string   `protobuf:"bytes,4,opt,name=Desc,proto3" json:"Desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{8}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type Equip struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Price                int32    `protobuf:"varint,3,opt,name=Price,proto3" json:"Price,omitempty"`
	Effect               string   `protobuf:"bytes,4,opt,name=Effect,proto3" json:"Effect,omitempty"`
	Desc                 string   `protobuf:"bytes,5,opt,name=Desc,proto3" json:"Desc,omitempty"`
	Mixs                 []*Mix   `protobuf:"bytes,6,rep,name=Mixs,proto3" json:"Mixs,omitempty"`
	ItemId               string   `protobuf:"bytes,7,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	HeroId               string   `protobuf:"bytes,8,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Equip) Reset()         { *m = Equip{} }
func (m *Equip) String() string { return proto.CompactTextString(m) }
func (*Equip) ProtoMessage()    {}
func (*Equip) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{9}
}

func (m *Equip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Equip.Unmarshal(m, b)
}
func (m *Equip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Equip.Marshal(b, m, deterministic)
}
func (m *Equip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Equip.Merge(m, src)
}
func (m *Equip) XXX_Size() int {
	return xxx_messageInfo_Equip.Size(m)
}
func (m *Equip) XXX_DiscardUnknown() {
	xxx_messageInfo_Equip.DiscardUnknown(m)
}

var xxx_messageInfo_Equip proto.InternalMessageInfo

func (m *Equip) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Equip) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Equip) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Equip) GetEffect() string {
	if m != nil {
		return m.Effect
	}
	return ""
}

func (m *Equip) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Equip) GetMixs() []*Mix {
	if m != nil {
		return m.Mixs
	}
	return nil
}

func (m *Equip) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *Equip) GetHeroId() string {
	if m != nil {
		return m.HeroId
	}
	return ""
}

type Consume struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Price                int32    `protobuf:"varint,3,opt,name=Price,proto3" json:"Price,omitempty"`
	Desc                 string   `protobuf:"bytes,4,opt,name=Desc,proto3" json:"Desc,omitempty"`
	ItemId               string   `protobuf:"bytes,5,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consume) Reset()         { *m = Consume{} }
func (m *Consume) String() string { return proto.CompactTextString(m) }
func (*Consume) ProtoMessage()    {}
func (*Consume) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{10}
}

func (m *Consume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consume.Unmarshal(m, b)
}
func (m *Consume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consume.Marshal(b, m, deterministic)
}
func (m *Consume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consume.Merge(m, src)
}
func (m *Consume) XXX_Size() int {
	return xxx_messageInfo_Consume.Size(m)
}
func (m *Consume) XXX_DiscardUnknown() {
	xxx_messageInfo_Consume.DiscardUnknown(m)
}

var xxx_messageInfo_Consume proto.InternalMessageInfo

func (m *Consume) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Consume) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Consume) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Consume) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Consume) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

type HeroChip struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Price                int32    `protobuf:"varint,3,opt,name=Price,proto3" json:"Price,omitempty"`
	Desc                 string   `protobuf:"bytes,4,opt,name=Desc,proto3" json:"Desc,omitempty"`
	ItemId               string   `protobuf:"bytes,5,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	Cnt                  int32    `protobuf:"varint,6,opt,name=Cnt,proto3" json:"Cnt,omitempty"`
	Compose              int32    `protobuf:"varint,7,opt,name=Compose,proto3" json:"Compose,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeroChip) Reset()         { *m = HeroChip{} }
func (m *HeroChip) String() string { return proto.CompactTextString(m) }
func (*HeroChip) ProtoMessage()    {}
func (*HeroChip) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{11}
}

func (m *HeroChip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroChip.Unmarshal(m, b)
}
func (m *HeroChip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroChip.Marshal(b, m, deterministic)
}
func (m *HeroChip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroChip.Merge(m, src)
}
func (m *HeroChip) XXX_Size() int {
	return xxx_messageInfo_HeroChip.Size(m)
}
func (m *HeroChip) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroChip.DiscardUnknown(m)
}

var xxx_messageInfo_HeroChip proto.InternalMessageInfo

func (m *HeroChip) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HeroChip) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HeroChip) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *HeroChip) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *HeroChip) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *HeroChip) GetCnt() int32 {
	if m != nil {
		return m.Cnt
	}
	return 0
}

func (m *HeroChip) GetCompose() int32 {
	if m != nil {
		return m.Compose
	}
	return 0
}

type Mix struct {
	ItemId               int32    `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	Num                  int32    `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mix) Reset()         { *m = Mix{} }
func (m *Mix) String() string { return proto.CompactTextString(m) }
func (*Mix) ProtoMessage()    {}
func (*Mix) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{12}
}

func (m *Mix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mix.Unmarshal(m, b)
}
func (m *Mix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mix.Marshal(b, m, deterministic)
}
func (m *Mix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mix.Merge(m, src)
}
func (m *Mix) XXX_Size() int {
	return xxx_messageInfo_Mix.Size(m)
}
func (m *Mix) XXX_DiscardUnknown() {
	xxx_messageInfo_Mix.DiscardUnknown(m)
}

var xxx_messageInfo_Mix proto.InternalMessageInfo

func (m *Mix) GetItemId() int32 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *Mix) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*ItemRequest)(nil), "msg.ItemRequest")
	proto.RegisterType((*ItemResponse)(nil), "msg.ItemResponse")
	proto.RegisterType((*EquipRequest)(nil), "msg.EquipRequest")
	proto.RegisterType((*EquipResponse)(nil), "msg.EquipResponse")
	proto.RegisterType((*ConsumeRequest)(nil), "msg.ConsumeRequest")
	proto.RegisterType((*ConsumeResponse)(nil), "msg.ConsumeResponse")
	proto.RegisterType((*HeroChipRequest)(nil), "msg.HeroChipRequest")
	proto.RegisterType((*HeroChipResponse)(nil), "msg.HeroChipResponse")
	proto.RegisterType((*Item)(nil), "msg.Item")
	proto.RegisterType((*Equip)(nil), "msg.Equip")
	proto.RegisterType((*Consume)(nil), "msg.Consume")
	proto.RegisterType((*HeroChip)(nil), "msg.HeroChip")
	proto.RegisterType((*Mix)(nil), "msg.Mix")
}

func init() { proto.RegisterFile("item.proto", fileDescriptor_6007f868cf6553df) }

var fileDescriptor_6007f868cf6553df = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x8a, 0xdb, 0x30,
	0x10, 0xc6, 0xb1, 0xe5, 0x4d, 0x26, 0x3f, 0x9b, 0x15, 0xa5, 0x88, 0xb2, 0xd0, 0x20, 0x28, 0x04,
	0x0a, 0x29, 0xa4, 0x8f, 0x90, 0x06, 0xea, 0x43, 0x96, 0x22, 0xfa, 0x02, 0xad, 0x3d, 0xbb, 0xeb,
	0x83, 0xa2, 0xac, 0xe4, 0x40, 0x8e, 0x3d, 0xf4, 0x31, 0xfa, 0x26, 0x7d, 0xb9, 0xa2, 0x91, 0x15,
	0xfb, 0x5e, 0xb3, 0x37, 0x7d, 0x33, 0x9f, 0xbf, 0xef, 0x9b, 0x19, 0x30, 0x40, 0xdd, 0xa0, 0xde,
	0x9c, 0xac, 0x69, 0x0c, 0x4f, 0xb5, 0x7b, 0x7a, 0x37, 0x41, 0x6b, 0x03, 0x96, 0x73, 0x98, 0x16,
	0x0d, 0x6a, 0x85, 0x2f, 0x67, 0x74, 0x8d, 0x6c, 0x60, 0x16, 0xa0, 0x3b, 0x99, 0xa3, 0x43, 0xfe,
	0x01, 0xb2, 0xd2, 0x54, 0x28, 0x92, 0x55, 0xb2, 0x5e, 0x6c, 0xef, 0x36, 0xda, 0x3d, 0x6d, 0x62,
	0x73, 0x67, 0x2a, 0x54, 0xd4, 0xe6, 0xf7, 0x90, 0xa2, 0xb5, 0x62, 0xb4, 0x4a, 0xd6, 0xd3, 0x2d,
	0x10, 0x6b, 0x6f, 0xad, 0xb1, 0xca, 0x97, 0xf9, 0x7b, 0x60, 0x3e, 0x81, 0x13, 0xe9, 0x2a, 0x5d,
	0x4f, 0xb7, 0x13, 0xea, 0x93, 0x4d, 0xa8, 0xcb, 0x05, 0xcc, 0xf6, 0x2f, 0xe7, 0xfa, 0x14, 0x53,
	0x5c, 0x60, 0xde, 0xe2, 0x21, 0x63, 0x48, 0xc8, 0xd1, 0xab, 0xc6, 0x1c, 0x2d, 0x81, 0x8c, 0xda,
	0x8e, 0x5c, 0xc2, 0x62, 0x67, 0x8e, 0xee, 0xac, 0x31, 0x66, 0xf9, 0x95, 0xc0, 0xed, 0xb5, 0x34,
	0x64, 0x9c, 0x35, 0x8c, 0xcb, 0xa0, 0x1b, 0x03, 0xcd, 0x88, 0x12, 0xcd, 0xae, 0x5d, 0x79, 0x07,
	0xb7, 0x5f, 0xd1, 0x9a, 0xdd, 0x73, 0xb7, 0xa1, 0xdf, 0x09, 0x2c, 0xbb, 0xda, 0x90, 0xb1, 0x3e,
	0xc2, 0xe4, 0xb9, 0x15, 0x8e, 0xb9, 0xe6, 0xc4, 0xb9, 0xda, 0x75, 0x7d, 0xf9, 0x1d, 0x32, 0x7f,
	0x47, 0xbe, 0x80, 0x51, 0x51, 0x91, 0x2f, 0x53, 0xa3, 0xa2, 0xe2, 0x1c, 0xb2, 0x87, 0x1f, 0x1a,
	0xc9, 0x63, 0xa2, 0xe8, 0xcd, 0xdf, 0x00, 0xfb, 0x66, 0xeb, 0x12, 0x45, 0x4a, 0xb4, 0x00, 0x3c,
	0xf3, 0x0b, 0xba, 0x52, 0x64, 0x81, 0xe9, 0xdf, 0xf2, 0x6f, 0x02, 0x8c, 0xce, 0xf2, 0x1f, 0xba,
	0x6f, 0x21, 0xdf, 0x3f, 0x3e, 0x62, 0xd9, 0xb4, 0xca, 0x2d, 0xba, 0xfa, 0xb1, 0xce, 0x8f, 0xdf,
	0x43, 0x76, 0xa8, 0x2f, 0x4e, 0xe4, 0x34, 0xed, 0x98, 0xa6, 0x3d, 0xd4, 0x17, 0x45, 0x55, 0xaf,
	0xe4, 0x67, 0x2c, 0x2a, 0x71, 0x13, 0x94, 0x02, 0xf2, 0x75, 0xbf, 0x92, 0xa2, 0x12, 0xe3, 0x50,
	0x0f, 0x48, 0x1a, 0xb8, 0x69, 0x4f, 0x38, 0xec, 0x5a, 0x7a, 0x41, 0x58, 0x3f, 0x88, 0xfc, 0x93,
	0xc0, 0x38, 0x1e, 0xe7, 0x75, 0x2c, 0xf9, 0x12, 0xd2, 0xdd, 0xb1, 0x11, 0x39, 0x7d, 0xef, 0x9f,
	0x5c, 0xf8, 0xa9, 0xf5, 0xc9, 0x38, 0xa4, 0x35, 0x31, 0x15, 0xa1, 0xfc, 0x04, 0xe9, 0xa1, 0xbe,
	0xf4, 0xa4, 0x42, 0xb8, 0x9e, 0xd4, 0xc3, 0x59, 0x53, 0x3e, 0xa6, 0xfc, 0xf3, 0x67, 0x4e, 0x7f,
	0xa6, 0xcf, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x45, 0x07, 0xa1, 0xb7, 0x04, 0x00, 0x00,
}
