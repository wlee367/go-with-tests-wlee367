package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	MINIMUM_OPTION = 1
	MAXIMUM_OPTION = 5
)

func main() {
	fmt.Println("Booting up calculator...")

	scanner := bufio.NewScanner(os.Stdin)

	calculator := Calculator{}

	for {
		fmt.Println("1: Add, 2: Subtract, 3: Multiply, 4: Divide, 5: History, q to quit")

		// Use scanner to read the user input
		scanner.Scan()
		userInput := scanner.Text()

		if userInput == "q" {
			fmt.Println("exiting...")
			break
		}

		// Convert user input to integer
		option, err := strconv.Atoi(userInput)
		if err != nil || option < MINIMUM_OPTION || option > MAXIMUM_OPTION {
			fmt.Println("Invalid input. Please enter a valid option (1, 2, 3, 4 or 5).")
			continue
		}

		historyList, err := doWorkBasedOnInput(option, calculator)
		if err == nil {
			calculator.AddToHistory(historyList)
		}
		fmt.Println("------- restarting ---------")
		fmt.Println()
	}
}
func doWorkBasedOnInput(userInput int, calculator Calculator) (HistoryList, error) {
	switch userInput {
	case 1:
		return PrintResult(AddNumbers)
	case 2:
		return PrintResult(SubtractNumbers)
	case 3:
		return PrintResult(MultiplyNumbers)
	case 4:
		return PrintResult(DivideNumbers)
	case 5:
		return PrintHistory(calculator.History)
	default:
		fmt.Println("Unrecognized value")
		return HistoryList{}, nil
	}
}
