package leetcode

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	stones := numJewelsInStones("aA", "aAAbbbb")
	if stones != 3 {
		t.Errorf("expected 3, got %d", stones)
	}

}
