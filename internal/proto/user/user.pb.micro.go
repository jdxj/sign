// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: internal/proto/user/user.proto

package user

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

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	AuthUser(ctx context.Context, in *AuthUserRequest, opts ...client.CallOption) (*AuthUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*CreateUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*emptypb.Empty, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) AuthUser(ctx context.Context, in *AuthUserRequest, opts ...client.CallOption) (*AuthUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.AuthUser", in)
	out := new(AuthUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*CreateUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.CreateUser", in)
	out := new(CreateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUser(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*GetUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetUser", in)
	out := new(GetUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateUser", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	AuthUser(context.Context, *AuthUserRequest, *AuthUserResponse) error
	CreateUser(context.Context, *CreateUserRequest, *CreateUserResponse) error
	GetUser(context.Context, *GetUserRequest, *GetUserResponse) error
	UpdateUser(context.Context, *UpdateUserRequest, *emptypb.Empty) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		AuthUser(ctx context.Context, in *AuthUserRequest, out *AuthUserResponse) error
		CreateUser(ctx context.Context, in *CreateUserRequest, out *CreateUserResponse) error
		GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error
		UpdateUser(ctx context.Context, in *UpdateUserRequest, out *emptypb.Empty) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) AuthUser(ctx context.Context, in *AuthUserRequest, out *AuthUserResponse) error {
	return h.UserServiceHandler.AuthUser(ctx, in, out)
}

func (h *userServiceHandler) CreateUser(ctx context.Context, in *CreateUserRequest, out *CreateUserResponse) error {
	return h.UserServiceHandler.CreateUser(ctx, in, out)
}

func (h *userServiceHandler) GetUser(ctx context.Context, in *GetUserRequest, out *GetUserResponse) error {
	return h.UserServiceHandler.GetUser(ctx, in, out)
}

func (h *userServiceHandler) UpdateUser(ctx context.Context, in *UpdateUserRequest, out *emptypb.Empty) error {
	return h.UserServiceHandler.UpdateUser(ctx, in, out)
}
