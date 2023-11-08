package leetcode

import (
	"container/list"
	"testing"
)

//Given the root of a binary tree, return the average value of the nodes on
//each level in the form of an array. Answers within 10⁻⁵ of the actual answer will
//be accepted.
//
//
// Example 1:
//
//
//Input: root = [3,9,20,null,null,15,7]
//Output: [3.00000,14.50000,11.00000]
//Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5,
//and on level 2 is 11.
//Hence return [3, 14.5, 11].
//
//
// Example 2:
//
//
//Input: root = [3,9,20,15,7]
//Output: [3.00000,14.50000,11.00000]
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 10⁴].
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 50
//32 👎 307

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)
	queue := list.New()
	queue.PushBack(root)
	var sum int
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
			sum += node.Val
		}
		res = append(res, float64(sum)/float64(length))
		sum = 0
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestAverageOfLevelsInBinaryTree(t *testing.T) {

}
