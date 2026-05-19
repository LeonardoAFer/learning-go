package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var exepectedReturnRate = 5.5
	var years float64 = 10
	const inflationRate = 1.2

	futureValue := investmentAmount * math.Pow(1+exepectedReturnRate/100, years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future value:", futureValue)
	fmt.Println("Real value:", realFutureValue)
}
