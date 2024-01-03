package main

const ADD = "add"
const SUBTRACT = "subtract"
const MULTIPLY = "multiply"
const DIVIDE = "divide"

func AddNumbers(historyList []HistoryList) (int, []HistoryList) {
	a, b, err := PromptUserForNumbers()
	CheckForErrors(err)

	result := AddNumsHelper(a, b)

	historyInputs := HistoryInputs{
		a, b,
	}
	historyList = saveToHistory(result, historyList, historyInputs, ADD)
	return result, historyList
}

func SubtractNumbers(historyList []HistoryList) (int, []HistoryList) {
	a, b, err := PromptUserForNumbers()
	CheckForErrors(err)

	result := SubtractNumbersHelper(a, b)
	historyInputs := HistoryInputs{
		a, b,
	}
	historyList = saveToHistory(result, historyList, historyInputs, SUBTRACT)

	return result, historyList
}

func MultiplyNumbers(historyList []HistoryList) (int, []HistoryList) {
	a, b, err := PromptUserForNumbers()
	CheckForErrors(err)

	result := MultiplyNumbersHelper(a, b)
	historyInputs := HistoryInputs{
		a, b,
	}
	historyList = saveToHistory(result, historyList, historyInputs, MULTIPLY)

	return result, historyList
}

func DivideNumbers(historyList []HistoryList) (int, []HistoryList) {
	a, b, err := PromptUserForNumbers()
	CheckForErrors(err)

	result := DivideNumbersHelper(a, b)
	historyInputs := HistoryInputs{
		a, b,
	}
	historyList = saveToHistory(result, historyList, historyInputs, DIVIDE)

	return result, historyList
}

func saveToHistory(result int, historyList []HistoryList, historyInputs HistoryInputs, operationType string) []HistoryList {
	input := HistoryList{
		OperationType:   operationType,
		OperationInputs: historyInputs,
		OperationResult: result,
	}
	return SaveResult(historyList, input)
}
