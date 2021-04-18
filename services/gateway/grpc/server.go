package grpc

import (
	"context"
	"fmt"

	authority "github.com/mercari/go-conference-2021-spring-office-hour/services/authority/proto"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/gateway/proto"
)

var _ proto.GatewayServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedGatewayServiceServer
	authorityClient authority.AuthorityServiceClient
}

func (s *server) Signin(ctx context.Context, req *proto.SigninRequest) (*proto.SigninResponse, error) {
	res, err := s.authorityClient.Signin(ctx, &authority.SigninRequest{Name: req.Name})
	if err != nil {
		return nil, fmt.Errorf("failed to call authority.Singin: %w", err) // TODO:
	}

	return &proto.SigninResponse{
		AccessToken: res.AccessToken,
	}, nil
}
