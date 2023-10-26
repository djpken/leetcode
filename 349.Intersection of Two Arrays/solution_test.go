package leetcode

import (
	"testing"
)

//Given two integer arrays nums1 and nums2, return an array of their
//intersection. Each element in the result must be unique and you may return the result in
//any order.
//
//
// Example 1:
//
//
//Input: nums1 = [1,2,2,1], nums2 = [2,2]
//Output: [2]
//
//
// Example 2:
//
//
//Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//Output: [9,4]
//Explanation: [4,9] is also accepted.
//
//
//
// Constraints:
//
//
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000
//
//
// Related Topics Array Hash Table Two Pointers Binary Search Sorting 👍 5335 👎
// 2199

// leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	result := make([]int, 0)
	for _, i2 := range nums1 {
		set[i2] = struct{}{}
	}
	for _, i3 := range nums2 {
		if _, exist := set[i3]; exist {
			result = append(result, i3)
			delete(set, i3)
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIntersectionOfTwoArrays(t *testing.T) {

}
