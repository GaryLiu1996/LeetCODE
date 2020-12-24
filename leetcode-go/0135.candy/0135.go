package _135_candy

import "math"

func candy(ratings []int) int {
	numLeft := make([]int, len(ratings))
	numRight := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		numLeft[i] = 1
		numRight[i] = 1
	}
	for i := 0; i+1 < len(ratings); i++ {
		if ratings[i+1] > ratings[i] {
			numLeft[i+1] = numLeft[i] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			numRight[i-1] = numRight[i] + 1
		}
	}
	result := 0
	for i := 0; i < len(ratings); i++ {
		result += int(math.Max(float64(numRight[i]), float64(numLeft[i])))
	}
	return result
}
