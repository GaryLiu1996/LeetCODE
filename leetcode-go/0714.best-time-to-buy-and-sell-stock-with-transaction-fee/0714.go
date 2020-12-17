package _714_best_time_to_buy_and_sell_stock_with_transaction_fee

import "math"

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := [2]int{}
	dp[0] = 0
	dp[1] = -prices[0]
	for i := 0; i < n; i++ {
		tmp := dp[0]
		dp[0] = int(math.Max(float64(dp[0]), float64(dp[1]+prices[i]-fee)))
		dp[1] = int(math.Max(float64(dp[1]), float64(tmp-prices[i])))
	}
	return dp[0]
}
