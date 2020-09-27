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

type attendanceGroup struct {
	logger core.Logger
	gs     *service.AttendanceGroup
}

// NewAttendanceGroup コンストラクタ
func NewAttendanceGroup(l core.Logger, gs *service.AttendanceGroup) pb.AttendanceGroupServer {
	return &attendanceGroup{logger: l, gs: gs}
}

func (ag *attendanceGroup) Create(ctx context.Context, req *proto.AttendanceGroupRequesut) (*proto.AttendanceGroupResponse, error) {
	if _, _, err := ag.gs.Create(ctx, req.GroupName, req.LoginId, req.Password, req.UserName); err != nil {
		ag.logger.WithUUID(ctx).Error("creatr group errpr: %s", err)
		status := status.New(codes.Internal, err.Error())
		return nil, status.Err()
	}

	return &pb.AttendanceGroupResponse{}, nil
}
