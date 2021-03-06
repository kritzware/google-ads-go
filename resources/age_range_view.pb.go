// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/age_range_view.proto

package resources

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

// An age range view.
type AgeRangeView struct {
	// The resource name of the age range view.
	// Age range view resource names have the form:
	//
	// `customers/{customer_id}/ageRangeViews/{ad_group_id}_{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgeRangeView) Reset()         { *m = AgeRangeView{} }
func (m *AgeRangeView) String() string { return proto.CompactTextString(m) }
func (*AgeRangeView) ProtoMessage()    {}
func (*AgeRangeView) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571fe1afb43562a, []int{0}
}

func (m *AgeRangeView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgeRangeView.Unmarshal(m, b)
}
func (m *AgeRangeView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgeRangeView.Marshal(b, m, deterministic)
}
func (m *AgeRangeView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgeRangeView.Merge(m, src)
}
func (m *AgeRangeView) XXX_Size() int {
	return xxx_messageInfo_AgeRangeView.Size(m)
}
func (m *AgeRangeView) XXX_DiscardUnknown() {
	xxx_messageInfo_AgeRangeView.DiscardUnknown(m)
}

var xxx_messageInfo_AgeRangeView proto.InternalMessageInfo

func (m *AgeRangeView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*AgeRangeView)(nil), "google.ads.googleads.v0.resources.AgeRangeView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/age_range_view.proto", fileDescriptor_8571fe1afb43562a)
}

var fileDescriptor_8571fe1afb43562a = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xa2, 0xd4,
	0xe2, 0xfc, 0xd2, 0xa2, 0xe4, 0xd4, 0x62, 0xfd, 0xc4, 0xf4, 0xd4, 0xf8, 0xa2, 0xc4, 0xbc, 0xf4,
	0xd4, 0xf8, 0xb2, 0xcc, 0xd4, 0x72, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x45, 0x88, 0x62,
	0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x3e, 0xbd, 0x32, 0x03, 0x3d, 0xb8, 0x3e, 0x25, 0x63, 0x2e,
	0x1e, 0xc7, 0xf4, 0xd4, 0x20, 0x90, 0xce, 0xb0, 0xcc, 0xd4, 0x72, 0x21, 0x65, 0x2e, 0x5e, 0x98,
	0x64, 0x7c, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x0f, 0x4c, 0xd0,
	0x2f, 0x31, 0x37, 0xd5, 0xe9, 0x1f, 0x23, 0x97, 0x6a, 0x72, 0x7e, 0xae, 0x1e, 0x41, 0xe3, 0x9d,
	0x04, 0x91, 0x0d, 0x0f, 0x00, 0x39, 0x2a, 0x80, 0x31, 0xca, 0x0b, 0xaa, 0x2f, 0x3d, 0x3f, 0x27,
	0x31, 0x2f, 0x5d, 0x2f, 0xbf, 0x28, 0x5d, 0x3f, 0x3d, 0x35, 0x0f, 0xec, 0x64, 0x98, 0xf7, 0x0a,
	0x32, 0x8b, 0xf1, 0xf8, 0xd6, 0x1a, 0xce, 0x5a, 0xc4, 0xc4, 0xec, 0xee, 0xe8, 0xb8, 0x8a, 0x49,
	0xd1, 0x1d, 0x62, 0xa4, 0x63, 0x4a, 0xb1, 0x1e, 0x84, 0x09, 0x62, 0x85, 0x19, 0xe8, 0x05, 0xc1,
	0x54, 0x9e, 0x82, 0xa9, 0x89, 0x71, 0x4c, 0x29, 0x8e, 0x81, 0xab, 0x89, 0x09, 0x33, 0x88, 0x81,
	0xab, 0x79, 0xc5, 0xa4, 0x0a, 0x91, 0xb0, 0xb2, 0x72, 0x4c, 0x29, 0xb6, 0xb2, 0x82, 0xab, 0xb2,
	0xb2, 0x0a, 0x33, 0xb0, 0xb2, 0x82, 0xab, 0x4b, 0x62, 0x03, 0x3b, 0xd6, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x54, 0xef, 0x65, 0x6b, 0x99, 0x01, 0x00, 0x00,
}
