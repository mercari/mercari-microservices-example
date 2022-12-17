package grpc

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"

	db "github.com/mercari/mercari-microservices-example/platform/db/proto"
	dbconnect "github.com/mercari/mercari-microservices-example/platform/db/proto/protoconnect"
	"github.com/mercari/mercari-microservices-example/services/item/proto"
	"github.com/mercari/mercari-microservices-example/services/item/proto/protoconnect"
)

var _ protoconnect.ItemServiceHandler = (*server)(nil)

type server struct {
	protoconnect.UnimplementedItemServiceHandler

	dbClient dbconnect.DBServiceClient
}

func (s *server) CreateItem(ctx context.Context, req *connect.Request[proto.CreateItemRequest]) (*connect.Response[proto.CreateItemResponse], error) {
	res, err := s.dbClient.CreateItem(ctx, connect.NewRequest(&db.CreateItemRequest{
		CustomerId: req.Msg.CustomerId,
		Title:      req.Msg.Title,
		Price:      req.Msg.Price,
	}))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	item := res.Msg.GetItem()

	return connect.NewResponse(&proto.CreateItemResponse{
		Item: &proto.Item{
			Id:         item.Id,
			CustomerId: item.CustomerId,
			Title:      item.Title,
			Price:      item.Price,
		},
	}), nil
}

func (s *server) GetItem(ctx context.Context, req *connect.Request[proto.GetItemRequest]) (*connect.Response[proto.GetItemResponse], error) {
	res, err := s.dbClient.GetItem(ctx, connect.NewRequest(&db.GetItemRequest{Id: req.Msg.Id}))
	if err != nil {
		if connect.CodeOf(err) == connect.CodeNotFound {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("not found: %w", err))
		}
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	item := res.Msg.GetItem()

	return connect.NewResponse(&proto.GetItemResponse{
		Item: &proto.Item{
			Id:         item.Id,
			CustomerId: item.CustomerId,
			Title:      item.Title,
			Price:      item.Price,
		},
	}), nil
}

func (s *server) ListItems(ctx context.Context, req *connect.Request[proto.ListItemsRequest]) (*connect.Response[proto.ListItemsResponse], error) {
	result, err := s.dbClient.ListItems(ctx, connect.NewRequest(&db.ListItemsRequest{}))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	res := &proto.ListItemsResponse{
		Items: make([]*proto.Item, len(result.Msg.Items)),
	}

	for i, item := range result.Msg.Items {
		res.Items[i] = &proto.Item{
			Id:         item.Id,
			CustomerId: item.CustomerId,
			Title:      item.Title,
			Price:      item.Price,
		}
	}

	return connect.NewResponse(res), nil
}
