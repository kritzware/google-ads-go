// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/vanity_pharma_text.proto

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

// Enum describing possible text.
type VanityPharmaTextEnum_VanityPharmaText int32

const (
	// Not specified.
	VanityPharmaTextEnum_UNSPECIFIED VanityPharmaTextEnum_VanityPharmaText = 0
	// Used for return value only. Represents value unknown in this version.
	VanityPharmaTextEnum_UNKNOWN VanityPharmaTextEnum_VanityPharmaText = 1
	// Prescription treatment website with website content in English.
	VanityPharmaTextEnum_PRESCRIPTION_TREATMENT_WEBSITE_EN VanityPharmaTextEnum_VanityPharmaText = 2
	// Prescription treatment website with website content in Spanish
	// (Sitio de tratamientos con receta).
	VanityPharmaTextEnum_PRESCRIPTION_TREATMENT_WEBSITE_ES VanityPharmaTextEnum_VanityPharmaText = 3
	// Prescription device website with website content in English.
	VanityPharmaTextEnum_PRESCRIPTION_DEVICE_WEBSITE_EN VanityPharmaTextEnum_VanityPharmaText = 4
	// Prescription device website with website content in Spanish (Sitio de
	// dispositivos con receta).
	VanityPharmaTextEnum_PRESCRIPTION_DEVICE_WEBSITE_ES VanityPharmaTextEnum_VanityPharmaText = 5
	// Medical device website with website content in English.
	VanityPharmaTextEnum_MEDICAL_DEVICE_WEBSITE_EN VanityPharmaTextEnum_VanityPharmaText = 6
	// Medical device website with website content in Spanish (Sitio de
	// dispositivos médicos).
	VanityPharmaTextEnum_MEDICAL_DEVICE_WEBSITE_ES VanityPharmaTextEnum_VanityPharmaText = 7
	// Preventative treatment website with website content in English.
	VanityPharmaTextEnum_PREVENTATIVE_TREATMENT_WEBSITE_EN VanityPharmaTextEnum_VanityPharmaText = 8
	// Preventative treatment website with website content in Spanish (Sitio de
	// tratamientos preventivos).
	VanityPharmaTextEnum_PREVENTATIVE_TREATMENT_WEBSITE_ES VanityPharmaTextEnum_VanityPharmaText = 9
	// Prescription contraception website with website content in English.
	VanityPharmaTextEnum_PRESCRIPTION_CONTRACEPTION_WEBSITE_EN VanityPharmaTextEnum_VanityPharmaText = 10
	// Prescription contraception website with website content in Spanish (Sitio
	// de anticonceptivos con receta).
	VanityPharmaTextEnum_PRESCRIPTION_CONTRACEPTION_WEBSITE_ES VanityPharmaTextEnum_VanityPharmaText = 11
	// Prescription vaccine website with website content in English.
	VanityPharmaTextEnum_PRESCRIPTION_VACCINE_WEBSITE_EN VanityPharmaTextEnum_VanityPharmaText = 12
	// Prescription vaccine website with website content in Spanish (Sitio de
	// vacunas con receta).
	VanityPharmaTextEnum_PRESCRIPTION_VACCINE_WEBSITE_ES VanityPharmaTextEnum_VanityPharmaText = 13
)

var VanityPharmaTextEnum_VanityPharmaText_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "PRESCRIPTION_TREATMENT_WEBSITE_EN",
	3:  "PRESCRIPTION_TREATMENT_WEBSITE_ES",
	4:  "PRESCRIPTION_DEVICE_WEBSITE_EN",
	5:  "PRESCRIPTION_DEVICE_WEBSITE_ES",
	6:  "MEDICAL_DEVICE_WEBSITE_EN",
	7:  "MEDICAL_DEVICE_WEBSITE_ES",
	8:  "PREVENTATIVE_TREATMENT_WEBSITE_EN",
	9:  "PREVENTATIVE_TREATMENT_WEBSITE_ES",
	10: "PRESCRIPTION_CONTRACEPTION_WEBSITE_EN",
	11: "PRESCRIPTION_CONTRACEPTION_WEBSITE_ES",
	12: "PRESCRIPTION_VACCINE_WEBSITE_EN",
	13: "PRESCRIPTION_VACCINE_WEBSITE_ES",
}

