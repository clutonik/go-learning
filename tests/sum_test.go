package tests

import "testing"

func TestSum(t *testing.T) {

	type test struct {
		data     []int
		expected int
	}

	// Table of tests
	tests := []test{
		test{[]int{2, 3, 4}, 9},
		test{[]int{2, 4}, 6},
		test{[]int{1, 3}, 4},
		test{[]int{4, 4, 4}, 12},
	}

	for _, v := range tests {
		got := Sum(v.data...)

		if got != v.expected {
			t.Error("Expected: ", v.expected, "Got: ", got)
		}
	}

}
