package dynamic

func longestCommonSubsequence(text1, text2 string) int {
	len1, len2 := len(text1), len(text2)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i, char1 := range text1 {
		for j, char2 := range text2 {
			if char1 == char2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[len1][len2]
}

// in latest version, we can use function Max()
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