var VanityPharmaTextEnum_VanityPharmaText_value = map[string]int32{
	"UNSPECIFIED":                           0,
	"UNKNOWN":                               1,
	"PRESCRIPTION_TREATMENT_WEBSITE_EN":     2,
	"PRESCRIPTION_TREATMENT_WEBSITE_ES":     3,
	"PRESCRIPTION_DEVICE_WEBSITE_EN":        4,
	"PRESCRIPTION_DEVICE_WEBSITE_ES":        5,
	"MEDICAL_DEVICE_WEBSITE_EN":             6,
	"MEDICAL_DEVICE_WEBSITE_ES":             7,
	"PREVENTATIVE_TREATMENT_WEBSITE_EN":     8,
	"PREVENTATIVE_TREATMENT_WEBSITE_ES":     9,
	"PRESCRIPTION_CONTRACEPTION_WEBSITE_EN": 10,
	"PRESCRIPTION_CONTRACEPTION_WEBSITE_ES": 11,
	"PRESCRIPTION_VACCINE_WEBSITE_EN":       12,
	"PRESCRIPTION_VACCINE_WEBSITE_ES":       13,
}

func (x VanityPharmaTextEnum_VanityPharmaText) String() string {
	return proto.EnumName(VanityPharmaTextEnum_VanityPharmaText_name, int32(x))
}

func (VanityPharmaTextEnum_VanityPharmaText) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d18ebee01472e81c, []int{0, 0}
}

// The text that will be displayed in display URL of the text ad when website
// description is the selected display mode for vanity pharma URLs.
type VanityPharmaTextEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VanityPharmaTextEnum) Reset()         { *m = VanityPharmaTextEnum{} }
func (m *VanityPharmaTextEnum) String() string { return proto.CompactTextString(m) }
func (*VanityPharmaTextEnum) ProtoMessage()    {}
func (*VanityPharmaTextEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d18ebee01472e81c, []int{0}
}

func (m *VanityPharmaTextEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VanityPharmaTextEnum.Unmarshal(m, b)
}
func (m *VanityPharmaTextEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VanityPharmaTextEnum.Marshal(b, m, deterministic)
}
func (m *VanityPharmaTextEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VanityPharmaTextEnum.Merge(m, src)
}
func (m *VanityPharmaTextEnum) XXX_Size() int {
	return xxx_messageInfo_VanityPharmaTextEnum.Size(m)
}
func (m *VanityPharmaTextEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_VanityPharmaTextEnum.DiscardUnknown(m)
}

var xxx_messageInfo_VanityPharmaTextEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.VanityPharmaTextEnum_VanityPharmaText", VanityPharmaTextEnum_VanityPharmaText_name, VanityPharmaTextEnum_VanityPharmaText_value)
	proto.RegisterType((*VanityPharmaTextEnum)(nil), "google.ads.googleads.v0.enums.VanityPharmaTextEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/vanity_pharma_text.proto", fileDescriptor_d18ebee01472e81c)
}

