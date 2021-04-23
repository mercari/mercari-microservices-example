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
