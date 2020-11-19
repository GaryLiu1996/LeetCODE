package _283_move_zeroes

import "fmt"

// 非 0 项覆盖到数组前头
//func moveZeroes(nums []int) {
//	i, j := 0, 0
//	for i < len(nums) {
//		if nums[i] != 0 {
//			nums[j] = nums[i]
//			j++
//		}
//		i++
//	}
//	for j<len(nums){
//		nums[j]=0
//		j++
//	}
//}

// 双指针交换
func moveZeroes(nums []int) {
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j] //GO语言直接对换变量值
			j++
		}
		i++
	}
	fmt.Println("nums", nums)
}
