// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package gads

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcespb "github.com/opteo/google-ads-go/resources"
	servicespb "github.com/opteo/google-ads-go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newBillingSetupClientHook clientHook

// BillingSetupCallOptions contains the retry settings for each method of BillingSetupClient.
type BillingSetupCallOptions struct {
	GetBillingSetup []gax.CallOption
	MutateBillingSetup []gax.CallOption
}

func defaultBillingSetupGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultBillingSetupCallOptions() *BillingSetupCallOptions {
	return &BillingSetupCallOptions{
		GetBillingSetup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		MutateBillingSetup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalBillingSetupClient is an interface that defines the methods availaible from Google Ads API.
type internalBillingSetupClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetBillingSetup(context.Context, *servicespb.GetBillingSetupRequest, ...gax.CallOption) (*resourcespb.BillingSetup, error)
	MutateBillingSetup(context.Context, *servicespb.MutateBillingSetupRequest, ...gax.CallOption) (*servicespb.MutateBillingSetupResponse, error)
}

// BillingSetupClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for designating the business entity responsible for accrued costs.
//
// A billing setup is associated with a payments account.  Billing-related
// activity for all billing setups associated with a particular payments account
// will appear on a single invoice generated monthly.
//
// Mutates:
// The REMOVE operation cancels a pending billing setup.
// The CREATE operation creates a new billing setup.
type BillingSetupClient struct {
	// The internal transport-dependent client.
	internalClient internalBillingSetupClient

	// The call options for this service.
	CallOptions *BillingSetupCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BillingSetupClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BillingSetupClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *BillingSetupClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetBillingSetup returns a billing setup.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *BillingSetupClient) GetBillingSetup(ctx context.Context, req *servicespb.GetBillingSetupRequest, opts ...gax.CallOption) (*resourcespb.BillingSetup, error) {
	return c.internalClient.GetBillingSetup(ctx, req, opts...)
}

// MutateBillingSetup creates a billing setup, or cancels an existing billing setup.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// BillingSetupError (at )
// DateError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *BillingSetupClient) MutateBillingSetup(ctx context.Context, req *servicespb.MutateBillingSetupRequest, opts ...gax.CallOption) (*servicespb.MutateBillingSetupResponse, error) {
	return c.internalClient.MutateBillingSetup(ctx, req, opts...)
}

// billingSetupGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type billingSetupGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing BillingSetupClient
	CallOptions **BillingSetupCallOptions

	// The gRPC API client.
	billingSetupClient servicespb.BillingSetupServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewBillingSetupClient creates a new billing setup service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for designating the business entity responsible for accrued costs.
//
// A billing setup is associated with a payments account.  Billing-related
// activity for all billing setups associated with a particular payments account
// will appear on a single invoice generated monthly.
//
// Mutates:
// The REMOVE operation cancels a pending billing setup.
// The CREATE operation creates a new billing setup.
func NewBillingSetupClient(ctx context.Context, opts ...option.ClientOption) (*BillingSetupClient, error) {
	clientOpts := defaultBillingSetupGRPCClientOptions()
	if newBillingSetupClientHook != nil {
		hookOpts, err := newBillingSetupClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := BillingSetupClient{CallOptions: defaultBillingSetupCallOptions()}

	c := &billingSetupGRPCClient{
		connPool:    connPool,
		disableDeadlines: disableDeadlines,
		billingSetupClient: servicespb.NewBillingSetupServiceClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *billingSetupGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *billingSetupGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *billingSetupGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *billingSetupGRPCClient) GetBillingSetup(ctx context.Context, req *servicespb.GetBillingSetupRequest, opts ...gax.CallOption) (*resourcespb.BillingSetup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetBillingSetup[0:len((*c.CallOptions).GetBillingSetup):len((*c.CallOptions).GetBillingSetup)], opts...)
	var resp *resourcespb.BillingSetup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.billingSetupClient.GetBillingSetup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *billingSetupGRPCClient) MutateBillingSetup(ctx context.Context, req *servicespb.MutateBillingSetupRequest, opts ...gax.CallOption) (*servicespb.MutateBillingSetupResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateBillingSetup[0:len((*c.CallOptions).MutateBillingSetup):len((*c.CallOptions).MutateBillingSetup)], opts...)
	var resp *servicespb.MutateBillingSetupResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.billingSetupClient.MutateBillingSetup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
