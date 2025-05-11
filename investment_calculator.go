package main

import (
	"fmt"
	"math"
)

func investmentCalculator() {
	const inflation = 7.5

	investmentAmount, err := getUserInput("Enter an investment amount: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expectedReturnRate, err := getUserInput("Enter expected return rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	years, err := getUserInput("Enter time period in years: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflation/100, years)

	fmt.Println("Real Value: ", futureRealValue)
	fmt.Println("Future Value: ", futureValue)
}
