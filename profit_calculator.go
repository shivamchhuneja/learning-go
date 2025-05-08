package main

import "fmt"

func profitCalculator() {
	var revenue, expense, tax, PBT, PAT float64

	fmt.Print("What is the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("What is the expense: ")
	fmt.Scan(&expense)

	fmt.Print("What is the tax rate in %: ")
	fmt.Scan(&tax)

	PBT = revenue - expense

	PAT = PBT * (1 - (tax / 100))

	fmt.Printf("Your profit before tax is: %.2f\n", PBT)
	fmt.Printf("Your profit after tax is: %.2f\n", PAT)
}
