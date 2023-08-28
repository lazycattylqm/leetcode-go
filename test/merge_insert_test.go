package test

import (
	"com.example.liqiming/leetcode/leetcode"
	"fmt"
	"testing"
)

func TestMergeInsert(t *testing.T) {
	t.Log("Test Merge Insert")

	insert := leetcode.Insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})

	fmt.Println(insert)

}
