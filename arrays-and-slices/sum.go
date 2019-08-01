package main

// Sum takes an array of 5 numbers and returns the sum
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}

	return sum
}
