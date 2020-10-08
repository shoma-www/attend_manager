package infra

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shoma-www/attend_manager/api/entity"
	pb "github.com/shoma-www/attend_manager/api/proto"
)

type groupGrpc struct {
	address string
}

func (gg *groupGrpc) Create(ctx context.Context, group entity.Group, user entity.User) error {
	con, err := createGrpcConn(gg.address)
	if err != nil {
		return errors.Wrap(err, "create grpc connection error")
	}
	defer con.Close()
	client := pb.NewAttendanceGroupClient(con)
	req := &pb.AttendanceGroupRequesut{
		GroupName: group.Name,
		LoginId:   user.LoginID,
		Password:  user.Password,
		UserName:  user.Name,
	}

	if _, err = client.Create(ctx, req); err != nil {
		return err
	}
	return nil
}
