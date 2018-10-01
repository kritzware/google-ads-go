// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/services/ad_group_ad_service.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	resources "github.com/kritzware/google-ads-go/resources"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [AdGroupAdService.GetAdGroupAd][].
type GetAdGroupAdRequest struct {
	// The resource name of the ad to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupAdRequest) Reset()         { *m = GetAdGroupAdRequest{} }
func (m *GetAdGroupAdRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupAdRequest) ProtoMessage()    {}
func (*GetAdGroupAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0977399cace737b, []int{0}
}
func (m *GetAdGroupAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupAdRequest.Unmarshal(m, b)
}
func (m *GetAdGroupAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupAdRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupAdRequest.Merge(m, src)
}
func (m *GetAdGroupAdRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupAdRequest.Size(m)
}
func (m *GetAdGroupAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupAdRequest proto.InternalMessageInfo

func (m *GetAdGroupAdRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupAdService.MutateAdGroupAds][].
type MutateAdGroupAdsRequest struct {
	// The ID of the customer whose ads are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual ads.
	Operations           []*AdGroupAdOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MutateAdGroupAdsRequest) Reset()         { *m = MutateAdGroupAdsRequest{} }
func (m *MutateAdGroupAdsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdsRequest) ProtoMessage()    {}
func (*MutateAdGroupAdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0977399cace737b, []int{1}
}
func (m *MutateAdGroupAdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupAdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdsRequest.Merge(m, src)
}
func (m *MutateAdGroupAdsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdsRequest.Size(m)
}
func (m *MutateAdGroupAdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdsRequest proto.InternalMessageInfo

func (m *MutateAdGroupAdsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupAdsRequest) GetOperations() []*AdGroupAdOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A single operation (create, update, remove) on an ad group ad.
type AdGroupAdOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupAdOperation_Create
	//	*AdGroupAdOperation_Update
	//	*AdGroupAdOperation_Remove
	Operation            isAdGroupAdOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AdGroupAdOperation) Reset()         { *m = AdGroupAdOperation{} }
func (m *AdGroupAdOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdOperation) ProtoMessage()    {}
func (*AdGroupAdOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0977399cace737b, []int{2}
}
func (m *AdGroupAdOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdOperation.Unmarshal(m, b)
}
func (m *AdGroupAdOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupAdOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdOperation.Merge(m, src)
}
func (m *AdGroupAdOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdOperation.Size(m)
}
func (m *AdGroupAdOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdOperation proto.InternalMessageInfo

func (m *AdGroupAdOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isAdGroupAdOperation_Operation interface {
	isAdGroupAdOperation_Operation()
}

type AdGroupAdOperation_Create struct {
	Create *resources.AdGroupAd `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupAdOperation_Update struct {
	Update *resources.AdGroupAd `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupAdOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupAdOperation_Create) isAdGroupAdOperation_Operation() {}

func (*AdGroupAdOperation_Update) isAdGroupAdOperation_Operation() {}

func (*AdGroupAdOperation_Remove) isAdGroupAdOperation_Operation() {}

func (m *AdGroupAdOperation) GetOperation() isAdGroupAdOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupAdOperation) GetCreate() *resources.AdGroupAd {
	if x, ok := m.GetOperation().(*AdGroupAdOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupAdOperation) GetUpdate() *resources.AdGroupAd {
	if x, ok := m.GetOperation().(*AdGroupAdOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdGroupAdOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupAdOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AdGroupAdOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AdGroupAdOperation_OneofMarshaler, _AdGroupAdOperation_OneofUnmarshaler, _AdGroupAdOperation_OneofSizer, []interface{}{
		(*AdGroupAdOperation_Create)(nil),
		(*AdGroupAdOperation_Update)(nil),
		(*AdGroupAdOperation_Remove)(nil),
	}
}

func _AdGroupAdOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AdGroupAdOperation)
	// operation
	switch x := m.Operation.(type) {
	case *AdGroupAdOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *AdGroupAdOperation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *AdGroupAdOperation_Remove:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("AdGroupAdOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _AdGroupAdOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AdGroupAdOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.AdGroupAd)
		err := b.DecodeMessage(msg)
		m.Operation = &AdGroupAdOperation_Create{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.AdGroupAd)
		err := b.DecodeMessage(msg)
		m.Operation = &AdGroupAdOperation_Update{msg}
		return true, err
	case 3: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &AdGroupAdOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _AdGroupAdOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AdGroupAdOperation)
	// operation
	switch x := m.Operation.(type) {
	case *AdGroupAdOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdGroupAdOperation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdGroupAdOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for an ad group ad mutate.
type MutateAdGroupAdsResponse struct {
	// All results for the mutate.
	Results              []*MutateAdGroupAdResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MutateAdGroupAdsResponse) Reset()         { *m = MutateAdGroupAdsResponse{} }
func (m *MutateAdGroupAdsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdsResponse) ProtoMessage()    {}
func (*MutateAdGroupAdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0977399cace737b, []int{3}
}
func (m *MutateAdGroupAdsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupAdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdsResponse.Merge(m, src)
}
func (m *MutateAdGroupAdsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdsResponse.Size(m)
}
func (m *MutateAdGroupAdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdsResponse proto.InternalMessageInfo

func (m *MutateAdGroupAdsResponse) GetResults() []*MutateAdGroupAdResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the ad mutate.
type MutateAdGroupAdResult struct {
	// The resource name returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupAdResult) Reset()         { *m = MutateAdGroupAdResult{} }
func (m *MutateAdGroupAdResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdResult) ProtoMessage()    {}
func (*MutateAdGroupAdResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0977399cace737b, []int{4}
}
func (m *MutateAdGroupAdResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdResult.Unmarshal(m, b)
}
func (m *MutateAdGroupAdResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdResult.Merge(m, src)
}
func (m *MutateAdGroupAdResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdResult.Size(m)
}
func (m *MutateAdGroupAdResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdResult proto.InternalMessageInfo

func (m *MutateAdGroupAdResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupAdRequest)(nil), "google.ads.googleads.v0.services.GetAdGroupAdRequest")
	proto.RegisterType((*MutateAdGroupAdsRequest)(nil), "google.ads.googleads.v0.services.MutateAdGroupAdsRequest")
	proto.RegisterType((*AdGroupAdOperation)(nil), "google.ads.googleads.v0.services.AdGroupAdOperation")
	proto.RegisterType((*MutateAdGroupAdsResponse)(nil), "google.ads.googleads.v0.services.MutateAdGroupAdsResponse")
	proto.RegisterType((*MutateAdGroupAdResult)(nil), "google.ads.googleads.v0.services.MutateAdGroupAdResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdGroupAdServiceClient is the client API for AdGroupAdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupAdServiceClient interface {
	// Returns the requested ad in full detail.
	GetAdGroupAd(ctx context.Context, in *GetAdGroupAdRequest, opts ...grpc.CallOption) (*resources.AdGroupAd, error)
	// Creates, updates, or removes ads. Operation statuses are returned.
	MutateAdGroupAds(ctx context.Context, in *MutateAdGroupAdsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdsResponse, error)
}

type adGroupAdServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdGroupAdServiceClient(cc *grpc.ClientConn) AdGroupAdServiceClient {
	return &adGroupAdServiceClient{cc}
}

func (c *adGroupAdServiceClient) GetAdGroupAd(ctx context.Context, in *GetAdGroupAdRequest, opts ...grpc.CallOption) (*resources.AdGroupAd, error) {
	out := new(resources.AdGroupAd)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AdGroupAdService/GetAdGroupAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupAdServiceClient) MutateAdGroupAds(ctx context.Context, in *MutateAdGroupAdsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdsResponse, error) {
	out := new(MutateAdGroupAdsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AdGroupAdService/MutateAdGroupAds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAdServiceServer is the server API for AdGroupAdService service.
type AdGroupAdServiceServer interface {
	// Returns the requested ad in full detail.
	GetAdGroupAd(context.Context, *GetAdGroupAdRequest) (*resources.AdGroupAd, error)
	// Creates, updates, or removes ads. Operation statuses are returned.
	MutateAdGroupAds(context.Context, *MutateAdGroupAdsRequest) (*MutateAdGroupAdsResponse, error)
}

func RegisterAdGroupAdServiceServer(s *grpc.Server, srv AdGroupAdServiceServer) {
	s.RegisterService(&_AdGroupAdService_serviceDesc, srv)
}

func _AdGroupAdService_GetAdGroupAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdServiceServer).GetAdGroupAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AdGroupAdService/GetAdGroupAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdServiceServer).GetAdGroupAd(ctx, req.(*GetAdGroupAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupAdService_MutateAdGroupAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupAdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdServiceServer).MutateAdGroupAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AdGroupAdService/MutateAdGroupAds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdServiceServer).MutateAdGroupAds(ctx, req.(*MutateAdGroupAdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupAdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.AdGroupAdService",
	HandlerType: (*AdGroupAdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupAd",
			Handler:    _AdGroupAdService_GetAdGroupAd_Handler,
		},
		{
			MethodName: "MutateAdGroupAds",
			Handler:    _AdGroupAdService_MutateAdGroupAds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/kritzware/google-ads-go/services/ad_group_ad_service.proto",
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/services/ad_group_ad_service.proto", fileDescriptor_b0977399cace737b)
}

var fileDescriptor_b0977399cace737b = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x0e, 0x0a, 0xea, 0xb8, 0x48, 0xd5, 0xa2, 0x0a, 0x2b, 0x42, 0x22, 0x32, 0x1c, 0xa2,
	0xa8, 0xf2, 0x46, 0x69, 0x10, 0xc2, 0xd0, 0x83, 0x73, 0x68, 0xe0, 0x50, 0x28, 0x06, 0xf5, 0x14,
	0xc9, 0xda, 0x66, 0xb7, 0x51, 0xd4, 0x38, 0x6b, 0xbc, 0xeb, 0x5c, 0xaa, 0x5e, 0x38, 0x20, 0x8e,
	0x48, 0xfc, 0x01, 0x37, 0x38, 0xf0, 0x17, 0x5c, 0xb8, 0xf2, 0x07, 0x88, 0x0f, 0x41, 0xeb, 0xf5,
	0xba, 0x21, 0x6d, 0x15, 0xda, 0xdb, 0x78, 0xfd, 0xde, 0x9b, 0x99, 0x37, 0x33, 0x10, 0x8c, 0x39,
	0x1f, 0x4f, 0x19, 0x26, 0x54, 0x60, 0x1d, 0xaa, 0x68, 0xde, 0xc1, 0x82, 0x65, 0xf3, 0xc9, 0x88,
	0x09, 0x4c, 0x68, 0x3c, 0xce, 0x78, 0x9e, 0xc6, 0x84, 0xc6, 0xe5, 0xa3, 0x9f, 0x66, 0x5c, 0x72,
	0xd4, 0xd4, 0x04, 0x9f, 0x50, 0xe1, 0x57, 0x5c, 0x7f, 0xde, 0xf1, 0x0d, 0xb7, 0xb1, 0x7d, 0x99,
	0x7a, 0xc6, 0x04, 0xcf, 0xb3, 0x25, 0x79, 0x2d, 0xdb, 0xb8, 0x67, 0x48, 0xe9, 0x04, 0x93, 0xd9,
	0x8c, 0x4b, 0x22, 0x27, 0x7c, 0x26, 0xca, 0xbf, 0x65, 0x52, 0x5c, 0x7c, 0x1d, 0xe6, 0x47, 0xf8,
	0x68, 0xc2, 0xa6, 0x34, 0x4e, 0x88, 0x38, 0xd6, 0x08, 0x2f, 0x80, 0x3b, 0x03, 0x26, 0x43, 0x3a,
	0x50, 0xb2, 0x21, 0x8d, 0xd8, 0xbb, 0x9c, 0x09, 0x89, 0x1e, 0xc0, 0x6d, 0x93, 0x35, 0x9e, 0x91,
	0x84, 0xb9, 0x56, 0xd3, 0x6a, 0xad, 0x45, 0xeb, 0xe6, 0xf1, 0x25, 0x49, 0x98, 0xf7, 0xc9, 0x82,
	0xbb, 0x7b, 0xb9, 0x24, 0x92, 0x55, 0x7c, 0x61, 0x04, 0xee, 0x83, 0x33, 0xca, 0x85, 0xe4, 0x09,
	0xcb, 0xe2, 0x09, 0x2d, 0xe9, 0x60, 0x9e, 0x5e, 0x50, 0xf4, 0x16, 0x80, 0xa7, 0x2c, 0xd3, 0xe5,
	0xba, 0x76, 0xb3, 0xd6, 0x72, 0xba, 0x3d, 0x7f, 0x95, 0x49, 0x7e, 0x95, 0xe9, 0x95, 0x21, 0x47,
	0x0b, 0x3a, 0xde, 0x47, 0x1b, 0xd0, 0x79, 0x08, 0x7a, 0x0a, 0x4e, 0x9e, 0x52, 0x22, 0x59, 0xd1,
	0xba, 0x7b, 0xb3, 0x69, 0xb5, 0x9c, 0x6e, 0xc3, 0x64, 0x33, 0xee, 0xf8, 0xbb, 0xca, 0x9d, 0x3d,
	0x22, 0x8e, 0x23, 0xd0, 0x70, 0x15, 0xa3, 0x5d, 0xa8, 0x8f, 0x32, 0x46, 0xa4, 0x36, 0xc1, 0xe9,
	0x6e, 0x5d, 0x5a, 0x65, 0x35, 0xa8, 0xb3, 0x32, 0x9f, 0xdf, 0x88, 0x4a, 0xb6, 0xd2, 0xd1, 0xaa,
	0xae, 0x7d, 0x3d, 0x1d, 0xcd, 0x46, 0x2e, 0xd4, 0x33, 0x96, 0xf0, 0x39, 0x73, 0x6b, 0xca, 0x55,
	0xf5, 0x47, 0x7f, 0xf7, 0x1d, 0x58, 0xab, 0xbc, 0xf0, 0x12, 0x70, 0xcf, 0x0f, 0x47, 0xa4, 0x7c,
	0x26, 0x18, 0x7a, 0x0d, 0xb7, 0x32, 0x26, 0xf2, 0xa9, 0x34, 0xce, 0x3f, 0x5e, 0xed, 0xfc, 0x92,
	0x58, 0x54, 0xf0, 0x23, 0xa3, 0xe3, 0x3d, 0x83, 0xcd, 0x0b, 0x11, 0xff, 0xb5, 0x4a, 0xdd, 0x0f,
	0x35, 0xd8, 0xa8, 0x88, 0x6f, 0x74, 0x4a, 0xf4, 0xd5, 0x82, 0xf5, 0xc5, 0xe5, 0x44, 0x8f, 0x56,
	0x57, 0x79, 0xc1, 0x32, 0x37, 0xae, 0x64, 0xb4, 0xd7, 0x7b, 0xff, 0xeb, 0xcf, 0x67, 0xdb, 0x47,
	0x5b, 0xea, 0xf4, 0x4e, 0xfe, 0x29, 0x7d, 0xc7, 0xec, 0xaf, 0xc0, 0x6d, 0x4c, 0x2a, 0x5b, 0x71,
	0xfb, 0x14, 0xfd, 0xb0, 0x60, 0x63, 0xd9, 0x6e, 0xf4, 0xe4, 0xca, 0xae, 0x9a, 0xfb, 0x69, 0x04,
	0xd7, 0xa1, 0xea, 0xe9, 0x7a, 0x41, 0xd1, 0x41, 0xcf, 0xc3, 0xaa, 0x83, 0xb3, 0x92, 0x4f, 0x16,
	0x0e, 0x72, 0xa7, 0x7d, 0xba, 0xd0, 0x40, 0x90, 0x14, 0x52, 0x81, 0xd5, 0xee, 0x7f, 0xb7, 0xe0,
	0xe1, 0x88, 0x27, 0x2b, 0xb3, 0xf7, 0x37, 0x97, 0xc7, 0xb5, 0xaf, 0xae, 0x68, 0xdf, 0xfa, 0x62,
	0xd7, 0x06, 0x61, 0xf8, 0xcd, 0x6e, 0x0e, 0xb4, 0x42, 0x48, 0x85, 0xaf, 0x43, 0x15, 0x1d, 0x74,
	0xfc, 0x12, 0x2e, 0x7e, 0x1a, 0xc8, 0x30, 0xa4, 0x62, 0x58, 0x41, 0x86, 0x07, 0x9d, 0xa1, 0x81,
	0xfc, 0x5e, 0x0d, 0x39, 0xac, 0x17, 0xd7, 0xbb, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xa6,
	0x35, 0x3d, 0x9c, 0x05, 0x00, 0x00,
}