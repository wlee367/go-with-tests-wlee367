package main

func AddNumbers() int {
	a, b, err := PromptUserForNumbers()
	CheckForErrors(err)

	result := AddNumsHelper(a, b)
	return result
}

func SubtractNumbers() int {
	a, b, err := PromptUserForNumbers()
	CheckForErrors(err)

	result := SubtractNumbersHelper(a, b)
	return result
}

func MultiplyNumbers() int {
	a, b, err := PromptUserForNumbers()
	CheckForErrors(err)

	result := MultiplyNumbersHelper(a, b)
	return result
}

func DivideNumbers() int {
	a, b, err := PromptUserForNumbers()
	CheckForErrors(err)

	result := DivideNumbersHelper(a, b)
	return result
}
