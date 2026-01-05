package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type HealthHandler struct {
	DB *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{DB: db}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	status := "ok"
	message := "Service is healthy"
	dbStatus := "connected"

	if err := h.DB.Ping(); err != nil {
		status = "error"
		message = "Service is unhealthy: database unreachable"
		dbStatus = "disconnected"
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	response := map[string]string{
		"status":    status,
		"message":   message,
		"database":  dbStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
