// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/affiliate_location_feed_relationship_type.proto

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

// Possible values for a relationship type for an affiliate location feed.
type AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType int32

const (
	// Not specified.
	AffiliateLocationFeedRelationshipTypeEnum_UNSPECIFIED AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType = 0
	// Used for return value only. Represents value unknown in this version.
	AffiliateLocationFeedRelationshipTypeEnum_UNKNOWN AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType = 1
	// General retailer relationship.
	AffiliateLocationFeedRelationshipTypeEnum_GENERAL_RETAILER AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType = 2
)

var AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "GENERAL_RETAILER",
}

var AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType_value = map[string]int32{
	"UNSPECIFIED":      0,
	"UNKNOWN":          1,
	"GENERAL_RETAILER": 2,
}

func (x AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType) String() string {
	return proto.EnumName(AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType_name, int32(x))
}

func (AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_15e58409f3245b86, []int{0, 0}
}

// Container for enum describing possible values for a relationship type for
// an affiliate location feed.
type AffiliateLocationFeedRelationshipTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AffiliateLocationFeedRelationshipTypeEnum) Reset() {
	*m = AffiliateLocationFeedRelationshipTypeEnum{}
}
func (m *AffiliateLocationFeedRelationshipTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AffiliateLocationFeedRelationshipTypeEnum) ProtoMessage()    {}
func (*AffiliateLocationFeedRelationshipTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e58409f3245b86, []int{0}
}

func (m *AffiliateLocationFeedRelationshipTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AffiliateLocationFeedRelationshipTypeEnum.Unmarshal(m, b)
}
func (m *AffiliateLocationFeedRelationshipTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AffiliateLocationFeedRelationshipTypeEnum.Marshal(b, m, deterministic)
}
func (m *AffiliateLocationFeedRelationshipTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AffiliateLocationFeedRelationshipTypeEnum.Merge(m, src)
}
func (m *AffiliateLocationFeedRelationshipTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AffiliateLocationFeedRelationshipTypeEnum.Size(m)
}
func (m *AffiliateLocationFeedRelationshipTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AffiliateLocationFeedRelationshipTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AffiliateLocationFeedRelationshipTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType", AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType_name, AffiliateLocationFeedRelationshipTypeEnum_AffiliateLocationFeedRelationshipType_value)
	proto.RegisterType((*AffiliateLocationFeedRelationshipTypeEnum)(nil), "google.ads.googleads.v0.enums.AffiliateLocationFeedRelationshipTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/affiliate_location_feed_relationship_type.proto", fileDescriptor_15e58409f3245b86)
}

var fileDescriptor_15e58409f3245b86 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x00, 0x86, 0x6d, 0x05, 0x85, 0xec, 0x60, 0x29, 0x5e, 0x77, 0xd8, 0xc0, 0x83, 0x1e, 0xd2, 0x82,
	0xb7, 0x78, 0xca, 0x34, 0x1b, 0xc3, 0x59, 0x47, 0xdd, 0x2a, 0x68, 0xa1, 0xc4, 0x25, 0x8d, 0x85,
	0xb6, 0x29, 0x4d, 0x37, 0xd8, 0x1b, 0xf8, 0x1c, 0x1e, 0x7d, 0x14, 0x1f, 0xc5, 0x97, 0x50, 0x9a,
	0xae, 0xc5, 0x8b, 0xb2, 0x4b, 0xf8, 0x93, 0xff, 0xe7, 0xe3, 0xcf, 0x0f, 0xee, 0x84, 0x94, 0x22,
	0xe5, 0x0e, 0x65, 0xca, 0x69, 0x64, 0xad, 0x36, 0xae, 0xc3, 0xf3, 0x75, 0xa6, 0x1c, 0x1a, 0xc7,
	0x49, 0x9a, 0xd0, 0x8a, 0x47, 0xa9, 0x5c, 0xd1, 0x2a, 0x91, 0x79, 0x14, 0x73, 0xce, 0xa2, 0x92,
	0xa7, 0xfa, 0xa6, 0x5e, 0x93, 0x22, 0xaa, 0xb6, 0x05, 0x87, 0x45, 0x29, 0x2b, 0x69, 0xf7, 0x1b,
	0x06, 0xa4, 0x4c, 0xc1, 0x0e, 0x07, 0x37, 0x2e, 0xd4, 0xb8, 0xe1, 0x9b, 0x01, 0xce, 0x71, 0x8b,
	0x9c, 0xed, 0x88, 0x63, 0xce, 0x99, 0xff, 0x8b, 0xb7, 0xd8, 0x16, 0x9c, 0xe4, 0xeb, 0x6c, 0xf8,
	0x0c, 0xce, 0xf6, 0x0a, 0xdb, 0x27, 0xa0, 0xb7, 0xf4, 0x1e, 0xe6, 0xe4, 0x7a, 0x3a, 0x9e, 0x92,
	0x1b, 0xeb, 0xc0, 0xee, 0x81, 0xe3, 0xa5, 0x77, 0xeb, 0xdd, 0x3f, 0x7a, 0x96, 0x61, 0x9f, 0x02,
	0x6b, 0x42, 0x3c, 0xe2, 0xe3, 0x59, 0xe4, 0x93, 0x05, 0x9e, 0xce, 0x88, 0x6f, 0x99, 0xa3, 0x6f,
	0x03, 0x0c, 0x56, 0x32, 0x83, 0xff, 0x16, 0x1e, 0x5d, 0xec, 0x55, 0x60, 0x5e, 0xff, 0x7d, 0x6e,
	0x3c, 0x8d, 0x76, 0x30, 0x21, 0x53, 0x9a, 0x0b, 0x28, 0x4b, 0xe1, 0x08, 0x9e, 0xeb, 0x65, 0xda,
	0x71, 0x8b, 0x44, 0xfd, 0xb1, 0xf5, 0x95, 0x3e, 0xdf, 0xcd, 0xc3, 0x09, 0xc6, 0x1f, 0x66, 0x7f,
	0xd2, 0xa0, 0x30, 0x53, 0xb0, 0x91, 0xb5, 0x0a, 0x5c, 0x58, 0x2f, 0xa3, 0x3e, 0x5b, 0x3f, 0xc4,
	0x4c, 0x85, 0x9d, 0x1f, 0x06, 0x6e, 0xa8, 0xfd, 0x2f, 0x73, 0xd0, 0x3c, 0x22, 0x84, 0x99, 0x42,
	0xa8, 0x4b, 0x20, 0x14, 0xb8, 0x08, 0xe9, 0xcc, 0xcb, 0x91, 0x2e, 0x76, 0xf9, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x4c, 0xe6, 0x2f, 0xd6, 0x03, 0x02, 0x00, 0x00,
}
