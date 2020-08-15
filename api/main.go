package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/shoma-www/attend_manager/core"
)

func main() {
	path := "./config.yaml"
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalln(err)
		return
	}

	c, err := LoadConfig(absPath)
	if err != nil {
		log.Fatalln(err)
		return
	}

	logger := core.NewLogger(core.Debug)
	logger.Info("Start Attend Manager API Server")

	r := mux.NewRouter()

	ch := NewCheckHandler(logger)
	r.HandleFunc("/healthcheck", ch.HealthCheck)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", c.Server.Addr, c.Server.Port),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			logger.Error(err.Error())
		}
	}()

	// Graceful Shutdownっぽいこと
	// https://blog.potproject.net/2019/08/29/golang-graceful-shutdown-queue-process
	var endWaiter sync.WaitGroup
	endWaiter.Add(1)
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)
	go func() {
		<-sigCh
		logger.Info("Graceful ShutDown...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = srv.Shutdown(ctx); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
		endWaiter.Done()
	}()
	endWaiter.Wait()

	log.Println("Exit API Server")
}
