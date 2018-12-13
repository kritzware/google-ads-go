// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/policy_topic_evidence_destination_mismatch_url_type.proto

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

// The possible policy topic evidence destination mismatch url types.
type PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType int32

const (
	// No value has been specified.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_UNSPECIFIED PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_UNKNOWN PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 1
	// The display url.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_DISPLAY_URL PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 2
	// The final url.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_FINAL_URL PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 3
	// The final mobile url.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_FINAL_MOBILE_URL PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 4
	// The tracking url template, with substituted desktop url.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_TRACKING_URL PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 5
	// The tracking url template, with substituted mobile url.
	PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_MOBILE_TRACKING_URL PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType = 6
)

var PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DISPLAY_URL",
	3: "FINAL_URL",
	4: "FINAL_MOBILE_URL",
	5: "TRACKING_URL",
	6: "MOBILE_TRACKING_URL",
}

var PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"DISPLAY_URL":         2,
	"FINAL_URL":           3,
	"FINAL_MOBILE_URL":    4,
	"TRACKING_URL":        5,
	"MOBILE_TRACKING_URL": 6,
}

func (x PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType) String() string {
	return proto.EnumName(PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType_name, int32(x))
}

func (PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33fc5ad251e25646, []int{0, 0}
}

// Container for enum describing possible policy topic evidence destination
// mismatch url types.
type PolicyTopicEvidenceDestinationMismatchUrlTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) Reset() {
	*m = PolicyTopicEvidenceDestinationMismatchUrlTypeEnum{}
}
func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) String() string {
	return proto.CompactTextString(m)
}
func (*PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) ProtoMessage() {}
func (*PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_33fc5ad251e25646, []int{0}
}

func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyTopicEvidenceDestinationMismatchUrlTypeEnum.Unmarshal(m, b)
}
func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyTopicEvidenceDestinationMismatchUrlTypeEnum.Marshal(b, m, deterministic)
}
func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyTopicEvidenceDestinationMismatchUrlTypeEnum.Merge(m, src)
}
func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyTopicEvidenceDestinationMismatchUrlTypeEnum.Size(m)
}
func (m *PolicyTopicEvidenceDestinationMismatchUrlTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyTopicEvidenceDestinationMismatchUrlTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyTopicEvidenceDestinationMismatchUrlTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType", PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType_name, PolicyTopicEvidenceDestinationMismatchUrlTypeEnum_PolicyTopicEvidenceDestinationMismatchUrlType_value)
	proto.RegisterType((*PolicyTopicEvidenceDestinationMismatchUrlTypeEnum)(nil), "google.ads.googleads.v0.enums.PolicyTopicEvidenceDestinationMismatchUrlTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/policy_topic_evidence_destination_mismatch_url_type.proto", fileDescriptor_33fc5ad251e25646)
}

var fileDescriptor_33fc5ad251e25646 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4e, 0x83, 0x30,
	0x18, 0xc7, 0x65, 0xd3, 0x19, 0x3b, 0x8d, 0x04, 0x4d, 0x3c, 0xed, 0xe0, 0xee, 0x16, 0xd4, 0xa3,
	0x27, 0xd8, 0xd8, 0x42, 0xc6, 0x18, 0xd9, 0xc6, 0x16, 0x0d, 0x49, 0x83, 0xd0, 0x60, 0x13, 0x68,
	0x09, 0x65, 0x4b, 0xf6, 0x3a, 0x1e, 0x3d, 0xf8, 0x20, 0xbe, 0x86, 0x3e, 0x88, 0xa1, 0xb0, 0x19,
	0x0f, 0x9a, 0xec, 0xd2, 0xfc, 0xbf, 0xef, 0xdf, 0xfe, 0x9a, 0xef, 0xff, 0x81, 0x65, 0xcc, 0x58,
	0x9c, 0x60, 0x35, 0x88, 0xb8, 0x5a, 0xc9, 0x52, 0xad, 0x35, 0x15, 0xd3, 0x55, 0xca, 0xd5, 0x8c,
	0x25, 0x24, 0xdc, 0xa0, 0x82, 0x65, 0x24, 0x44, 0x78, 0x4d, 0x22, 0x4c, 0x43, 0x8c, 0x22, 0xcc,
	0x0b, 0x42, 0x83, 0x82, 0x30, 0x8a, 0x52, 0xc2, 0xd3, 0xa0, 0x08, 0x5f, 0xd0, 0x2a, 0x4f, 0x50,
	0xb1, 0xc9, 0x30, 0xcc, 0x72, 0x56, 0x30, 0xa5, 0x53, 0xd1, 0x60, 0x10, 0x71, 0xb8, 0x03, 0xc3,
	0xb5, 0x06, 0x05, 0xb8, 0xfb, 0x25, 0x81, 0x5b, 0x57, 0xc0, 0xe7, 0x25, 0xdb, 0xac, 0xd1, 0xfd,
	0x1f, 0xf2, 0xb8, 0x06, 0x7b, 0x79, 0x32, 0xdf, 0x64, 0xd8, 0xa4, 0xab, 0xb4, 0xfb, 0x2e, 0x81,
	0x9b, 0xbd, 0x5e, 0x29, 0xe7, 0xa0, 0xed, 0x39, 0x33, 0xd7, 0xec, 0x59, 0x03, 0xcb, 0xec, 0xcb,
	0x07, 0x4a, 0x1b, 0x1c, 0x7b, 0xce, 0xc8, 0x99, 0x2c, 0x1d, 0x59, 0x2a, 0xdd, 0xbe, 0x35, 0x73,
	0x6d, 0xfd, 0x11, 0x79, 0x53, 0x5b, 0x6e, 0x28, 0x67, 0xe0, 0x64, 0x60, 0x39, 0xba, 0x2d, 0xca,
	0xa6, 0x72, 0x09, 0xe4, 0xaa, 0x1c, 0x4f, 0x0c, 0xcb, 0x36, 0x45, 0xf7, 0x50, 0x91, 0xc1, 0xe9,
	0x7c, 0xaa, 0xf7, 0x46, 0x96, 0x33, 0x14, 0x9d, 0x23, 0xe5, 0x0a, 0x5c, 0xd4, 0x37, 0x7e, 0x19,
	0x2d, 0xe3, 0x53, 0x02, 0xd7, 0x21, 0x4b, 0xe1, 0xbf, 0x61, 0x18, 0x77, 0x7b, 0xcd, 0xe4, 0x96,
	0xf9, 0xba, 0xd2, 0x93, 0x51, 0x43, 0x63, 0x96, 0x04, 0x34, 0x86, 0x2c, 0x8f, 0xd5, 0x18, 0x53,
	0x91, 0xfe, 0x76, 0x95, 0x19, 0xe1, 0x7f, 0x6c, 0xf6, 0x41, 0x9c, 0xaf, 0x8d, 0xe6, 0x50, 0xd7,
	0xdf, 0x1a, 0x9d, 0x61, 0x85, 0xd2, 0x23, 0x0e, 0x2b, 0x59, 0xaa, 0x85, 0x06, 0xcb, 0xd4, 0xf9,
	0xc7, 0xd6, 0xf7, 0xf5, 0x88, 0xfb, 0x3b, 0xdf, 0x5f, 0x68, 0xbe, 0xf0, 0x9f, 0x5b, 0xe2, 0xd3,
	0xfb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x46, 0x2f, 0x08, 0x4d, 0x02, 0x00, 0x00,
}