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
// source: google/ads/googleads/v8/services/customer_user_access_invitation_service.proto

package services

import (
	context "context"
	resources "github.com/opteo/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for
// [CustomerUserAccessInvitation.GetCustomerUserAccessInvitation][]
type GetCustomerUserAccessInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of the access invitation.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetCustomerUserAccessInvitationRequest) Reset() {
	*x = GetCustomerUserAccessInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_customer_user_access_invitation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerUserAccessInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerUserAccessInvitationRequest) ProtoMessage() {}

func (x *GetCustomerUserAccessInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_customer_user_access_invitation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerUserAccessInvitationRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerUserAccessInvitationRequest) Descriptor() ([]byte, []int) {
	return file_services_customer_user_access_invitation_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomerUserAccessInvitationRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for
// [CustomerUserAccessInvitation.MutateCustomerUserAccessInvitation][]
type MutateCustomerUserAccessInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose access invitation is being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The operation to perform on the access invitation
	Operation *CustomerUserAccessInvitationOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *MutateCustomerUserAccessInvitationRequest) Reset() {
	*x = MutateCustomerUserAccessInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_customer_user_access_invitation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerUserAccessInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerUserAccessInvitationRequest) ProtoMessage() {}

func (x *MutateCustomerUserAccessInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_customer_user_access_invitation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerUserAccessInvitationRequest.ProtoReflect.Descriptor instead.
func (*MutateCustomerUserAccessInvitationRequest) Descriptor() ([]byte, []int) {
	return file_services_customer_user_access_invitation_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateCustomerUserAccessInvitationRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCustomerUserAccessInvitationRequest) GetOperation() *CustomerUserAccessInvitationOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

// A single operation (create or remove) on customer user access invitation.
type CustomerUserAccessInvitationOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mutate operation
	//
	// Types that are assignable to Operation:
	//	*CustomerUserAccessInvitationOperation_Create
	//	*CustomerUserAccessInvitationOperation_Remove
	Operation isCustomerUserAccessInvitationOperation_Operation `protobuf_oneof:"operation"`
}

func (x *CustomerUserAccessInvitationOperation) Reset() {
	*x = CustomerUserAccessInvitationOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_customer_user_access_invitation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerUserAccessInvitationOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerUserAccessInvitationOperation) ProtoMessage() {}

