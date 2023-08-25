package leetcode

import "testing"

func TestMaxAbsSum(t *testing.T) {
	sum := MaxAbsoluteSum([]int{1, -3, 2, 3, -4})
	if sum != 5 {
		t.Errorf("Expected 5, got %d", sum)
	}

}
