package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/quick"
)

// Example 1:

// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:

// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:

// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
//              Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func BenchmarkLengthOfLonestSubstring0(b *testing.B) {
	gen := func(count int) string {
		result := make([]rune, count)
		for i := 0; i < count; i++ {
			result[i] = rune(i)
		}
		return string(result)
	}

	b.Run("length 0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lengthOfLongestSubstring("")
		}
	})
	b.Run("length 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lengthOfLongestSubstring(" ")
		}
	})

	s10 := gen(10)
	b.Run("length 10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lengthOfLongestSubstring(s10)
		}
	})

	s100 := gen(100)
	b.Run("length 100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lengthOfLongestSubstring(s100)
		}
	})
}

func TestLengthOfLongestSubstring(t *testing.T) {
	reverse := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	assertExample := func(s string, l int) {
		assert.Equal(t, l, lengthOfLongestSubstring(s))
	}

	assertExample(" ", 1)
	assertExample("abcabcbb", 3)
	assertExample("bbbbb", 1)
	assertExample("pwwkew", 3)
	assertExample("\U0010260e𫩼罠\U0007d6f2\U000ff9a8\U000a8c63\U0009dba7\U000dabfd\U00102f3a\U0003e6e3\U0002fde9\U000a8e83\U000e4ea7\U000a2c7b\U000c846d죋\U00092638\U0007d183\U000dbdbd\U000f7b45\U0006172f\U0010f5bc\U000f83e7🀤\U000e17b5�\U000f4368\U00014026�\U000d77f3銀ᓶ乯\U000bb5a0\U0008b550\U000a7906𧢞\U0009772a\U000a5cf0\U0003bd6c\U000aa4ba\U000f143d\U000a4ab4\U00056db3\U000fd8a9", 28)
	assertExample(reverse("\U0010260e𫩼罠\U0007d6f2\U000ff9a8\U000a8c63\U0009dba7\U000dabfd\U00102f3a\U0003e6e3\U0002fde9\U000a8e83\U000e4ea7\U000a2c7b\U000c846d죋\U00092638\U0007d183\U000dbdbd\U000f7b45\U0006172f\U0010f5bc\U000f83e7🀤\U000e17b5�\U000f4368\U00014026�\U000d77f3銀ᓶ乯\U000bb5a0\U0008b550\U000a7906𧢞\U0009772a\U000a5cf0\U0003bd6c\U000aa4ba\U000f143d\U000a4ab4\U00056db3\U000fd8a9"), 28)

	f := func(s string) bool {
		return lengthOfLongestSubstring(s) >= 0
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}

	f = func(s string) bool {
		return lengthOfLongestSubstring(s) == lengthOfLongestSubstring(reverse(s))
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
