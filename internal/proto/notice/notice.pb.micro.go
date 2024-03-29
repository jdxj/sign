// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: internal/proto/notice/notice.proto

package notice

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for NoticeService service

func NewNoticeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for NoticeService service

type NoticeService interface {
	SendNotice(ctx context.Context, in *SendNoticeRequest, opts ...client.CallOption) (*emptypb.Empty, error)
}

type noticeService struct {
	c    client.Client
	name string
}

func NewNoticeService(name string, c client.Client) NoticeService {
	return &noticeService{
		c:    c,
		name: name,
	}
}

func (c *noticeService) SendNotice(ctx context.Context, in *SendNoticeRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "NoticeService.SendNotice", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NoticeService service

type NoticeServiceHandler interface {
	SendNotice(context.Context, *SendNoticeRequest, *emptypb.Empty) error
}

func RegisterNoticeServiceHandler(s server.Server, hdlr NoticeServiceHandler, opts ...server.HandlerOption) error {
	type noticeService interface {
		SendNotice(ctx context.Context, in *SendNoticeRequest, out *emptypb.Empty) error
	}
	type NoticeService struct {
		noticeService
	}
	h := &noticeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&NoticeService{h}, opts...))
}

type noticeServiceHandler struct {
	NoticeServiceHandler
}

func (h *noticeServiceHandler) SendNotice(ctx context.Context, in *SendNoticeRequest, out *emptypb.Empty) error {
	return h.NoticeServiceHandler.SendNotice(ctx, in, out)
}
