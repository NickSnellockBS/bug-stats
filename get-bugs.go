package main

import (
	retrievebugtickets "debug-stats/retrieve-bug-tickets"
	"fmt"

	"golang.org/x/exp/slices"
)

func main1() {
	bugTickets := retrievebugtickets.GetTickets()

	ticketIds := []int{}

	for i := 0; i < len(bugTickets); i++ {
		ticketIds = append(ticketIds, bugTickets[i].ID)
	}

	slices.Sort(ticketIds)

	for i := 0; i < len(ticketIds); i++ {
		if i > 4800 {
			fmt.Printf("%d %d\n", i, ticketIds[i])
		}
	}

	fmt.Printf("Returned %d sorted %d\n", len(bugTickets), len(ticketIds))
}