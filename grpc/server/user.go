package server

import (
	"context"

	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/proto"
	pb "github.com/shoma-www/attend_manager/grpc/proto"
)

type user struct {
	logger core.Logger
}

// NewUser コンストラクタ
func NewUser(l core.Logger) pb.UserServer {
	return &user{logger: l}
}

func (u *user) Register(ctx context.Context, req *proto.RegisterRequesut) (*proto.RegisterResponse, error) {
	u.logger.WithUUID(ctx).Info("Hello World!")
	u.logger.WithUUID(ctx).Info("%v", req)
	return &pb.RegisterResponse{
		Status:       pb.RegisterStatus_SUCCESS,
		ErrorMessage: "",
	}, nil
}
