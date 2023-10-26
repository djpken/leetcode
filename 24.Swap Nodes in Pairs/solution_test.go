package leetcode

import (
	"testing"
)

//Given a linked list, swap every two adjacent nodes and return its head. You
//must solve the problem without modifying the values in the list's nodes (i.e.,
//only nodes themselves may be changed.)
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4]
//Output: [2,1,4,3]
//
//
// Example 2:
//
//
//Input: head = []
//Output: []
//
//
// Example 3:
//
//
//Input: head = [1]
//Output: [1]
//
//
//
// Constraints:
//
//
// The number of nodes in the list is in the range [0, 100].
// 0 <= Node.val <= 100
//
//
// Related Topics Linked List Recursion 👍 11347 👎 414

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{-1, head}
	pre := dummy
	if head == nil || head.Next == nil {
		return head
	}
	temp := head
	cur := head.Next
	head = cur
	for cur != nil {
		temp.Next = cur.Next
		cur.Next = temp
		pre.Next = cur

		if temp.Next != nil {
			cur = temp.Next.Next
			pre = temp
			temp = temp.Next
		} else {
			cur = temp.Next
			pre = temp
			temp = temp.Next
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSwapNodesInPairs(t *testing.T) {

}
