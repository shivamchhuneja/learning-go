package main

import "fmt"

func main() {
	runApp()
}

func runApp() {
	var option int

	fmt.Println("Select the calculator")
	fmt.Println("1. Investment Calculator")
	fmt.Println("2. Profit Calculator")
	fmt.Print("Your choice: ")

	fmt.Scan(&option)

	if option == 1 {
		investmentCalculator()
	} else if option == 2 {
		profitCalculator()
	} else {
		fmt.Println("Try again")
		runApp()
	}
}
