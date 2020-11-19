package _283_move_zeroes

func moveZeroes(nums []int) {
	j := len(nums) - 1
	for i := 0; i < j; i++ {
		if nums[i] == 0 {
			k := nums[i]
			nums[i] = nums[i+1]
			nums[i+1] = k
		}
		if i+1 == j {
			i = 0
			j--
		}
	}
}
