package server

import (
	"context"

	"github.com/rs/xid"
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/entity"
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

func (u *user) Register(ctx context.Context, req *proto.UserRegisterRequesut) (*proto.UserRegisterResponse, error) {
	groupID, err := xid.FromString(req.GetAttendanceGroupId())
	if err != nil {
		u.logger.WithUUID(ctx).Error("invalid group id: %s", req.GetAttendanceGroupId())
		status := status.New(codes.Internal, err.Error())
		return nil, status.Err()
	}

	if _, err := u.us.Register(ctx, groupID, req.LoginId, req.Password, req.Name); err != nil {
		if err == entity.ErrDuplicatedUser {
			return &pb.UserRegisterResponse{
				Status: pb.UserRegisterStatus_ERROR,
			}, nil
		}
		u.logger.WithUUID(ctx).Error("register errpr: %s", err)
		status := status.New(codes.Internal, err.Error())
		return nil, status.Err()
	}

	res := &pb.UserRegisterResponse{
		Status: pb.UserRegisterStatus_SUCCESS,
	}

	return res, nil
}
