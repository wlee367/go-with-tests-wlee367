package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Booting up calculator...")

	scanner := bufio.NewScanner(os.Stdin)

	historyList := []HistoryList{}

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
		if err != nil || option < 1 || option > 5 {
			fmt.Println("Invalid input. Please enter a valid option (1, 2, 3, 4 or 5).")
			continue
		}

		historyList = doWorkBasedOnInput(option, historyList)
		fmt.Println("------- restarting ---------")
		fmt.Println()
	}
}
func doWorkBasedOnInput(userInput int, history []HistoryList) []HistoryList {
	switch userInput {
	case 1:
		return PrintResult(AddNumbers, history)
	case 2:
		return PrintResult(SubtractNumbers, history)
	case 3:
		return PrintResult(MultiplyNumbers, history)
	case 4:
		return PrintResult(DivideNumbers, history)
	case 5:
		PrintHistory(history)
	default:
		fmt.Println("Unrecognized value")
		return history
	}
	return nil
}
