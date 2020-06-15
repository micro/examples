// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: micro/examples/blog/posts/proto/posts/post.proto

package go_micro_service_posts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Posts service

func NewPostsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Posts service

type PostsService interface {
	// Query currently only supports read by slug or timestamp, no listing.
	Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error)
	Post(ctx context.Context, in *PostRequest, opts ...client.CallOption) (*PostResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type postsService struct {
	c    client.Client
	name string
}

func NewPostsService(name string, c client.Client) PostsService {
	return &postsService{
		c:    c,
		name: name,
	}
}

func (c *postsService) Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error) {
	req := c.c.NewRequest(c.name, "Posts.Query", in)
	out := new(QueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsService) Post(ctx context.Context, in *PostRequest, opts ...client.CallOption) (*PostResponse, error) {
	req := c.c.NewRequest(c.name, "Posts.Post", in)
	out := new(PostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Posts.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Posts service

type PostsHandler interface {
	// Query currently only supports read by slug or timestamp, no listing.
	Query(context.Context, *QueryRequest, *QueryResponse) error
	Post(context.Context, *PostRequest, *PostResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterPostsHandler(s server.Server, hdlr PostsHandler, opts ...server.HandlerOption) error {
	type posts interface {
		Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error
		Post(ctx context.Context, in *PostRequest, out *PostResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type Posts struct {
		posts
	}
	h := &postsHandler{hdlr}
	return s.Handle(s.NewHandler(&Posts{h}, opts...))
}

type postsHandler struct {
	PostsHandler
}

func (h *postsHandler) Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error {
	return h.PostsHandler.Query(ctx, in, out)
}

func (h *postsHandler) Post(ctx context.Context, in *PostRequest, out *PostResponse) error {
	return h.PostsHandler.Post(ctx, in, out)
}

func (h *postsHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.PostsHandler.Delete(ctx, in, out)
}
