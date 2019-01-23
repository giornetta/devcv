package developers

import (
	"errors"

	"github.com/giornetta/devcv/auth"
	"github.com/giornetta/devcv/devcv"

	"github.com/giornetta/devcv/valid"

	"golang.org/x/crypto/bcrypt"
)

// Service provides methods which can be accessed from the API
type Service interface {
	// Login will be moved out to an auth service
	Login(*LoginRequest) (*TokenResponse, error)
	Create(*CreateRequest) (*TokenResponse, error)
	Get(*UsernameRequest) (Developer, error)
	Update(Developer) error
	Delete(*UsernameRequest) error
}

// Developer is an alias for simplicity, some of its fields will stay
// unused in the case of Service.Update
type Developer *devcv.Developer

// LoginRequest contains the fields needed for Service.Login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateRequest contains the fields needed for Service.CreateRequest
type CreateRequest struct {
	Username  string
	FirstName string
	LastName  string
	Password  string
}

// TokenResponse contains the authorization token that will be sent in our methods after authenticating
type TokenResponse struct {
	Token string
}

// UsernameRequest contains only a field, an username
// could be a string alias but it's a struct just for consistency.
type UsernameRequest struct {
	Username string
}

type service struct {
	developers devcv.DeveloperRepository
	authSvc    auth.Service
}

// New returns an implementation of Service
func New(dr devcv.DeveloperRepository, authSvc auth.Service) Service {
	return &service{
		developers: dr,
		authSvc:    authSvc,
	}
}

// Login implements Service
func (s *service) Login(req *LoginRequest) (*TokenResponse, error) {
	hash, err := s.developers.GetHash(req.Username)
	if err != nil {
		return nil, errors.New("wrong username and password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)); err != nil {
		return nil, errors.New("wrong username and password")
	}

	token, err := s.authSvc.GenerateToken(req.Username)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{token}, nil
}

// Lookup implements Service
func (s *service) Get(req *UsernameRequest) (Developer, error) {
	if !valid.Username(req.Username) {
		return nil, errors.New("invalid username")
	}

	dev, err := s.developers.Lookup(req.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return dev, nil
}

// Create implements Service
func (s *service) Create(req *CreateRequest) (*TokenResponse, error) {
	if !valid.Username(req.Username) {
		return nil, errors.New("invalid username")
	}

	if len(req.FirstName) <= 0 || len(req.LastName) <= 0 {
		return nil, errors.New("invalid name")
	}

	if !valid.Password(req.Password) {
		return nil, errors.New("invalid password")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	err := s.developers.Create(req.Username, req.FirstName, req.LastName, string(hashed))
	if err != nil {
		return nil, errors.New("user with same credentials already exists")
	}

	token, err := s.authSvc.GenerateToken(req.Username)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{token}, nil
}

// Update implements Service
func (s *service) Update(req Developer) error {
	if !valid.Username(req.Username) {
		return errors.New("invalid username")
	}

	err := s.developers.Update(req)
	if err != nil {
		return errors.New("could not update")
	}

	return nil
}

// Delete implements Service
func (s *service) Delete(req *UsernameRequest) error {
	if !valid.Username(req.Username) {
		return errors.New("invalid username")
	}

	err := s.developers.Delete(req.Username)
	if err != nil {
		return err
	}

	return nil
}
