package developers

import (
	"context"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/giornetta/devcv/proto"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/giornetta/devcv/valid"

	"github.com/giornetta/devcv/devcv"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	developers devcv.DeveloperRepository
}

// New returns an implementation of devcv.DeveloperService
func New(dr devcv.DeveloperRepository) proto.DeveloperServiceServer {
	return &service{
		developers: dr,
	}
}

// Login implements proto.DeveloperServiceServer
func (s *service) Login(ctx context.Context, req *proto.LoginRequest) (*proto.Token, error) {
	hash, err := s.developers.GetHash(req.Username)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "wrong username and password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "wrong username and password")
	}

	return &proto.Token{}, nil
}

// Lookup implements proto.DeveloperServiceServer
func (s *service) Get(ctx context.Context, req *proto.UsernameRequest) (*proto.Developer, error) {
	if !valid.Username(req.Username) {
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}

	dev, err := s.developers.Lookup(req.Username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return toProto(dev), nil
}

// Register implements proto.DeveloperServiceServer
func (s *service) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.Token, error) {
	if !valid.Username(req.Username) {
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}

	if len(req.FirstName) <= 0 || len(req.LastName) <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid name")
	}

	if !valid.Password(req.Password) {
		return nil, status.Error(codes.InvalidArgument, "invalid password")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	err := s.developers.Create(req.Username, req.FirstName, req.LastName, string(hashed))
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "user with same credentials already exists")
	}

	return &proto.Token{}, nil
}

// Update implements devcv.DeveloperService
func (s *service) Update(ctx context.Context, req *proto.Developer) (*empty.Empty, error) {
	if !valid.Username(req.Username) {
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}

	err := s.developers.Update(toDeveloper(req))
	if err != nil {
		return nil, status.Error(codes.Internal, "could not update")
	}

	return &empty.Empty{}, nil
}

// Delete implements devcv.DeveloperService
func (s *service) Delete(ctx context.Context, req *proto.UsernameRequest) (*empty.Empty, error) {
	if !valid.Username(req.Username) {
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}

	err := s.developers.Delete(req.Username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &empty.Empty{}, nil
}
