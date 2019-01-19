require "test/unit"

extend Test::Unit::Assertions

# @param {Integer[]} nums1
# @param {Integer[]} nums2
# @return {Float}
# @param {Integer[]} nums1
# @param {Integer[]} nums2
# @return {Float}
def find_median_sorted_arrays(nums1, nums2)
  s = (nums1 + nums2).sort
  l = s.length
  if l.odd?
    s[(l / 2)]
  else
    (s[l / 2 - 1] + s[l / 2]) / 2.0
  end
end

assert_equal find_median_sorted_arrays([1, 2, 3], []), 2
assert_equal find_median_sorted_arrays([1, 2], [3, 4]), 2.5
assert_equal find_median_sorted_arrays([1, 3], [2]), 2
assert_equal find_median_sorted_arrays([1, 2], [1, 2]), 1.5
assert_equal find_median_sorted_arrays([1, 3, 5], [2, 4, 6]), 3.5
assert_equal find_median_sorted_arrays([100, 101, 109], [1, 2, 9, 10, 100]), 55.0
assert_equal find_median_sorted_arrays([1], [1]), 1.0
assert_equal find_median_sorted_arrays([1], []), 1.0
assert_equal find_median_sorted_arrays([], [1, 2, 3]), 2
assert_equal find_median_sorted_arrays([], [1]), 1.0
