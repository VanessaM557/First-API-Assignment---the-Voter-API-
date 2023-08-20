package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//handling GET /voters/:id
func handleGetVoterByID(w http.ResponseWriter, r *http.Request, voterID uint, voterList *VoterList) {
	voter, err := GetVoterByID(voterID, *voterList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	respondWithJSON(w, http.StatusOK, voter)
}

//handling PUT /voters/:id
func handleUpdateVoterByID(w http.ResponseWriter, r *http.Request, voterID uint, voterList *VoterList) {
	var updatedVoter Voter
	err := json.NewDecoder(r.Body).Decode(&updatedVoter)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = UpdateVoterByID(voterID, updatedVoter.FirstName, updatedVoter.LastName, voterList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//handling DELETE /voters/:id
func handleDeleteVoterByID(w http.ResponseWriter, r *http.Request, voterID uint, voterList *VoterList) {
	err := DeleteVoterByID(voterID, *voterList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//get path variable
func getPathVariable(r *http.Request, paramName string) string {
	return r.URL.Query().Get(paramName)
}

// responding with JSON
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
