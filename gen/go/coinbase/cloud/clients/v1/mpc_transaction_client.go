// Copyright 2023 Google LLC
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


package v1

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	mpc_transactionspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newMPCTransactionClientHook clientHook

// MPCTransactionCallOptions contains the retry settings for each method of MPCTransactionClient.
type MPCTransactionCallOptions struct {
	CreateMPCTransaction []gax.CallOption
	GetMPCTransaction []gax.CallOption
	ListMPCTransactions []gax.CallOption
}

func defaultMPCTransactionGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("api.developer.coinbase.com/waas/mpc_transactions:443"),
		internaloption.WithDefaultMTLSEndpoint("api.developer.coinbase.com/waas/mpc_transactions:443"),
		internaloption.WithDefaultAudience("https://api.developer.coinbase.com/waas/mpc_transactions/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultMPCTransactionCallOptions() *MPCTransactionCallOptions {
	return &MPCTransactionCallOptions{
		CreateMPCTransaction: []gax.CallOption{
		},
		GetMPCTransaction: []gax.CallOption{
		},
		ListMPCTransactions: []gax.CallOption{
		},
	}
}

func defaultMPCTransactionRESTCallOptions() *MPCTransactionCallOptions {
	return &MPCTransactionCallOptions{
		CreateMPCTransaction: []gax.CallOption{
		},
		GetMPCTransaction: []gax.CallOption{
		},
		ListMPCTransactions: []gax.CallOption{
		},
	}
}

// internalMPCTransactionClient is an interface that defines the methods available from .
type internalMPCTransactionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateMPCTransaction(context.Context, *mpc_transactionspb.CreateMPCTransactionRequest, ...gax.CallOption) (*CreateMPCTransactionOperation, error)
	CreateMPCTransactionOperation(name string) *CreateMPCTransactionOperation
	GetMPCTransaction(context.Context, *mpc_transactionspb.GetMPCTransactionRequest, ...gax.CallOption) (*mpc_transactionspb.MPCTransaction, error)
	ListMPCTransactions(context.Context, *mpc_transactionspb.ListMPCTransactionsRequest, ...gax.CallOption) *MPCTransactionIterator
}

