package main

import "fmt"

func main() {
	arr := []int{4, 2, 5, 7}
	fmt.Println("arr", runningSum(arr))
}
func runningSum(nums []int) []int {
	arr := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			arr[i] = nums[i]
		} else {
			arr[i] = arr[i-1] + nums[i]
		}
	}
	return arr
}
