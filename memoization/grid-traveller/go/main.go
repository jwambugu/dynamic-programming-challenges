package main

import "fmt"

func gridTraveller(m, n uint64, memo map[string]uint64) uint64 {
	key := fmt.Sprintf("%d,%d", m, n)

	if _, exists := memo[key]; exists {
		return memo[key]
	}

	if m == 1 && n == 1 {
		return 1
	}

	if m == 0 || n == 0 {
		return 0
	}

	memo[key] = gridTraveller(m-1, n, memo) + gridTraveller(m, n-1, memo)

	return memo[key]
}

func main() {
	fmt.Println(gridTraveller(1, 1, map[string]uint64{}))   // 1
	fmt.Println(gridTraveller(2, 3, map[string]uint64{}))   // 3
	fmt.Println(gridTraveller(3, 2, map[string]uint64{}))   // 3
	fmt.Println(gridTraveller(18, 18, map[string]uint64{})) // 2333606220
}
