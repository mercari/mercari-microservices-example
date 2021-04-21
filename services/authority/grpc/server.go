package grpc

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpccontext "github.com/mercari/go-conference-2021-spring-office-hour/pkg/grpc/context"
	"github.com/mercari/go-conference-2021-spring-office-hour/services/authority/proto"
	customer "github.com/mercari/go-conference-2021-spring-office-hour/services/customer/proto"
)

var (
	rsaPrivateKey *rsa.PrivateKey
	_             proto.AuthorityServiceServer = (*server)(nil)
)

const (
	issuer = "authority"
	kid    = "aa7c6287-c45d-4966-84b4-a1633e4e3a64"
)

type server struct {
	proto.UnimplementedAuthorityServiceServer

	customerClient customer.CustomerServiceClient
	logger         logr.Logger
}

func (s *server) Signin(ctx context.Context, req *proto.SigninRequest) (*proto.SigninResponse, error) {
	res, err := s.customerClient.GetCustomerByName(ctx, &customer.GetCustomerByNameRequest{Name: req.Name})
	if err != nil {
		s.log(ctx).Info(fmt.Sprintf("failed to authenticate the customer: %s", err))
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	token := jwt.New()

	if err = token.Set(jwt.IssuerKey, issuer); err != nil {
		s.log(ctx).Error(err, "failed to set issuer key to the token")
		return nil, status.Error(codes.Internal, "failed to create access token")
	}

	if err = token.Set(jwt.SubjectKey, res.GetCustomer().Id); err != nil {
		s.log(ctx).Error(err, "failed to set subject key to the token")
		return nil, status.Error(codes.Internal, "failed to create access token")
	}

	headers := jws.NewHeaders()
	if err = headers.Set(jws.KeyIDKey, kid); err != nil {
		s.log(ctx).Error(err, "failed to create jws header")
		return nil, status.Error(codes.Internal, "failed to create access token")
	}

	if err = headers.Set(jws.AlgorithmKey, jwa.RS256.String()); err != nil {
		s.log(ctx).Error(err, "failed to set alg key to the token")
		return nil, status.Error(codes.Internal, "failed to create access token")
	}

	if err = headers.Set(jws.TypeKey, "JWT"); err != nil {
		s.log(ctx).Error(err, "failed to set typ key to the token")
		return nil, status.Error(codes.Internal, "failed to create access token")
	}

	b, err := json.Marshal(token)
	if err != nil {
		s.log(ctx).Error(err, "failed to marshal the token to json")
		return nil, status.Error(codes.Internal, "failed to create access token")
	}

	signedToken, err := jws.Sign(b, jwa.RS256, rsaPrivateKey, jws.WithHeaders(headers))
	if err != nil {
		s.log(ctx).Error(err, "failed to sing the token")
		return nil, status.Error(codes.Internal, "failed to create access token")
	}

	return &proto.SigninResponse{
		AccessToken: string(signedToken),
	}, nil
}

func (s *server) log(ctx context.Context) logr.Logger {
	reqid := grpccontext.GetRequestID(ctx)

	return s.logger.WithValues("request_id", reqid)
}

func init() {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(fmt.Sprintf("failed to generate private key: %s", err))
	}
	rsaPrivateKey = key
}
