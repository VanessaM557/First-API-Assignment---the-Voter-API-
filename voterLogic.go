package main
import (
	"fmt"
	"time"
)

//adding in UPDATE

func UpdateVoterByID(id uint, firstName, lastName string, voterList *VoterList) error {
	voterIndex, exists := findVoterIndexByID(id, voterList)
	if !exists {
		return fmt.Errorf("voter not found")
	}

	voterList.Voters[voterIndex].FirstName = firstName
	voterList.Voters[voterIndex].LastName = lastName

	return nil
}

func findVoterIndexByID(id uint, voterList *VoterList) (int, bool) {
	for i, voter := range voterList.Voters {
		if voter.VoterID == id {
			return i, true
		}
	}
	return -1, false
}

type HealthData struct {
	BootTime              time.Time `json:"bootTime"`
	TotalAPICalls         int       `json:"totalApiCalls"`
	TotalAPICallsWithError int       `json:"totalApiErrors"`
}

type VoterList struct {
	Voters    []Voter      
	HealthData HealthData 
}

// Adding in DELETE:
func DeleteVoterByID(id uint, voterList VoterList) error {
	_, exists := voterList.Voters[id]
	if !exists {
		return fmt.Errorf("voter not found")
	}

	delete(voterList.Voters, id)
	return nil
}

// Function to create a new unique voter 

func CreateNewVoter(voterID uint, firstName, lastName string) Voter {
	return Voter{
		VoterID:     voterID,
		FirstName:   firstName,
		LastName:    lastName,
		VoteHistory: make([]voterPoll, 0),
	}
}

// retrieving voters

func GetAllVoters(voterList VoterList) []Voter {
	voters := make([]Voter, 0, len(voterList.Voters))
	for _, voter := range voterList.Voters {
		voters = append(voters, voter)
	}
	return voters
}

//to record a vote for a specific voter 

func RecordVote(voterID, pollID uint, voteDate time.Time, voterList VoterList) error {
	voter, exists := voterList.Voters[voterID]
	if !exists {
		return fmt.Errorf("voter cannot be found")
	}
	vote := voterPoll{
		PollID:   pollID,
		VoteDate: voteDate,
	}
	voter.VoteHistory = append(voter.VoteHistory, vote)
	voterList.Voters[voterID] = voter

	return nil
}

//retrieving voter by ID
func GetVoterByID(id uint, voterList VoterList) (Voter, error) {
	voter, exists := voterList.Voters[id]
	if !exists {
		return Voter{}, fmt.Errorf("voter cannot be found")
	}
	return voter, nil
}

//retrieving voter history by ID

func GetVoterPollsByID(id uint, voterList VoterList) ([]voterPoll, error) {
	voter, exists := voterList.Voters[id]
	if !exists {
		return nil, fmt.Errorf("voter cannot be found")
	}
	return voter.VoteHistory, nil
}

// retrieving vote with voter ID and poll ID

func GetVoterPollByID(voterID, pollID uint, voterList VoterList) (voterPoll, error) {
	voter, exists := voterList.Voters[voterID]
	if !exists {
		return voterPoll{}, fmt.Errorf("voter cannot be found")
	}

	for _, vote := range voter.VoteHistory {
		if vote.PollID == pollID {
			return vote, nil
		}
	}
	return voterPoll{}, fmt.Errorf("poll not found for this voter")
}

func NewVoterList() *VoterList {
	return &VoterList{
		Voters:    make([]Voter, 0),
		HealthData: InitializeHealth(),
	}
}
