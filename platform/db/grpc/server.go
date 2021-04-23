package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mercari/go-conference-2021-spring-office-hour/platform/db/db"
	"github.com/mercari/go-conference-2021-spring-office-hour/platform/db/model"
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

func (s *server) CreateItem(ctx context.Context, req *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	item := &model.Item{
		CustomerID: req.CustomerId,
		Title:      req.Title,
		Price:      uint64(req.Price),
	}

	i, err := s.db.CreateItem(ctx, item)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.CreateItemResponse{
		Item: &proto.Item{
			Id:         i.ID,
			CustomerId: i.CustomerID,
			Title:      i.Title,
			Price:      int64(i.Price),
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

func (s *server) ListItems(ctx context.Context, req *proto.ListItemsRequest) (*proto.ListItemsResponse, error) {
	items, err := s.db.GetAllItems(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	res := &proto.ListItemsResponse{
		Items: make([]*proto.Item, len(items)),
	}

	for i, item := range items {
		res.Items[i] = &proto.Item{
			Id:         item.ID,
			CustomerId: item.CustomerID,
			Title:      item.Title,
			Price:      int64(item.Price),
		}
	}

	return res, nil
}
