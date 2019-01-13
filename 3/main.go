package main

import (
	"fmt"
)

var _ = fmt.Printf

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	length := len(runes)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return 1
	}

	max := 0
	start := -1
	m := make(map[rune]int)

	for current, r := range runes {
		old, ok := m[r]
		m[r] = current

		if ok && old > start {
			if delta := current - start - 1; delta > max {
				max = delta
			}
			start = old
		} else if current == length-1 {
			if delta := length - start - 1; delta > max {
				max = delta
			}
		}
	}

	return max
}
