// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: protocolbuffers/greet/v1/greet.proto

package greetv1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/Hirochon/connect-go-test/server/protocolbuffers/greet/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// GreetServiceName is the fully-qualified name of the GreetService service.
	GreetServiceName = "protocolbuffers.greet.v1.GreetService"
)

// GreetServiceClient is a client for the protocolbuffers.greet.v1.GreetService service.
type GreetServiceClient interface {
	GreetUnary(context.Context, *connect_go.Request[v1.GreetUnaryRequest]) (*connect_go.Response[v1.GreetUnaryResponse], error)
	GreetServerStream(context.Context, *connect_go.Request[v1.GreetServerStreamRequest]) (*connect_go.ServerStreamForClient[v1.GreetServerStreamResponse], error)
}

// NewGreetServiceClient constructs a client for the protocolbuffers.greet.v1.GreetService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGreetServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) GreetServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &greetServiceClient{
		greetUnary: connect_go.NewClient[v1.GreetUnaryRequest, v1.GreetUnaryResponse](
			httpClient,
			baseURL+"/protocolbuffers.greet.v1.GreetService/GreetUnary",
			opts...,
		),
		greetServerStream: connect_go.NewClient[v1.GreetServerStreamRequest, v1.GreetServerStreamResponse](
			httpClient,
			baseURL+"/protocolbuffers.greet.v1.GreetService/GreetServerStream",
			opts...,
		),
	}
}

// greetServiceClient implements GreetServiceClient.
type greetServiceClient struct {
	greetUnary        *connect_go.Client[v1.GreetUnaryRequest, v1.GreetUnaryResponse]
	greetServerStream *connect_go.Client[v1.GreetServerStreamRequest, v1.GreetServerStreamResponse]
}

// GreetUnary calls protocolbuffers.greet.v1.GreetService.GreetUnary.
func (c *greetServiceClient) GreetUnary(ctx context.Context, req *connect_go.Request[v1.GreetUnaryRequest]) (*connect_go.Response[v1.GreetUnaryResponse], error) {
	return c.greetUnary.CallUnary(ctx, req)
}

// GreetServerStream calls protocolbuffers.greet.v1.GreetService.GreetServerStream.
func (c *greetServiceClient) GreetServerStream(ctx context.Context, req *connect_go.Request[v1.GreetServerStreamRequest]) (*connect_go.ServerStreamForClient[v1.GreetServerStreamResponse], error) {
	return c.greetServerStream.CallServerStream(ctx, req)
}

// GreetServiceHandler is an implementation of the protocolbuffers.greet.v1.GreetService service.
type GreetServiceHandler interface {
	GreetUnary(context.Context, *connect_go.Request[v1.GreetUnaryRequest]) (*connect_go.Response[v1.GreetUnaryResponse], error)
	GreetServerStream(context.Context, *connect_go.Request[v1.GreetServerStreamRequest], *connect_go.ServerStream[v1.GreetServerStreamResponse]) error
}

// NewGreetServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGreetServiceHandler(svc GreetServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/protocolbuffers.greet.v1.GreetService/GreetUnary", connect_go.NewUnaryHandler(
		"/protocolbuffers.greet.v1.GreetService/GreetUnary",
		svc.GreetUnary,
		opts...,
	))
	mux.Handle("/protocolbuffers.greet.v1.GreetService/GreetServerStream", connect_go.NewServerStreamHandler(
		"/protocolbuffers.greet.v1.GreetService/GreetServerStream",
		svc.GreetServerStream,
		opts...,
	))
	return "/protocolbuffers.greet.v1.GreetService/", mux
}

// UnimplementedGreetServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGreetServiceHandler struct{}

func (UnimplementedGreetServiceHandler) GreetUnary(context.Context, *connect_go.Request[v1.GreetUnaryRequest]) (*connect_go.Response[v1.GreetUnaryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("protocolbuffers.greet.v1.GreetService.GreetUnary is not implemented"))
}

func (UnimplementedGreetServiceHandler) GreetServerStream(context.Context, *connect_go.Request[v1.GreetServerStreamRequest], *connect_go.ServerStream[v1.GreetServerStreamResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("protocolbuffers.greet.v1.GreetService.GreetServerStream is not implemented"))
}
