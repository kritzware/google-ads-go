// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/conversion_action_status.proto

package enums

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

// Possible statuses of a conversion action.
type ConversionActionStatusEnum_ConversionActionStatus int32

const (
	// Not specified.
	ConversionActionStatusEnum_UNSPECIFIED ConversionActionStatusEnum_ConversionActionStatus = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionActionStatusEnum_UNKNOWN ConversionActionStatusEnum_ConversionActionStatus = 1
	// Conversions will be recorded.
	ConversionActionStatusEnum_ENABLED ConversionActionStatusEnum_ConversionActionStatus = 2
	// Conversions will not be recorded.
	ConversionActionStatusEnum_REMOVED ConversionActionStatusEnum_ConversionActionStatus = 3
	// Conversions will not be recorded and the conversion action will not
	// appear in the UI.
	ConversionActionStatusEnum_HIDDEN ConversionActionStatusEnum_ConversionActionStatus = 4
)

var ConversionActionStatusEnum_ConversionActionStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
	4: "HIDDEN",
}

var ConversionActionStatusEnum_ConversionActionStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
	"HIDDEN":      4,
}

func (x ConversionActionStatusEnum_ConversionActionStatus) String() string {
	return proto.EnumName(ConversionActionStatusEnum_ConversionActionStatus_name, int32(x))
}

func (ConversionActionStatusEnum_ConversionActionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ff1eb7eb51c9f116, []int{0, 0}
}

// Container for enum describing possible statuses of a conversion action.
type ConversionActionStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionActionStatusEnum) Reset()         { *m = ConversionActionStatusEnum{} }
func (m *ConversionActionStatusEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionActionStatusEnum) ProtoMessage()    {}
func (*ConversionActionStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff1eb7eb51c9f116, []int{0}
}

func (m *ConversionActionStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionActionStatusEnum.Unmarshal(m, b)
}
func (m *ConversionActionStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionActionStatusEnum.Marshal(b, m, deterministic)
}
func (m *ConversionActionStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionActionStatusEnum.Merge(m, src)
}
func (m *ConversionActionStatusEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionActionStatusEnum.Size(m)
}
func (m *ConversionActionStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionActionStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionActionStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.ConversionActionStatusEnum_ConversionActionStatus", ConversionActionStatusEnum_ConversionActionStatus_name, ConversionActionStatusEnum_ConversionActionStatus_value)
	proto.RegisterType((*ConversionActionStatusEnum)(nil), "google.ads.googleads.v0.enums.ConversionActionStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/conversion_action_status.proto", fileDescriptor_ff1eb7eb51c9f116)
}

var fileDescriptor_ff1eb7eb51c9f116 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x2d, 0x35, 0x35, 0xd9, 0x1e, 0x24, 0x1c, 0x3c, 0x68, 0x7a, 0x68, 0x1f, 0x60, 0x21,
	0xf1, 0xb6, 0x7a, 0x59, 0xca, 0x5a, 0x89, 0xba, 0x25, 0x36, 0xc5, 0xc4, 0x90, 0x18, 0x04, 0xb2,
	0x69, 0x52, 0x76, 0x1b, 0x06, 0x38, 0xf8, 0x38, 0x1e, 0x7d, 0x14, 0x1f, 0xc5, 0x93, 0x8f, 0x60,
	0x76, 0xb1, 0x9c, 0xaa, 0x17, 0xf8, 0x67, 0xff, 0xf9, 0x26, 0x33, 0x3f, 0xba, 0x16, 0x4a, 0x89,
	0x6d, 0xe1, 0xa6, 0x39, 0xb8, 0x9d, 0xd4, 0xaa, 0xf5, 0xdc, 0x42, 0x36, 0x25, 0xb8, 0x99, 0x92,
	0x6d, 0x51, 0xc1, 0x46, 0xc9, 0x97, 0x34, 0xab, 0xf5, 0x0f, 0xea, 0xb4, 0x6e, 0x00, 0xef, 0x2a,
	0x55, 0x2b, 0x67, 0xd2, 0x21, 0x38, 0xcd, 0x01, 0xf7, 0x34, 0x6e, 0x3d, 0x6c, 0xe8, 0xd9, 0x1b,
	0x3a, 0x9f, 0xf7, 0x03, 0xa8, 0xe1, 0x57, 0x06, 0x67, 0xb2, 0x29, 0x67, 0x09, 0x3a, 0x3b, 0xec,
	0x3a, 0xa7, 0x68, 0xbc, 0xe6, 0xab, 0x88, 0xcd, 0xc3, 0x9b, 0x90, 0x05, 0xf6, 0x91, 0x33, 0x46,
	0x27, 0x6b, 0x7e, 0xc7, 0x97, 0x4f, 0xdc, 0x1e, 0xe8, 0x82, 0x71, 0xea, 0xdf, 0xb3, 0xc0, 0xb6,
	0x74, 0xf1, 0xc8, 0x1e, 0x96, 0x31, 0x0b, 0xec, 0xa1, 0x83, 0xd0, 0xe8, 0x36, 0x0c, 0x02, 0xc6,
	0xed, 0x63, 0xff, 0x7b, 0x80, 0xa6, 0x99, 0x2a, 0xf1, 0xbf, 0x1b, 0xfa, 0x17, 0x87, 0x37, 0x88,
	0xf4, 0x75, 0xd1, 0xe0, 0xd9, 0xff, 0xa5, 0x85, 0xda, 0xa6, 0x52, 0x60, 0x55, 0x09, 0x57, 0x14,
	0xd2, 0xdc, 0xbe, 0x4f, 0x6b, 0xb7, 0x81, 0x3f, 0xc2, 0xbb, 0x32, 0xdf, 0x77, 0x6b, 0xb8, 0xa0,
	0xf4, 0xc3, 0x9a, 0x2c, 0xba, 0x51, 0x34, 0x07, 0xdc, 0x49, 0xad, 0x62, 0x0f, 0xeb, 0x2c, 0xe0,
	0x73, 0xef, 0x27, 0x34, 0x87, 0xa4, 0xf7, 0x93, 0xd8, 0x4b, 0x8c, 0xff, 0x65, 0x4d, 0xbb, 0x47,
	0x42, 0x68, 0x0e, 0x84, 0xf4, 0x1d, 0x84, 0xc4, 0x1e, 0x21, 0xa6, 0xe7, 0x75, 0x64, 0x16, 0xbb,
	0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x75, 0xa9, 0x2c, 0x21, 0xd4, 0x01, 0x00, 0x00,
}
