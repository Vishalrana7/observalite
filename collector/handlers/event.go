package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"log"

	"github.com/jackc/pgx/v5"
)

type Event struct {
	Service    string            `json:"service"`
	InstanceID string            `json:"instance_id"`
	Endpoint   string            `json:"endpoint"`
	Method     string            `json:"method"`
	Status     int               `json:"status"`
	LatencyMs  int               `json:"latency_ms"`
	SizeBytes  int               `json:"size_bytes"`
	Tags       map[string]string `json:"tags"`
}

func PostEventHandler(conn *pgx.Conn) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("===> Hit /v1/events")
		var evt Event

        // Log that request arrived
        log.Println("POST /v1/events called")

        if err := json.NewDecoder(r.Body).Decode(&evt); err != nil {
            log.Printf("JSON decode failed: %v\n", err)
            http.Error(w, "invalid json", http.StatusBadRequest)
            return
        }

        log.Printf("Decoded event: %+v\n", evt)

        _, err := conn.Exec(context.Background(),
            `INSERT INTO events (service, instance_id, endpoint, method, status, latency_ms, size_bytes, tags)
             VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
            evt.Service, evt.InstanceID, evt.Endpoint, evt.Method,
            evt.Status, evt.LatencyMs, evt.SizeBytes, evt.Tags,
        )

        if err != nil {
            log.Printf("DB insert failed: %v\n", err)
            http.Error(w, "db insert failed", http.StatusInternalServerError)
            return
        }

        log.Println("Insert successful")
        w.WriteHeader(http.StatusAccepted)
    }
}

