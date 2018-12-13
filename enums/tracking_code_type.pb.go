// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/tracking_code_type.proto

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

// The type of the generated tag snippets for tracking conversions.
type TrackingCodeTypeEnum_TrackingCodeType int32

const (
	// Not specified.
	TrackingCodeTypeEnum_UNSPECIFIED TrackingCodeTypeEnum_TrackingCodeType = 0
	// Used for return value only. Represents value unknown in this version.
	TrackingCodeTypeEnum_UNKNOWN TrackingCodeTypeEnum_TrackingCodeType = 1
	// The snippet that is fired as a result of a website page loading.
	TrackingCodeTypeEnum_WEBPAGE TrackingCodeTypeEnum_TrackingCodeType = 2
	// The snippet contains a JavaScript function which fires the tag. This
	// function is typically called from an onClick handler added to a link or
	// button element on the page.
	TrackingCodeTypeEnum_WEBPAGE_ONCLICK TrackingCodeTypeEnum_TrackingCodeType = 3
	// For embedding on a mobile webpage. The snippet contains a JavaScript
	// function which fires the tag.
	TrackingCodeTypeEnum_CLICK_TO_CALL TrackingCodeTypeEnum_TrackingCodeType = 4
)

var TrackingCodeTypeEnum_TrackingCodeType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "WEBPAGE",
	3: "WEBPAGE_ONCLICK",
	4: "CLICK_TO_CALL",
}

var TrackingCodeTypeEnum_TrackingCodeType_value = map[string]int32{
	"UNSPECIFIED":     0,
	"UNKNOWN":         1,
	"WEBPAGE":         2,
	"WEBPAGE_ONCLICK": 3,
	"CLICK_TO_CALL":   4,
}

func (x TrackingCodeTypeEnum_TrackingCodeType) String() string {
	return proto.EnumName(TrackingCodeTypeEnum_TrackingCodeType_name, int32(x))
}

func (TrackingCodeTypeEnum_TrackingCodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b23766bb1613558, []int{0, 0}
}

// Container for enum describing the type of the generated tag snippets for
// tracking conversions.
type TrackingCodeTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackingCodeTypeEnum) Reset()         { *m = TrackingCodeTypeEnum{} }
func (m *TrackingCodeTypeEnum) String() string { return proto.CompactTextString(m) }
func (*TrackingCodeTypeEnum) ProtoMessage()    {}
func (*TrackingCodeTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b23766bb1613558, []int{0}
}

func (m *TrackingCodeTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackingCodeTypeEnum.Unmarshal(m, b)
}
func (m *TrackingCodeTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackingCodeTypeEnum.Marshal(b, m, deterministic)
}
func (m *TrackingCodeTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackingCodeTypeEnum.Merge(m, src)
}
func (m *TrackingCodeTypeEnum) XXX_Size() int {
	return xxx_messageInfo_TrackingCodeTypeEnum.Size(m)
}
func (m *TrackingCodeTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackingCodeTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_TrackingCodeTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.TrackingCodeTypeEnum_TrackingCodeType", TrackingCodeTypeEnum_TrackingCodeType_name, TrackingCodeTypeEnum_TrackingCodeType_value)
	proto.RegisterType((*TrackingCodeTypeEnum)(nil), "google.ads.googleads.v0.enums.TrackingCodeTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/tracking_code_type.proto", fileDescriptor_1b23766bb1613558)
}

var fileDescriptor_1b23766bb1613558 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0x92, 0xa2, 0xc4, 0xe4, 0xec, 0xcc, 0xbc, 0xf4, 0xf8, 0xe4, 0xfc, 0x94,
	0xd4, 0xf8, 0x92, 0xca, 0x82, 0x54, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x62,
	0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x3e, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x3e, 0xa5, 0x5a, 0x2e,
	0x91, 0x10, 0xa8, 0x56, 0xe7, 0xfc, 0x94, 0xd4, 0x90, 0xca, 0x82, 0x54, 0xd7, 0xbc, 0xd2, 0x5c,
	0xa5, 0x54, 0x2e, 0x01, 0x74, 0x71, 0x21, 0x7e, 0x2e, 0xee, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67,
	0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01, 0x06, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff,
	0x70, 0x3f, 0x01, 0x46, 0x10, 0x27, 0xdc, 0xd5, 0x29, 0xc0, 0xd1, 0xdd, 0x55, 0x80, 0x49, 0x48,
	0x98, 0x8b, 0x1f, 0xca, 0x89, 0xf7, 0xf7, 0x73, 0xf6, 0xf1, 0x74, 0xf6, 0x16, 0x60, 0x16, 0x12,
	0xe4, 0xe2, 0x05, 0x33, 0xe3, 0x43, 0xfc, 0xe3, 0x9d, 0x1d, 0x7d, 0x7c, 0x04, 0x58, 0x9c, 0x8e,
	0x31, 0x72, 0x29, 0x26, 0xe7, 0xe7, 0xea, 0xe1, 0x75, 0xa4, 0x93, 0x28, 0xba, 0x53, 0x02, 0x40,
	0x5e, 0x0b, 0x60, 0x8c, 0x72, 0x82, 0xea, 0x4b, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f,
	0x4a, 0xd7, 0x4f, 0x4f, 0xcd, 0x03, 0x7b, 0x1c, 0x16, 0x48, 0x05, 0x99, 0xc5, 0x38, 0xc2, 0xcc,
	0x1a, 0x4c, 0x2e, 0x62, 0x62, 0x76, 0x77, 0x74, 0x5c, 0xc5, 0x24, 0xeb, 0x0e, 0x31, 0xca, 0x31,
	0xa5, 0x58, 0x0f, 0xc2, 0x04, 0xb1, 0xc2, 0x0c, 0xf4, 0x40, 0xc1, 0x51, 0x7c, 0x0a, 0x26, 0x1f,
	0xe3, 0x98, 0x52, 0x1c, 0x03, 0x97, 0x8f, 0x09, 0x33, 0x88, 0x01, 0xcb, 0x27, 0xb1, 0x81, 0x2d,
	0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x3a, 0x69, 0xa9, 0xa7, 0x01, 0x00, 0x00,
}