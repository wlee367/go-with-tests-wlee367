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

func PrintHistory(historyList []HistoryList) (HistoryList, error) {
	listToReturn := HistoryList{
		OperationType:   "history",
		OperationInputs: HistoryInputs{0, 0},
		OperationResult: 0,
	}

	if len(historyList) == 0 {
		fmt.Println("\nno history found yet")
		fmt.Println()
		return listToReturn, nil
	}

	fmt.Println("printing history...")
	for _, item := range historyList {
		fmt.Printf("{ \n Operation Type: %s, \n Operation Inputs: %d, \n Operation Result: %d \n}",
			item.OperationType, item.OperationInputs, item.OperationResult)
		fmt.Println("")
	}

	return listToReturn, nil
}
