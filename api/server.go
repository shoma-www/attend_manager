package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/shoma-www/attend_manager/core"
)

// Server hhtp.Serverをラップした構造体
type Server struct {
	server *http.Server
	conf   *Config
	logger *core.Logger
	repof  *RepoFactory
}

// NewServer コンストラクタ
func NewServer(c *Config, l *core.Logger, repof *RepoFactory) *Server {
	return &Server{
		conf:   c,
		logger: l,
		repof:  repof,
	}
}

// Init サーバーの初期化。ハンドラーとか設定
func (s *Server) Init() {
	r := mux.NewRouter()

	ch := NewCheckHandler(s.logger, s.repof)
	r.HandleFunc("/healthcheck", ch.HealthCheck)
	m := NewMiddleware(s.logger)
	r.Use(
		m.AddUUIDWithContext,
		m.Logger,
	)

	s.server = &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", s.conf.Server.Addr, s.conf.Server.Port),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

// ListenAndServe Serveします
func (s *Server) ListenAndServe() {
	if err := s.server.ListenAndServe(); err != nil {
		s.logger.Error(err.Error())
	}
}

// Shutdown Serverのシャットダウン
func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error(err.Error())
		os.Exit(1)
	}
}
