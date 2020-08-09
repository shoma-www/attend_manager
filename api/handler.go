package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HealthCheckHandler ヘルスチェック用API
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := map[string]string{
		"status": "success",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Fprint(w, string(jsonBody))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(jsonBody))
	w.WriteHeader(http.StatusOK)
}
