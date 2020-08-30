package server

import (
	"context"

	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/proto"
	pb "github.com/shoma-www/attend_manager/grpc/proto"
	"github.com/shoma-www/attend_manager/grpc/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type user struct {
	logger core.Logger
	us     *service.User
}

// NewUser コンストラクタ
func NewUser(l core.Logger, us *service.User) pb.UserServer {
	return &user{logger: l, us: us}
}

func (u *user) Register(ctx context.Context, req *proto.RegisterRequesut) (*proto.RegisterResponse, error) {
	st := pb.RegisterStatus_SUCCESS
	_, err := u.us.Register(ctx, req.UserId, req.Password)
	if err != nil {
		u.logger.WithUUID(ctx).Error("register errpr: %s", err)
		status := status.New(codes.Internal, err.Error())
		return nil, status.Err()
	}

	res := &pb.RegisterResponse{
		Status: st,
	}

	return res, nil
}
