package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	idx := len(nums) / 2

	if idx % 2 == 0 {
		n := float64(nums[idx] + nums[idx - 1])
		return  n / 2.0
	}

	return float64(nums[idx])
}

func TestFindMedianSortedArrays(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, 2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
	})

	t.Run("Example 2", func(t *testing.T) {
		assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	})

	t.Run("Empty argument", func(t *testing.T) {
		num1 := []int{1, 2, 3}
		nums2 := []int{}
		assert.Equal(t, float64(2), findMedianSortedArrays(nums2, num1))
		assert.Equal(t, float64(2), findMedianSortedArrays(num1, nums2))
	})
}
