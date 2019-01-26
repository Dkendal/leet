package main

import (
	"sync"
)


func longestPalindrome(s string) string {
	var result struct {
		m sync.RWMutex
		v int
	}

	for i := 0; i <= len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			println(s[i:j])
		}
	}
}

func isPalindrome(s string) bool {
	length := len(s)
	even := length%2 == 0
	lower := length / 2
	upper := lower

	if even {
		lower--
	} else {
		lower--
		upper++
	}

	for ; lower >= 0; lower, upper = lower-1, upper+1 {
		if s[lower] != s[upper] {
			return false
		}
	}

	return true
}
