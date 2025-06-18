package calculators

import (
	"fmt"

	"github.com/shivamchhuneja/learning-go/helpers"
)

func ProfitCalculator() {
	revenue, err := helpers.GetUserInput("What is the revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expense, err := helpers.GetUserInput("What is the expense: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	tax, err := helpers.GetUserInput("What is the tax rate in %: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	pbt, pat := calculateFinancials(revenue, expense, tax)

	fmt.Printf("Your profit before tax is %.2f & profit after tax is %.2f\n", pbt, pat)
}

func calculateFinancials(revenue, expense, tax float64) (float64, float64) {
	pbt := revenue - expense
	pat := pbt * (1 - (tax / 100))
	return pbt, pat
}
