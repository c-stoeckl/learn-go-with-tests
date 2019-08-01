package main

// Sum takes an array of 5 numbers and returns the sum
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}
