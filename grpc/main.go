package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"cloud.google.com/go/profiler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/config"
	"github.com/shoma-www/attend_manager/grpc/ent"
	"github.com/shoma-www/attend_manager/grpc/infra"
	pb "github.com/shoma-www/attend_manager/grpc/proto"
	"github.com/shoma-www/attend_manager/grpc/server"
	"github.com/shoma-www/attend_manager/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var revision = "unknown"

func main() {
	path := "./config/config.yaml"
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c, err := config.LoadConfig(absPath)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = profiler.Start(profiler.Config{
		ProjectID:      c.ProjectID,
		Service:        c.Service,
		ServiceVersion: revision,
		DebugLogging:   false,
		MutexProfiling: true,
	})
	if err != nil {
		log.Fatalf("failed to start the profiler: %v", err)
	}

	logger := core.NewLogger(core.Debug)
	logger.Info("Start gRPC Server")
	lis, err := net.Listen("tcp", c.Server.Address)
	if err != nil {
		logger.Error("%s", err.Error())
		os.Exit(1)
		return
	}

	dbName := os.Getenv("ATTEND_DB_NAME")
	dbUser := os.Getenv("ATTEND_DB_USER")
	dbPassword := os.Getenv("ATTEND_DB_PASSWARD")
	cl, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?parseTime=true", dbUser, dbPassword, dbName))
	if err != nil {
		logger.Error("%s", err.Error())
		os.Exit(1)
		return
	}
	defer cl.Close()
	rFactory := infra.NewFactory(logger, cl)

	s := grpc.NewServer(grpc.UnaryInterceptor(LoggingInterceptor(logger)))
	Register(s, logger, rFactory)

	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Println(err)
	}

	logger.Info("Exit gRPC Server")
	s.GracefulStop()
}

// Register サーバーの登録
func Register(s *grpc.Server, l core.Logger, factory *infra.Factory) {
	pb.RegisterCheckServer(s, server.NewCheck(l))
	tr := factory.CreateTransaction()
	ur := factory.CreateUserRepository()
	gr := factory.CreateAttendanceGroupRepository()

	us := service.NewUser(l, tr, ur, gr)
	gs := service.NewAttendanceGroup(l, tr, gr, ur)

	pb.RegisterUserServer(s, server.NewUser(l, us))
	pb.RegisterAttendanceGroupServer(s, server.NewAttendanceGroup(l, gs))
}
