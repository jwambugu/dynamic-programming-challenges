package main

import "fmt"

type memo map[int][]int

type combinations [][]int

func bestSum(targetSum int, numbers []int, m memo) ([]int, combinations) {
	if _, exists := m[targetSum]; exists {
		return m[targetSum], nil
	}

	if targetSum == 0 {
		return []int{}, nil
	}

	if targetSum < 0 {
		return nil, nil
	}

	var shortestCombination []int
	var possibleCombinations combinations

	for _, number := range numbers {
		remainder := targetSum - number

		remainderCombination, _ := bestSum(remainder, numbers, m)

		if remainderCombination != nil {
			combination := append(remainderCombination, number)

			if shortestCombination == nil || len(combination) < len(shortestCombination) {
				possibleCombinations = append(possibleCombinations, combination)

				shortestCombination = combination
			}
		}

	}

	m[targetSum] = shortestCombination
	return m[targetSum], possibleCombinations
}

func main() {
	fmt.Println(bestSum(7, []int{5, 3, 4, 7}, memo{})) // [7] [[4 3] [7]]
	fmt.Println(bestSum(8, []int{2, 3, 5}, memo{}))    // [5 3] [[3 3 2] [5 3]]
	fmt.Println(bestSum(8, []int{1, 4, 5}, memo{}))    // [4 4] [[5 1 1 1] [4 4]]
	fmt.Println(bestSum(100, []int{1, 2, 5, 25}, memo{}))
	// [25 25 25 25] [[25 25 25 5 5 5 5 5 2 1] [25 25 25 5 5 5 5 5] [25 25 25 25]]
}
