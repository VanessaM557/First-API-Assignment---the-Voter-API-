package main
import (
	"fmt"
	"time"
)

// Function to create a new unique voter 

func CreateNewVoter(voterID uint, firstName, lastName string) Voter {
	return Voter{
		VoterID:     voterID,
		FirstName:   firstName,
		LastName:    lastName,
		VoteHistory: make([]voterPoll, 0),
	}
}

