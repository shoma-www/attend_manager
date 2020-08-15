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
}

// NewCheckHandler コンストラクタ
func NewCheckHandler(l *core.Logger) *CheckHandler {
	return &CheckHandler{logger: l}
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

	cs := NewHealthCheckService()
	if err := cs.Grpc(ctx); err != nil {
		ch.logger.Error(err.Error())
		status = http.StatusServiceUnavailable
		body["status"] = "failed"
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		ch.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
