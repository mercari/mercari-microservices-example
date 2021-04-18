package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mercari/go-conference-2021-spring-office-hour/services/catalog/proto"
	customer "github.com/mercari/go-conference-2021-spring-office-hour/services/customer/proto"
	item "github.com/mercari/go-conference-2021-spring-office-hour/services/item/proto"
)

var _ proto.CatalogServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedCatalogServiceServer
	itemClient     item.ItemServiceClient
	customerClient customer.CustomerServiceClient
}

func (s *server) GetItem(ctx context.Context, req *proto.GetItemRequest) (*proto.GetItemResponse, error) {
	ires, err := s.itemClient.GetItem(ctx, &item.GetItemRequest{Id: req.Id})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	i := ires.GetItem()
	if i == nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	cres, err := s.customerClient.GetCustomer(ctx, &customer.GetCustomerRequest{Id: i.CustomerId})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	c := cres.GetCustomer()
	if c == nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.GetItemResponse{
		Item: &proto.Item{
			Id:           i.Id,
			CustomerId:   i.CustomerId,
			CustomerName: c.Name,
			Title:        i.Title,
			Price:        int64(i.Price),
		},
	}, nil
}
