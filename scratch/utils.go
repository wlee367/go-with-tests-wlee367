package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func PromptUserForNumbers() (int, int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	// Prompt the user for the first number
	fmt.Print("Enter the first number: ")
	scanner.Scan()
	firstInput := scanner.Text()

	// Prompt the user for the second number
	fmt.Print("Enter the second number: ")
	scanner.Scan()
	secondInput := scanner.Text()

	// Convert user inputs to integers
	firstNumber, err1 := strconv.Atoi(firstInput)
	secondNumber, err2 := strconv.Atoi(secondInput)

	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("invalid input. Please enter valid numbers")
	}

	return firstNumber, secondNumber, nil
}

func AddNumsHelper(a, b int) int {
	return a + b
}

func SubtractNumbersHelper(a, b int) int {
	return a - b
}

func MultiplyNumbersHelper(a, b int) int {
	return a * b
}

func DivideNumbersHelper(a, b int) int {
	return a / b
}

func PrintResult(operation func() CalculatorResult) (HistoryList, error) {
	result := operation()

	if result.CalculationError != nil {
		fmt.Println("Error:", result.CalculationError)
		return HistoryList{}, result.CalculationError
	}

	fmt.Println("")
	fmt.Println("=======================")
	fmt.Println("The result is: ", result.HistoryList.OperationResult)
	fmt.Println("=======================")
	fmt.Println("")
	return result.HistoryList, nil
}
