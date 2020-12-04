package _659_split_array_into_consecutive_subsequences

import "fmt"

//参考
//执行用时：
//100 ms
//, 在所有 Go 提交中击败了
//65.63%
//的用户
//内存消耗：
//7 MB
//, 在所有 Go 提交中击败了
//65.63%
//的用户
func isPossible(nums []int) bool {
	numsMap := make(map[int]int, 0)
	tailMap := make(map[int]int, 0)
	for _, num := range nums {
		numsMap[num]++
	}
	for i := 0; i < len(nums); i++ {
		if numsMap[nums[i]] == 0 {
			continue
		}
		if numsMap[nums[i]] > 0 && tailMap[nums[i]-1] > 0 {
			numsMap[nums[i]]--
			tailMap[nums[i]-1]--
			tailMap[nums[i]]++
		} else if numsMap[nums[i]] > 0 && numsMap[nums[i]+1] > 0 && numsMap[nums[i]+2] > 0 {
			numsMap[nums[i]]--
			numsMap[nums[i]+1]--
			numsMap[nums[i]+2]--
			tailMap[nums[i]+2]++
		} else {
			fmt.Println("false")
			return false
		}
	}
	fmt.Println("true")
	return true
}
