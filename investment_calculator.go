package main

import (
	"fmt"
	"math"
)

func main() {
	const inflation = 7.5
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflation/100, years)

	fmt.Println(futureRealValue)
	fmt.Println(futureValue)
}