// MPCTransactionClient is a client for interacting with .
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service that orchestrates on-chain transactions originating from MPCWallets. The service handles
// nonce management, transaction construction, signing, broadcasting, and confirmation detection including
// on-chain re-organizations.
type MPCTransactionClient struct {
	// The internal transport-dependent client.
	internalClient internalMPCTransactionClient

	// The call options for this service.
	CallOptions *MPCTransactionCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MPCTransactionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MPCTransactionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *MPCTransactionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateMPCTransaction creates an MPCTransaction. The long-running operation returned from this API contains
// information about the state of the MPCTransaction that can be used to complete the operation.
// The LRO is considered Done once the MPCTransaction reaches a state of `CONFIRMING`` (i.e.,
// broadcast on-chain). See the MPCTransaction documentation for its lifecycle.
func (c *MPCTransactionClient) CreateMPCTransaction(ctx context.Context, req *mpc_transactionspb.CreateMPCTransactionRequest, opts ...gax.CallOption) (*CreateMPCTransactionOperation, error) {
	return c.internalClient.CreateMPCTransaction(ctx, req, opts...)
}

// CreateMPCTransactionOperation returns a new CreateMPCTransactionOperation from a given name.
// The name must be that of a previously created CreateMPCTransactionOperation, possibly from a different process.
func (c *MPCTransactionClient) CreateMPCTransactionOperation(name string) *CreateMPCTransactionOperation {
	return c.internalClient.CreateMPCTransactionOperation(name)
}

// GetMPCTransaction gets an MPCTransaction. There can be a delay between when CreateMPCTransaction is called
// and when this API returns an MPCTransaction in the CREATED state.
func (c *MPCTransactionClient) GetMPCTransaction(ctx context.Context, req *mpc_transactionspb.GetMPCTransactionRequest, opts ...gax.CallOption) (*mpc_transactionspb.MPCTransaction, error) {
	return c.internalClient.GetMPCTransaction(ctx, req, opts...)
}

// ListMPCTransactions returns a list of MPCTransactions in an MPCWallet.
func (c *MPCTransactionClient) ListMPCTransactions(ctx context.Context, req *mpc_transactionspb.ListMPCTransactionsRequest, opts ...gax.CallOption) *MPCTransactionIterator {
	return c.internalClient.ListMPCTransactions(ctx, req, opts...)
}

// mPCTransactionGRPCClient is a client for interacting with  over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type mPCTransactionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing MPCTransactionClient
	CallOptions **MPCTransactionCallOptions

	// The gRPC API client.
	mPCTransactionClient mpc_transactionspb.MPCTransactionServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewMPCTransactionClient creates a new mpc transaction service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service that orchestrates on-chain transactions originating from MPCWallets. The service handles
// nonce management, transaction construction, signing, broadcasting, and confirmation detection including
// on-chain re-organizations.
func NewMPCTransactionClient(ctx context.Context, opts ...option.ClientOption) (*MPCTransactionClient, error) {
	clientOpts := defaultMPCTransactionGRPCClientOptions()
	if newMPCTransactionClientHook != nil {
		hookOpts, err := newMPCTransactionClientHook(ctx, clientHookParams{})
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
	client := MPCTransactionClient{CallOptions: defaultMPCTransactionCallOptions()}

	c := &mPCTransactionGRPCClient{
		connPool:    connPool,
		disableDeadlines: disableDeadlines,
		mPCTransactionClient: mpc_transactionspb.NewMPCTransactionServiceClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *mPCTransactionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *mPCTransactionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *mPCTransactionGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type mPCTransactionRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing MPCTransactionClient
	CallOptions **MPCTransactionCallOptions
}

// NewMPCTransactionRESTClient creates a new mpc transaction service rest client.
//
// A service that orchestrates on-chain transactions originating from MPCWallets. The service handles
// nonce management, transaction construction, signing, broadcasting, and confirmation detection including
// on-chain re-organizations.
func NewMPCTransactionRESTClient(ctx context.Context, opts ...option.ClientOption) (*MPCTransactionClient, error) {
	clientOpts := append(defaultMPCTransactionRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultMPCTransactionRESTCallOptions()
	c := &mPCTransactionRESTClient{
		endpoint: endpoint,
		httpClient: httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &MPCTransactionClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultMPCTransactionRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://api.developer.coinbase.com/waas/mpc_transactions"),
		internaloption.WithDefaultMTLSEndpoint("https://api.developer.coinbase.com/waas/mpc_transactions"),
		internaloption.WithDefaultAudience("https://api.developer.coinbase.com/waas/mpc_transactions/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}
// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *mPCTransactionRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *mPCTransactionRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *mPCTransactionRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *mPCTransactionGRPCClient) CreateMPCTransaction(ctx context.Context, req *mpc_transactionspb.CreateMPCTransactionRequest, opts ...gax.CallOption) (*CreateMPCTransactionOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateMPCTransaction[0:len((*c.CallOptions).CreateMPCTransaction):len((*c.CallOptions).CreateMPCTransaction)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.mPCTransactionClient.CreateMPCTransaction(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateMPCTransactionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *mPCTransactionGRPCClient) GetMPCTransaction(ctx context.Context, req *mpc_transactionspb.GetMPCTransactionRequest, opts ...gax.CallOption) (*mpc_transactionspb.MPCTransaction, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetMPCTransaction[0:len((*c.CallOptions).GetMPCTransaction):len((*c.CallOptions).GetMPCTransaction)], opts...)
	var resp *mpc_transactionspb.MPCTransaction
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.mPCTransactionClient.GetMPCTransaction(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *mPCTransactionGRPCClient) ListMPCTransactions(ctx context.Context, req *mpc_transactionspb.ListMPCTransactionsRequest, opts ...gax.CallOption) *MPCTransactionIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListMPCTransactions[0:len((*c.CallOptions).ListMPCTransactions):len((*c.CallOptions).ListMPCTransactions)], opts...)
	it := &MPCTransactionIterator{}
	req = proto.Clone(req).(*mpc_transactionspb.ListMPCTransactionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*mpc_transactionspb.MPCTransaction, string, error) {
		resp := &mpc_transactionspb.ListMPCTransactionsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.mPCTransactionClient.ListMPCTransactions(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetMpcTransactions(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// CreateMPCTransaction creates an MPCTransaction. The long-running operation returned from this API contains
// information about the state of the MPCTransaction that can be used to complete the operation.
// The LRO is considered Done once the MPCTransaction reaches a state of `CONFIRMING`` (i.e.,
// broadcast on-chain). See the MPCTransaction documentation for its lifecycle.
func (c *mPCTransactionRESTClient) CreateMPCTransaction(ctx context.Context, req *mpc_transactionspb.CreateMPCTransactionRequest, opts ...gax.CallOption) (*CreateMPCTransactionOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/mpcTransactions", req.GetParent())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil{
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &CreateMPCTransactionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// GetMPCTransaction gets an MPCTransaction. There can be a delay between when CreateMPCTransaction is called
// and when this API returns an MPCTransaction in the CREATED state.
func (c *mPCTransactionRESTClient) GetMPCTransaction(ctx context.Context, req *mpc_transactionspb.GetMPCTransactionRequest, opts ...gax.CallOption) (*mpc_transactionspb.MPCTransaction, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetMPCTransaction[0:len((*c.CallOptions).GetMPCTransaction):len((*c.CallOptions).GetMPCTransaction)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &mpc_transactionspb.MPCTransaction{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil{
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
// ListMPCTransactions returns a list of MPCTransactions in an MPCWallet.
func (c *mPCTransactionRESTClient) ListMPCTransactions(ctx context.Context, req *mpc_transactionspb.ListMPCTransactionsRequest, opts ...gax.CallOption) *MPCTransactionIterator {
	it := &MPCTransactionIterator{}
	req = proto.Clone(req).(*mpc_transactionspb.ListMPCTransactionsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*mpc_transactionspb.MPCTransaction, string, error) {
		resp := &mpc_transactionspb.ListMPCTransactionsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1/%v/mpcTransactions", req.GetParent())

		params := url.Values{}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil{
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return maybeUnknownEnum(err)
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetMpcTransactions(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}
// CreateMPCTransactionOperation manages a long-running operation from CreateMPCTransaction.
type CreateMPCTransactionOperation struct {
	lro *longrunning.Operation
	pollPath string
}

// CreateMPCTransactionOperation returns a new CreateMPCTransactionOperation from a given name.
// The name must be that of a previously created CreateMPCTransactionOperation, possibly from a different process.
func (c *mPCTransactionGRPCClient) CreateMPCTransactionOperation(name string) *CreateMPCTransactionOperation {
	return &CreateMPCTransactionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// CreateMPCTransactionOperation returns a new CreateMPCTransactionOperation from a given name.
// The name must be that of a previously created CreateMPCTransactionOperation, possibly from a different process.
func (c *mPCTransactionRESTClient) CreateMPCTransactionOperation(name string) *CreateMPCTransactionOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &CreateMPCTransactionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateMPCTransactionOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*mpc_transactionspb.MPCTransaction, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp mpc_transactionspb.MPCTransaction
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateMPCTransactionOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*mpc_transactionspb.MPCTransaction, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp mpc_transactionspb.MPCTransaction
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateMPCTransactionOperation) Metadata() (*mpc_transactionspb.CreateMPCTransactionMetadata, error) {
	var meta mpc_transactionspb.CreateMPCTransactionMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateMPCTransactionOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateMPCTransactionOperation) Name() string {
	return op.lro.Name()
}

// MPCTransactionIterator manages a stream of *mpc_transactionspb.MPCTransaction.
type MPCTransactionIterator struct {
	items    []*mpc_transactionspb.MPCTransaction
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*mpc_transactionspb.MPCTransaction, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *MPCTransactionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *MPCTransactionIterator) Next() (*mpc_transactionspb.MPCTransaction, error) {
	var item *mpc_transactionspb.MPCTransaction
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *MPCTransactionIterator) bufLen() int {
	return len(it.items)
}

func (it *MPCTransactionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}