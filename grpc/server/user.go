package server

import (
	"context"

	"github.com/rs/xid"
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/entity"
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

func (u *user) Register(ctx context.Context, req *pb.UserRegisterRequesut) (*pb.UserRegisterResponse, error) {
	groupID, err := xid.FromString(req.GetAttendanceGroupId())
	if err != nil {
		u.logger.WithUUID(ctx).Error("invalid group id: %s", req.GetAttendanceGroupId())
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	if _, err := u.us.Register(ctx, groupID, req.LoginId, req.Password, req.Name); err != nil {
		if err == entity.ErrDuplicatedUser {
			return &pb.UserRegisterResponse{
				Status: pb.UserRegisterStatus_ERROR,
			}, nil
		}
		u.logger.WithUUID(ctx).Error("register errpr: %s", err)
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	res := &pb.UserRegisterResponse{
		Status: pb.UserRegisterStatus_SUCCESS,
	}

	return res, nil
}

// SignIn サインイン
func (u *user) SignIn(ctx context.Context, req *pb.UserSignInRequest) (*pb.UserSignInResponse, error) {
	groupName := req.GetGroupName()
	loginID, password := req.GetLoginId(), req.GetPassword()
	userName, err := u.us.SignIn(ctx, groupName, loginID, password)
	if err != nil {
		if err == service.ErrUnauthorized {
			st := status.New(codes.Unauthenticated, err.Error())
			return nil, st.Err()
		}
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}
	return &pb.UserSignInResponse{
		GroupName: groupName,
		UserName:  userName,
	}, nil
}
