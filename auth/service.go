package auth

import (
	context "context"
	"errors"
	"log"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUnknown      = errors.New("Unknown")
)

type AuthSvc struct {
}

func (AuthSvc) Validate(ctx context.Context, req *ValidateRequest, resp *ValidateResponse) error {
	log.Printf("Got validate request with token %s", req.Token)
	if len(req.Token) > 0 && req.Token[0] == 'a' {
		resp.UserId = "acasajus"
		resp.IsAdmin = true
		return nil
	}
	return ErrUnauthorized
}

func (AuthSvc) Register(ctx context.Context, req *RegisterRequest, resp *RegisterResponse) error {
	if len(req.Username) > 0 {
		resp.Ok = true
		return nil
	}
	return ErrUnknown
}

func (AuthSvc) Login(ctx context.Context, req *LoginRequest, resp *LoginResponse) error {
	resp.Token = "aa"
	return nil
}
