package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/shoma-www/attend_manager/grpc/server"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		return
	}

	s := grpc.NewServer()
	server.Register(s)

	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Println(err)
	}
}
