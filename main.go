package main

import (
	"fmt"
	"net/http"
)

// main function
func main() {
	// Initializing the VoterList
	voterList := NewVoterList()

  // handeling for route /voters
  
	http.HandleFunc("/voters", func(w http.ResponseWriter, r *http.Request) {
    //switch statements  for GET and POST Request methods
		switch r.Method {
		case http.MethodGet:
			handleGetAllVoters(w, r, voterList)
		case http.MethodPost:
			handleCreateNewVoter(w, r, voterList)
		default:
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		}
	})

  http.HandleFunc("/voters/", func(w http.ResponseWriter, r *http.Request) {
		handleGetVoterByID(w, r, voterList)
	})

	http.HandleFunc("/voters/:id/polls", func(w http.ResponseWriter, r *http.Request) {
		handleGetVoterPollsByID(w, r, voterList)
	})

	http.HandleFunc("/voters/:id/polls/:pollid", func(w http.ResponseWriter, r *http.Request) {
		handleGetVoterPollByID(w, r, voterList)
	})

	http.HandleFunc("/voters/health", func(w http.ResponseWriter, r *http.Request) {
		handleHealthCheck(w)
	})

	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
