// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.2
// source: google/ads/googleads/v8/enums/policy_review_status.proto

package enums

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The possible policy review statuses.
type PolicyReviewStatusEnum_PolicyReviewStatus int32

const (
	// No value has been specified.
	PolicyReviewStatusEnum_UNSPECIFIED PolicyReviewStatusEnum_PolicyReviewStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	PolicyReviewStatusEnum_UNKNOWN PolicyReviewStatusEnum_PolicyReviewStatus = 1
	// Currently under review.
	PolicyReviewStatusEnum_REVIEW_IN_PROGRESS PolicyReviewStatusEnum_PolicyReviewStatus = 2
	// Primary review complete. Other reviews may be continuing.
	PolicyReviewStatusEnum_REVIEWED PolicyReviewStatusEnum_PolicyReviewStatus = 3
	// The resource has been resubmitted for approval or its policy decision has
	// been appealed.
	PolicyReviewStatusEnum_UNDER_APPEAL PolicyReviewStatusEnum_PolicyReviewStatus = 4
	// The resource is eligible and may be serving but could still undergo
	// further review.
	PolicyReviewStatusEnum_ELIGIBLE_MAY_SERVE PolicyReviewStatusEnum_PolicyReviewStatus = 5
)

// Enum value maps for PolicyReviewStatusEnum_PolicyReviewStatus.
var (
	PolicyReviewStatusEnum_PolicyReviewStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "REVIEW_IN_PROGRESS",
		3: "REVIEWED",
		4: "UNDER_APPEAL",
		5: "ELIGIBLE_MAY_SERVE",
	}
	PolicyReviewStatusEnum_PolicyReviewStatus_value = map[string]int32{
		"UNSPECIFIED":        0,
		"UNKNOWN":            1,
		"REVIEW_IN_PROGRESS": 2,
		"REVIEWED":           3,
		"UNDER_APPEAL":       4,
		"ELIGIBLE_MAY_SERVE": 5,
	}
)

func (x PolicyReviewStatusEnum_PolicyReviewStatus) Enum() *PolicyReviewStatusEnum_PolicyReviewStatus {
	p := new(PolicyReviewStatusEnum_PolicyReviewStatus)
	*p = x
	return p
}

func (x PolicyReviewStatusEnum_PolicyReviewStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PolicyReviewStatusEnum_PolicyReviewStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_policy_review_status_proto_enumTypes[0].Descriptor()
}

func (PolicyReviewStatusEnum_PolicyReviewStatus) Type() protoreflect.EnumType {
	return &file_enums_policy_review_status_proto_enumTypes[0]
}

func (x PolicyReviewStatusEnum_PolicyReviewStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PolicyReviewStatusEnum_PolicyReviewStatus.Descriptor instead.
func (PolicyReviewStatusEnum_PolicyReviewStatus) EnumDescriptor() ([]byte, []int) {
	return file_enums_policy_review_status_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible policy review statuses.
type PolicyReviewStatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PolicyReviewStatusEnum) Reset() {
	*x = PolicyReviewStatusEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_policy_review_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyReviewStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyReviewStatusEnum) ProtoMessage() {}

func (x *PolicyReviewStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_policy_review_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyReviewStatusEnum.ProtoReflect.Descriptor instead.
func (*PolicyReviewStatusEnum) Descriptor() ([]byte, []int) {
	return file_enums_policy_review_status_proto_rawDescGZIP(), []int{0}
}

var File_enums_policy_review_status_proto protoreflect.FileDescriptor

var file_enums_policy_review_status_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x16, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0x82, 0x01, 0x0a, 0x12, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x56, 0x49, 0x45,
	0x57, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x45, 0x44, 0x10, 0x03, 0x12, 0x10, 0x0a,
	0x0c, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x50, 0x50, 0x45, 0x41, 0x4c, 0x10, 0x04, 0x12,
	0x16, 0x0a, 0x12, 0x45, 0x4c, 0x49, 0x47, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x4d, 0x41, 0x59, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x10, 0x05, 0x42, 0xec, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x17, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x38, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_policy_review_status_proto_rawDescOnce sync.Once
	file_enums_policy_review_status_proto_rawDescData = file_enums_policy_review_status_proto_rawDesc
)

func file_enums_policy_review_status_proto_rawDescGZIP() []byte {
	file_enums_policy_review_status_proto_rawDescOnce.Do(func() {
		file_enums_policy_review_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_policy_review_status_proto_rawDescData)
	})
	return file_enums_policy_review_status_proto_rawDescData
}

var file_enums_policy_review_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_policy_review_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enums_policy_review_status_proto_goTypes = []interface{}{
	(PolicyReviewStatusEnum_PolicyReviewStatus)(0), // 0: google.ads.googleads.v8.enums.PolicyReviewStatusEnum.PolicyReviewStatus
	(*PolicyReviewStatusEnum)(nil),                 // 1: google.ads.googleads.v8.enums.PolicyReviewStatusEnum
}
var file_enums_policy_review_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_policy_review_status_proto_init() }
func file_enums_policy_review_status_proto_init() {
	if File_enums_policy_review_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_policy_review_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyReviewStatusEnum); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_policy_review_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_policy_review_status_proto_goTypes,
		DependencyIndexes: file_enums_policy_review_status_proto_depIdxs,
		EnumInfos:         file_enums_policy_review_status_proto_enumTypes,
		MessageInfos:      file_enums_policy_review_status_proto_msgTypes,
	}.Build()
	File_enums_policy_review_status_proto = out.File
	file_enums_policy_review_status_proto_rawDesc = nil
	file_enums_policy_review_status_proto_goTypes = nil
	file_enums_policy_review_status_proto_depIdxs = nil
}
