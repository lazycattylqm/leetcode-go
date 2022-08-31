package leetcode

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	unique := permuteUnique([]int{1, 1, 2})
	fmt.Println(unique)
}

func TestCase2(t *testing.T) {
	unique := permuteUnique([]int{1, 2, 3})
	fmt.Println(unique)
}
