package handlers

import (
	"encoding/json"
	"net/http"
)

type Mission struct {
	Name     string `json:"name"`
	Topology string `json:"topology"`
	Nodes    int    `json:"nodes"`
	GPU      int    `json:"gpu"`
}

var MissionQueue []Mission

func SubmitMission(w http.ResponseWriter, r *http.Request) {
	var m Mission
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	MissionQueue = append(MissionQueue, m)
	json.NewEncoder(w).Encode(m)
}

func ListMissions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(MissionQueue)
}
