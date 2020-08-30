package infra

import (
	"context"

	pb "github.com/shoma-www/attend_manager/api/proto"
)

type userGrpc struct {
	address string
}

func (ug *userGrpc) Resister(ctx context.Context, userID string, password string) error {
	con, err := createGrpcConn(ug.address)
	if err != nil {
		return err
	}
	defer con.Close()
	client := pb.NewUserClient(con)
	req := &pb.RegisterRequesut{
		UserId:   userID,
		Password: password,
	}
	_, err = client.Register(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
