package main

import (
	retrievebugtickets "debug-stats/retrieve-bug-tickets"
	"fmt"
	"time"
)

func main() {
	totalTickets, bugTickets := retrievebugtickets.GetTickets()
	totalTimeToFix := 0.0
	totalWaitTime := 0.0
	totalCompleted := 0
	totalUncompleted := 0

	today := time.Now()

	for i := 0; i < totalTickets; i++ {
		createdAt := bugTickets[i].CreatedAt
		if bugTickets[i].CompletedAt != nil {
			completedAt := bugTickets[i].CompletedAt.(string)
			completedAtTime, error := time.Parse(time.RFC3339, completedAt)
			if error != nil {
				fmt.Println(error)
				return
			}
			difference := completedAtTime.Sub(createdAt)
			totalTimeToFix += difference.Hours()/24
			totalCompleted++
		} else {
			difference := today.Sub(createdAt)
			totalWaitTime += difference.Hours()/24
			totalUncompleted++
		}
	}

	fmt.Printf("Total completed: %d Mean time to completion: %d days\n", totalCompleted, int64(totalTimeToFix/float64(totalCompleted)))
	fmt.Printf("Total uncompleted: %d Mean time waiting: %d days\n", totalUncompleted, int64(totalWaitTime/float64(totalUncompleted)))
}