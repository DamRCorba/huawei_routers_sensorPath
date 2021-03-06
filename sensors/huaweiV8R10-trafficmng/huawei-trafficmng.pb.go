// Code generated by protoc-gen-go. DO NOT EDIT.
// source: huawei-trafficmng.proto

package huaweiV8R10_trafficmng

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

type Trafficmng struct {
	TmSlotSFUs           *Trafficmng_TmSlotSFUs `protobuf:"bytes,4,opt,name=tmSlotSFUs,proto3" json:"tmSlotSFUs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Trafficmng) Reset()         { *m = Trafficmng{} }
func (m *Trafficmng) String() string { return proto.CompactTextString(m) }
func (*Trafficmng) ProtoMessage()    {}
func (*Trafficmng) Descriptor() ([]byte, []int) {
	return fileDescriptor_97c1b2fe3d27ed83, []int{0}
}

func (m *Trafficmng) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trafficmng.Unmarshal(m, b)
}
func (m *Trafficmng) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trafficmng.Marshal(b, m, deterministic)
}
func (m *Trafficmng) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trafficmng.Merge(m, src)
}
func (m *Trafficmng) XXX_Size() int {
	return xxx_messageInfo_Trafficmng.Size(m)
}
func (m *Trafficmng) XXX_DiscardUnknown() {
	xxx_messageInfo_Trafficmng.DiscardUnknown(m)
}

var xxx_messageInfo_Trafficmng proto.InternalMessageInfo

func (m *Trafficmng) GetTmSlotSFUs() *Trafficmng_TmSlotSFUs {
	if m != nil {
		return m.TmSlotSFUs
	}
	return nil
}

type Trafficmng_TmSlotSFUs struct {
	TmSlotSFU            []*Trafficmng_TmSlotSFUs_TmSlotSFU `protobuf:"bytes,1,rep,name=tmSlotSFU,proto3" json:"tmSlotSFU,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *Trafficmng_TmSlotSFUs) Reset()         { *m = Trafficmng_TmSlotSFUs{} }
func (m *Trafficmng_TmSlotSFUs) String() string { return proto.CompactTextString(m) }
func (*Trafficmng_TmSlotSFUs) ProtoMessage()    {}
func (*Trafficmng_TmSlotSFUs) Descriptor() ([]byte, []int) {
	return fileDescriptor_97c1b2fe3d27ed83, []int{0, 0}
}

func (m *Trafficmng_TmSlotSFUs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trafficmng_TmSlotSFUs.Unmarshal(m, b)
}
func (m *Trafficmng_TmSlotSFUs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trafficmng_TmSlotSFUs.Marshal(b, m, deterministic)
}
func (m *Trafficmng_TmSlotSFUs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trafficmng_TmSlotSFUs.Merge(m, src)
}
func (m *Trafficmng_TmSlotSFUs) XXX_Size() int {
	return xxx_messageInfo_Trafficmng_TmSlotSFUs.Size(m)
}
func (m *Trafficmng_TmSlotSFUs) XXX_DiscardUnknown() {
	xxx_messageInfo_Trafficmng_TmSlotSFUs.DiscardUnknown(m)
}

var xxx_messageInfo_Trafficmng_TmSlotSFUs proto.InternalMessageInfo

func (m *Trafficmng_TmSlotSFUs) GetTmSlotSFU() []*Trafficmng_TmSlotSFUs_TmSlotSFU {
	if m != nil {
		return m.TmSlotSFU
	}
	return nil
}

type Trafficmng_TmSlotSFUs_TmSlotSFU struct {
	SfuStatisticss       *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss `protobuf:"bytes,1,opt,name=sfuStatisticss,proto3" json:"sfuStatisticss,omitempty"`
	SlotId               string                                          `protobuf:"bytes,2,opt,name=slotId,proto3" json:"slotId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU) Reset()         { *m = Trafficmng_TmSlotSFUs_TmSlotSFU{} }
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU) String() string { return proto.CompactTextString(m) }
func (*Trafficmng_TmSlotSFUs_TmSlotSFU) ProtoMessage()    {}
func (*Trafficmng_TmSlotSFUs_TmSlotSFU) Descriptor() ([]byte, []int) {
	return fileDescriptor_97c1b2fe3d27ed83, []int{0, 0, 0}
}

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU.Unmarshal(m, b)
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU.Marshal(b, m, deterministic)
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU.Merge(m, src)
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU) XXX_Size() int {
	return xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU.Size(m)
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU) XXX_DiscardUnknown() {
	xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU.DiscardUnknown(m)
}

var xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU proto.InternalMessageInfo

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU) GetSfuStatisticss() *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss {
	if m != nil {
		return m.SfuStatisticss
	}
	return nil
}

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU) GetSlotId() string {
	if m != nil {
		return m.SlotId
	}
	return ""
}

type Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss struct {
	SfuStatistics        []*Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics `protobuf:"bytes,1,rep,name=sfuStatistics,proto3" json:"sfuStatistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                        `json:"-"`
	XXX_unrecognized     []byte                                                          `json:"-"`
	XXX_sizecache        int32                                                           `json:"-"`
}

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss) Reset() {
	*m = Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss{}
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss) String() string {
	return proto.CompactTextString(m)
}
func (*Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss) ProtoMessage() {}
func (*Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss) Descriptor() ([]byte, []int) {
	return fileDescriptor_97c1b2fe3d27ed83, []int{0, 0, 0, 0}
}

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss.Unmarshal(m, b)
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss.Marshal(b, m, deterministic)
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss.Merge(m, src)
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss) XXX_Size() int {
	return xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss.Size(m)
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss) XXX_DiscardUnknown() {
	xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss.DiscardUnknown(m)
}

var xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss proto.InternalMessageInfo

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss) GetSfuStatistics() []*Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics {
	if m != nil {
		return m.SfuStatistics
	}
	return nil
}

type Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics struct {
	ChipId               uint32   `protobuf:"varint,1,opt,name=chipId,proto3" json:"chipId,omitempty"`
	DiscardCells         uint64   `protobuf:"varint,2,opt,name=discardCells,proto3" json:"discardCells,omitempty"`
	PassCells            uint64   `protobuf:"varint,3,opt,name=passCells,proto3" json:"passCells,omitempty"`
	SendCells            uint64   `protobuf:"varint,4,opt,name=sendCells,proto3" json:"sendCells,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) Reset() {
	*m = Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics{}
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) String() string {
	return proto.CompactTextString(m)
}
func (*Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) ProtoMessage() {}
func (*Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_97c1b2fe3d27ed83, []int{0, 0, 0, 0, 0}
}

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics.Unmarshal(m, b)
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics.Marshal(b, m, deterministic)
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics.Merge(m, src)
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) XXX_Size() int {
	return xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics.Size(m)
}
func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics proto.InternalMessageInfo

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) GetChipId() uint32 {
	if m != nil {
		return m.ChipId
	}
	return 0
}

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) GetDiscardCells() uint64 {
	if m != nil {
		return m.DiscardCells
	}
	return 0
}

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) GetPassCells() uint64 {
	if m != nil {
		return m.PassCells
	}
	return 0
}

func (m *Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics) GetSendCells() uint64 {
	if m != nil {
		return m.SendCells
	}
	return 0
}

func init() {
	proto.RegisterType((*Trafficmng)(nil), "huaweiV8R10_trafficmng.Trafficmng")
	proto.RegisterType((*Trafficmng_TmSlotSFUs)(nil), "huaweiV8R10_trafficmng.Trafficmng.TmSlotSFUs")
	proto.RegisterType((*Trafficmng_TmSlotSFUs_TmSlotSFU)(nil), "huaweiV8R10_trafficmng.Trafficmng.TmSlotSFUs.TmSlotSFU")
	proto.RegisterType((*Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss)(nil), "huaweiV8R10_trafficmng.Trafficmng.TmSlotSFUs.TmSlotSFU.SfuStatisticss")
	proto.RegisterType((*Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics)(nil), "huaweiV8R10_trafficmng.Trafficmng.TmSlotSFUs.TmSlotSFU.SfuStatisticss.SfuStatistics")
}

func init() { proto.RegisterFile("huawei-trafficmng.proto", fileDescriptor_97c1b2fe3d27ed83) }

var fileDescriptor_97c1b2fe3d27ed83 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x66, 0x9b, 0x50, 0xc8, 0xd4, 0x14, 0xdc, 0x83, 0x86, 0xe0, 0x21, 0xf4, 0x94, 0x8b, 0x39,
	0xc4, 0x27, 0x10, 0x41, 0xec, 0x49, 0xd9, 0xd4, 0xb3, 0xac, 0xf9, 0xb1, 0x0b, 0x69, 0x13, 0x32,
	0x1b, 0x7c, 0x00, 0x0f, 0x9e, 0x7d, 0x05, 0x9f, 0xcb, 0x87, 0x91, 0x6c, 0xc2, 0x6e, 0x57, 0x2f,
	0x4a, 0x6f, 0xf3, 0xcd, 0xf7, 0x33, 0x33, 0xcb, 0xc2, 0xf9, 0xb6, 0xe7, 0xaf, 0xa5, 0xb8, 0x94,
	0x1d, 0xaf, 0x2a, 0x91, 0xef, 0xf6, 0x2f, 0x49, 0xdb, 0x35, 0xb2, 0xa1, 0xa7, 0x23, 0xf1, 0x64,
	0x88, 0xd5, 0x97, 0x0b, 0xb0, 0xd1, 0x90, 0xde, 0x01, 0xc8, 0x5d, 0x56, 0x37, 0x32, 0xbb, 0x7d,
	0xc4, 0xc0, 0x8d, 0x48, 0xbc, 0x48, 0xe3, 0xe4, 0x97, 0x2d, 0xd9, 0x1c, 0x94, 0x5a, 0xcf, 0x0e,
	0xbc, 0xe1, 0xdb, 0x10, 0xac, 0x21, 0x7d, 0x00, 0x4f, 0x93, 0x01, 0x89, 0x9c, 0x78, 0x91, 0xa6,
	0x7f, 0xcd, 0x35, 0x25, 0x33, 0x21, 0xe1, 0xa7, 0x03, 0x9e, 0x26, 0xa8, 0x80, 0x25, 0x56, 0x7d,
	0x26, 0xb9, 0x14, 0x28, 0x45, 0x8e, 0x18, 0x10, 0xb5, 0xfc, 0xf5, 0xff, 0x87, 0x24, 0x99, 0x15,
	0xc4, 0x7e, 0x04, 0xd3, 0x33, 0x98, 0x63, 0xdd, 0xc8, 0x75, 0x11, 0xcc, 0x22, 0x12, 0x7b, 0x6c,
	0x42, 0xe1, 0xc7, 0x0c, 0x96, 0xb6, 0x95, 0xf6, 0xe0, 0x5b, 0xe6, 0xe9, 0xf2, 0xfb, 0xa3, 0x97,
	0xb2, 0x21, 0xb3, 0xa7, 0x84, 0xef, 0x04, 0x7c, 0x4b, 0x30, 0xec, 0x9c, 0x6f, 0x45, 0xbb, 0x2e,
	0xd4, 0xb3, 0xf8, 0x6c, 0x42, 0x74, 0x05, 0x27, 0x85, 0xc0, 0x9c, 0x77, 0xc5, 0x4d, 0x59, 0xd7,
	0xa8, 0x2e, 0x72, 0x99, 0xd5, 0xa3, 0x17, 0xe0, 0xb5, 0x1c, 0x71, 0x14, 0x38, 0x4a, 0x60, 0x1a,
	0x03, 0x8b, 0xe5, 0x7e, 0xb2, 0xbb, 0x23, 0xab, 0x1b, 0xcf, 0x73, 0xf5, 0xf1, 0xae, 0xbe, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x95, 0x61, 0x53, 0x2e, 0x93, 0x02, 0x00, 0x00,
}
