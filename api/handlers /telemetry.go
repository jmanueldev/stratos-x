package handlers

import (
	"encoding/json"
	"net/http"
)

type Telemetry struct {
	NodeID string  `json:"node_id"`
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
}

var TelemetryData = []Telemetry{
	{"node-1", 12.5, 2048},
	{"node-2", 45.2, 4096},
}

func ListTelemetry(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(TelemetryData)
}
