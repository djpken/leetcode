package leetcode

import (
	"testing"
)

//Given two strings s and t, return true if t is an anagram of s, and false
//otherwise.
//
// An Anagram is a word or phrase formed by rearranging the letters of a
//different word or phrase, typically using all the original letters exactly once.
//
//
// Example 1:
// Input: s = "anagram", t = "nagaram"
//Output: true
//
// Example 2:
// Input: s = "rat", t = "car"
//Output: false
//
//
// Constraints:
//
//
// 1 <= s.length, t.length <= 5 * 10⁴
// s and t consist of lowercase English letters.
//
//
//
// Follow up: What if the inputs contain Unicode characters? How would you
//adapt your solution to such a case?
//
// Related Topics Hash Table String Sorting 👍 10836 👎 341

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	record := [26]int{}
	for _, r := range s {
		record[r-rune('a')]++
	}
	for _, r := range t {
		record[r-rune('a')]--
	}
	return record == [26]int{}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidAnagram(t *testing.T) {

}
