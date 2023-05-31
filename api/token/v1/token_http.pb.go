// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.17.3
// source: token/v1/token.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTokenGetToken = "/token.v1.Token/GetToken"
const OperationTokenListTokens = "/token.v1.Token/ListTokens"

type TokenHTTPServer interface {
	GetToken(context.Context, *GetTokenRequest) (*TokenReply, error)
	ListTokens(context.Context, *ListTokenRequest) (*ListTokenReply, error)
}

func RegisterTokenHTTPServer(s *http.Server, srv TokenHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/tokens/{tick}/{token_id}", _Token_GetToken0_HTTP_Handler(srv))
	r.GET("/v1/tokens", _Token_ListTokens0_HTTP_Handler(srv))
}

func _Token_GetToken0_HTTP_Handler(srv TokenHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTokenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTokenGetToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetToken(ctx, req.(*GetTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TokenReply)
		return ctx.Result(200, reply)
	}
}

func _Token_ListTokens0_HTTP_Handler(srv TokenHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTokenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTokenListTokens)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTokens(ctx, req.(*ListTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTokenReply)
		return ctx.Result(200, reply)
	}
}

type TokenHTTPClient interface {
	GetToken(ctx context.Context, req *GetTokenRequest, opts ...http.CallOption) (rsp *TokenReply, err error)
	ListTokens(ctx context.Context, req *ListTokenRequest, opts ...http.CallOption) (rsp *ListTokenReply, err error)
}

type TokenHTTPClientImpl struct {
	cc *http.Client
}

func NewTokenHTTPClient(client *http.Client) TokenHTTPClient {
	return &TokenHTTPClientImpl{client}
}

func (c *TokenHTTPClientImpl) GetToken(ctx context.Context, in *GetTokenRequest, opts ...http.CallOption) (*TokenReply, error) {
	var out TokenReply
	pattern := "/v1/tokens/{tick}/{token_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTokenGetToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TokenHTTPClientImpl) ListTokens(ctx context.Context, in *ListTokenRequest, opts ...http.CallOption) (*ListTokenReply, error) {
	var out ListTokenReply
	pattern := "/v1/tokens"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTokenListTokens))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}