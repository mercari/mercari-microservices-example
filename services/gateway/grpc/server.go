package grpc

import (
	"context"

	authority "github.com/mercari/go-conference-2021-spring-office-hour/services/authority/proto"
	catalog "github.com/mercari/go-conference-2021-spring-office-hour/services/catalog/proto"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/gateway/proto"
)

var _ proto.GatewayServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedGatewayServiceServer
	authorityClient authority.AuthorityServiceClient
	catalogClient   catalog.CatalogServiceClient
}

func (s *server) Signin(ctx context.Context, req *authority.SigninRequest) (*authority.SigninResponse, error) {
	return s.authorityClient.Signin(ctx, req)
}

func (s *server) GetItem(ctx context.Context, req *catalog.GetItemRequest) (*catalog.GetItemResponse, error) {
	return s.catalogClient.GetItem(ctx, req)
}
