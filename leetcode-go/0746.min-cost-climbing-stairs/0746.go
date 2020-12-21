package _746_min_cost_climbing_stairs

import "math"

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	var minCost, minCost0, minCost1 float64
	minCost, minCost0, minCost1 = 0, 0, math.Min(float64(cost[0]), float64(cost[1]))
	for i := 2; i < n; i++ {
		minCost = math.Min(minCost1+float64(cost[i]), minCost0+float64(cost[i-1]))
		minCost0 = minCost1
		minCost1 = minCost
	}
	return int(minCost)
}
