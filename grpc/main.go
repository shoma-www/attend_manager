package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/server"
)

const (
	port = ":50051"
)

func main() {
	logger := core.NewLogger(core.Debug)
	logger.Info("Start gRPC Server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		return
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(LoggingInterceptor(logger)))
	server.Register(s, logger)

	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Println(err)
	}

	logger.Info("Exit gRPC Server")
	s.Stop()
}
