package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//handling GET /voters/health
func handleHealthCheck(w http.ResponseWriter, r *http.Request, voterList *VoterList) {
	voterList.HealthData.TotalAPICalls++
	uptimeDuration := time.Since(voterList.HealthData.BootTime)
	voterList.HealthData.Uptime = uptimeDuration.String()

	respondWithJSON(w, http.StatusOK, voterList.HealthData)
}

//responding with JSON

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}



