package _034_find_first_and_last_position_of_element_in_sorted_array

import (
	"sort"
)

//Mysoluion 暴力法：
//执行用时： 8 ms, 在所有 Go 提交中击败了 90.40% 的用户
//内存消耗： 3.9 MB, 在所有 Go 提交中击败了 100.00% 的用户
//func searchRange(nums []int, target int) []int {
//	position := []int{}
//	tag:=0
//	for i:=0;i<len(nums);i++{
//		if nums[i] == target{
//			tag = 1
//			position = append(position,i)
//			break
//		}
//	}
//	if tag == 0{
//		return []int{-1,-1}
//	}
//	for i:=len(nums)-1;i>=0;i--{
//		if nums[i] == target{
//			position = append(position,i)
//			break
//		}
//	}
//	return position
//}

//官方解法 ： 二分查找
//执行用时：
//8 ms
//, 在所有 Go 提交中击败了
//90.40%
//的用户
//内存消耗：
//3.9 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户
func searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}
