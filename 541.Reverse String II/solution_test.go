package leetcode

import (
	"testing"
)

//Given a string s and an integer k, reverse the first k characters for every 2
//k characters counting from the start of the string.
//
// If there are fewer than k characters left, reverse all of them. If there are
//less than 2k but greater than or equal to k characters, then reverse the first
//k characters and leave the other as original.
//
//
// Example 1:
// Input: s = "abcdefg", k = 2
//Output: "bacdfeg"
//
// Example 2:
// Input: s = "abcd", k = 2
//Output: "bacd"
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁴
// s consists of only lowercase English letters.
// 1 <= k <= 10⁴
//
//
// Related Topics Two Pointers String 👍 1810 👎 3542

// leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	ss := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		if i+k <= len(s) {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:len(s)])
		}

	}
	return string(ss)
}
func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseStringIi(t *testing.T) {

}
