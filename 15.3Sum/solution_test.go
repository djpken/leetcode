package leetcode

import (
	"sort"
	"testing"
)

//Given an integer array nums, return all the triplets [nums[i], nums[j], nums[
//k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.
//
//
// Example 1:
//
//
//Input: nums = [-1,0,1,2,-1,-4]
//Output: [[-1,-1,2],[-1,0,1]]
//Explanation:
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
//The distinct triplets are [-1,0,1] and [-1,-1,2].
//Notice that the order of the output and the order of the triplets does not
//matter.
//
//
// Example 2:
//
//
//Input: nums = [0,1,1]
//Output: []
//Explanation: The only possible triplet does not sum up to 0.
//
//
// Example 3:
//
//
//Input: nums = [0,0,0]
//Output: [[0,0,0]]
//Explanation: The only possible triplet sums up to 0.
//
//
//
// Constraints:
//
//
// 3 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
//
// Related Topics Array Two Pointers Sorting 👍 28897 👎 2601

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	i := make([][]int, 0)
	for cur := 0; cur < len(nums)-2; cur++ {
		if cur > 0 && nums[cur] == nums[cur-1] {
			continue
		}
		left := cur + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[cur] + nums[left] + nums[right]
			if sum == 0 {
				i = append(i, []int{nums[cur], nums[left], nums[right]})
				right--
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			}
			if sum < 0 {
				left++
			}
			if sum > 0 {
				right--
			}
		}

	}
	return i
}

//leetcode submit region end(Prohibit modification and deletion)

func TestThreeSum(t *testing.T) {

}
