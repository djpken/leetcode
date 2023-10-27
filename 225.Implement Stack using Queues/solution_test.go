package leetcode

import (
	"testing"
)

//Implement a last-in-first-out (LIFO) stack using only two queues. The
//implemented stack should support all the functions of a normal stack (push, top, pop,
//and empty).
//
// Implement the MyStack class:
//
//
// void push(int x) Pushes element x to the top of the stack.
// int pop() Removes the element on the top of the stack and returns it.
// int top() Returns the element on the top of the stack.
// boolean empty() Returns true if the stack is empty, false otherwise.
//
//
// Notes:
//
//
// You must use only standard operations of a queue, which means that only push
//to back, peek/pop from front, size and is empty operations are valid.
// Depending on your language, the queue may not be supported natively. You may
//simulate a queue using a list or deque (double-ended queue) as long as you use
//only a queue's standard operations.
//
//
//
// Example 1:
//
//
//Input
//["MyStack", "push", "push", "top", "pop", "empty"]
//[[], [1], [2], [], [], []]
//Output
//[null, null, null, 2, 2, false]
//
//Explanation
//MyStack myStack = new MyStack();
//myStack.push(1);
//myStack.push(2);
//myStack.top(); // return 2
//myStack.pop(); // return 2
//myStack.empty(); // return False
//
//
//
// Constraints:
//
//
// 1 <= x <= 9
// At most 100 calls will be made to push, pop, top, and empty.
// All the calls to pop and top are valid.
//
//
//
// Follow-up: Can you implement the stack using only one queue?
//
// Related Topics Stack Design Queue 👍 5658 👎 1161

// leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	queueIn  []int
	queueOut []int
}

func Constructor() MyStack {
	return MyStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.queueIn = append(this.queueIn, x)
}

func (this *MyStack) Pop() int {
	if len(this.queueIn) == 0 {
		return -1
	}
	for len(this.queueIn) > 1 {
		this.queueOut = append(this.queueOut, this.queueIn[0])
		this.queueIn = this.queueIn[1:]
	}
	val := this.queueIn[0]
	this.queueIn = []int{}
	for len(this.queueOut) > 0 {
		this.queueIn = append(this.queueIn, this.queueOut[0])
		this.queueOut = this.queueOut[1:]
	}
	return val
}

func (this *MyStack) Top() int {
	if len(this.queueIn) == 0 {
		return -1
	}
	return this.queueIn[len(this.queueIn)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.queueIn) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestImplementStackUsingQueues(t *testing.T) {

}
