package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mercari/go-conference-2021-spring-office-hour/services/authority/proto"
	customer "github.com/mercari/go-conference-2021-spring-office-hour/services/customer/proto"
)

var _ proto.AuthorityServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedAuthorityServiceServer
	customerClient customer.CustomerServiceClient
}

func (s *server) Signin(ctx context.Context, req *proto.SigninRequest) (*proto.SigninResponse, error) {
	res, err := s.customerClient.GetCustomerByName(ctx, &customer.GetCustomerByNameRequest{Name: req.Name})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated") // TODO:
	}

	return &proto.SigninResponse{
		AccessToken: res.GetCustomer().Id, // TODO:
	}, nil
}
