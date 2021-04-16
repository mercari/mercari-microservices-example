package grpc

import (
	"context"

	"github.com/mercari/go-conference-2021-spring-office-hour/services/balance/proto"
)

var _ proto.BalanceServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedBalanceServiceServer
}

func (s *server) GetBalance(context.Context, *proto.GetBalanceRequest) (*proto.GetBalanceResponse, error) {
	return &proto.GetBalanceResponse{}, nil
}
