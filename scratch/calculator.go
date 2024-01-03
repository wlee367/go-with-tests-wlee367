package main

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

func (c *Calculator) SaveHistory(historyList HistoryList) {
	c.History = append(c.History, historyList)
}

func AddNumbers() (int, HistoryList) {
	a, b, err := PromptUserForNumbers()
	CheckForErrors(err)

	result := AddNumsHelper(a, b)
	historyInputs := HistoryInputs{
		a, b,
	}
	historyList := HistoryList{
		OperationType:   ADD,
		OperationInputs: historyInputs,
		OperationResult: result,
	}

	return result, historyList
}

func SubtractNumbers() (int, HistoryList) {
	a, b, err := PromptUserForNumbers()
	CheckForErrors(err)

	result := SubtractNumbersHelper(a, b)
	historyInputs := HistoryInputs{
		a, b,
	}
	historyList := HistoryList{
		OperationType:   SUBTRACT,
		OperationInputs: historyInputs,
		OperationResult: result,
	}

	return result, historyList
}

func MultiplyNumbers() (int, HistoryList) {
	a, b, err := PromptUserForNumbers()
	CheckForErrors(err)

	result := MultiplyNumbersHelper(a, b)
	historyInputs := HistoryInputs{
		a, b,
	}
	historyList := HistoryList{
		OperationType:   MULTIPLY,
		OperationInputs: historyInputs,
		OperationResult: result,
	}

	return result, historyList
}

func DivideNumbers() (int, HistoryList) {
	a, b, err := PromptUserForNumbers()
	CheckForErrors(err)

	result := DivideNumbersHelper(a, b)
	historyInputs := HistoryInputs{
		a, b,
	}
	historyList := HistoryList{
		OperationType:   DIVIDE,
		OperationInputs: historyInputs,
		OperationResult: result,
	}
	return result, historyList
}
