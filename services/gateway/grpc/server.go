package grpc

import (
	"context"

	"github.com/mercari/go-conference-2021-spring-office-hour/services/gateway/proto"
)

var _ proto.GatewayServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedGatewayServiceServer
}

func (s *server) Signin(context.Context, *proto.SigninRequest) (*proto.SigninResponse, error) {
	return &proto.SigninResponse{
		AccessToken: "foo",
	}, nil
}
