package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		j, ok := m[target-n]

		if ok {
			return []int{j, i}
		}

		m[nums[i]] = i
	}

	panic("Expected two members of nums to sum to target")
}

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))

	in := []int{2, 7, 11, 15}
	for i := 0; i < 100; i++ {
		t.Logf("%v", in)
		shuffle(in)
		assert.Equal(t, 2, len(twoSum(in, 9)))
	}
}

func shuffle(a interface{}) {
	sort.Slice(a, func(i, j int) bool {
		return rand.Float64() >= 0.5
	})
}
