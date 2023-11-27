package main

import (
	retrievebugtickets "debug-stats/retrieve-bug-tickets"
	retrieveworkflows "debug-stats/retrieve-workflows"
	"encoding/csv"
	"fmt"
	"sort"
	"strconv"

	"log"
	"os"
	"time"
)

type OutstandingTickets struct {
	count int
	totalWaitTime float64
}

type CompletedTickets struct {
	count int
	totalTimeToFix float64
}

func main() {
	totalTickets, bugTickets := retrievebugtickets.GetTickets()
	totalTimeToFix := 0.0
	totalWaitTime := 0.0
	totalCompleted := 0
	totalUncompleted := 0

	today := time.Now()

	workflows := retrieveworkflows.GetWorkflows()

	workflowNames := make(map[int]string)
	workflowStates := make(map[int]string)

	for i:=0; i < len(workflows); i++ {
		workflowNames[workflows[i].ID] = workflows[i].Name
		for j := 0; j < len(workflows[i].States); j++ {
			workflowStates[workflows[i].States[j].ID] = workflows[i].States[j].Name
		}
	}

	outstandingTickets := make(map[int]map[int]OutstandingTickets)

	completedTickets := make(map[int]CompletedTickets)
	
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
			monthCompleted := completedAtTime.Year() * 100 + int(completedAtTime.Month())
			fmt.Println(monthCompleted)
			if completedInMonth, ok := completedTickets[monthCompleted] ; ok {
				completedInMonth.count++
				completedInMonth.totalTimeToFix += difference.Hours()/24
				completedTickets[monthCompleted] = completedInMonth
			} else {
				completedInMonth := CompletedTickets{1, difference.Hours()/24}
				completedTickets[monthCompleted] = completedInMonth
			}
			totalCompleted++
		} else {
			if bugTickets[i].WorkflowID == 500053632 {
				fmt.Println(bugTickets[i].AppURL)
			}
			difference := today.Sub(createdAt)
			waitTime := difference.Hours()/24
			totalWaitTime += waitTime
			totalUncompleted++
			if workflowTickets, ok := outstandingTickets[bugTickets[i].WorkflowID] ; ok {
				if workflowState, ok1 := workflowTickets[bugTickets[i].WorkflowStateID] ; ok1 {
					workflowState.count++
					workflowState.totalWaitTime += waitTime
					outstandingTickets[bugTickets[i].WorkflowID][bugTickets[i].WorkflowStateID] = workflowState
				} else {
					workflowState := OutstandingTickets{1, waitTime}
					outstandingTickets[bugTickets[i].WorkflowID][bugTickets[i].WorkflowStateID] = workflowState
				}
			} else {
				workflowState := OutstandingTickets{1, waitTime}
				workflowStateMap := make(map[int]OutstandingTickets)
				workflowStateMap[bugTickets[i].WorkflowStateID] = workflowState
				outstandingTickets[bugTickets[i].WorkflowID] = workflowStateMap
				
			}
			if bugTickets[i].WorkflowID == 0 {
				fmt.Println(bugTickets[i].ID)
			}
		}
	}

	for workflowId, workflowTickets := range outstandingTickets {
		workflowName := workflowNames[workflowId]
		fmt.Printf("\n%d %s\n", workflowId, workflowName)
		for workflowStateId, outstandingTickets := range workflowTickets {
			workflowStateName := workflowStates[workflowStateId]
			fmt.Printf("%-20s \t%d\t%d\n", workflowStateName, outstandingTickets.count, int64(outstandingTickets.totalWaitTime/float64(outstandingTickets.count)))
		}
	} 

	completedTicketsFile, err := os.Create("completed-tickets.csv")
	defer completedTicketsFile.Close()
	
	if err != nil {
		log.Fatal("Failed creating completed-tickets.csv")
	}
	fileWriter := csv.NewWriter(completedTicketsFile)
	defer fileWriter.Flush()

	monthsCompleted := make([]int, 0, len(completedTickets))

	for monthCompleted := range completedTickets {
		monthsCompleted = append(monthsCompleted, monthCompleted)
	}

	sort.Ints(monthsCompleted)

	for _, monthCompleted := range monthsCompleted {
		completedDetails := completedTickets[monthCompleted]
		meanTimeToFix := int(completedDetails.totalTimeToFix/float64(completedDetails.count))
		fmt.Printf("%d : %d\n", monthCompleted, meanTimeToFix)
		row := []string{strconv.Itoa(monthCompleted), strconv.Itoa(meanTimeToFix)}
		if err := fileWriter.Write(row); err != nil {
            log.Fatalln("error writing record to file", err)
        }
	}

	fmt.Printf("Total completed: %d Mean time to completion: %d days\n", totalCompleted, int64(totalTimeToFix/float64(totalCompleted)))
	fmt.Printf("Total uncompleted: %d Mean time waiting: %d days\n", totalUncompleted, int64(totalWaitTime/float64(totalUncompleted)))
}