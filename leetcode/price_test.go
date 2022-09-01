package leetcode

import (
	"fmt"
	"testing"
)

func TestCase_1(t *testing.T) {
	prices := finalPrices([]int{8, 4, 6, 2, 3})
	fmt.Println(prices)
}

func TestCase_2(t *testing.T) {
	prices := finalPrices([]int{1, 2, 3, 4, 5})
	fmt.Println(prices)
}

func TestCase_3(t *testing.T) {
	prices := finalPrices([]int{10, 1, 1, 6})
	fmt.Println(prices)
}
