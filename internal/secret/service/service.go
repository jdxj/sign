package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jdxj/sign/internal/proto/secret"
	"github.com/jdxj/sign/internal/secret/model"
)

type Service struct {
	secret.UnimplementedSecretServiceServer
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
	return model.CreateSecret(req)
}

func (srv *Service) GetSecret(ctx context.Context, req *secret.GetSecretReq) (*secret.GetSecretRsp, error) {
	if req.SecretID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty secret id")
	}
	return model.GetSecret(req)
}

func (srv *Service) UpdateSecret(ctx context.Context, req *secret.UpdateSecretReq) (*emptypb.Empty, error) {
	if req.SecretID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty secret id")
	}
	return &emptypb.Empty{}, model.UpdateSecret(req)
}

func (srv *Service) DeleteSecret(ctx context.Context, req *secret.DeleteSecretReq) (*emptypb.Empty, error) {
	if req.SecretID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty secret id")
	}
	return &emptypb.Empty{}, model.DeleteSecret(req)
}
