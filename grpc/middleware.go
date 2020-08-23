package main

import (
	"context"

	"github.com/rs/xid"
	"github.com/shoma-www/attend_manager/core"
	"google.golang.org/grpc"
)

// LoggingInterceptor 処理の前後にログを仕込む
func LoggingInterceptor(logger core.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		uuid := xid.New().String()
		ctx = context.WithValue(ctx, core.UUIDContextKey, uuid)

		logger.WithUUID(ctx).Info("method: %s, request: %s", info.FullMethod, req)
		resp, err := handler(ctx, req)
		logger.WithUUID(ctx).Info("method: %s, response: %s", info.FullMethod, resp)
		return resp, err
	}
}
