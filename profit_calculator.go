package main

import "fmt"

func profitCalculator() {
	revenue := getUserInput("What is the revenue: ")
	expense := getUserInput("What is the expense: ")
	tax := getUserInput("What is the tax rate in %: ")

	pbt, pat := CalculateFinancials(revenue, expense, tax)

	fmt.Printf("Your profit before tax is %.2f & profit after tax is %.2f\n", pbt, pat)
}

func CalculateFinancials(revenue, expense, tax float64) (float64, float64) {
	pbt := revenue - expense
	pat := pbt * (1 - (tax / 100))
	return pbt, pat
}
