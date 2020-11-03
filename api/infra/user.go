package infra

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shoma-www/attend_manager/api/entity"
	pb "github.com/shoma-www/attend_manager/api/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userGrpc struct {
	address string
}

func (ug *userGrpc) Resister(ctx context.Context, user entity.User) error {
	con, err := createGrpcConn(ug.address)
	if err != nil {
		return errors.Wrap(err, "create grpc connection error")
	}
	defer con.Close()
	client := pb.NewUserClient(con)
	req := &pb.UserRegisterRequesut{
		AttendanceGroupId: user.GroupID,
		LoginId:           user.LoginID,
		Password:          user.Password,
		Name:              user.Name,
	}
	_, err = client.Register(ctx, req)
	if err != nil {
		return errors.Wrap(err, "failed user register")
	}
	return nil
}

func (ug *userGrpc) SigIn(ctx context.Context, groupName, loginID, password string) (*entity.SigninData, error) {
	con, err := createGrpcConn(ug.address)
	if err != nil {
		return nil, errors.Wrap(err, "create grpc connection error")
	}
	defer con.Close()

	client := pb.NewUserClient(con)
	req := &pb.UserSignInRequest{
		GroupName: groupName,
		LoginId:   loginID,
		Password:  password,
	}
	res, err := client.SignIn(ctx, req)
	if err != nil {
		st := status.Convert(err)
		if st.Code() == codes.Unauthenticated {
			return nil, entity.ErrUnauthenticated
		}
		return nil, err
	}

	return &entity.SigninData{
		GroupName: groupName,
		UserID:    loginID,
		UserName:  res.GetUserName(),
	}, nil
}
