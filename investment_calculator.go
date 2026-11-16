// Investment Calculator is a command-line program that estimates the future
// value of a one-time investment using compound interest, both in nominal
// terms and adjusted for inflation.
//
// The user is prompted for the initial amount, the expected annual return
// rate (in percent) and the number of years. Inflation is fixed at 2.5% a year.
package main

import (
	"fmt"
	"math"
)

// main reads the investment parameters from standard input, computes the
// future value and its inflation-adjusted equivalent, and prints both
// rounded to two decimal places.
func main() {
	// inflationRate is the assumed annual inflation, in percent.
	const inflationRate float64 = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate (%): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	// Compound interest: FV = P * (1 + r)^n
	// Real value: FV / (1 + i)^n, where i is the inflation rate.
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
}
