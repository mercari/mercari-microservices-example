package grpc

import (
	"context"

	"github.com/mercari/go-conference-2021-spring-office-hour/services/category/proto"
)

var _ proto.CategoryServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedCategoryServiceServer
}

func (s *server) ListCategories(ctx context.Context, req *proto.ListCategoriesRequest) (*proto.ListCategoriesResponse, error) {
	return &proto.ListCategoriesResponse{
		Categories: []*proto.Category{
			{
				Id:   "9e893b41-65e5-4ab0-94dd-bb7f6cd06da4",
				Name: "Books",
			},
			{
				Id:   "35c51c6f-7c5b-4bb3-9093-3acce1ccf3e5",
				Name: "Gadgets",
			},
			{
				Id:   "3d450ed6-94ac-4f34-9f24-1b934735c06c",
				Name: "Toys",
			},
		},
	}, nil
}
