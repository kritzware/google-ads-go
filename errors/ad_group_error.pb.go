// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/ad_group_error.proto

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

// Enum describing possible ad group errors.
type AdGroupErrorEnum_AdGroupError int32

const (
	// Enum unspecified.
	AdGroupErrorEnum_UNSPECIFIED AdGroupErrorEnum_AdGroupError = 0
	// The received error code is not known in this version.
	AdGroupErrorEnum_UNKNOWN AdGroupErrorEnum_AdGroupError = 1
	// AdGroup with the same name already exists for the campaign.
	AdGroupErrorEnum_DUPLICATE_ADGROUP_NAME AdGroupErrorEnum_AdGroupError = 2
	// AdGroup name is not valid.
	AdGroupErrorEnum_INVALID_ADGROUP_NAME AdGroupErrorEnum_AdGroupError = 3
	// Advertiser is not allowed to target sites or set site bids that are not
	// on the Google Search Network.
	AdGroupErrorEnum_ADVERTISER_NOT_ON_CONTENT_NETWORK AdGroupErrorEnum_AdGroupError = 5
	// Bid amount is too big.
	AdGroupErrorEnum_BID_TOO_BIG AdGroupErrorEnum_AdGroupError = 6
	// AdGroup bid does not match the campaign's bidding strategy.
	AdGroupErrorEnum_BID_TYPE_AND_BIDDING_STRATEGY_MISMATCH AdGroupErrorEnum_AdGroupError = 7
	// AdGroup name is required for Add.
	AdGroupErrorEnum_MISSING_ADGROUP_NAME AdGroupErrorEnum_AdGroupError = 8
	// No link found between the ad group and the label.
	AdGroupErrorEnum_ADGROUP_LABEL_DOES_NOT_EXIST AdGroupErrorEnum_AdGroupError = 9
	// The label has already been attached to the ad group.
	AdGroupErrorEnum_ADGROUP_LABEL_ALREADY_EXISTS AdGroupErrorEnum_AdGroupError = 10
	// The CriterionTypeGroup is not supported for the content bid dimension.
	AdGroupErrorEnum_INVALID_CONTENT_BID_CRITERION_TYPE_GROUP AdGroupErrorEnum_AdGroupError = 11
	// The ad group type is not compatible with the campaign channel type.
	AdGroupErrorEnum_AD_GROUP_TYPE_NOT_VALID_FOR_ADVERTISING_CHANNEL_TYPE AdGroupErrorEnum_AdGroupError = 12
	// The ad group type is not supported in the country of sale of the
	// campaign.
	AdGroupErrorEnum_ADGROUP_TYPE_NOT_SUPPORTED_FOR_CAMPAIGN_SALES_COUNTRY AdGroupErrorEnum_AdGroupError = 13
	// Ad groups of AdGroupType.SEARCH_DYNAMIC_ADS can only be added to
	// campaigns that have DynamicSearchAdsSetting attached.
	AdGroupErrorEnum_CANNOT_ADD_ADGROUP_OF_TYPE_DSA_TO_CAMPAIGN_WITHOUT_DSA_SETTING AdGroupErrorEnum_AdGroupError = 14
)

var AdGroupErrorEnum_AdGroupError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "DUPLICATE_ADGROUP_NAME",
	3:  "INVALID_ADGROUP_NAME",
	5:  "ADVERTISER_NOT_ON_CONTENT_NETWORK",
	6:  "BID_TOO_BIG",
	7:  "BID_TYPE_AND_BIDDING_STRATEGY_MISMATCH",
	8:  "MISSING_ADGROUP_NAME",
	9:  "ADGROUP_LABEL_DOES_NOT_EXIST",
	10: "ADGROUP_LABEL_ALREADY_EXISTS",
	11: "INVALID_CONTENT_BID_CRITERION_TYPE_GROUP",
	12: "AD_GROUP_TYPE_NOT_VALID_FOR_ADVERTISING_CHANNEL_TYPE",
	13: "ADGROUP_TYPE_NOT_SUPPORTED_FOR_CAMPAIGN_SALES_COUNTRY",
	14: "CANNOT_ADD_ADGROUP_OF_TYPE_DSA_TO_CAMPAIGN_WITHOUT_DSA_SETTING",
}

