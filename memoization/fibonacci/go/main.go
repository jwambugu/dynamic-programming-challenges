package main

import "fmt"

// fibonacci returns the number in the sequence
func fibonacci(n uint64, memo map[uint64]uint64) uint64 {
	if _, exists := memo[n]; exists {
		return memo[n]
	}

	if n <= 2 {
		return 1
	}

	memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)

	return memo[n]
}

func main() {
	//memo := map[uint64]uint64{}

	fmt.Println(fibonacci(6, map[uint64]uint64{}))  // 8
	fmt.Println(fibonacci(7, map[uint64]uint64{}))  // 13
	fmt.Println(fibonacci(8, map[uint64]uint64{}))  // 21
	fmt.Println(fibonacci(50, map[uint64]uint64{})) // 12586269025
}
