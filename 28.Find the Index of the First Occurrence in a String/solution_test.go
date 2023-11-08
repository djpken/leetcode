package leetcode

import (
	"log"
	"testing"
)

//Given two strings needle and haystack, return the index of the first
//occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//
// Example 1:
//
//
//Input: haystack = "sadbutsad", needle = "sad"
//Output: 0
//Explanation: "sad" occurs at index 0 and 6.
//The first occurrence is at index 0, so we return 0.
//
//
// Example 2:
//
//
//Input: haystack = "leetcode", needle = "leeto"
//Output: -1
//Explanation: "leeto" did not occur in "leetcode", so we return -1.
//
//
//
// Constraints:
//
//
// 1 <= haystack.length, needle.length <= 10⁴
// haystack and needle consist of only lowercase English characters.
//
//
// Related Topics Two Pointers String String Matching 👍 4981 👎 285

// leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	next := prefixTable(needle)
	j := -1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j]
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j+1 == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}
func prefixTable(str string) []int {
	res := make([]int, len(str))
	j := -1
	res[0] = j
	for i := 1; i < len(str); i++ {
		for j >= 0 && str[i] != str[j+1] {
			j = res[j]
		}
		if str[i] == str[j+1] {
			j++
		}
		res[i] = j
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindTheIndexOfTheFirstOccurrenceInAString(t *testing.T) {
	log.Printf("%v", prefixTable("aabaaac"))
	log.Printf("%v", getNext("aabaaac"))
}

func getNext(s string) []int {
	j := -1 // j表示 最长相等前后缀长度
	next := make([]int, len(s))
	next[0] = j

	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j] // 回退前一位
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j // next[i]是i（包括i）之前的最长相等前后缀长度
	}
	return next
}
