// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/keyword_plan_negative_keyword.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "github.com/kritzware/google-ads-go/enums"
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

// A Keyword Plan negative keyword.
type KeywordPlanNegativeKeyword struct {
	// The resource name of the Keyword Plan negative keyword.
	// KeywordPlanNegativeKeyword resource names have the form:
	//
	//
	// `customers/{customer_id}/keywordPlanNegativeKeywords/{kp_negative_keyword_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The Keyword Plan campaign to which this negative keyword belongs.
	KeywordPlanCampaign *wrappers.StringValue `protobuf:"bytes,2,opt,name=keyword_plan_campaign,json=keywordPlanCampaign,proto3" json:"keyword_plan_campaign,omitempty"`
	// The ID of the Keyword Plan negative keyword.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The keyword text.
	Text *wrappers.StringValue `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	// The keyword match type.
	MatchType            enums.KeywordMatchTypeEnum_KeywordMatchType `protobuf:"varint,5,opt,name=match_type,json=matchType,proto3,enum=google.ads.googleads.v0.enums.KeywordMatchTypeEnum_KeywordMatchType" json:"match_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *KeywordPlanNegativeKeyword) Reset()         { *m = KeywordPlanNegativeKeyword{} }
func (m *KeywordPlanNegativeKeyword) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanNegativeKeyword) ProtoMessage()    {}
func (*KeywordPlanNegativeKeyword) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b897ee4c51adca9, []int{0}
}

func (m *KeywordPlanNegativeKeyword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanNegativeKeyword.Unmarshal(m, b)
}
func (m *KeywordPlanNegativeKeyword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanNegativeKeyword.Marshal(b, m, deterministic)
}
func (m *KeywordPlanNegativeKeyword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanNegativeKeyword.Merge(m, src)
}
func (m *KeywordPlanNegativeKeyword) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanNegativeKeyword.Size(m)
}
func (m *KeywordPlanNegativeKeyword) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanNegativeKeyword.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanNegativeKeyword proto.InternalMessageInfo

func (m *KeywordPlanNegativeKeyword) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlanNegativeKeyword) GetKeywordPlanCampaign() *wrappers.StringValue {
	if m != nil {
		return m.KeywordPlanCampaign
	}
	return nil
}

func (m *KeywordPlanNegativeKeyword) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlanNegativeKeyword) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *KeywordPlanNegativeKeyword) GetMatchType() enums.KeywordMatchTypeEnum_KeywordMatchType {
	if m != nil {
		return m.MatchType
	}
	return enums.KeywordMatchTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*KeywordPlanNegativeKeyword)(nil), "google.ads.googleads.v0.resources.KeywordPlanNegativeKeyword")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/keyword_plan_negative_keyword.proto", fileDescriptor_4b897ee4c51adca9)
}

