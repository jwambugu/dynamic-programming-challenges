package main

import "fmt"

type memo map[int][]int

func howSum(targetSum int, numbers []int, m memo) []int {

	if _, exists := m[targetSum]; exists {
		return m[targetSum]
	}

	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	for _, number := range numbers {
		remainder := targetSum - number

		result := howSum(remainder, numbers, m)

		if result != nil {
			//fmt.Println(targetSum, "-", number, "=", remainder)

			result = append(result, number)
			m[targetSum] = result

			return m[targetSum]
		}
	}

	m[targetSum] = nil
	return m[targetSum]
}

func main() {
	fmt.Println(howSum(7, []int{2, 3}, memo{}))       // [3,2,2]
	fmt.Println(howSum(7, []int{5, 3, 4, 7}, memo{})) // [4,3]
	fmt.Println(howSum(7, []int{2, 4}, memo{}))       // []
	fmt.Println(howSum(8, []int{2, 3, 5}, memo{}))    // [2,2,2,2]
	fmt.Println(howSum(300, []int{7, 14}, memo{}))    // []
}
