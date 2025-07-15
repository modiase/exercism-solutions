package interest

import (
    m "math"
)

type amountWithRate struct {
	amount float64
	rate   float32
}
var interestRates = []amountWithRate{
	{m.Inf(-1), 3.213},
	{0, 0.5},
	{1000, 1.621},
	{5000, 2.475},
}
// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	low, high := 0, len(interestRates)-1
	ans := 0

	for low <= high {
		mid := low + (high-low)/2
		if interestRates[mid].amount <= balance {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return interestRates[ans].rate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance))*balance / 100.0
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	currentBalance, year := balance, 0
    for true{
        if currentBalance >= targetBalance { 
        	return year
        }
		currentBalance = AnnualBalanceUpdate(currentBalance)
        year += 1
        
    }
    panic("Never")
}
