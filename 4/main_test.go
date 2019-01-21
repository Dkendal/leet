package main

import (
	"sort"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArraysSlow(t *testing.T) {
	Suite(t, findMedianSortedArraysSlow)
}

func Test_findMedianSortedArrays(t *testing.T) {
	Suite(t, findMedianSortedArrays)

	t.Run("Equivalancy test", func(t *testing.T) {
		f := func(a []int, b []int) bool {
			// Atleast one must have a length greater than zero
			if len(a)+len(b) == 0 {
				return true
			}

			sort.Ints(a)
			sort.Ints(b)

			return assert.Equal(t,
				findMedianSortedArraysSlow(a, b),
				findMedianSortedArrays(a, b),
			)
		}

		quick.Check(f, nil)
	})
}

func Suite(t *testing.T, fn func([]int, []int) float64) {
	test := func(t *testing.T, expected float64, nums1 []int, nums2 []int) bool {
		return assert.Equal(
			t,
			expected,
			fn(nums1, nums2),
		) && assert.Equal(
			t,
			expected,
			fn(nums2, nums1),
		)
	}

	test(t, 1.0, []int{1}, []int{1})
	test(t, 1.0, []int{1}, []int{})
	test(t, 1.0, []int{1}, []int{})
	test(t, 1.0, []int{}, []int{1})
	test(t, 1.5, []int{1, 2}, []int{1, 2})
	test(t, 1.5, []int{1, 2}, []int{})
	test(t, 2, []int{}, []int{1, 2, 3})
	test(t, 2.0, []int{1, 2, 3}, []int{})
	test(t, 2.0, []int{1, 3}, []int{2})
	test(t, 2.5, []int{1, 2}, []int{3, 4})
	test(t, 3.0, []int{1, 3, 9}, []int{})
	test(t, 3.5, []int{1, 3, 5}, []int{2, 4, 6})
	test(t, 55.0, []int{100, 101, 109}, []int{1, 2, 9, 10, 100})

	t.Run("Example 1", func(t *testing.T) {
		test(t, 2.0, []int{1, 3}, []int{2})
	})

	t.Run("Example 2", func(t *testing.T) {
		test(t, 2.5, []int{1, 2}, []int{3, 4})
	})

	t.Run("Property test", func(t *testing.T) {
		f := func(a []int, b []int) bool {
			// At least one must have a length greater than zero
			if len(a)+len(b) == 0 {
				return true
			}

			sort.Ints(a)
			sort.Ints(b)

			return assert.NotPanicsf(t, func() {
				fn(a, b)
			},
				"input: %v %v", a, b,
			) && assert.NotPanicsf(t, func() {
				fn(b, a)
			}, "input %v %v", b, a,
			) && assert.Equal(t,
				fn(a, b),
				fn(b, a),
			)
		}

		quick.Check(f, nil)
	})
}
