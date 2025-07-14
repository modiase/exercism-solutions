package chance

import (
        "math/rand"
        "time"
)

func RollADie() int {
    return 1 + rand.Intn(6)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12.0
}

// Factorial computes the factorial of n.
func Factorial(n int) int {
        if n == 0 {
                return 1
        }
        result := 1
        for i := 1; i <= n; i++ {
                result *= i
        }
        return result
}


var animals = []string { "ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog" }
// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
		rand.Seed(time.Now().UnixNano())
        n := len(animals)
        permutation := make([]string, n)
        elements := make([]int, n)
        for i := 0; i < n; i++ {
                elements[i] = i
        }


        randomNumber := rand.Intn(Factorial(n))

        for i := 0; i < n; i++ {
                factorial := Factorial(n - 1 - i)
                index := randomNumber / factorial
                randomNumber %= factorial

                permutation[i] = animals[elements[index]]
                elements = append(elements[:index], elements[index+1:]...)
        }

        return permutation
}

