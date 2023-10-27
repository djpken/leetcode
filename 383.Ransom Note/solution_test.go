package leetcode

import (
	"testing"
)

//Given two strings ransomNote and magazine, return true if ransomNote can be
//constructed by using the letters from magazine and false otherwise.
//
// Each letter in magazine can only be used once in ransomNote.
//
//
// Example 1:
// Input: ransomNote = "a", magazine = "b"
//Output: false
//
// Example 2:
// Input: ransomNote = "aa", magazine = "ab"
//Output: false
//
// Example 3:
// Input: ransomNote = "aa", magazine = "aab"
//Output: true
//
//
// Constraints:
//
//
// 1 <= ransomNote.length, magazine.length <= 10⁵
// ransomNote and magazine consist of lowercase English letters.
//
//
// Related Topics Hash Table String Counting 👍 4590 👎 463

// leetcode submit region begin(Prohibit modification and deletion)
func canConstruct(ransomNote string, magazine string) bool {
	record := make([]int, 26)
	for _, s := range magazine {
		record[s-'a']++
	}
	for _, s := range ransomNote {
		record[s-'a']--
		if record[s-'a'] < 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRansomNote(t *testing.T) {

}
