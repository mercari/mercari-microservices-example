package grpc

import (
	"context"

	"github.com/mercari/go-conference-2021-spring-office-hour/services/authority/proto"
)

var _ proto.AuthorityServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedAuthorityServiceServer
}

func (s *server) Signin(context.Context, *proto.SigninRequest) (*proto.SigninResponse, error) {
	return &proto.SigninResponse{
		AccessToken: "52b554dc-3619-4f7f-9d67-1c35e39f6340", // TODO:
	}, nil
}
