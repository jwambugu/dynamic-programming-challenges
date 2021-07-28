package main

import (
	"fmt"
	"strings"
)

type memo map[string]int

func countConstruct(target string, wordBank []string, m memo) int {
	if _, exists := m[target]; exists {
		return m[target]
	}

	if target == "" {
		return 1
	}

	totalCount := 0

	for _, s := range wordBank {
		if strings.HasPrefix(target, s) {
			suffix := target[len(s):]

			possibleWays := countConstruct(suffix, wordBank, m)
			totalCount += possibleWays
		}
	}

	m[target] = totalCount
	return m[target]
}

func main() {
	fmt.Println(countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, memo{})) // 2

	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, memo{})) // 1

	fmt.Println(
		countConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, memo{}),
	) // 0

	fmt.Println(
		countConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, memo{}),
	) // 4

	fmt.Println(
		countConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
			"e",
			"ee",
			"eee",
			"eeee",
			"eeeee",
		}, memo{}),
	) // 0
}
