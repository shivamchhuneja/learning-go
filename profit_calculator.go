package main

import (
	"fmt"
	// "errors"
)

func profitCalculator() {
	revenue, err := getUserInput("What is the revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expense, err := getUserInput("What is the expense: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	tax, err := getUserInput("What is the tax rate in %: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	pbt, pat := CalculateFinancials(revenue, expense, tax)

	fmt.Printf("Your profit before tax is %.2f & profit after tax is %.2f\n", pbt, pat)
}

func CalculateFinancials(revenue, expense, tax float64) (float64, float64) {
	pbt := revenue - expense
	pat := pbt * (1 - (tax / 100))
	return pbt, pat
}
