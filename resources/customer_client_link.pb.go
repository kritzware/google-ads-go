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
// source: google/ads/googleads/v8/resources/customer_client_link.proto

package resources

import (
	enums "github.com/opteo/google-ads-go/enums"
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

// Represents customer client link relationship.
type CustomerClientLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Name of the resource.
	// CustomerClientLink resource names have the form:
	// `customers/{customer_id}/customerClientLinks/{client_customer_id}~{manager_link_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The client customer linked to this customer.
	ClientCustomer *string `protobuf:"bytes,7,opt,name=client_customer,json=clientCustomer,proto3,oneof" json:"client_customer,omitempty"`
	// Output only. This is uniquely identifies a customer client link. Read only.
	ManagerLinkId *int64 `protobuf:"varint,8,opt,name=manager_link_id,json=managerLinkId,proto3,oneof" json:"manager_link_id,omitempty"`
	// This is the status of the link between client and manager.
	Status enums.ManagerLinkStatusEnum_ManagerLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v8.enums.ManagerLinkStatusEnum_ManagerLinkStatus" json:"status,omitempty"`
	// The visibility of the link. Users can choose whether or not to see hidden
	// links in the Google Ads UI.
	// Default value is false
	Hidden *bool `protobuf:"varint,9,opt,name=hidden,proto3,oneof" json:"hidden,omitempty"`
}

func (x *CustomerClientLink) Reset() {
	*x = CustomerClientLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_customer_client_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerClientLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerClientLink) ProtoMessage() {}

func (x *CustomerClientLink) ProtoReflect() protoreflect.Message {
	mi := &file_resources_customer_client_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerClientLink.ProtoReflect.Descriptor instead.
func (*CustomerClientLink) Descriptor() ([]byte, []int) {
	return file_resources_customer_client_link_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerClientLink) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CustomerClientLink) GetClientCustomer() string {
	if x != nil && x.ClientCustomer != nil {
		return *x.ClientCustomer
	}
	return ""
}

func (x *CustomerClientLink) GetManagerLinkId() int64 {
	if x != nil && x.ManagerLinkId != nil {
		return *x.ManagerLinkId
	}
	return 0
}

func (x *CustomerClientLink) GetStatus() enums.ManagerLinkStatusEnum_ManagerLinkStatus {
	if x != nil {
		return x.Status
	}
	return enums.ManagerLinkStatusEnum_UNSPECIFIED
}

func (x *CustomerClientLink) GetHidden() bool {
	if x != nil && x.Hidden != nil {
		return *x.Hidden
	}
	return false
}

var File_resources_customer_client_link_proto protoreflect.FileDescriptor

var file_resources_customer_client_link_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x04, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x58, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x57, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29,
	0xe0, 0x41, 0x05, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x30,
	0x0a, 0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x0d,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x5e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69,
	0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1b, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x02, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x3a, 0x85, 0x01,
	0xea, 0x41, 0x81, 0x01, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x52, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x7b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x42, 0x84, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x42, 0x17, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_customer_client_link_proto_rawDescOnce sync.Once
	file_resources_customer_client_link_proto_rawDescData = file_resources_customer_client_link_proto_rawDesc
)

func file_resources_customer_client_link_proto_rawDescGZIP() []byte {
	file_resources_customer_client_link_proto_rawDescOnce.Do(func() {
		file_resources_customer_client_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_customer_client_link_proto_rawDescData)
	})
	return file_resources_customer_client_link_proto_rawDescData
}

var file_resources_customer_client_link_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_customer_client_link_proto_goTypes = []interface{}{
	(*CustomerClientLink)(nil),                         // 0: google.ads.googleads.v8.resources.CustomerClientLink
	(enums.ManagerLinkStatusEnum_ManagerLinkStatus)(0), // 1: google.ads.googleads.v8.enums.ManagerLinkStatusEnum.ManagerLinkStatus
}
var file_resources_customer_client_link_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v8.resources.CustomerClientLink.status:type_name -> google.ads.googleads.v8.enums.ManagerLinkStatusEnum.ManagerLinkStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_resources_customer_client_link_proto_init() }
func file_resources_customer_client_link_proto_init() {
	if File_resources_customer_client_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_customer_client_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerClientLink); i {
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
	file_resources_customer_client_link_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_customer_client_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_customer_client_link_proto_goTypes,
		DependencyIndexes: file_resources_customer_client_link_proto_depIdxs,
		MessageInfos:      file_resources_customer_client_link_proto_msgTypes,
	}.Build()
	File_resources_customer_client_link_proto = out.File
	file_resources_customer_client_link_proto_rawDesc = nil
	file_resources_customer_client_link_proto_goTypes = nil
	file_resources_customer_client_link_proto_depIdxs = nil
}