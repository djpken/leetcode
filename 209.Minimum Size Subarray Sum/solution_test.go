package leetcode

import (
	"math"
	"testing"
)

//Given an array of positive integers nums and a positive integer target,
//return the minimal length of a subarray whose sum is greater than or equal to target.
//If there is no such subarray, return 0 instead.
//
//
// Example 1:
//
//
//Input: target = 7, nums = [2,3,1,2,4,3]
//Output: 2
//Explanation: The subarray [4,3] has the minimal length under the problem
//constraint.
//
//
// Example 2:
//
//
//Input: target = 4, nums = [1,4,4]
//Output: 1
//
//
// Example 3:
//
//
//Input: target = 11, nums = [1,1,1,1,1,1,1,1]
//Output: 0
//
//
//
// Constraints:
//
//
// 1 <= target <= 10⁹
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁴
//
//
//
//Follow up: If you have figured out the
//O(n) solution, try coding another solution of which the time complexity is
//O(n log(n)).
//
// Related Topics Array Binary Search Sliding Window Prefix Sum 👍 11826 👎 357

// leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	left, right, sum, result := 0, 0, 0, math.MaxInt
	for right < length {
		sum += nums[right]
		right++
		for sum >= target {
			if right-left < result {
				result = right - left
			}
			sum -= nums[left]
			left++
		}
	}
	if result == math.MaxInt {
		return 0
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumSizeSubarraySum(t *testing.T) {
	tests := []struct {
		target int
		nums   []int
		want   int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}
	for _, tt := range tests {
		got := minSubArrayLen(tt.target, tt.nums)
		if got != tt.want {
			t.Errorf("target = %v, nums = %v, got %v, want %v", tt.target, tt.nums, got, tt.want)
		}
	}
}
