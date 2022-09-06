package leetcode

func maxProfit(prices []int) int {
	sell := make([][]int, 0)
	buy := make([][]int, 0)
	sell = append(sell, []int{0, 0})
	buy = append(buy, []int{-prices[0], -prices[0]})
	for index, price := range prices {
		if index == 0 {
			continue
		}
		buy1 := max(buy[index-1][0], -price)
		sell1 := max(sell[index-1][0], buy[index-1][0]+price)
		buy2 := max(buy[index-1][1], sell[index-1][0]-price)
		sell2 := max(sell[index-1][1], buy[index-1][1]+price)
		buy = append(buy, []int{buy1, buy2})
		sell = append(sell, []int{sell1, sell2})
	}
	buyMaxProfit := maxNumber(buy[len(buy)-1])
	sellMaxProfit := maxNumber(sell[len(sell)-1])
	return max(buyMaxProfit, sellMaxProfit)
}

func maxNumber(nums1 []int) (ans int) {
	ans = 0
	for _, val := range nums1 {
		if val > ans {
			ans = val
		}
	}
	return
}