func (x *CustomerUserAccessInvitationOperation) ProtoReflect() protoreflect.Message {
	mi := &file_services_customer_user_access_invitation_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerUserAccessInvitationOperation.ProtoReflect.Descriptor instead.
func (*CustomerUserAccessInvitationOperation) Descriptor() ([]byte, []int) {
	return file_services_customer_user_access_invitation_service_proto_rawDescGZIP(), []int{2}
}

func (m *CustomerUserAccessInvitationOperation) GetOperation() isCustomerUserAccessInvitationOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *CustomerUserAccessInvitationOperation) GetCreate() *resources.CustomerUserAccessInvitation {
	if x, ok := x.GetOperation().(*CustomerUserAccessInvitationOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *CustomerUserAccessInvitationOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*CustomerUserAccessInvitationOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isCustomerUserAccessInvitationOperation_Operation interface {
	isCustomerUserAccessInvitationOperation_Operation()
}

type CustomerUserAccessInvitationOperation_Create struct {
	// Create operation: No resource name is expected for the new access
	// invitation.
	Create *resources.CustomerUserAccessInvitation `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerUserAccessInvitationOperation_Remove struct {
	// Remove operation: A resource name for the revoke invitation is
	// expected, in this format:
	//
	// `customers/{customer_id}/customerUserAccessInvitations/{invitation_id}`
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*CustomerUserAccessInvitationOperation_Create) isCustomerUserAccessInvitationOperation_Operation() {
}

func (*CustomerUserAccessInvitationOperation_Remove) isCustomerUserAccessInvitationOperation_Operation() {
}

// Response message for access invitation mutate.
type MutateCustomerUserAccessInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result for the mutate.
	Result *MutateCustomerUserAccessInvitationResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MutateCustomerUserAccessInvitationResponse) Reset() {
	*x = MutateCustomerUserAccessInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_customer_user_access_invitation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerUserAccessInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerUserAccessInvitationResponse) ProtoMessage() {}

func (x *MutateCustomerUserAccessInvitationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_customer_user_access_invitation_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerUserAccessInvitationResponse.ProtoReflect.Descriptor instead.
func (*MutateCustomerUserAccessInvitationResponse) Descriptor() ([]byte, []int) {
	return file_services_customer_user_access_invitation_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCustomerUserAccessInvitationResponse) GetResult() *MutateCustomerUserAccessInvitationResult {
	if x != nil {
		return x.Result
	}
	return nil
}

// The result for the access invitation mutate.
type MutateCustomerUserAccessInvitationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateCustomerUserAccessInvitationResult) Reset() {
	*x = MutateCustomerUserAccessInvitationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_customer_user_access_invitation_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerUserAccessInvitationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerUserAccessInvitationResult) ProtoMessage() {}

func (x *MutateCustomerUserAccessInvitationResult) ProtoReflect() protoreflect.Message {
	mi := &file_services_customer_user_access_invitation_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerUserAccessInvitationResult.ProtoReflect.Descriptor instead.
func (*MutateCustomerUserAccessInvitationResult) Descriptor() ([]byte, []int) {
	return file_services_customer_user_access_invitation_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateCustomerUserAccessInvitationResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_services_customer_user_access_invitation_service_proto protoreflect.FileDescriptor

var file_services_customer_user_access_invitation_service_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c,
	0x01, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x62, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x3d, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x37, 0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xbd, 0x01,
	0x0a, 0x29, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x6a, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa9, 0x01,
	0x0a, 0x25, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x90, 0x01, 0x0a, 0x2a, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x4f, 0x0a, 0x28,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x9d, 0x05,
	0x0a, 0x23, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x02, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x57, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x12, 0x3f, 0x2f, 0x76,
	0x38, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xa6, 0x02,
	0x0a, 0x22, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x47, 0x22, 0x42, 0x2f, 0x76, 0x38, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41,
	0x15, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x8f, 0x02,
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x38, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x28, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x38, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x38, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x38, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x38, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_customer_user_access_invitation_service_proto_rawDescOnce sync.Once
	file_services_customer_user_access_invitation_service_proto_rawDescData = file_services_customer_user_access_invitation_service_proto_rawDesc
)

func file_services_customer_user_access_invitation_service_proto_rawDescGZIP() []byte {
	file_services_customer_user_access_invitation_service_proto_rawDescOnce.Do(func() {
		file_services_customer_user_access_invitation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_customer_user_access_invitation_service_proto_rawDescData)
	})
	return file_services_customer_user_access_invitation_service_proto_rawDescData
}

var file_services_customer_user_access_invitation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_services_customer_user_access_invitation_service_proto_goTypes = []interface{}{
	(*GetCustomerUserAccessInvitationRequest)(nil),     // 0: google.ads.googleads.v8.services.GetCustomerUserAccessInvitationRequest
	(*MutateCustomerUserAccessInvitationRequest)(nil),  // 1: google.ads.googleads.v8.services.MutateCustomerUserAccessInvitationRequest
	(*CustomerUserAccessInvitationOperation)(nil),      // 2: google.ads.googleads.v8.services.CustomerUserAccessInvitationOperation
	(*MutateCustomerUserAccessInvitationResponse)(nil), // 3: google.ads.googleads.v8.services.MutateCustomerUserAccessInvitationResponse
	(*MutateCustomerUserAccessInvitationResult)(nil),   // 4: google.ads.googleads.v8.services.MutateCustomerUserAccessInvitationResult
	(*resources.CustomerUserAccessInvitation)(nil),     // 5: google.ads.googleads.v8.resources.CustomerUserAccessInvitation
}
var file_services_customer_user_access_invitation_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v8.services.MutateCustomerUserAccessInvitationRequest.operation:type_name -> google.ads.googleads.v8.services.CustomerUserAccessInvitationOperation
	5, // 1: google.ads.googleads.v8.services.CustomerUserAccessInvitationOperation.create:type_name -> google.ads.googleads.v8.resources.CustomerUserAccessInvitation
	4, // 2: google.ads.googleads.v8.services.MutateCustomerUserAccessInvitationResponse.result:type_name -> google.ads.googleads.v8.services.MutateCustomerUserAccessInvitationResult
	0, // 3: google.ads.googleads.v8.services.CustomerUserAccessInvitationService.GetCustomerUserAccessInvitation:input_type -> google.ads.googleads.v8.services.GetCustomerUserAccessInvitationRequest
	1, // 4: google.ads.googleads.v8.services.CustomerUserAccessInvitationService.MutateCustomerUserAccessInvitation:input_type -> google.ads.googleads.v8.services.MutateCustomerUserAccessInvitationRequest
	5, // 5: google.ads.googleads.v8.services.CustomerUserAccessInvitationService.GetCustomerUserAccessInvitation:output_type -> google.ads.googleads.v8.resources.CustomerUserAccessInvitation
	3, // 6: google.ads.googleads.v8.services.CustomerUserAccessInvitationService.MutateCustomerUserAccessInvitation:output_type -> google.ads.googleads.v8.services.MutateCustomerUserAccessInvitationResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_services_customer_user_access_invitation_service_proto_init()
}
func file_services_customer_user_access_invitation_service_proto_init() {
	if File_services_customer_user_access_invitation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_customer_user_access_invitation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerUserAccessInvitationRequest); i {
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
		file_services_customer_user_access_invitation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerUserAccessInvitationRequest); i {
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
		file_services_customer_user_access_invitation_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerUserAccessInvitationOperation); i {
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
		file_services_customer_user_access_invitation_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerUserAccessInvitationResponse); i {
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
		file_services_customer_user_access_invitation_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerUserAccessInvitationResult); i {
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
	file_services_customer_user_access_invitation_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CustomerUserAccessInvitationOperation_Create)(nil),
		(*CustomerUserAccessInvitationOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_customer_user_access_invitation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_customer_user_access_invitation_service_proto_goTypes,
		DependencyIndexes: file_services_customer_user_access_invitation_service_proto_depIdxs,
		MessageInfos:      file_services_customer_user_access_invitation_service_proto_msgTypes,
	}.Build()
	File_services_customer_user_access_invitation_service_proto = out.File
	file_services_customer_user_access_invitation_service_proto_rawDesc = nil
	file_services_customer_user_access_invitation_service_proto_goTypes = nil
	file_services_customer_user_access_invitation_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerUserAccessInvitationServiceClient is the client API for CustomerUserAccessInvitationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerUserAccessInvitationServiceClient interface {
	// Returns the requested access invitation in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomerUserAccessInvitation(ctx context.Context, in *GetCustomerUserAccessInvitationRequest, opts ...grpc.CallOption) (*resources.CustomerUserAccessInvitation, error)
	// Creates or removes an access invitation.
	//
	// List of thrown errors:
	//   [AccessInvitationError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerUserAccessInvitation(ctx context.Context, in *MutateCustomerUserAccessInvitationRequest, opts ...grpc.CallOption) (*MutateCustomerUserAccessInvitationResponse, error)
}

type customerUserAccessInvitationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerUserAccessInvitationServiceClient(cc grpc.ClientConnInterface) CustomerUserAccessInvitationServiceClient {
	return &customerUserAccessInvitationServiceClient{cc}
}

func (c *customerUserAccessInvitationServiceClient) GetCustomerUserAccessInvitation(ctx context.Context, in *GetCustomerUserAccessInvitationRequest, opts ...grpc.CallOption) (*resources.CustomerUserAccessInvitation, error) {
	out := new(resources.CustomerUserAccessInvitation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.CustomerUserAccessInvitationService/GetCustomerUserAccessInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerUserAccessInvitationServiceClient) MutateCustomerUserAccessInvitation(ctx context.Context, in *MutateCustomerUserAccessInvitationRequest, opts ...grpc.CallOption) (*MutateCustomerUserAccessInvitationResponse, error) {
	out := new(MutateCustomerUserAccessInvitationResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v8.services.CustomerUserAccessInvitationService/MutateCustomerUserAccessInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerUserAccessInvitationServiceServer is the server API for CustomerUserAccessInvitationService service.
type CustomerUserAccessInvitationServiceServer interface {
	// Returns the requested access invitation in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomerUserAccessInvitation(context.Context, *GetCustomerUserAccessInvitationRequest) (*resources.CustomerUserAccessInvitation, error)
	// Creates or removes an access invitation.
	//
	// List of thrown errors:
	//   [AccessInvitationError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerUserAccessInvitation(context.Context, *MutateCustomerUserAccessInvitationRequest) (*MutateCustomerUserAccessInvitationResponse, error)
}

// UnimplementedCustomerUserAccessInvitationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerUserAccessInvitationServiceServer struct {
}

func (*UnimplementedCustomerUserAccessInvitationServiceServer) GetCustomerUserAccessInvitation(context.Context, *GetCustomerUserAccessInvitationRequest) (*resources.CustomerUserAccessInvitation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerUserAccessInvitation not implemented")
}
func (*UnimplementedCustomerUserAccessInvitationServiceServer) MutateCustomerUserAccessInvitation(context.Context, *MutateCustomerUserAccessInvitationRequest) (*MutateCustomerUserAccessInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerUserAccessInvitation not implemented")
}

func RegisterCustomerUserAccessInvitationServiceServer(s *grpc.Server, srv CustomerUserAccessInvitationServiceServer) {
	s.RegisterService(&_CustomerUserAccessInvitationService_serviceDesc, srv)
}

func _CustomerUserAccessInvitationService_GetCustomerUserAccessInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerUserAccessInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerUserAccessInvitationServiceServer).GetCustomerUserAccessInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.CustomerUserAccessInvitationService/GetCustomerUserAccessInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerUserAccessInvitationServiceServer).GetCustomerUserAccessInvitation(ctx, req.(*GetCustomerUserAccessInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerUserAccessInvitationService_MutateCustomerUserAccessInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerUserAccessInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerUserAccessInvitationServiceServer).MutateCustomerUserAccessInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v8.services.CustomerUserAccessInvitationService/MutateCustomerUserAccessInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerUserAccessInvitationServiceServer).MutateCustomerUserAccessInvitation(ctx, req.(*MutateCustomerUserAccessInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerUserAccessInvitationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v8.services.CustomerUserAccessInvitationService",
	HandlerType: (*CustomerUserAccessInvitationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerUserAccessInvitation",
			Handler:    _CustomerUserAccessInvitationService_GetCustomerUserAccessInvitation_Handler,
		},
		{
			MethodName: "MutateCustomerUserAccessInvitation",
			Handler:    _CustomerUserAccessInvitationService_MutateCustomerUserAccessInvitation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v8/services/customer_user_access_invitation_service.proto",
}
