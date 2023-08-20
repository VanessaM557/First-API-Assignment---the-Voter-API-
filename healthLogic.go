package main
import (
	"time"
)

//initialize the health data
func InitializeHealth() HealthData {
	return HealthData{
		BootTime:              time.Now(),
		TotalAPICalls:         0,
		TotalAPICallsWithError: 0,
	}
}

func IncrementTotalAPICalls(healthData *HealthData) {
	healthData.TotalAPICalls++
}

func IncrementTotalAPICallsWithError(healthData *HealthData) {
	healthData.TotalAPICallsWithError++
}
