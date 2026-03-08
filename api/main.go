package main

import (
	"net/http"
	"log"
	"stratos-x/api/handlers"
)

func main() {
	http.HandleFunc("/missions", handlers.SubmitMission)
	http.HandleFunc("/missions/list", handlers.ListMissions)
	http.HandleFunc("/nodes/list", handlers.ListNodes)
	http.HandleFunc("/telemetry", handlers.ListTelemetry)

	log.Println("STRATOS-X API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
