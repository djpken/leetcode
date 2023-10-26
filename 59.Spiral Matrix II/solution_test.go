package leetcode

import (
	"reflect"
	"testing"
)

//Given a positive integer n, generate an n x n matrix filled with elements
//from 1 to n² in spiral order.
//
//
// Example 1:
//
//
//Input: n = 3
//Output: [[1,2,3],[8,9,4],[7,6,5]]
//
//
// Example 2:
//
//
//Input: n = 1
//Output: [[1]]
//
//
//
// Constraints:
//
//
// 1 <= n <= 20
//
//
// Related Topics Array Matrix Simulation 👍 6070 👎 244

// leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	time := n / 2
	end := n - 1
	num := 1

	for i := 0; i < time; i++ {
		for j := i; j < end; j++ {
			matrix[i][j] = num
			num++
		}
		for j := i; j < end; j++ {
			matrix[j][end] = num
			num++
		}
		for j := end; j > i; j-- {
			matrix[end][j] = num
			num++
		}
		for j := end; j > i; j-- {
			matrix[j][i] = num
			num++
		}
		end--
	}
	if n%2 == 1 {
		matrix[time][time] = n * n
	}
	return matrix
}

//leetcode submit region end(Prohibit modification and deletion)

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		n      int
		result [][]int
	}{
		{
			1,
			[][]int{
				{1},
			},
		},
		{
			2,
			[][]int{
				{1, 2},
				{4, 3},
			},
		},
		{
			3,
			[][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			4,
			[][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
		// ... you can add more cases if needed
	}

	for _, test := range tests {
		got := generateMatrix(test.n)
		if !reflect.DeepEqual(got, test.result) {
			t.Errorf("For n = %v, got %v, want %v", test.n, got, test.result)
		}
	}
}
