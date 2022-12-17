package grpc

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	db "github.com/mercari/mercari-microservices-example/platform/db/proto"
	dbconnect "github.com/mercari/mercari-microservices-example/platform/db/proto/protoconnect"
	"github.com/mercari/mercari-microservices-example/services/customer/proto"
	"github.com/mercari/mercari-microservices-example/services/customer/proto/protoconnect"
)

var _ protoconnect.CustomerServiceHandler = (*server)(nil)

type server struct {
	protoconnect.UnimplementedCustomerServiceHandler

	dbClient dbconnect.DBServiceClient
}

func (s *server) CreateCustomer(ctx context.Context, req *connect.Request[proto.CreateCustomerRequest]) (*connect.Response[proto.CreateCustomerResponse], error) {
	r := connect.NewRequest(&db.CreateCustomerRequest{Name: req.Msg.Name})
	res, err := s.dbClient.CreateCustomer(ctx, r)
	if err != nil {
		if connect.CodeOf(err) == connect.CodeAlreadyExists {
			return nil, connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("already exists: %w", err))
		}

		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	customer := res.Msg.GetCustomer()

	return connect.NewResponse(&proto.CreateCustomerResponse{
		Customer: &proto.Customer{
			Id:   customer.Id,
			Name: customer.Name,
		},
	}), nil
}

func (s *server) GetCustomer(ctx context.Context, req *connect.Request[proto.GetCustomerRequest]) (*connect.Response[proto.GetCustomerResponse], error) {
	res, err := s.dbClient.GetCustomer(ctx, connect.NewRequest(&db.GetCustomerRequest{Id: req.Msg.Id}))
	if err != nil {
		if connect.CodeOf(err) == connect.CodeNotFound {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("not found: %w", err))
		}

		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	customer := res.Msg.GetCustomer()

	return connect.NewResponse(&proto.GetCustomerResponse{
		Customer: &proto.Customer{
			Id:   customer.Id,
			Name: customer.Name,
		},
	}), nil
}

func (s *server) GetCustomerByName(ctx context.Context, req *connect.Request[proto.GetCustomerByNameRequest]) (*connect.Response[proto.GetCustomerByNameResponse], error) {
	res, err := s.dbClient.GetCustomerByName(ctx, connect.NewRequest(&db.GetCustomerByNameRequest{Name: req.Msg.Name}))
	if err != nil {
		if connect.CodeOf(err) == connect.CodeAlreadyExists {
			return nil, connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("already exists: %w", err))
		}

		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	customer := res.Msg.GetCustomer()

	return connect.NewResponse(&proto.GetCustomerByNameResponse{
		Customer: &proto.Customer{
			Id:   customer.Id,
			Name: customer.Name,
		},
	}), nil
}
