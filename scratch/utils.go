package main

import (
	"bufio"
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
		fmt.Println("Invalid input. Please enter valid numbers.")
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

func CheckForErrors(err error) (int, error) {
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return 0, nil
}

func PrintResult(operation func() int, operationType string, history History) {
	result := operation()
	fmt.Println("")
	fmt.Println("=======================")
	fmt.Println("The result is: ", result)
	fmt.Println("=======================")
	fmt.Println("")

	history.saveResult(operationType, result)
}

func PrintHistory(history History) {
	fmt.Println(history)
}
