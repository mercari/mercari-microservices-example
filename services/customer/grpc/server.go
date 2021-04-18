package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mercari/go-conference-2021-spring-office-hour/services/customer/db"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/customer/proto"
)

var _ proto.CustomerServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedCustomerServiceServer
	db db.DB
}

func (s *server) GetCustomer(ctx context.Context, req *proto.GetCustomerRequest) (*proto.GetCustomerResponse, error) {
	c, err := s.db.GetCustomer(ctx, req.Name)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.GetCustomerResponse{
		Customer: &proto.Customer{
			Id:   c.ID,
			Name: c.Name,
		},
	}, nil
}
