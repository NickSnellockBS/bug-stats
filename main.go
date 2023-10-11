package main

import (
	"fmt"
	"time"
)

func main1() {
	completedTickets := GetTickets()
	totalTickets := completedTickets[0].Total
	totalWaitTime := 0.0

	for j := 0; j < len(completedTickets); j++ {
		for i:=0; i < len(completedTickets[j].Data); i++ {
			createdAt := completedTickets[j].Data[i].CreatedAt
			completedAt := completedTickets[j].Data[i].CompletedAt.(string)
			completedAtTime, error := time.Parse(time.RFC3339, completedAt)
			if error != nil {
				fmt.Println(error)
				return
			}
			difference := completedAtTime.Sub(createdAt)
			totalWaitTime += difference.Hours()/24
		}
	}

	fmt.Printf("Mean time to completion: %d\n", int64(totalWaitTime/float64(totalTickets)))

	outstandingTickets := GetTickets()
	totalTicketsWaiting := outstandingTickets[0].Total
	totalWaitTimeWaiting := 0.0

	now := time.Now()

	for j := 0; j < len(outstandingTickets); j++ {
		for i:=0; i < len(outstandingTickets[j].Data); i++ {
			createdAt := outstandingTickets[j].Data[i].CreatedAt
			difference := now.Sub(createdAt)
			totalWaitTimeWaiting += difference.Hours()/24
			fmt.Printf("%d %d-%d-%d\n", outstandingTickets[j].Data[i].ID, createdAt.Day(), createdAt.Month(), createdAt.Year())
		}
	}

	fmt.Printf("Mean time waiting: %d\n", int64(totalWaitTimeWaiting/float64(totalTicketsWaiting)))
}