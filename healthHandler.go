package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//handling GET /voters/health
func handleHealthCheck(w http.ResponseWriter, r *http.Request, voterList VoterList) {
	healthData := voterList.HealthData

	respondWithJSON(w, http.StatusOK, healthData)
}

//responding with JSON

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