var fileDescriptor_4b897ee4c51adca9 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0x25, 0xf9, 0xaa, 0xd0, 0xf1, 0x67, 0x11, 0x11, 0x42, 0x15, 0x6d, 0x15, 0xa1, 0x20, 0x4c,
	0x42, 0x95, 0x6e, 0x5c, 0xa5, 0x5a, 0x8a, 0x8a, 0x25, 0x44, 0xe9, 0x42, 0x02, 0x61, 0x9a, 0x8c,
	0x63, 0x68, 0xe6, 0x87, 0x99, 0x49, 0x6b, 0x5f, 0xc7, 0xa5, 0x8f, 0xe2, 0x0b, 0xf8, 0x02, 0x3e,
	0x88, 0x24, 0x33, 0x49, 0x91, 0x52, 0xfd, 0x76, 0x27, 0x37, 0xe7, 0x9c, 0x7b, 0xcf, 0xbd, 0x03,
	0x96, 0x84, 0x73, 0x52, 0xe1, 0x00, 0x15, 0x2a, 0x30, 0xb0, 0x41, 0xfb, 0x30, 0x90, 0x58, 0xf1,
	0x5a, 0xe6, 0x58, 0x05, 0x3b, 0x7c, 0x3c, 0x70, 0x59, 0x64, 0xa2, 0x42, 0x2c, 0x63, 0x98, 0x20,
	0x5d, 0xee, 0x71, 0x66, 0xab, 0x50, 0x48, 0xae, 0xb9, 0x37, 0x31, 0x5a, 0x88, 0x0a, 0x05, 0x7b,
	0x1b, 0xb8, 0x0f, 0x61, 0x6f, 0x33, 0x9a, 0x5f, 0xea, 0x84, 0x59, 0x4d, 0x4f, 0x5d, 0x28, 0xd2,
	0xf9, 0xd7, 0x4c, 0x1f, 0x05, 0x36, 0xd6, 0xa3, 0x47, 0x56, 0xd7, 0x7e, 0x6d, 0xeb, 0x2f, 0xc1,
	0x41, 0x22, 0x21, 0xb0, 0x54, 0xe6, 0xff, 0x93, 0x5f, 0x2e, 0x18, 0xbd, 0x37, 0xe2, 0xb8, 0x42,
	0x6c, 0x6d, 0x07, 0xb4, 0x25, 0xef, 0x29, 0xb8, 0xd3, 0xcd, 0x90, 0x31, 0x44, 0xb1, 0xef, 0x8c,
	0x9d, 0xe9, 0x30, 0xb9, 0xdd, 0x15, 0xd7, 0x88, 0x62, 0x2f, 0x06, 0xf7, 0xff, 0x4a, 0x99, 0x23,
	0x2a, 0x50, 0x49, 0x98, 0xef, 0x8e, 0x9d, 0xe9, 0xad, 0xd9, 0x43, 0x9b, 0x09, 0x76, 0x33, 0xc0,
	0x8f, 0x5a, 0x96, 0x8c, 0x6c, 0x50, 0x55, 0xe3, 0xe4, 0xde, 0xee, 0xd4, 0xfd, 0xb5, 0x15, 0x7a,
	0xcf, 0x81, 0x5b, 0x16, 0xfe, 0x55, 0x2b, 0x7f, 0x70, 0x26, 0x7f, 0xcb, 0xf4, 0xfc, 0xa5, 0x51,
	0xbb, 0x65, 0xe1, 0x85, 0x60, 0xa0, 0xf1, 0x37, 0xed, 0x0f, 0xae, 0xd1, 0xad, 0x65, 0x7a, 0x39,
	0x00, 0xa7, 0x45, 0xf9, 0x37, 0xc6, 0xce, 0xf4, 0xee, 0xec, 0x0d, 0xbc, 0x74, 0x84, 0x76, 0xc3,
	0xd0, 0x6e, 0xe4, 0x43, 0xa3, 0xfb, 0x74, 0x14, 0x78, 0xc9, 0x6a, 0x7a, 0x56, 0x4c, 0x86, 0xb4,
	0x83, 0x8b, 0xdf, 0x0e, 0x78, 0x96, 0x73, 0x0a, 0xff, 0x7b, 0xdb, 0xc5, 0xe3, 0xcb, 0x07, 0x88,
	0x9b, 0x10, 0xb1, 0xf3, 0xf9, 0x9d, 0x75, 0x21, 0xbc, 0x42, 0x8c, 0x40, 0x2e, 0x49, 0x40, 0x30,
	0x6b, 0x23, 0x76, 0xcf, 0x41, 0x94, 0xea, 0x1f, 0xef, 0xf0, 0x55, 0x8f, 0xbe, 0xbb, 0x57, 0xab,
	0x28, 0xfa, 0xe1, 0x4e, 0x56, 0xc6, 0x32, 0x2a, 0x14, 0x34, 0xb0, 0x41, 0x9b, 0x10, 0x26, 0x1d,
	0xf3, 0x67, 0xc7, 0x49, 0xa3, 0x42, 0xa5, 0x3d, 0x27, 0xdd, 0x84, 0x69, 0xcf, 0xd9, 0xde, 0x6c,
	0x87, 0x78, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x28, 0xf2, 0x8f, 0xb8, 0x0b, 0x03, 0x00, 0x00,
}