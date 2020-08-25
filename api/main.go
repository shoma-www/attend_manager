package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/shoma-www/attend_manager/api/config"
	"github.com/shoma-www/attend_manager/core"
)

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

	logger := core.NewLogger(core.Debug)
	logger.Info("Start Attend Manager API Server")

	repof := NewRepoFactory(c)

	s := NewServer(c, logger, repof)
	s.Init()
	go s.ListenAndServe()

	// Graceful Shutdownっぽいこと
	// https://blog.potproject.net/2019/08/29/golang-graceful-shutdown-queue-process
	var endWaiter sync.WaitGroup
	endWaiter.Add(1)
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		<-sigCh
		logger.Info("Graceful ShutDown...")
		s.Shutdown()
		endWaiter.Done()
	}()
	endWaiter.Wait()

	logger.Info("Exit API Server")
}
