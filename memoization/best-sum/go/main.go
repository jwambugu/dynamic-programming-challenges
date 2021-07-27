package main

import "fmt"

type memo map[int][]int

func bestSum(targetSum int, numbers []int, m memo) []int {
	if _, exists := m[targetSum]; exists {
		return m[targetSum]
	}

	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	var shortestCombination []int

	for _, number := range numbers {
		remainder := targetSum - number
		remainderCombination := bestSum(remainder, numbers, m)

		if remainderCombination != nil {
			combination := append(remainderCombination, number)

			if shortestCombination == nil || len(combination) < len(shortestCombination) {
				shortestCombination = combination
			}
		}

	}

	m[targetSum] = shortestCombination
	return m[targetSum]
}

func main() {
	fmt.Println(bestSum(7, []int{5, 3, 4, 7}, memo{}))    // [7]
	fmt.Println(bestSum(8, []int{2, 3, 5}, memo{}))       // [3,5]
	fmt.Println(bestSum(8, []int{1, 4, 5}, memo{}))       // [4,4]
	fmt.Println(bestSum(100, []int{1, 2, 5, 25}, memo{})) // [25,25,25,25]
}