var AdGroupErrorEnum_AdGroupError_value = map[string]int32{
	"UNSPECIFIED":                                                    0,
	"UNKNOWN":                                                        1,
	"DUPLICATE_ADGROUP_NAME":                                         2,
	"INVALID_ADGROUP_NAME":                                           3,
	"ADVERTISER_NOT_ON_CONTENT_NETWORK":                              5,
	"BID_TOO_BIG":                                                    6,
	"BID_TYPE_AND_BIDDING_STRATEGY_MISMATCH":                         7,
	"MISSING_ADGROUP_NAME":                                           8,
	"ADGROUP_LABEL_DOES_NOT_EXIST":                                   9,
	"ADGROUP_LABEL_ALREADY_EXISTS":                                   10,
	"INVALID_CONTENT_BID_CRITERION_TYPE_GROUP":                       11,
	"AD_GROUP_TYPE_NOT_VALID_FOR_ADVERTISING_CHANNEL_TYPE":           12,
	"ADGROUP_TYPE_NOT_SUPPORTED_FOR_CAMPAIGN_SALES_COUNTRY":          13,
	"CANNOT_ADD_ADGROUP_OF_TYPE_DSA_TO_CAMPAIGN_WITHOUT_DSA_SETTING": 14,
}

func (x AdGroupErrorEnum_AdGroupError) String() string {
	return proto.EnumName(AdGroupErrorEnum_AdGroupError_name, int32(x))
}

func (AdGroupErrorEnum_AdGroupError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d2823512e8e4828, []int{0, 0}
}

// Container for enum describing possible ad group errors.
type AdGroupErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupErrorEnum) Reset()         { *m = AdGroupErrorEnum{} }
func (m *AdGroupErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupErrorEnum) ProtoMessage()    {}
func (*AdGroupErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2823512e8e4828, []int{0}
}

func (m *AdGroupErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupErrorEnum.Unmarshal(m, b)
}
func (m *AdGroupErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupErrorEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupErrorEnum.Merge(m, src)
}
func (m *AdGroupErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupErrorEnum.Size(m)
}
func (m *AdGroupErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.AdGroupErrorEnum_AdGroupError", AdGroupErrorEnum_AdGroupError_name, AdGroupErrorEnum_AdGroupError_value)
	proto.RegisterType((*AdGroupErrorEnum)(nil), "google.ads.googleads.v0.errors.AdGroupErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/ad_group_error.proto", fileDescriptor_0d2823512e8e4828)
}

