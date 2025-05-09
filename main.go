package main

import (
	"fmt"
	"os"
)

func main() {
	runApp()
}

func runApp() {
	var option int

	fmt.Println("Select the calculator")
	fmt.Println("1. Investment Calculator")
	fmt.Println("2. Profit Calculator")
	fmt.Println("3. Bank Service")
	fmt.Println("4. Exit")
	fmt.Print("Your choice: ")

	fmt.Scan(&option)

	switch option {

	case 1:
		investmentCalculator()
	case 2:
		profitCalculator()
	case 3:
		bankingMode()
	case 4:
		fmt.Println("Exiting now...")
		os.Exit(0)
	default:
		fmt.Println("Try again")
		runApp()
	}
}

func getUserInput(infoText string) float64 {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
