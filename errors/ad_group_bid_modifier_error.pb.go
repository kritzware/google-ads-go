// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/ad_group_bid_modifier_error.proto

package errors

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

// Enum describing possible ad group bid modifier errors.
type AdGroupBidModifierErrorEnum_AdGroupBidModifierError int32

const (
	// Enum unspecified.
	AdGroupBidModifierErrorEnum_UNSPECIFIED AdGroupBidModifierErrorEnum_AdGroupBidModifierError = 0
	// The received error code is not known in this version.
	AdGroupBidModifierErrorEnum_UNKNOWN AdGroupBidModifierErrorEnum_AdGroupBidModifierError = 1
	// The criterion ID does not support bid modification.
	AdGroupBidModifierErrorEnum_CRITERION_ID_NOT_SUPPORTED AdGroupBidModifierErrorEnum_AdGroupBidModifierError = 2
	// Cannot override the bid modifier for the given criterion ID if the parent
	// campaign is opted out of the same criterion.
	AdGroupBidModifierErrorEnum_CANNOT_OVERRIDE_OPTED_OUT_CAMPAIGN_CRITERION_BID_MODIFIER AdGroupBidModifierErrorEnum_AdGroupBidModifierError = 3
)

var AdGroupBidModifierErrorEnum_AdGroupBidModifierError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CRITERION_ID_NOT_SUPPORTED",
	3: "CANNOT_OVERRIDE_OPTED_OUT_CAMPAIGN_CRITERION_BID_MODIFIER",
}

var AdGroupBidModifierErrorEnum_AdGroupBidModifierError_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"CRITERION_ID_NOT_SUPPORTED": 2,
	"CANNOT_OVERRIDE_OPTED_OUT_CAMPAIGN_CRITERION_BID_MODIFIER": 3,
}

func (x AdGroupBidModifierErrorEnum_AdGroupBidModifierError) String() string {
	return proto.EnumName(AdGroupBidModifierErrorEnum_AdGroupBidModifierError_name, int32(x))
}

func (AdGroupBidModifierErrorEnum_AdGroupBidModifierError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4a6d661fb90a168, []int{0, 0}
}

// Container for enum describing possible ad group bid modifier errors.
type AdGroupBidModifierErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupBidModifierErrorEnum) Reset()         { *m = AdGroupBidModifierErrorEnum{} }
func (m *AdGroupBidModifierErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupBidModifierErrorEnum) ProtoMessage()    {}
func (*AdGroupBidModifierErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a6d661fb90a168, []int{0}
}

func (m *AdGroupBidModifierErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupBidModifierErrorEnum.Unmarshal(m, b)
}
func (m *AdGroupBidModifierErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupBidModifierErrorEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupBidModifierErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupBidModifierErrorEnum.Merge(m, src)
}
func (m *AdGroupBidModifierErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupBidModifierErrorEnum.Size(m)
}
func (m *AdGroupBidModifierErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupBidModifierErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupBidModifierErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.AdGroupBidModifierErrorEnum_AdGroupBidModifierError", AdGroupBidModifierErrorEnum_AdGroupBidModifierError_name, AdGroupBidModifierErrorEnum_AdGroupBidModifierError_value)
	proto.RegisterType((*AdGroupBidModifierErrorEnum)(nil), "google.ads.googleads.v0.errors.AdGroupBidModifierErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/ad_group_bid_modifier_error.proto", fileDescriptor_f4a6d661fb90a168)
}

