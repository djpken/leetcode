package leetcode

import (
	"testing"
)

//Given an input string s, reverse the order of the words.
//
// A word is defined as a sequence of non-space characters. The words in s will
//be separated by at least one space.
//
// Return a string of the words in reverse order concatenated by a single space.
//
//
// Note that s may contain leading or trailing spaces or multiple spaces
//between two words. The returned string should only have a single space separating the
//words. Do not include any extra spaces.
//
//
// Example 1:
//
//
//Input: s = "the sky is blue"
//Output: "blue is sky the"
//
//
// Example 2:
//
//
//Input: s = "  hello world  "
//Output: "world hello"
//Explanation: Your reversed string should not contain leading or trailing
//spaces.
//
//
// Example 3:
//
//
//Input: s = "a good   example"
//Output: "example good a"
//Explanation: You need to reduce multiple spaces between two words to a single
//space in the reversed string.
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁴
// s contains English letters (upper-case and lower-case), digits, and spaces '
//'.
// There is at least one word in s.
//
//
//
// Follow-up: If the string data type is mutable in your language, can you
//solve it in-place with O(1) extra space?
//
// Related Topics Two Pointers String 👍 7518 👎 4928

// leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	b := []byte(s)
	slowIndex, fastIndex := 0, 0
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}

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

func TestReverseWordsInAString(t *testing.T) {

}
