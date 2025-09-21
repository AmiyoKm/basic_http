package main

import (
	"net/http"
	"time"

	"github.com/AmiyoKm/basic_http/utils"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	type HealthStatus struct {
		Status    string    `json:"status"`
		Timestamp time.Time `json:"timestamp"`
		Version   string    `json:"version"`
	}

	utils.WriteJSON(w, HealthStatus{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   "1.0.0",
	})
}
