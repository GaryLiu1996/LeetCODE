package _034_find_first_and_last_position_of_element_in_sorted_array

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

//官方解法 ：
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
//func searchRange(nums []int, target int) []int {
//	leftmost := sort.SearchInts(nums, target)
//	if leftmost == len(nums) || nums[leftmost] != target {
//		return []int{-1, -1}
//	}
//	rightmost := sort.SearchInts(nums, target+1) - 1
//	return []int{leftmost, rightmost}
//}

func searchRange(nums []int, target int) []int {
	return []int{searchFirstMatchElement(nums, target), searchLastMatchElement(nums, target)}
}

//霜神：二分查找
//执行用时：
//8 ms
//, 在所有 Go 提交中击败了
//90.40%
//的用户
//内存消耗：
//4 MB
//, 在所有 Go 提交中击败了
//10.19%
//的用户
//----求第一个与target相等的元素
func searchFirstMatchElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1 //*二分法只要取到中间的一个点即可，并不是非要正在中间，比如长度是偶数就没办法取到中间*
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

//----求最后一个与target相等的元素
func searchLastMatchElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}