var fileDescriptor_f4a6d661fb90a168 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x8a, 0xdb, 0x30,
	0x14, 0x86, 0x6b, 0x07, 0x5a, 0x50, 0x16, 0x35, 0xde, 0x14, 0xda, 0x92, 0x85, 0x0f, 0x20, 0x1b,
	0xba, 0xaa, 0x4a, 0xa1, 0xb2, 0xa5, 0x1a, 0x51, 0x22, 0x19, 0xc7, 0x76, 0xa1, 0x18, 0x84, 0x53,
	0xb9, 0xc6, 0x90, 0x44, 0xc1, 0x6a, 0x72, 0x94, 0x39, 0xc0, 0x2c, 0xe7, 0x00, 0x73, 0x88, 0x39,
	0xca, 0x5c, 0x60, 0xb6, 0x83, 0xad, 0x24, 0xb3, 0xca, 0xac, 0xfc, 0xf3, 0xde, 0xff, 0xbe, 0xf7,
	0xfc, 0x0b, 0xfc, 0xe8, 0xb4, 0xee, 0x36, 0x6d, 0xd8, 0x28, 0x13, 0x5a, 0x39, 0xaa, 0x63, 0x14,
	0xb6, 0xc3, 0xa0, 0x07, 0x13, 0x36, 0x4a, 0x76, 0x83, 0x3e, 0xec, 0xe5, 0xba, 0x57, 0x72, 0xab,
	0x55, 0xff, 0xaf, 0x6f, 0x07, 0x39, 0x35, 0xe1, 0x7e, 0xd0, 0xff, 0xb5, 0xbf, 0xb0, 0x63, 0xb0,
	0x51, 0x06, 0x5e, 0x08, 0xf0, 0x18, 0x41, 0x4b, 0x08, 0xee, 0x1d, 0xf0, 0x09, 0xab, 0x74, 0x84,
	0xc4, 0xbd, 0x5a, 0x9e, 0x10, 0x74, 0x6c, 0xd2, 0xdd, 0x61, 0x1b, 0xdc, 0x38, 0xe0, 0xc3, 0x95,
	0xbe, 0xff, 0x1e, 0xcc, 0x4b, 0xbe, 0xca, 0x68, 0xc2, 0x7e, 0x32, 0x4a, 0xbc, 0x37, 0xfe, 0x1c,
	0xbc, 0x2b, 0xf9, 0x2f, 0x2e, 0x7e, 0x73, 0xcf, 0xf1, 0x17, 0xe0, 0x63, 0x92, 0xb3, 0x82, 0xe6,
	0x4c, 0x70, 0xc9, 0x88, 0xe4, 0xa2, 0x90, 0xab, 0x32, 0xcb, 0x44, 0x5e, 0x50, 0xe2, 0xb9, 0xfe,
	0x77, 0xf0, 0x35, 0xc1, 0x7c, 0xac, 0x8a, 0x8a, 0xe6, 0x39, 0x23, 0x54, 0x8a, 0xac, 0xa0, 0x44,
	0x8a, 0xb2, 0x90, 0x09, 0x5e, 0x66, 0x98, 0xa5, 0x5c, 0xbe, 0x20, 0x62, 0x46, 0xe4, 0x52, 0x90,
	0x71, 0x57, 0xee, 0xcd, 0xe2, 0x27, 0x07, 0x04, 0x7f, 0xf5, 0x16, 0xbe, 0xfe, 0x7f, 0xf1, 0xe7,
	0x2b, 0xc7, 0x67, 0x63, 0x3a, 0x99, 0xf3, 0x87, 0x9c, 0xe6, 0x3b, 0xbd, 0x69, 0x76, 0x1d, 0xd4,
	0x43, 0x17, 0x76, 0xed, 0x6e, 0xca, 0xee, 0x9c, 0xf8, 0xbe, 0x37, 0xd7, 0x1e, 0xe0, 0x9b, 0xfd,
	0xdc, 0xba, 0xb3, 0x14, 0xe3, 0x3b, 0x77, 0x91, 0x5a, 0x18, 0x56, 0x06, 0x5a, 0x39, 0xaa, 0x2a,
	0x82, 0xd3, 0x4a, 0xf3, 0x70, 0x36, 0xd4, 0x58, 0x99, 0xfa, 0x62, 0xa8, 0xab, 0xa8, 0xb6, 0x86,
	0x47, 0x37, 0xb0, 0x55, 0x84, 0xb0, 0x32, 0x08, 0x5d, 0x2c, 0x08, 0x55, 0x11, 0x42, 0xd6, 0xb4,
	0x7e, 0x3b, 0x5d, 0xf7, 0xe5, 0x39, 0x00, 0x00, 0xff, 0xff, 0xac, 0x26, 0x85, 0xb6, 0x1d, 0x02,
	0x00, 0x00,
}
