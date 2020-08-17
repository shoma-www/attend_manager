package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/shoma-www/attend_manager/core"
)

// CheckHandler handler
type CheckHandler struct {
	logger *core.Logger
	grpc   *Grpc
}

// NewCheckHandler コンストラクタ
func NewCheckHandler(l *core.Logger, g *Grpc) *CheckHandler {
	return &CheckHandler{logger: l, grpc: g}
}

// HealthCheck ヘルスチェック用API
func (ch *CheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	status := http.StatusOK
	body := map[string]string{
		"status": "success",
	}

	client := ch.grpc.CreateCheckClient()
	defer ch.grpc.Close()
	cs := NewHealthCheckService(ch.logger.WithUUID(r.Context()), client)
	if err := cs.Grpc(ctx); err != nil {
		ch.logger.Error(err.Error())
		status = http.StatusServiceUnavailable
		body["status"] = "failed"
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		ch.logger.WithUUID(r.Context()).Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
