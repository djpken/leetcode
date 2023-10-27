package leetcode

import (
	"testing"
)

//Given a string s, check if it can be constructed by taking a substring of it
//and appending multiple copies of the substring together.
//
//
// Example 1:
//
//
//Input: s = "abab"
//Output: true
//Explanation: It is the substring "ab" twice.
//
//
// Example 2:
//
//
//Input: s = "aba"
//Output: false
//
//
// Example 3:
//
//
//Input: s = "abcabcabcabc"
//Output: true
//Explanation: It is the substring "abc" four times or the substring "abcabc"
//twice.
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁴
// s consists of lowercase English letters.
//
//
// Related Topics String String Matching 👍 6166 👎 489

// leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
	for i := 0; i < (len(s) / 2); i++ {
		if (len(s) % i) != 0 {
			continue
		}
		for j := 1; j < len(s)/i; i++ {
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRepeatedSubstringPattern(t *testing.T) {

}
