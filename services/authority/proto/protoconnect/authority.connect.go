// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: services/authority/proto/authority.proto

package protoconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	proto "github.com/mercari/mercari-microservices-example/services/authority/proto"
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
	// AuthorityServiceName is the fully-qualified name of the AuthorityService service.
	AuthorityServiceName = "mercari.mercari_microservices_example.authority.AuthorityService"
)

// AuthorityServiceClient is a client for the
// mercari.mercari_microservices_example.authority.AuthorityService service.
type AuthorityServiceClient interface {
	Signup(context.Context, *connect_go.Request[proto.SignupRequest]) (*connect_go.Response[proto.SignupResponse], error)
	Signin(context.Context, *connect_go.Request[proto.SigninRequest]) (*connect_go.Response[proto.SigninResponse], error)
	ListPublicKeys(context.Context, *connect_go.Request[proto.ListPublicKeysRequest]) (*connect_go.Response[proto.ListPublicKeysResponse], error)
}

// NewAuthorityServiceClient constructs a client for the
// mercari.mercari_microservices_example.authority.AuthorityService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthorityServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AuthorityServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authorityServiceClient{
		signup: connect_go.NewClient[proto.SignupRequest, proto.SignupResponse](
			httpClient,
			baseURL+"/mercari.mercari_microservices_example.authority.AuthorityService/Signup",
			opts...,
		),
		signin: connect_go.NewClient[proto.SigninRequest, proto.SigninResponse](
			httpClient,
			baseURL+"/mercari.mercari_microservices_example.authority.AuthorityService/Signin",
			opts...,
		),
		listPublicKeys: connect_go.NewClient[proto.ListPublicKeysRequest, proto.ListPublicKeysResponse](
			httpClient,
			baseURL+"/mercari.mercari_microservices_example.authority.AuthorityService/ListPublicKeys",
			opts...,
		),
	}
}

// authorityServiceClient implements AuthorityServiceClient.
type authorityServiceClient struct {
	signup         *connect_go.Client[proto.SignupRequest, proto.SignupResponse]
	signin         *connect_go.Client[proto.SigninRequest, proto.SigninResponse]
	listPublicKeys *connect_go.Client[proto.ListPublicKeysRequest, proto.ListPublicKeysResponse]
}

// Signup calls mercari.mercari_microservices_example.authority.AuthorityService.Signup.
func (c *authorityServiceClient) Signup(ctx context.Context, req *connect_go.Request[proto.SignupRequest]) (*connect_go.Response[proto.SignupResponse], error) {
	return c.signup.CallUnary(ctx, req)
}

// Signin calls mercari.mercari_microservices_example.authority.AuthorityService.Signin.
func (c *authorityServiceClient) Signin(ctx context.Context, req *connect_go.Request[proto.SigninRequest]) (*connect_go.Response[proto.SigninResponse], error) {
	return c.signin.CallUnary(ctx, req)
}

// ListPublicKeys calls
// mercari.mercari_microservices_example.authority.AuthorityService.ListPublicKeys.
func (c *authorityServiceClient) ListPublicKeys(ctx context.Context, req *connect_go.Request[proto.ListPublicKeysRequest]) (*connect_go.Response[proto.ListPublicKeysResponse], error) {
	return c.listPublicKeys.CallUnary(ctx, req)
}

// AuthorityServiceHandler is an implementation of the
// mercari.mercari_microservices_example.authority.AuthorityService service.
type AuthorityServiceHandler interface {
	Signup(context.Context, *connect_go.Request[proto.SignupRequest]) (*connect_go.Response[proto.SignupResponse], error)
	Signin(context.Context, *connect_go.Request[proto.SigninRequest]) (*connect_go.Response[proto.SigninResponse], error)
	ListPublicKeys(context.Context, *connect_go.Request[proto.ListPublicKeysRequest]) (*connect_go.Response[proto.ListPublicKeysResponse], error)
}

// NewAuthorityServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthorityServiceHandler(svc AuthorityServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/mercari.mercari_microservices_example.authority.AuthorityService/Signup", connect_go.NewUnaryHandler(
		"/mercari.mercari_microservices_example.authority.AuthorityService/Signup",
		svc.Signup,
		opts...,
	))
	mux.Handle("/mercari.mercari_microservices_example.authority.AuthorityService/Signin", connect_go.NewUnaryHandler(
		"/mercari.mercari_microservices_example.authority.AuthorityService/Signin",
		svc.Signin,
		opts...,
	))
	mux.Handle("/mercari.mercari_microservices_example.authority.AuthorityService/ListPublicKeys", connect_go.NewUnaryHandler(
		"/mercari.mercari_microservices_example.authority.AuthorityService/ListPublicKeys",
		svc.ListPublicKeys,
		opts...,
	))
	return "/mercari.mercari_microservices_example.authority.AuthorityService/", mux
}

// UnimplementedAuthorityServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthorityServiceHandler struct{}

func (UnimplementedAuthorityServiceHandler) Signup(context.Context, *connect_go.Request[proto.SignupRequest]) (*connect_go.Response[proto.SignupResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mercari.mercari_microservices_example.authority.AuthorityService.Signup is not implemented"))
}

func (UnimplementedAuthorityServiceHandler) Signin(context.Context, *connect_go.Request[proto.SigninRequest]) (*connect_go.Response[proto.SigninResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mercari.mercari_microservices_example.authority.AuthorityService.Signin is not implemented"))
}

func (UnimplementedAuthorityServiceHandler) ListPublicKeys(context.Context, *connect_go.Request[proto.ListPublicKeysRequest]) (*connect_go.Response[proto.ListPublicKeysResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mercari.mercari_microservices_example.authority.AuthorityService.ListPublicKeys is not implemented"))
}