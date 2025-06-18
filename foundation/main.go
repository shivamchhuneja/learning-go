package main

import (
	"fmt"
	"os"

	"github.com/shivamchhuneja/learning-go/calculators"
	"github.com/shivamchhuneja/learning-go/structs"
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
	fmt.Println("4. Structs")
	fmt.Println("5. Exit")
	fmt.Print("Your choice: ")

	fmt.Scan(&option)

	switch option {

	case 1:
		calculators.InvestmentCalculator()
	case 2:
		calculators.ProfitCalculator()
	case 3:
		calculators.BankingMode()
	case 4:
		structs.Struct_fn()
	case 5:
		fmt.Println("Exiting now...")
		os.Exit(0)
	default:
		fmt.Println("Try again")
		runApp()
	}
}