var fileDescriptor_0d2823512e8e4828 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x71, 0x6b, 0xd3, 0x40,
	0x1c, 0x75, 0xdd, 0xdc, 0xf4, 0x3a, 0xf5, 0x3c, 0x44, 0x44, 0x64, 0x68, 0x41, 0x11, 0x91, 0xb4,
	0x30, 0x05, 0x8d, 0x20, 0xfc, 0x92, 0xbb, 0xa6, 0xc7, 0xd2, 0xbb, 0x90, 0xbb, 0xb4, 0x56, 0x0a,
	0x47, 0x35, 0x25, 0x08, 0x5b, 0x53, 0x12, 0xb7, 0x2f, 0xe1, 0x97, 0x10, 0xff, 0xf4, 0xa3, 0xf8,
	0x51, 0xc4, 0x0f, 0x21, 0x97, 0x33, 0xdd, 0x86, 0xe8, 0x5f, 0xf9, 0xdd, 0xbb, 0xf7, 0x7e, 0xef,
	0x25, 0x79, 0xe8, 0xb0, 0x28, 0xcb, 0xe2, 0x78, 0xd9, 0x5f, 0xe4, 0x75, 0xdf, 0x8d, 0x76, 0x3a,
	0x1b, 0xf4, 0x97, 0x55, 0x55, 0x56, 0x75, 0x7f, 0x91, 0x9b, 0xa2, 0x2a, 0x4f, 0xd7, 0xa6, 0x39,
	0x7b, 0xeb, 0xaa, 0xfc, 0x5c, 0x92, 0x03, 0xc7, 0xf4, 0x16, 0x79, 0xed, 0x6d, 0x44, 0xde, 0xd9,
	0xc0, 0x73, 0xa2, 0xde, 0xd7, 0x1d, 0x84, 0x21, 0x8f, 0xac, 0x8e, 0x59, 0x84, 0xad, 0x4e, 0x4f,
	0x7a, 0x5f, 0x76, 0xd0, 0xfe, 0x45, 0x90, 0xdc, 0x42, 0xdd, 0x4c, 0xa8, 0x84, 0x85, 0x7c, 0xc8,
	0x19, 0xc5, 0x57, 0x48, 0x17, 0xed, 0x65, 0xe2, 0x48, 0xc8, 0xa9, 0xc0, 0x5b, 0xe4, 0x3e, 0xba,
	0x4b, 0xb3, 0x24, 0xe6, 0x21, 0x68, 0x66, 0x80, 0x46, 0xa9, 0xcc, 0x12, 0x23, 0x60, 0xcc, 0x70,
	0x87, 0xdc, 0x43, 0x77, 0xb8, 0x98, 0x40, 0xcc, 0xe9, 0xe5, 0x9b, 0x6d, 0xf2, 0x18, 0x3d, 0x02,
	0x3a, 0x61, 0xa9, 0xe6, 0x8a, 0xa5, 0x46, 0x48, 0x6d, 0xa4, 0x30, 0xa1, 0x14, 0x9a, 0x09, 0x6d,
	0x04, 0xd3, 0x53, 0x99, 0x1e, 0xe1, 0xab, 0xd6, 0x3a, 0xe0, 0xd4, 0x68, 0x29, 0x4d, 0xc0, 0x23,
	0xbc, 0x4b, 0x9e, 0xa1, 0x27, 0x0d, 0x30, 0x4b, 0x98, 0x01, 0x41, 0x4d, 0xc0, 0x29, 0xe5, 0x22,
	0x32, 0x4a, 0xa7, 0xa0, 0x59, 0x34, 0x33, 0x63, 0xae, 0xc6, 0xa0, 0xc3, 0x11, 0xde, 0xb3, 0xee,
	0x63, 0xae, 0x94, 0xbd, 0xbe, 0xe4, 0x7e, 0x8d, 0x3c, 0x44, 0x0f, 0x5a, 0x24, 0x86, 0x80, 0xc5,
	0x86, 0x4a, 0xa6, 0x9a, 0x14, 0xec, 0x1d, 0x57, 0x1a, 0x5f, 0xff, 0x9b, 0x01, 0x71, 0xca, 0x80,
	0xce, 0x1c, 0x41, 0x61, 0x44, 0x9e, 0xa3, 0xa7, 0xed, 0xbb, 0xb5, 0xb9, 0x6d, 0xb2, 0x30, 0xe5,
	0x9a, 0xa5, 0x5c, 0x0a, 0x97, 0xb1, 0x59, 0x81, 0xbb, 0xe4, 0x15, 0x7a, 0x01, 0xd4, 0x9d, 0xdc,
	0x85, 0x35, 0x73, 0xea, 0xa1, 0x4c, 0x4d, 0xfb, 0x2d, 0x6c, 0xd6, 0x70, 0x04, 0x42, 0xb0, 0xb8,
	0xa1, 0xe1, 0x7d, 0xf2, 0x1a, 0xbd, 0x6c, 0x93, 0x6c, 0x84, 0x2a, 0x4b, 0x12, 0x99, 0x6a, 0xe6,
	0xc4, 0x21, 0x8c, 0x13, 0xe0, 0x91, 0x30, 0x0a, 0x62, 0xa6, 0x4c, 0x28, 0x33, 0xa1, 0xd3, 0x19,
	0xbe, 0x41, 0x02, 0xf4, 0x36, 0x04, 0x61, 0x05, 0x40, 0xcf, 0xff, 0x80, 0x1c, 0xba, 0x45, 0x54,
	0x81, 0xd1, 0xf2, 0x5c, 0x3d, 0xe5, 0x7a, 0x24, 0x33, 0xdd, 0xe0, 0x8a, 0x69, 0xcd, 0x45, 0x84,
	0x6f, 0x06, 0xbf, 0xb6, 0x50, 0xef, 0x63, 0x79, 0xe2, 0xfd, 0xbf, 0x49, 0xc1, 0xed, 0x8b, 0x8d,
	0x49, 0x6c, 0xf9, 0x92, 0xad, 0xf7, 0xf4, 0x8f, 0xa8, 0x28, 0x8f, 0x17, 0xab, 0xc2, 0x2b, 0xab,
	0xa2, 0x5f, 0x2c, 0x57, 0x4d, 0x35, 0xdb, 0x0e, 0xaf, 0x3f, 0xd5, 0xff, 0xaa, 0xf4, 0x1b, 0xf7,
	0xf8, 0xd6, 0xd9, 0x8e, 0x00, 0xbe, 0x77, 0x0e, 0x22, 0xb7, 0x0c, 0xf2, 0xda, 0x73, 0xa3, 0x9d,
	0x26, 0x03, 0xaf, 0xb1, 0xac, 0x7f, 0xb4, 0x84, 0x39, 0xe4, 0xf5, 0x7c, 0x43, 0x98, 0x4f, 0x06,
	0x73, 0x47, 0xf8, 0xd9, 0xe9, 0x39, 0xd4, 0xf7, 0x21, 0xaf, 0x7d, 0x7f, 0x43, 0xf1, 0xfd, 0xc9,
	0xc0, 0xf7, 0x1d, 0xe9, 0xc3, 0x6e, 0x93, 0xee, 0xf0, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd,
	0x2c, 0x14, 0x00, 0x6f, 0x03, 0x00, 0x00,
}
