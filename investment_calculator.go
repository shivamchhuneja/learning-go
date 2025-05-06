package main

import (
	"fmt"
	"math"
)

func main() {
	const inflation = 7.5
	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Enter an investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter time period in years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflation/100, years)

	fmt.Println("Real Value: ", futureRealValue)
	fmt.Println("Future Value: ", futureValue)
}
