package main

import (
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
)

func main() {
	path := "./config.yaml"
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Println(err)
		return
	}

	c, err := LoadConfig(absPath)

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", HealthCheckHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", c.Server.Addr, c.Server.Port),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Println(err)
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
		log.Println("Graceful ShutDown...")
		endWaiter.Done()
	}()
	endWaiter.Wait()

	log.Println("Exit API Server")
}
