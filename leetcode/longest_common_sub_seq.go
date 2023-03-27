package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	i := len(text1)
	j := len(text2)
	dp := make([][]int, i+1)
	for k := range dp {
		dp[k] = make([]int, j+1)
	}
	for rowIndex, rowCharacter := range text1 {
		for columnIndex, columnCharacter := range text2 {
			if rowCharacter == columnCharacter {
				dp[rowIndex+1][columnIndex+1] = dp[rowIndex][columnIndex] + 1
			} else {
				dp[rowIndex+1][columnIndex+1] = max(dp[rowIndex][columnIndex+1], dp[rowIndex+1][columnIndex])
			}
		}
	}

	return dp[i][j]

}

func LongestCommonSubsequence(text1, text2 string) int {
	return longestCommonSubsequence(text1, text2)
}
