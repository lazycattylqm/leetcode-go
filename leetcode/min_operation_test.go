package leetcode

import "testing"

func TestMinOPCase1(t *testing.T) {
	logs := []string{"d1/", "d2/", "../", "d21/", "./"}
	ans := minOperations(logs)
	if ans != 2 {
		t.Errorf("expected 2, got %d", ans)
	}
}

func TestMinOperationCase2(t *testing.T) {
	logs := []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}
	ans := minOperations(logs)
	if ans != 3 {
		t.Errorf("expected 3, got %d", ans)
	}

}

func TestMinOperationCase3(t *testing.T) {
	s := "0100"
	ans := minOperationsStr(s)
	if ans != 1 {
		t.Errorf("expected 1, got %d", ans)
	}
}
