package main

import (
	"fmt"
	"learnpackages/simpleinterest"
)

func init() {
	fmt.Printf("Main package Initialized\n")
}

func main() {

	simpleinterest.CalculateExample()
	fmt.Println("Simple Interest Calculation start from main")

	// Take input from user for the Principal Amount
	principalAmount := 4000.25

	// Take input from user for Rate Of Interest Per Annum
	rateOfInterest := 11.6

	// Take input from user for number for years
	period := 2.8

	// Calculates Simple Interest
	simpleInterestAmount := simpleinterest.Calculate(principalAmount, rateOfInterest, period)

	fmt.Printf("Simple Interest you will get on Principal Amount: %+v at ROI: %+v for %+v Years will be %+v\n", principalAmount, rateOfInterest,
		period, simpleInterestAmount)
}
