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

	history := History{}

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

		doWorkBasedOnInput(option, history)
		fmt.Println("------- restarting ---------")
		fmt.Println()
	}
}
func doWorkBasedOnInput(userInput int, history History) {
	switch userInput {
	case 1:
		PrintResult(AddNumbers, "add", history)
	case 2:
		PrintResult(SubtractNumbers, "substract", history)
	case 3:
		PrintResult(MultiplyNumbers, "multiply", history)
	case 4:
		PrintResult(DivideNumbers, "divide", history)
	case 5:
		PrintHistory(history)
	default:
		fmt.Println("Unrecognized value")
	}
}
