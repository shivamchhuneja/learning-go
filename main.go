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
	fmt.Println("3. Exit")
	fmt.Print("Your choice: ")

	fmt.Scan(&option)

	if option == 1 {
		investmentCalculator()
	} else if option == 2 {
		profitCalculator()
	} else if option == 3 {
		fmt.Println("Exiting now...")
		os.Exit(0)
	} else {
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
