package main

import (
	decodeticket "debug-stats/decode-ticket"
	retrievebugtickets "debug-stats/retrieve-bug-tickets"
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	bugTickets := GetTickets()

	ticketIds := []int{}

	fmt.Printf("Total tickets : %d\n", bugTickets[0].Total)

	for i := 0; i < len(bugTickets); i++ {
		ticketIds = append(ticketIds, bugTickets[i].Data[0].ID)
	}

	slices.Sort(ticketIds)

	for i := 0; i < len(ticketIds); i++ {
		fmt.Printf("%d %d\n", i, ticketIds[i])
	}
}	

func GetTickets() []decodeticket.Ticket {
	tickets := []decodeticket.Ticket{}
	url := "/api/v3/search/stories"
	ticket := GetTicket1(url)

	totalTickets := ticket.Total

	tickets = append(tickets, ticket)

	url = ticket.Next

	for i := 2; i <= totalTickets; i++ {
		ticket := GetTicket1(url)
		tickets = append(tickets, ticket)
		url = ticket.Next
	}

	return tickets
}

func GetTicket1(url string) decodeticket.Ticket {
	ticketJson := retrievebugtickets.RetrieveTickets(url)
	ticket := decodeticket.DecodeTicket(ticketJson)

	return ticket
}