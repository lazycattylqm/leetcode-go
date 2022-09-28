package leetcode

func getKthMagicNumber(k int) int {
	dp := make([]int, k+1)
	dp[1] = 1
	p3, p5, p7 := 1, 1, 1
	for i := 2; i <= k; i++ {
		dp[i] = min(dp[p3]*3, min(dp[p5]*5, dp[p7]*7))
		if dp[i] == dp[p3]*3 {
			p3++
		}
		if dp[i] == dp[p5]*5 {
			p5++
		}
		if dp[i] == dp[p7]*7 {
			p7++
		}
	}
	return dp[k]
}
