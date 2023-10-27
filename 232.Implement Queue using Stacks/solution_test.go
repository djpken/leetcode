package leetcode

import (
	"testing"
)

//Implement a first in first out (FIFO) queue using only two stacks. The
//implemented queue should support all the functions of a normal queue (push, peek, pop,
//and empty).
//
// Implement the MyQueue class:
//
//
// void push(int x) Pushes element x to the back of the queue.
// int pop() Removes the element from the front of the queue and returns it.
// int peek() Returns the element at the front of the queue.
// boolean empty() Returns true if the queue is empty, false otherwise.
//
//
// Notes:
//
//
// You must use only standard operations of a stack, which means only push to
//top, peek/pop from top, size, and is empty operations are valid.
// Depending on your language, the stack may not be supported natively. You may
//simulate a stack using a list or deque (double-ended queue) as long as you use
//only a stack's standard operations.
//
//
//
// Example 1:
//
//
//Input
//["MyQueue", "push", "push", "peek", "pop", "empty"]
//[[], [1], [2], [], [], []]
//Output
//[null, null, null, 1, 1, false]
//
//Explanation
//MyQueue myQueue = new MyQueue();
//myQueue.push(1); // queue is: [1]
//myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
//myQueue.peek(); // return 1
//myQueue.pop(); // return 1, queue is [2]
//myQueue.empty(); // return false
//
//
//
// Constraints:
//
//
// 1 <= x <= 9
// At most 100 calls will be made to push, pop, peek, and empty.
// All the calls to pop and peek are valid.
//
//
//
// Follow-up: Can you implement the queue such that each operation is amortized
//O(1) time complexity? In other words, performing n operations will take overall
//O(n) time even if one of those operations may take longer.
//
// Related Topics Stack Design Queue 👍 6774 👎 384

// leetcode submit region begin(Prohibit modification and deletion)
type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stackOut) == 0 {
		if len(this.stackIn) == 0 {
			return -1
		}
		for i := len(this.stackIn) - 1; i >= 0; i-- {
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = []int{}
	}
	val := this.stackOut[len(this.stackOut)-1]
	this.stackOut = this.stackOut[:len(this.stackOut)-1]
	return val
}

func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == -1 {
		return -1
	}
	this.stackOut = append(this.stackOut, val)
	return val

}

func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestImplementQueueUsingStacks(t *testing.T) {

}
