package service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/proto/secret"
	"github.com/jdxj/sign/internal/secret/model"
)

var (
	ErrInvalidKey = errors.New("invalid key")
)

func New(conf config.Secret) *Service {
	if conf.Key == "" {
		panic(ErrInvalidKey)
	}
	srv := &Service{
		key: conf.Key,
	}
	return srv
}

type Service struct {
	secret.UnimplementedSecretServiceServer

	key string
}

func (srv *Service) CreateSecret(ctx context.Context, req *secret.CreateSecretReq) (*secret.CreateSecretRsp, error) {
	if req.UserID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty user id")
	}
	if req.Domain == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty domain")
	}
	if req.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty key")
	}
	return model.CreateSecret(srv.key, req)
}

func (srv *Service) GetSecret(ctx context.Context, req *secret.GetSecretReq) (*secret.GetSecretRsp, error) {
	if req.SecretID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty secret id")
	}
	if req.UserID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty user id")
	}
	return model.GetSecret(srv.key, req)
}

func (srv *Service) GetSecretList(ctx context.Context, req *secret.GetSecretListReq) (*secret.GetSecretListRsp, error) {
	if req.UserID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty user id")
	}
	return model.GetSecretList(srv.key, req)
}

func (srv *Service) UpdateSecret(ctx context.Context, req *secret.UpdateSecretReq) (*emptypb.Empty, error) {
	if req.SecretID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty secret id")
	}
	if req.Key == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty key")
	}
	return &emptypb.Empty{}, model.UpdateSecret(req)
}

func (srv *Service) DeleteSecret(ctx context.Context, req *secret.DeleteSecretReq) (*emptypb.Empty, error) {
	if req.SecretID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty secret id")
	}
	if req.UserID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty user id")
	}
	return &emptypb.Empty{}, model.DeleteSecret(req)
}
