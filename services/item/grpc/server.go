package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	db "github.com/mercari/go-conference-2021-spring-office-hour/platform/db/proto"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/item/proto"
)

var _ proto.ItemServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedItemServiceServer

	dbClient db.DBServiceClient
}

func (s *server) CreateItem(ctx context.Context, req *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	res, err := s.dbClient.CreateItem(ctx, &db.CreateItemRequest{
		CustomerId: req.CustomerId,
		Title:      req.Title,
		Price:      req.Price,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	item := res.GetItem()

	return &proto.CreateItemResponse{
		Item: &proto.Item{
			Id:         item.Id,
			CustomerId: item.CustomerId,
			Title:      item.Title,
			Price:      item.Price,
		},
	}, nil
}

func (s *server) GetItem(ctx context.Context, req *proto.GetItemRequest) (*proto.GetItemResponse, error) {
	res, err := s.dbClient.GetItem(ctx, &db.GetItemRequest{Id: req.Id})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	item := res.GetItem()

	return &proto.GetItemResponse{
		Item: &proto.Item{
			Id:         item.Id,
			CustomerId: item.CustomerId,
			Title:      item.Title,
			Price:      int64(item.Price),
		},
	}, nil
}

func (s *server) ListItems(ctx context.Context, req *proto.ListItemsRequest) (*proto.ListItemsResponse, error) {
	result, err := s.dbClient.ListItems(ctx, &db.ListItemsRequest{})
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	res := &proto.ListItemsResponse{
		Items: make([]*proto.Item, len(result.Items)),
	}

	for i, item := range result.Items {
		res.Items[i] = &proto.Item{
			Id:         item.Id,
			CustomerId: item.CustomerId,
			Title:      item.Title,
			Price:      item.Price,
		}
	}

	return res, nil
}
