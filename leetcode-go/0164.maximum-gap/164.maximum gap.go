package _164_maximum_gap

import (
	"fmt"
	"sort"
)

//My solution
//执行用时：
//8 ms
//, 在所有 Go 提交中击败了
//41.67%的用户
//内存消耗：
//3 MB
//, 在所有 Go 提交中击败了
//100.00%
//的用户
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums) //sort.Ints的时间/空间复杂度？
	fmt.Println("--", nums)
	diff := 0
	for i := 0; i+1 < len(nums); i++ {
		if nums[i+1]-nums[i] > diff {
			diff = nums[i+1] - nums[i]
		}
	}
	fmt.Println("--", diff)
	return diff
}
