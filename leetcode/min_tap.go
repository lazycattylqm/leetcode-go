package leetcode

func minTaps(n int, ranges []int) int {
	dp := make([]int, n+1)
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < n+1; i++ {
		left := max(0, i-ranges[i])
		right := min(n, i+ranges[i])
		dp[left] = max(dp[left], right)
	}
	var count = 0
	var maxRight = 0
	var pre = 0
	for i := 0; i < n; i++ {
		maxRight = max(maxRight, dp[i])
		if i == pre {
			count++
			pre = maxRight
		}
	}
	if pre >= n {
		return count
	}
	return -1
}
