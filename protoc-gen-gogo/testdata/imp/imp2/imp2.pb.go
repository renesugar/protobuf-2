// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imp/imp2/imp2.proto

package imp2 // import "github.com/gogo/protobuf/protoc-gen-gogo/testdata/imp/imp2"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PubliclyImportedEnum int32

const (
	PubliclyImportedEnum_GLASSES PubliclyImportedEnum = 1
	PubliclyImportedEnum_HAIR    PubliclyImportedEnum = 2
)

var PubliclyImportedEnum_name = map[int32]string{
	1: "GLASSES",
	2: "HAIR",
}
var PubliclyImportedEnum_value = map[string]int32{
	"GLASSES": 1,
	"HAIR":    2,
}

func (x PubliclyImportedEnum) Enum() *PubliclyImportedEnum {
	p := new(PubliclyImportedEnum)
	*p = x
	return p
}
func (x PubliclyImportedEnum) String() string {
	return proto.EnumName(PubliclyImportedEnum_name, int32(x))
}
func (x *PubliclyImportedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PubliclyImportedEnum_value, data, "PubliclyImportedEnum")
	if err != nil {
		return err
	}
	*x = PubliclyImportedEnum(value)
	return nil
}
func (PubliclyImportedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_imp2_c0fd690e537b2345, []int{0}
}

type PubliclyImportedMessage struct {
	Field *int64 `protobuf:"varint,1,opt,name=field" json:"field,omitempty"`
	// Field using a type in the same Go package, but a different source file.
	Field2               *EnumPB_Type `protobuf:"varint,2,opt,name=field2,enum=imp.EnumPB_Type" json:"field2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PubliclyImportedMessage) Reset()         { *m = PubliclyImportedMessage{} }
func (m *PubliclyImportedMessage) String() string { return proto.CompactTextString(m) }
func (*PubliclyImportedMessage) ProtoMessage()    {}
func (*PubliclyImportedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_imp2_c0fd690e537b2345, []int{0}
}
func (m *PubliclyImportedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubliclyImportedMessage.Unmarshal(m, b)
}
func (m *PubliclyImportedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubliclyImportedMessage.Marshal(b, m, deterministic)
}
func (dst *PubliclyImportedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubliclyImportedMessage.Merge(dst, src)
}
func (m *PubliclyImportedMessage) XXX_Size() int {
	return xxx_messageInfo_PubliclyImportedMessage.Size(m)
}
func (m *PubliclyImportedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PubliclyImportedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PubliclyImportedMessage proto.InternalMessageInfo

func (m *PubliclyImportedMessage) GetField() int64 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

func (m *PubliclyImportedMessage) GetField2() EnumPB_Type {
	if m != nil && m.Field2 != nil {
		return *m.Field2
	}
	return EnumPB_TYPE_1
}

type PubliclyImportedMessage2 struct {
	Field                *PubliclyImportedMessage `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PubliclyImportedMessage2) Reset()         { *m = PubliclyImportedMessage2{} }
func (m *PubliclyImportedMessage2) String() string { return proto.CompactTextString(m) }
func (*PubliclyImportedMessage2) ProtoMessage()    {}
func (*PubliclyImportedMessage2) Descriptor() ([]byte, []int) {
	return fileDescriptor_imp2_c0fd690e537b2345, []int{1}
}
func (m *PubliclyImportedMessage2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubliclyImportedMessage2.Unmarshal(m, b)
}
func (m *PubliclyImportedMessage2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubliclyImportedMessage2.Marshal(b, m, deterministic)
}
func (dst *PubliclyImportedMessage2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubliclyImportedMessage2.Merge(dst, src)
}
func (m *PubliclyImportedMessage2) XXX_Size() int {
	return xxx_messageInfo_PubliclyImportedMessage2.Size(m)
}
func (m *PubliclyImportedMessage2) XXX_DiscardUnknown() {
	xxx_messageInfo_PubliclyImportedMessage2.DiscardUnknown(m)
}

var xxx_messageInfo_PubliclyImportedMessage2 proto.InternalMessageInfo

func (m *PubliclyImportedMessage2) GetField() *PubliclyImportedMessage {
	if m != nil {
		return m.Field
	}
	return nil
}

func init() {
	proto.RegisterType((*PubliclyImportedMessage)(nil), "imp.PubliclyImportedMessage")
	proto.RegisterType((*PubliclyImportedMessage2)(nil), "imp.PubliclyImportedMessage2")
	proto.RegisterEnum("imp.PubliclyImportedEnum", PubliclyImportedEnum_name, PubliclyImportedEnum_value)
}

func init() { proto.RegisterFile("imp/imp2/imp2.proto", fileDescriptor_imp2_c0fd690e537b2345) }

var fileDescriptor_imp2_c0fd690e537b2345 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xcc, 0x2d, 0xd0,
	0xcf, 0xcc, 0x2d, 0x30, 0x02, 0x13, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xcc, 0x99, 0xb9,
	0x05, 0x52, 0x12, 0x28, 0x32, 0xba, 0xa9, 0x79, 0xa5, 0xb9, 0x10, 0x69, 0xa5, 0x48, 0x2e, 0xf1,
	0x80, 0xd2, 0xa4, 0x9c, 0xcc, 0xe4, 0x9c, 0x4a, 0xcf, 0xdc, 0x82, 0xfc, 0xa2, 0x92, 0xd4, 0x14,
	0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x21, 0x11, 0x2e, 0xd6, 0xb4, 0xcc, 0xd4, 0x9c, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x08, 0x47, 0x48, 0x83, 0x8b, 0x0d, 0xcc, 0x30, 0x92,
	0x60, 0x52, 0x60, 0xd4, 0xe0, 0x33, 0x12, 0xd0, 0xcb, 0xcc, 0x2d, 0xd0, 0x73, 0xcd, 0x2b, 0xcd,
	0x0d, 0x70, 0xd2, 0x0b, 0xa9, 0x2c, 0x48, 0x0d, 0x82, 0xca, 0x2b, 0xf9, 0x71, 0x49, 0xe0, 0x30,
	0xda, 0x48, 0xc8, 0x08, 0xd9, 0x6c, 0x6e, 0x23, 0x19, 0xb0, 0x21, 0x38, 0x54, 0x43, 0x6d, 0xd6,
	0xd2, 0xe5, 0x12, 0x41, 0x57, 0x01, 0xb2, 0x56, 0x88, 0x9b, 0x8b, 0xdd, 0xdd, 0xc7, 0x31, 0x38,
	0xd8, 0x35, 0x58, 0x80, 0x51, 0x88, 0x83, 0x8b, 0xc5, 0xc3, 0xd1, 0x33, 0x48, 0x80, 0xc9, 0xc9,
	0x26, 0xca, 0x2a, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x3d, 0x3f,
	0x3d, 0x5f, 0x1f, 0xec, 0xe5, 0xa4, 0xd2, 0x34, 0x08, 0x23, 0x59, 0x37, 0x3d, 0x35, 0x4f, 0x17,
	0x2c, 0x51, 0x92, 0x5a, 0x5c, 0x92, 0x92, 0x58, 0x92, 0xa8, 0x0f, 0x0b, 0x27, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa8, 0x49, 0xb3, 0x82, 0x4c, 0x01, 0x00, 0x00,
}
