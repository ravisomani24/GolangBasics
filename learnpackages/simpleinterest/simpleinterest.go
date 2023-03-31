package simpleinterest

import (
	"fmt"
)

var principal, rateOfInterest, period float64

func init() {
	fmt.Printf("SimpleInterest Package Initialized\n")
	principal = 5000.0
	rateOfInterest = 10.0
	period = 2.0
}

// Calculate takes Principal Amount, Rate Of Interest and Time and return Simple Interest for these values
// This function is capitalised because we want this function as exported function, so that now this function can we called from other packages
func Calculate(principal, rateOfInterest, period float64) float64 {
	var finalInterest float64
	finalInterest = principal * (rateOfInterest / 100) * period
	return finalInterest
}

// CalculateExample takes value of Principal Amount, Rate Of Interest and Time from Global Variables and these Variables are initialized in init()
// And this init is called during package initialization
func CalculateExample() {
	finalInterest := principal * (rateOfInterest / 100) * period
	fmt.Printf("Simple Interest you will get on Principal Amount: %+v at ROI: %+v for %+v Years will be %+v\n", principal, rateOfInterest,
		period, finalInterest)
}
