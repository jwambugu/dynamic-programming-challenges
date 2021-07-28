package main

import (
	"fmt"
	"strings"
)

type memo map[string]bool

func canConstruct(target string, wordBank []string, m memo) bool {
	if _, exists := m[target]; exists {
		return m[target]
	}

	if target == "" {
		return true
	}

	for _, s := range wordBank {
		if strings.HasPrefix(target, s) {
			suffix := target[len(s):]

			if canConstruct(suffix, wordBank, m) {
				m[target] = true
				return m[target]
			}
		}
	}

	m[target] = false
	return m[target]
}

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, memo{})) // true

	fmt.Println(
		canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, memo{}),
	) // false

	fmt.Println(
		canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, memo{}),
	) // true

	fmt.Println(
		canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
			"e",
			"ee",
			"eee",
			"eeee",
			"eeeee",
		}, memo{}),
	) // false
}
