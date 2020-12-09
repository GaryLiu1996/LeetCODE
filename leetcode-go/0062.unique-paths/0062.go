package _062_unique_paths

//参考霜神代码
//DP思想+边界特殊值处理
//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//2 MB
//, 在所有 Go 提交中击败了
//36.48%
//的用户
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
