package thefarm

import (
    "errors"
    "fmt"
)

var ErrSomethingWrong = errors.New("Something went wrong")
// Divides food for cows
func DivideFood (calc FodderCalculator, num_cows int) (float64, error){
    val, err := calc.FodderAmount(num_cows)
    if err != nil{
        return 0, err
    }
    ff, err := calc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    return val * ff / float64(num_cows), nil
    
}

var ErrInvalidCows = errors.New("invalid number of cows")
// Validates no of cows and calls divide food
func ValidateInputAndDivideFood (calc FodderCalculator, num_cows int) (float64, error){
    if (num_cows <= 0){
        return 0, ErrInvalidCows
    }
    return DivideFood(calc, num_cows)
}


type ErrCowCount struct {
    count int
    message string
}
func (e *ErrCowCount) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.count, e.message)
}
// Verify the number of cows is valid
func ValidateNumberOfCows(num_cows int) error{
    if (num_cows < 0){
        return &ErrCowCount { count: num_cows, message: "there are no negative cows" }
    }
    if (num_cows == 0){
        return &ErrCowCount { count: 0, message: "no cows don't need food" }
    }
    return nil
}
