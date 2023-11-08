package leetcode

import (
	"container/list"
	"testing"
)

//Given the root of a binary tree, imagine yourself standing on the right side
//of it, return the values of the nodes you can see ordered from top to bottom.
//
//
// Example 1:
//
//
//Input: root = [1,2,3,null,5,null,4]
//Output: [1,3,4]
//
//
// Example 2:
//
//
//Input: root = [1,null,3]
//Output: [1,3]
//
//
// Example 3:
//
//
//Input: root = []
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 11
//236 👎 764

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if i == length-1 {
				result = append(result, node.Val)
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestBinaryTreeRightSideView(t *testing.T) {

}
