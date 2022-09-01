package leetcode

func finalPrices(prices []int) []int {
	length := len(prices)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if prices[j] <= prices[i] {
				prices[i] -= prices[j]
				break
			}
		}
	}
	return prices
}
