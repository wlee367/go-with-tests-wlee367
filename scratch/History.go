package main

import (
	"fmt"
)

type HistoryInputs []int

type HistoryList struct {
	OperationType   string
	OperationInputs HistoryInputs
	OperationResult int
}

func SaveResult(history []HistoryList, historyToSave HistoryList) []HistoryList {
	history = append(history, historyToSave)
	return history
}

func PrintHistory(historyList []HistoryList) {

	if len(historyList) == 0 {
		fmt.Println("\nno history found yet")
		fmt.Println()
		return
	}

	fmt.Println("printing history...")
	for _, item := range historyList {
		fmt.Printf("{ \n Operation Type: %s, \n Operation Inputs: %d, \n Operation Result: %d \n}",
			item.OperationType, item.OperationInputs, item.OperationResult)
		fmt.Println("")
	}
}
