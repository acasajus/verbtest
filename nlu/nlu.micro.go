// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: nlu/nlu.proto

/*
Package nlu is a generated protocol buffer package.

It is generated from these files:
	nlu/nlu.proto

It has these top-level messages:
	ProcessRequest
	ProcessResponse
*/
package nlu

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

// Client API for NLU service

type NLUService interface {
	Process(ctx context.Context, in *ProcessRequest, opts ...client.CallOption) (*ProcessResponse, error)
}

type nLUService struct {
	c    client.Client
	name string
}

func NewNLUService(name string, c client.Client) NLUService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "nlu"
	}
	return &nLUService{
		c:    c,
		name: name,
	}
}

func (c *nLUService) Process(ctx context.Context, in *ProcessRequest, opts ...client.CallOption) (*ProcessResponse, error) {
	req := c.c.NewRequest(c.name, "NLU.Process", in)
	out := new(ProcessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NLU service

type NLUHandler interface {
	Process(context.Context, *ProcessRequest, *ProcessResponse) error
}

func RegisterNLUHandler(s server.Server, hdlr NLUHandler, opts ...server.HandlerOption) error {
	type nLU interface {
		Process(ctx context.Context, in *ProcessRequest, out *ProcessResponse) error
	}
	type NLU struct {
		nLU
	}
	h := &nLUHandler{hdlr}
	return s.Handle(s.NewHandler(&NLU{h}, opts...))
}

type nLUHandler struct {
	NLUHandler
}

func (h *nLUHandler) Process(ctx context.Context, in *ProcessRequest, out *ProcessResponse) error {
	return h.NLUHandler.Process(ctx, in, out)
}
