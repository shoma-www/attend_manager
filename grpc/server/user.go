package server

import (
	"context"

	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/proto"
	pb "github.com/shoma-www/attend_manager/grpc/proto"
	"github.com/shoma-www/attend_manager/grpc/service"
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
	var message string
	status := pb.RegisterStatus_SUCCESS
	_, err := u.us.Register(ctx, req.UserId, req.Password)
	if err != nil {
		message = err.Error()
		status = pb.RegisterStatus_ERROR
		u.logger.WithUUID(ctx).Error("register err: %s", err)
	}
	return &pb.RegisterResponse{
		Status:       status,
		ErrorMessage: message,
	}, nil
}
