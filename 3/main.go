package main

import (
	"fmt"
)

var _ = fmt.Printf

func lengthOfLongestSubstring(s string) int {
	arr := make([]rune, 0, len(s))
	max := 0

	indexOf := func(t rune) (int, bool) {
		for i, r := range arr {
			if r == t {
				return i, true
			}
		}
		return -1, false
	}

	for _, r := range s {
		if i, ok := indexOf(r); ok {
			if m := len(arr); m > max {
				max = m
			}

			arr = arr[i+1:]
		}

		arr = append(arr, r)
	}

	if m := len(arr); m > max {
		max = m
	}

	return max
}
