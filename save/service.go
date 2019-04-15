package auth

import (
	context "context"
	"errors"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUnknown      = errors.New("Unknown")
)

// Define service interface
type Service interface {
	Validate(ctx context.Context, req ValidateRequest) (*ValidateResponse, error)
	Register(ctx context.Context, req RegisterRequest) (*RegisterResponse, error)
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
}

type AuthService struct {
}

func (AuthService) Validate(ctx context.Context, req ValidateRequest) (ValidateResponse, error) {
	if len(req.Token) > 0 && req.Token[0] == 'a' {
		return ValidateResponse{
			UserId:  "acasajus",
			IsAdmin: true,
		}, nil
	}
	return ValidateResponse{}, ErrUnauthorized
}

func (AuthService) Register(ctx context.Context, req RegisterRequest) (RegisterResponse, error) {
	if len(req.Username) > 0 {
		return RegisterResponse{
			Ok: true,
		}, nil
	}
	return RegisterResponse{}, ErrUnknown
}

func (AuthService) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	return LoginResponse{Token: "aa"}, nil
}
