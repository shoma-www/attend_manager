package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shoma-www/attend_manager/api/config"
	"github.com/shoma-www/attend_manager/api/handler"
	"github.com/shoma-www/attend_manager/api/infra"
	"github.com/shoma-www/attend_manager/api/service"
	"github.com/shoma-www/attend_manager/core"
)

// Server hhtp.Serverをラップした構造体
type Server struct {
	server  *http.Server
	conf    *config.Server
	logger  core.Logger
	factory *infra.Factory
}

// NewServer コンストラクタ
func NewServer(c *config.Server, l core.Logger, f *infra.Factory) *Server {
	return &Server{
		conf:    c,
		logger:  l,
		factory: f,
	}
}

// Init サーバーの初期化。ハンドラーとか設定
func (s *Server) Init() {
	ss := s.factory.CreateSession()
	r := mux.NewRouter()
	cc := service.NewCheck(s.logger, s.factory.CreateCheck())
	ch := handler.NewCheckHandler(s.logger, cc)
	r.HandleFunc("/healthcheck", ch.HealthCheck)

	cs := service.NewUser(s.factory.CreateUser())
	u := handler.NewUser(s.logger, cs, ss)
	ru := r.
		PathPrefix("/user").
		Subrouter()
	ru.
		HandleFunc("/register", u.Register).
		Methods(http.MethodPost, http.MethodOptions)
	ru.
		HandleFunc("/signin", u.SignIn).
		Methods(http.MethodPost, http.MethodOptions)

	gs := service.NewGroup(s.factory.CreateGroup())
	g := handler.NewGroup(gs)
	gu := r.PathPrefix("/group").
		Subrouter()
	gu.HandleFunc("/create", g.Create).
		Methods(http.MethodPost, http.MethodOptions)

	m := NewMiddleware(s.logger)
	r.Use(
		m.AddUUIDWithContext,
		m.Logger,
		CORSMiddleware,
	)

	s.server = &http.Server{
		Handler:      r,
		Addr:         s.conf.Addr,
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

const ShutdownTimes = 5

// Shutdown Serverのシャットダウン
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimes*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}
