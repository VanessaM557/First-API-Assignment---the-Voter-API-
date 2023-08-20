package main
import "time"

//initialize health information

func InitializeHealth() HealthData{
	return HealthData{
		BootTime:           time.Now(),
		TotalAPICalls:      0,
		TotalAPICallsWithError: 0,
	}
}

//update total API calls
func UpdateTotalAPICalls(healthData *HealthData) {
	healthData.TotalAPICalls++
}

//update total API calls with errors
func UpdateTotalAPICallsWithError(healthData *HealthData) {
	healthData.TotalAPICallsWithError++
}

//struct for Health Data
type HealthData struct {
	BootTime           time.Time `json:"bootTime"`
	TotalAPICalls      int       `json:"totalApiCalls"`
	TotalAPICallsWithError int    `json:"totalApiErrors"`
}

