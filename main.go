package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Initializing the VoterList
	voterList := NewVoterList()

	//Switch to choose between methods of GET and POST
	
	http.HandleFunc("/voters", func(w http.ResponseWriter, r *http.Request) {
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

	//incorporating extra credit of PUT/DELETE
	
	http.HandleFunc("/voters/:id", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetVoterByID(w, r, voterList)
		case http.MethodPut:
			handleUpdateVoterByID(w, r, voterList)
		case http.MethodDelete:
			handleDeleteVoterByID(w, r, voterList)
		default:
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/voters/:id/polls/:pollid", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGetVoterPollByID(w, r, voterList)
		case http.MethodPut:
			handleUpdateVoterPollByID(w, r, voterList)
		case http.MethodDelete:
			handleDeleteVoterPollByID(w, r, voterList)
		default:
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/voters/health", func(w http.ResponseWriter, r *http.Request) {
		handleHealthCheck(w)
	})

	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

