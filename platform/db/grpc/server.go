package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mercari/go-conference-2021-spring-office-hour/platform/db/db"
	"github.com/mercari/go-conference-2021-spring-office-hour/platform/db/proto"
)

var _ proto.DBServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedDBServiceServer
	db db.DB
}

func (s *server) CreateCustomer(ctx context.Context, req *proto.CreateCustomerRequest) (*proto.CreateCustomerResponse, error) {
	c, err := s.db.CreateCustomer(ctx, req.Name)
	if err != nil {
		if errors.Is(err, db.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.CreateCustomerResponse{
		Customer: &proto.Customer{
			Id:   c.ID,
			Name: c.Name,
		},
	}, nil
}

func (s *server) GetCustomer(ctx context.Context, req *proto.GetCustomerRequest) (*proto.GetCustomerResponse, error) {
	c, err := s.db.GetCustomer(ctx, req.Id)
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

func (s *server) GetCustomerByName(ctx context.Context, req *proto.GetCustomerByNameRequest) (*proto.GetCustomerByNameResponse, error) {
	c, err := s.db.GetCustomerByName(ctx, req.Name)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.GetCustomerByNameResponse{
		Customer: &proto.Customer{
			Id:   c.ID,
			Name: c.Name,
		},
	}, nil
}

func (s *server) GetItem(ctx context.Context, req *proto.GetItemRequest) (*proto.GetItemResponse, error) {
	i, err := s.db.GetItem(ctx, req.Id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.GetItemResponse{
		Item: &proto.Item{
			Id:         i.ID,
			CustomerId: i.CustomerID,
			Title:      i.Title,
			Price:      int64(i.Price),
		},
	}, nil
}
