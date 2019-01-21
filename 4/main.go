package main

import (
	"sort"
)

type ref struct {
	dat *[]int
	idx int
}

func (r *ref) get() int {
	return (*r.dat)[r.idx]
}

func findMedianSortedArraysSlow(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	idx := n / 2

	if n%2 == 0 {
		x := float64(nums[idx] + nums[idx-1])
		return x / 2.0
	}

	return float64(nums[idx])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	medianIdx := (n / 2) + 1

	var n1, n2 int
	var penultimate, ultimate *ref

	push := func(nums *[]int, n *int) {
		penultimate = ultimate
		ultimate = &ref{nums, *n}
		(*n)++
	}

	for i := 0; i < medianIdx; i++ {
		if n1 >= len(nums1) {
			push(&nums2, &n2)
		} else if n2 >= len(nums2) {
			push(&nums1, &n1)
		} else if a, b := nums1[n1], nums2[n2]; a < b {
			push(&nums1, &n1)
		} else {
			push(&nums2, &n2)
		}
	}

	if n%2 == 0 {
		x := float64(ultimate.get() + penultimate.get())
		return x / 2.0
	}

	return float64(ultimate.get())
}
