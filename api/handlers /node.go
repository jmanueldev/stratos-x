package handlers

import (
	"encoding/json"
	"net/http"
)

type Node struct {
	ID          string `json:"id"`
	GPUCount    int    `json:"gpu_count"`
	RDMAEnabled bool   `json:"rdma_enabled"`
	Status      string `json:"status"`
}

var NodeList = []Node{
	{"node-1", 8, true, "idle"},
	{"node-2", 4, true, "idle"},
	{"node-3", 8, false, "idle"},
}

func ListNodes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(NodeList)
}
