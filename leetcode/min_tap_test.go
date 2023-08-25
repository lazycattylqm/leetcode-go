package leetcode

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	taps := minTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1})
	if taps != 3 {
		_ = fmt.Errorf("taps: %d but should be 3", taps)
	}
}
