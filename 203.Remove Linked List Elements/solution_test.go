package leetcode

import (
	"testing"
)

//Given the head of a linked list and an integer val, remove all the nodes of
//the linked list that has Node.val == val, and return the new head.
//
//
// Example 1:
//
//
//Input: head = [1,2,6,3,4,5,6], val = 6
//Output: [1,2,3,4,5]
//
//
// Example 2:
//
//
//Input: head = [], val = 1
//Output: []
//
//
// Example 3:
//
//
//Input: head = [7,7,7,7], val = 7
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the list is in the range [0, 10⁴].
// 1 <= Node.val <= 50
// 0 <= val <= 50
//
//
// Related Topics Linked List Recursion 👍 7875 👎 219

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{-1, head}
	pre := dummyNode
	cur := dummyNode.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return dummyNode.Next
}

// leetcode submit region end(Prohibit modification and deletion)
type ListNode struct {
	Val  int
	Next *ListNode
}

func TestRemoveLinkedListElements(t *testing.T) {

}
