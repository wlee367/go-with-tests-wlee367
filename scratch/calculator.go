package main

import "fmt"

const (
	ADD      = "add"
	SUBTRACT = "subtract"
	MULTIPLY = "multiply"
	DIVIDE   = "divide"
	HISTORY  = "history"
)

type Calculator struct {
	History []HistoryList
}

type CalculatorResult struct {
	Result           int
	HistoryList      HistoryList
	CalculationError error
}

func (c *Calculator) AddToHistory(historyList HistoryList) {
	c.History = append(c.History, historyList)
}

func AddNumbers() CalculatorResult {
	a, b, err := PromptUserForNumbers()
	if err != nil {
		fmt.Println(err)
		return CalculatorResult{0, HistoryList{}, err}
	}

	result := AddNumsHelper(a, b)
	historyInputs := HistoryInputs{
		a, b,
	}
	historyList := HistoryList{
		OperationType:   ADD,
		OperationInputs: historyInputs,
		OperationResult: result,
	}

	return CalculatorResult{result, historyList, nil}
}

func SubtractNumbers() CalculatorResult {
	a, b, err := PromptUserForNumbers()
	if err != nil {
		fmt.Println(err)
		return CalculatorResult{0, HistoryList{}, err}
	}

	result := SubtractNumbersHelper(a, b)
	historyInputs := HistoryInputs{
		a, b,
	}
	historyList := HistoryList{
		OperationType:   SUBTRACT,
		OperationInputs: historyInputs,
		OperationResult: result,
	}

	return CalculatorResult{result, historyList, nil}
}

func MultiplyNumbers() CalculatorResult {
	a, b, err := PromptUserForNumbers()
	if err != nil {
		fmt.Println(err)
		return CalculatorResult{0, HistoryList{}, err}
	}

	result := MultiplyNumbersHelper(a, b)
	historyInputs := HistoryInputs{
		a, b,
	}
	historyList := HistoryList{
		OperationType:   MULTIPLY,
		OperationInputs: historyInputs,
		OperationResult: result,
	}

	return CalculatorResult{result, historyList, nil}
}

func DivideNumbers() CalculatorResult {
	a, b, err := PromptUserForNumbers()
	if err != nil {
		fmt.Println(err)
		return CalculatorResult{0, HistoryList{}, err}
	}

	result := DivideNumbersHelper(a, b)
	historyInputs := HistoryInputs{
		a, b,
	}
	historyList := HistoryList{
		OperationType:   DIVIDE,
		OperationInputs: historyInputs,
		OperationResult: result,
	}
	return CalculatorResult{result, historyList, nil}
}
