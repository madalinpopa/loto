package generator

import (
	"time"

	"math/rand"
)

// GenerateNumbers generates n unique random numbers between 1 and 49
func GenerateNumbers(n int) []int {
	// Seed the random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a slice to hold the numbers of n lenght
	numbers := make([]int, 0, n)

	// Create a map to keep track of used numbers
	used := make(map[int]bool)

	for len(numbers) < n {
		num := r.Intn(49) + 1
		if !used[num] {
			numbers = append(numbers, num)
			used[num] = true
		}

	}
	return numbers
}
