package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//Handling GET /voters/:id
func handleGetVoterByID(w http.ResponseWriter, r *http.Request, voterList VoterList) {
	voterIDStr := getPathVariable(r, "id")
	voterID, err := strconv.ParseUint(voterIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid voter ID", http.StatusBadRequest)
		return
	}

	voter, err := GetVoterByID(uint(voterID), voterList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, voter)
}


// Handling GET /voters

func handleGetAllVoters(w http.ResponseWriter, r *http.Request, voterList VoterList) {
	voters := GetAllVoters(voterList)
	respondWithJSON(w, http.StatusOK, voters)
}

// handling POST /voters
func handleCreateNewVoter(w http.ResponseWriter, r *http.Request, voterList VoterList) {
	var newVoter struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	err := json.NewDecoder(r.Body).Decode(&newVoter)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	voterID := uint(len(voterList.Voters) + 1) // Just for example, you should use a proper ID mechanism
	voter := CreateNewVoter(voterID, newVoter.FirstName, newVoter.LastName)
	voterList.Voters[voterID] = voter

	w.WriteHeader(http.StatusCreated)
}

// Handling GET /voters/:id/polls
func handleGetVoterPollsByID(w http.ResponseWriter, r *http.Request, voterList VoterList) {
	voterIDStr := getPathVariable(r, "id")
	voterID, err := strconv.ParseUint(voterIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid voter ID", http.StatusBadRequest)
		return
	}

	voteHistory, err := GetVoterPollsByID(uint(voterID), voterList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, voteHistory)
}

// Handling GET /voters/:id/polls/:pollid
func handleGetVoterPollByID(w http.ResponseWriter, r *http.Request, voterList VoterList) {
	voterIDStr := getPathVariable(r, "id")
	pollIDStr := getPathVariable(r, "pollid")

	voterID, err := strconv.ParseUint(voterIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid voter ID", http.StatusBadRequest)
		return
	}

	pollID, err := strconv.ParseUint(pollIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid poll ID", http.StatusBadRequest)
		return
	}

	vote, err := GetVoterPollByID(uint(voterID), uint(pollID), voterList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, vote)
}

// Handling PUT /voters/:id

func handleUpdateVoterByID(w http.ResponseWriter, r *http.Request, voterList VoterList) {
	voterIDStr := getPathVariable(r, "id")
	voterID, err := strconv.ParseUint(voterIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid voter ID", http.StatusBadRequest)
		return
	}

	var updateRequest struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
	err = json.NewDecoder(r.Body).Decode(&updateRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = UpdateVoterByID(uint(voterID), updateRequest.FirstName, updateRequest.LastName, voterList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Handling DELETE /voters/:id

func handleDeleteVoterByID(w http.ResponseWriter, r *http.Request, voterList VoterList) {
	voterIDStr := getPathVariable(r, "id")
	voterID, err := strconv.ParseUint(voterIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid voter ID", http.StatusBadRequest)
		return
	}

	err = DeleteVoterByID(uint(voterID), voterList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// grabbing the path variable from URL
func getPathVariable(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

// responding with JSON
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