var fileDescriptor_d18ebee01472e81c = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6a, 0xd4, 0x40,
	0x1c, 0xc6, 0xdd, 0xdd, 0xda, 0xea, 0xac, 0xe2, 0x30, 0xe8, 0xc1, 0xc3, 0x2a, 0xbb, 0xd2, 0x83,
	0x97, 0x49, 0x40, 0xf0, 0x30, 0x9e, 0x26, 0xb3, 0x63, 0x19, 0xb4, 0xb3, 0x21, 0x33, 0x9d, 0x82,
	0x04, 0x42, 0x34, 0x21, 0x16, 0x9a, 0x64, 0xd9, 0x49, 0x97, 0xfa, 0x0c, 0xbe, 0x85, 0x47, 0x1f,
	0xc5, 0x47, 0xe9, 0xd1, 0x27, 0x90, 0x4c, 0xda, 0xa5, 0x29, 0xb5, 0xdd, 0x4b, 0xf8, 0x67, 0xbe,
	0xdf, 0xf7, 0xf1, 0xff, 0xc3, 0x07, 0xde, 0x17, 0x75, 0x5d, 0x9c, 0xe6, 0x5e, 0x9a, 0x59, 0xaf,
	0x1b, 0xdb, 0x69, 0xed, 0x7b, 0x79, 0x75, 0x56, 0x5a, 0x6f, 0x9d, 0x56, 0x27, 0xcd, 0x8f, 0x64,
	0xf9, 0x3d, 0x5d, 0x95, 0x69, 0xd2, 0xe4, 0xe7, 0x0d, 0x5e, 0xae, 0xea, 0xa6, 0x46, 0x93, 0x0e,
	0xc6, 0x69, 0x66, 0xf1, 0xc6, 0x87, 0xd7, 0x3e, 0x76, 0xbe, 0xd9, 0xcf, 0x1d, 0xf0, 0xdc, 0x38,
	0x6f, 0xe8, 0xac, 0x3a, 0x3f, 0x6f, 0x78, 0x75, 0x56, 0xce, 0xfe, 0x8e, 0x00, 0xbc, 0x29, 0xa0,
	0x67, 0x60, 0x7c, 0x24, 0x55, 0xc8, 0x99, 0xf8, 0x28, 0xf8, 0x1c, 0x3e, 0x40, 0x63, 0xb0, 0x77,
	0x24, 0x3f, 0xc9, 0xc5, 0xb1, 0x84, 0x03, 0xb4, 0x0f, 0xa6, 0x61, 0xc4, 0x15, 0x8b, 0x44, 0xa8,
	0xc5, 0x42, 0x26, 0x3a, 0xe2, 0x54, 0x1f, 0x72, 0xa9, 0x93, 0x63, 0x1e, 0x28, 0xa1, 0x79, 0xc2,
	0x25, 0x1c, 0x6e, 0x83, 0x29, 0x38, 0x42, 0x33, 0xf0, 0xaa, 0x87, 0xcd, 0xb9, 0x11, 0x8c, 0x5f,
	0x8f, 0xda, 0xb9, 0x97, 0x51, 0xf0, 0x21, 0x9a, 0x80, 0x97, 0x87, 0x7c, 0x2e, 0x18, 0xfd, 0x7c,
	0x4b, 0xc4, 0xee, 0x5d, 0xb2, 0x82, 0x7b, 0x97, 0xcb, 0x1a, 0x2e, 0x35, 0xd5, 0xc2, 0xf0, 0xdb,
	0x6f, 0x7a, 0xb4, 0x0d, 0xa6, 0xe0, 0x63, 0xf4, 0x16, 0xec, 0xf7, 0xf6, 0x65, 0x0b, 0xa9, 0x23,
	0xca, 0x78, 0xf7, 0x77, 0x2d, 0x11, 0x6c, 0x8b, 0x2a, 0x38, 0x46, 0x6f, 0xc0, 0xeb, 0x1e, 0x6a,
	0x28, 0x63, 0x42, 0xf6, 0xee, 0x7c, 0x72, 0x3f, 0xa4, 0xe0, 0xd3, 0xe0, 0x62, 0x00, 0xa6, 0xdf,
	0xea, 0x12, 0xdf, 0xd9, 0x99, 0xe0, 0xc5, 0xcd, 0x5e, 0x84, 0x6d, 0xd3, 0xc2, 0xc1, 0x97, 0xe0,
	0xd2, 0x57, 0xd4, 0xa7, 0x69, 0x55, 0xe0, 0x7a, 0x55, 0x78, 0x45, 0x5e, 0xb9, 0x1e, 0x5e, 0x75,
	0x76, 0x79, 0x62, 0xff, 0x53, 0xe1, 0x0f, 0xee, 0xfb, 0x6b, 0x38, 0x3a, 0xa0, 0xf4, 0xf7, 0x70,
	0x72, 0xd0, 0x45, 0xd1, 0xcc, 0xe2, 0x6e, 0x6c, 0x27, 0xe3, 0xe3, 0xb6, 0x9c, 0xf6, 0xcf, 0x95,
	0x1e, 0xd3, 0xcc, 0xc6, 0x1b, 0x3d, 0x36, 0x7e, 0xec, 0xf4, 0x8b, 0xe1, 0xb4, 0x7b, 0x24, 0x84,
	0x66, 0x96, 0x90, 0x0d, 0x41, 0x88, 0xf1, 0x09, 0x71, 0xcc, 0xd7, 0x5d, 0xb7, 0xd8, 0xbb, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x9d, 0xc1, 0xdd, 0x5a, 0x03, 0x00, 0x00,
}
