package developers

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/giornetta/devcv/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/status"

	"github.com/giornetta/devcv/proto"
)

type authService struct {
	next proto.DeveloperServiceServer
	auth *auth.Service
}

// NewAuthenticator returns an implementation of devcv.DeveloperService that will handle authentication for specific methods
// before calling the real service
func NewAuthenticator(svc proto.DeveloperServiceServer, auth *auth.Service) proto.DeveloperServiceServer {
	return &authService{
		next: svc,
		auth: auth,
	}
}

// Login implements proto.DeveloperServiceServer
func (s *authService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.Token, error) {
	_, err := s.next.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	token, err := s.auth.GenerateToken(req.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not generate token")
	}
	return &proto.Token{
		Token: token,
	}, nil
}

// Lookup implements proto.DeveloperServiceServer
func (s *authService) Get(ctx context.Context, req *proto.UsernameRequest) (*proto.Developer, error) {
	return s.next.Get(ctx, req)
}

// Register implements proto.DeveloperServiceServer
func (s *authService) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.Token, error) {
	_, err := s.next.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	token, err := s.auth.GenerateToken(req.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not generate token")
	}
	return &proto.Token{
		Token: token,
	}, nil
}

// Update implements devcv.DeveloperService
func (s *authService) Update(ctx context.Context, req *proto.Developer) (*empty.Empty, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "no token")
	}

	if len(meta["authorization"]) != 1 {
		return nil, grpc.Errorf(codes.Unauthenticated, "no token")
	}

	if err := s.auth.Authenticate(meta["authorization"][0], req.Username); err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid token")
	}

	return s.next.Update(ctx, req)
}

// Delete implements devcv.DeveloperService
func (s *authService) Delete(ctx context.Context, req *proto.UsernameRequest) (*empty.Empty, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "no token")
	}

	if len(meta["authorization"]) != 1 {
		return nil, grpc.Errorf(codes.Unauthenticated, "no token")
	}

	if err := s.auth.Authenticate(meta["authorization"][0], req.Username); err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid token")
	}

	return s.next.Delete(ctx, req)
}
