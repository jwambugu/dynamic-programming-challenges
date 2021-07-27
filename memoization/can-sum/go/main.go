package main

import "fmt"

type memo map[int]bool

func canSum(targetSum int, numbers []int, m memo) bool {
	if _, exists := m[targetSum]; exists {
		return m[targetSum]
	}

	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, number := range numbers {
		remainder := targetSum - number

		if canSum(remainder, numbers, m) {
			m[remainder] = true
			return true
		}
	}

	m[targetSum] = false
	return false
}

func main() {
	fmt.Println(canSum(7, []int{2, 3}, memo{}))       // true
	fmt.Println(canSum(7, []int{5, 3, 4, 7}, memo{})) // true
	fmt.Println(canSum(7, []int{2, 4}, memo{}))       // false
	fmt.Println(canSum(8, []int{2, 3, 5}, memo{}))    // true
	fmt.Println(canSum(300, []int{7, 14}, memo{}))    // false
}
