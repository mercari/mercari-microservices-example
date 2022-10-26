package grpc

import (
	"bytes"
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/go-logr/logr"
	auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/lestrrat-go/jwx/jwt"

	"github.com/mercari/mercari-microservices-example/services/catalog/proto"
	"github.com/mercari/mercari-microservices-example/services/catalog/proto/protoconnect"
	customer "github.com/mercari/mercari-microservices-example/services/customer/proto"
	customerconnect "github.com/mercari/mercari-microservices-example/services/customer/proto/protoconnect"
	item "github.com/mercari/mercari-microservices-example/services/item/proto"
	itemconnect "github.com/mercari/mercari-microservices-example/services/item/proto/protoconnect"
)

var _ protoconnect.CatalogServiceHandler = (*server)(nil)

type server struct {
	protoconnect.UnimplementedCatalogServiceHandler
	itemClient     itemconnect.ItemServiceClient
	customerClient customerconnect.CustomerServiceClient
	logger         logr.Logger
}

func (s *server) CreateItem(ctx context.Context, req *connect.Request[proto.CreateItemRequest]) (*connect.Response[proto.CreateItemResponse], error) {
	tokenStr, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("token not found: %w", err))
	}

	token, err := jwt.Parse(bytes.NewBufferString(tokenStr).Bytes())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("failed to parse access token: %w", err))
	}

	r := connect.NewRequest(&item.CreateItemRequest{
		CustomerId: token.Subject(),
		Title:      req.Msg.Title,
		Price:      req.Msg.Price,
	})
	res, err := s.itemClient.CreateItem(ctx, r)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to call crate item api: %w", err))
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
	r := connect.NewRequest(&item.GetItemRequest{
		Id: req.Msg.Id,
	})
	ires, err := s.itemClient.GetItem(ctx, r)
	if err != nil {
		if connect.CodeOf(err) == connect.CodeNotFound {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("not found: %w", err))
		}

		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	i := ires.Msg.GetItem()
	if i == nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	cres, err := s.customerClient.GetCustomer(ctx, connect.NewRequest(&customer.GetCustomerRequest{Id: i.CustomerId}))
	if err != nil {
		if connect.CodeOf(err) == connect.CodeNotFound {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("not found: %w", err))
		}

		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	c := cres.Msg.GetCustomer()
	if c == nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	return connect.NewResponse(&proto.GetItemResponse{
		Item: &proto.Item{
			Id:           i.Id,
			CustomerId:   i.CustomerId,
			CustomerName: c.Name,
			Title:        i.Title,
			Price:        int64(i.Price),
		},
	}), nil
}

func (s *server) ListItems(ctx context.Context, req *connect.Request[proto.ListItemsRequest]) (*connect.Response[proto.ListItemsResponse], error) {
	result, err := s.itemClient.ListItems(ctx, connect.NewRequest(&item.ListItemsRequest{}))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	res := &proto.ListItemsResponse{
		Items: make([]*proto.Item, len(result.Msg.Items)),
	}

	for i, item := range result.Msg.Items {
		res.Items[i] = &proto.Item{
			Id:    item.Id,
			Title: item.Title,
			Price: item.Price,
		}
	}

	return connect.NewResponse(res), nil
}
