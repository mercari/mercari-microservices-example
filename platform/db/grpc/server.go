package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/bufbuild/connect-go"

	"github.com/mercari/mercari-microservices-example/platform/db/db"
	"github.com/mercari/mercari-microservices-example/platform/db/model"
	"github.com/mercari/mercari-microservices-example/platform/db/proto"
	"github.com/mercari/mercari-microservices-example/platform/db/proto/protoconnect"
)

var _ protoconnect.DBServiceHandler = (*server)(nil)

type server struct {
	protoconnect.UnimplementedDBServiceHandler
	db db.DB
}

func (s *server) CreateCustomer(ctx context.Context, req *connect.Request[proto.CreateCustomerRequest]) (*connect.Response[proto.CreateCustomerResponse], error) {
	c, err := s.db.CreateCustomer(ctx, req.Msg.Name)
	if err != nil {
		if errors.Is(err, db.ErrAlreadyExists) {
			return nil, connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("customer[%s] is already exists: %w", req.Msg.Name, err))
		}

		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	return connect.NewResponse(&proto.CreateCustomerResponse{
		Customer: &proto.Customer{
			Id:   c.ID,
			Name: c.Name,
		},
	}), nil
}

func (s *server) GetCustomer(ctx context.Context, req *connect.Request[proto.GetCustomerRequest]) (*connect.Response[proto.GetCustomerResponse], error) {
	c, err := s.db.GetCustomer(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("not found: %w", err))
		}

		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	return connect.NewResponse(&proto.GetCustomerResponse{
		Customer: &proto.Customer{
			Id:   c.ID,
			Name: c.Name,
		},
	}), nil
}

func (s *server) GetCustomerByName(ctx context.Context, req *connect.Request[proto.GetCustomerByNameRequest]) (*connect.Response[proto.GetCustomerByNameResponse], error) {
	c, err := s.db.GetCustomerByName(ctx, req.Msg.Name)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("not found: %w", err))
		}

		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	return connect.NewResponse(&proto.GetCustomerByNameResponse{
		Customer: &proto.Customer{
			Id:   c.ID,
			Name: c.Name,
		},
	}), nil
}

func (s *server) CreateItem(ctx context.Context, req *connect.Request[proto.CreateItemRequest]) (*connect.Response[proto.CreateItemResponse], error) {
	item := &model.Item{
		CustomerID: req.Msg.CustomerId,
		Title:      req.Msg.Title,
		Price:      uint64(req.Msg.Price),
	}

	i, err := s.db.CreateItem(ctx, item)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	return connect.NewResponse(&proto.CreateItemResponse{
		Item: &proto.Item{
			Id:         i.ID,
			CustomerId: i.CustomerID,
			Title:      i.Title,
			Price:      int64(i.Price),
		},
	}), nil
}

func (s *server) GetItem(ctx context.Context, req *connect.Request[proto.GetItemRequest]) (*connect.Response[proto.GetItemResponse], error) {
	i, err := s.db.GetItem(ctx, req.Msg.Id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("not found: %w", err))
		}

		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	return connect.NewResponse(&proto.GetItemResponse{
		Item: &proto.Item{
			Id:         i.ID,
			CustomerId: i.CustomerID,
			Title:      i.Title,
			Price:      int64(i.Price),
		},
	}), nil
}

func (s *server) ListItems(ctx context.Context, req *connect.Request[proto.ListItemsRequest]) (*connect.Response[proto.ListItemsResponse], error) {
	items, err := s.db.GetAllItems(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
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

	return connect.NewResponse(res), nil
}
