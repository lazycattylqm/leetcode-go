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

func maxProfit2(prices []int) int {
	statusA := 0 // 表示持有股票
	statusB := 0 //表示不持有 但是在冷冻期
	statusC := 0 //表示不持有 也不在冷冻期
	for index, price := range prices {
		if index == 0 {
			statusA = -price
			statusB = 0
			statusC = 0
			continue
		}
		tempA := max(statusA, statusC-price)

		tempC := max(statusC, statusB)

		tempB := statusA + price

		statusA = tempA
		statusB = tempB
		statusC = tempC
	}
	return max(statusA, max(statusB, statusC))
}
