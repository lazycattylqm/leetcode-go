package leetcode

import "testing"

func TestCountVowelStrings1(t *testing.T) {

	// ...
	i := countVowelStrings(1)
	if i != 5 {
		t.Errorf("expected 5, got %d", i)
	}
}

func TestCountVowelStrings2(t *testing.T) {

	// ...
	i := countVowelStrings(2)
	if i != 15 {
		t.Errorf("expected 15, got %d", i)
	}
}

func TestCountVowelStrings3(t *testing.T) {

	// ...
	i := countVowelStrings(33)
	if i != 66045 {
		t.Errorf("expected 66045, got %d", i)
	}

}
