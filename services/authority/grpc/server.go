package grpc

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/json"
	"encoding/pem"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/go-logr/logr"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"

	grpccontext "github.com/mercari/mercari-microservices-example/pkg/grpc/context"
	"github.com/mercari/mercari-microservices-example/services/authority/proto"
	"github.com/mercari/mercari-microservices-example/services/authority/proto/protoconnect"
	customer "github.com/mercari/mercari-microservices-example/services/customer/proto"
	customerconnect "github.com/mercari/mercari-microservices-example/services/customer/proto/protoconnect"
)

var (
	rsaPrivateKey *rsa.PrivateKey
	_             protoconnect.AuthorityServiceHandler = (*server)(nil)
)

//go:embed private-key.pem
var privateKeyFile []byte

const (
	issuer = "authority"
	kid    = "aa7c6287-c45d-4966-84b4-a1633e4e3a64"
)

type server struct {
	protoconnect.UnimplementedAuthorityServiceHandler

	customerClient customerconnect.CustomerServiceClient
	logger         logr.Logger
}

func (s *server) Signup(ctx context.Context, req *connect.Request[proto.SignupRequest]) (*connect.Response[proto.SignupResponse], error) {
	c, err := s.customerClient.CreateCustomer(ctx, connect.NewRequest(&customer.CreateCustomerRequest{Name: req.Msg.Name}))
	if err != nil {
		if connect.CodeOf(err)  == connect.CodeAlreadyExists{
			return nil, connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("customer already exisits by given name: %w", err))
		}
		s.log(ctx).Error(err, "failed to call customer service")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("internal error: %w", err))
	}

	return connect.NewResponse(&proto.SignupResponse{
		Customer: &customer.Customer{
			Id:   c.Msg.GetCustomer().Id,
			Name: c.Msg.GetCustomer().Name,
		},
	}), nil
}

func (s *server) Signin(ctx context.Context, req *connect.Request[proto.SigninRequest]) (*connect.Response[proto.SigninResponse], error) {
	res, err := s.customerClient.GetCustomerByName(ctx, connect.NewRequest(&customer.GetCustomerByNameRequest{Name: req.Msg.Name}))
	if err != nil {
		s.log(ctx).Info(fmt.Sprintf("failed to authenticate the customer: %s", err))
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("unauthenticated: %w", err))
	}

	token, err := createAccessToken(res.Msg.GetCustomer().Id)
	if err != nil {
		s.log(ctx).Error(err, "failed to create the access token")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create access token: %w", err))
	}

	return connect.NewResponse(&proto.SigninResponse{
		AccessToken: string(token),
	}), nil
}

func (s *server) ListPublicKeys(ctx context.Context, _ *connect.Request[proto.ListPublicKeysRequest]) (*connect.Response[proto.ListPublicKeysResponse], error) {
	key, err := jwk.New(rsaPrivateKey.PublicKey)
	if err != nil {
		s.log(ctx).Error(err, "failed to create jwk")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create jwks: %w", err))
	}

	if err = key.Set(jws.KeyIDKey, kid); err != nil {
		s.log(ctx).Error(err, "failed to set the kid to the jwk")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create jwks: %w", err))
	}

	if err = key.Set(jws.AlgorithmKey, jwa.RS256); err != nil {
		s.log(ctx).Error(err, "failed to set the alg to the jwk")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create jwks: %w", err))
	}

	set := jwk.NewSet()
	set.Add(key)

	buf, err := json.Marshal(set)
	if err != nil {
		s.log(ctx).Error(err, "failed to marshal the jwk")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create jwks: %w", err))
	}

	return connect.NewResponse(&proto.ListPublicKeysResponse{Jwks: string(buf)}), nil
}

func createAccessToken(sub string) ([]byte, error) {
	token := jwt.New()

	if err := token.Set(jwt.IssuerKey, issuer); err != nil {
		return nil, fmt.Errorf("failed to set the issuer key to the token: %w", err)
	}

	if err := token.Set(jwt.SubjectKey, sub); err != nil {
		return nil, fmt.Errorf("failed to set the subject key to the token: %w", err)
	}

	headers := jws.NewHeaders()
	if err := headers.Set(jws.KeyIDKey, kid); err != nil {
		return nil, fmt.Errorf("failed to create jws headers: %w", err)
	}

	if err := headers.Set(jws.AlgorithmKey, jwa.RS256); err != nil {
		return nil, fmt.Errorf("failed to set the alg key to the token: %w", err)
	}

	if err := headers.Set(jws.TypeKey, "JWT"); err != nil {
		return nil, fmt.Errorf("failed to set the typ key to the token: %w", err)
	}

	b, err := json.Marshal(token)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the token: %w", err)
	}

	signedToken, err := jws.Sign(b, jwa.RS256, rsaPrivateKey, jws.WithHeaders(headers))
	if err != nil {
		return nil, fmt.Errorf("failed to sign the token: %w", err)
	}
	return signedToken, nil
}

func (s *server) log(ctx context.Context) logr.Logger {
	reqid := grpccontext.GetRequestID(ctx)

	return s.logger.WithValues("request_id", reqid)
}

func init() {
	block, _ := pem.Decode(privateKeyFile)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Sprintf("failed to parse private key: %s", err))
	}

	rsaPrivateKey = key
}
