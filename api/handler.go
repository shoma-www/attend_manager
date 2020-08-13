package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// HealthCheckHandler ヘルスチェック用API
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	status := http.StatusOK
	body := map[string]string{
		"status": "success",
	}

	cs := NewHealthCheckService()
	if err := cs.Grpc(ctx); err != nil {
		log.Println(err)
		status = http.StatusServiceUnavailable
		body["status"] = "failed"
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
