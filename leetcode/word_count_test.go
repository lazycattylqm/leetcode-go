package leetcode

import "testing"

func Test_countWords(t *testing.T) {
	words1 := []string{"leetcode", "is", "amazing", "as", "is"}
	words2 := []string{"amazing", "leetcode", "is"}
	res := countWords(words1, words2)
	if res != 2 {
		t.Errorf("Expected 2, got %d", res)
	}

}
