package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mercari/go-conference-2021-spring-office-hour/services/item/db"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/item/proto"
)

var _ proto.ItemServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedItemServiceServer
	db db.DB
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
