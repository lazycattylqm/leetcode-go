package leetcode

import "testing"

func TestMaxProfitCase1(t *testing.T) {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	profit := maxProfit(prices)
	if profit != 6 {
		t.Errorf("max profit should be 6, but got %d", profit)
	}
}

func TestMaxProfitCase2(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	profit := maxProfit(prices)
	if profit != 4 {
		t.Errorf("max profit should be 4, but got %d", profit)
	}
}

func TestMaxProfitCase3(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	profit := maxProfit(prices)
	if profit != 0 {
		t.Errorf("max profit should be 0, but got %d", profit)
	}
}

func TestMaxProfitCase4(t *testing.T) {
	prices := []int{1}
	profit := maxProfit(prices)
	if profit != 0 {
		t.Errorf("max profit should be 0, but got %d", profit)
	}
}

func TestMaxProfit2Case1(t *testing.T) {
	prices := []int{1, 2, 3, 0, 2}
	profit := maxProfit2(prices)
	if profit != 3 {
		t.Errorf("max profit should be 3, but got %d", profit)
	}
}
