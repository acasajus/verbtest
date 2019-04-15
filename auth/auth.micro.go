// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: auth/auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth/auth.proto

It has these top-level messages:
	ValidateRequest
	ValidateResponse
	RegisterRequest
	RegisterResponse
	LoginRequest
	LoginResponse
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthService interface {
	Validate(ctx context.Context, in *ValidateRequest, opts ...client.CallOption) (*ValidateResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "auth"
	}
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Validate(ctx context.Context, in *ValidateRequest, opts ...client.CallOption) (*ValidateResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Validate", in)
	out := new(ValidateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Validate(context.Context, *ValidateRequest, *ValidateResponse) error
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		Validate(ctx context.Context, in *ValidateRequest, out *ValidateResponse) error
		Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Validate(ctx context.Context, in *ValidateRequest, out *ValidateResponse) error {
	return h.AuthHandler.Validate(ctx, in, out)
}

func (h *authHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.AuthHandler.Register(ctx, in, out)
}

func (h *authHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.AuthHandler.Login(ctx, in, out)
}
