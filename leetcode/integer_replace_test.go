package leetcode

import "testing"

func TestIntegerReplaceCase1(t *testing.T) {
	replacement := integerReplacement(8)
	if replacement != 3 {
		t.Errorf("expected 3, got %d", replacement)
	}
}

func TestIntegerReplaceCase2(t *testing.T) {
	replacement := integerReplacement(7)
	if replacement != 4 {
		t.Errorf("expected 4, got %d", replacement)
	}
}
