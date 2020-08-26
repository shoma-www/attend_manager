package infra

import (
	"context"

	"github.com/shoma-www/attend_manager/api/entity"
	pb "github.com/shoma-www/attend_manager/api/proto"
)

// checkGrpc Check系のGrpc通信するやつ
type checkGrpc struct {
	address string
}

// HealthCheck CheckServerへのClietnを生成
func (cg *checkGrpc) HealthCheck(ctx context.Context) (*entity.HealthCheckStatus, error) {
	con, err := createGrpcConn(cg.address)
	if err != nil {
		return nil, err
	}
	defer con.Close()
	client := pb.NewCheckClient(con)
	pbst, err := client.HealthCheck(ctx, &pb.HealthRequest{})
	if err != nil {
		return nil, err
	}
	st := &entity.HealthCheckStatus{
		Status: pbst.GetStatus(),
	}
	return st, nil
}
