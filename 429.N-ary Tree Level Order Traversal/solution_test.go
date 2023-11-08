package leetcode

import (
	"container/list"
	"testing"
)

//Given an n-ary tree, return the level order traversal of its nodes' values.
//
// Nary-Tree input serialization is represented in their level order traversal,
//each group of children is separated by the null value (See examples).
//
//
// Example 1:
//
//
//
//
//Input: root = [1,null,3,2,4,null,5,6]
//Output: [[1],[3,2,4],[5,6]]
//
//
// Example 2:
//
//
//
//
//Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,
//null,12,null,13,null,null,14]
//Output: [[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
//
//
//
// Constraints:
//
//
// The height of the n-ary tree is less than or equal to 1000
// The total number of nodes is between [0, 10⁴]
//
//
// Related Topics Tree Breadth-First Search 👍 3533 👎 135

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    var result [][]int
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		var tmp []int
		for i := 0; i < length; i++ {
			cur := queue.Remove(queue.Front()).(*Node)
            tmp = append(tmp, cur.Val)
			for j := 0; j < len(cur.Children); j++ {
				queue.PushBack(cur.Children[j])
			}
		}
		result = append(result, tmp)
	}
	return result
}

// leetcode submit region end(Prohibit modification and deletion)
type Node struct {
	Val      int
	Children []*Node
}

func TestNAryTreeLevelOrderTraversal(t *testing.T) {

}
