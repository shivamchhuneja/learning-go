package calculators

import (
	"fmt"
	"math"

	"github.com/shivamchhuneja/learning-go/helpers"
)

func InvestmentCalculator() {
	const inflation = 7.5

	investmentAmount, err := helpers.GetUserInput("Enter an investment amount: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expectedReturnRate, err := helpers.GetUserInput("Enter expected return rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	years, err := helpers.GetUserInput("Enter time period in years: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflation/100, years)

	fmt.Println("Real Value: ", futureRealValue)
	fmt.Println("Future Value: ", futureValue)
}
