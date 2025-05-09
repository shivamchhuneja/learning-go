package main

import (
	"fmt"
	"math"
)

func investmentCalculator() {
	const inflation = 7.5

	investmentAmount := getUserInput("Enter an investment amount: ")

	expectedReturnRate := getUserInput("Enter expected return rate: ")

	years := getUserInput("Enter time period in years: ")

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflation/100, years)

	fmt.Println("Real Value: ", futureRealValue)
	fmt.Println("Future Value: ", futureValue)
}
