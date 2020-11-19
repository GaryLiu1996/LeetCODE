//Authors: Liusyuan
//Title: 167-two-sum-ii-input-array-is-sorted

package main

import "fmt"

func main() {
	array := []int{5, 25, 75}
	fmt.Println("===", twoSum(array, 100))
}

//twoSum: way1,时间复杂度O(m*n)
//func twoSum(numbers []int, target int) []int {
//	numbersCopy := numbers
//	array := make([]int, 0)
//	for i := 0; i < len(numbers); i++ {
//		for j := i + 1; j < len(numbersCopy); j++ {
//			if numbers[i]+numbers[j] == target {
//				array = append(array, i+1, j+1)
//			}
//		}
//	}
//	return array
//}

//twoSum: way2,时间复杂度O(n)
func twoSum(numbers []int, target int) []int {
	array := make([]int, 0)
	k := len(numbers) - 1
	for i := 0; i < len(numbers); {
		if numbers[i]+numbers[k]-target > 0 {
			k--
		} else if numbers[i]+numbers[k]-target < 0 {
			i++
		} else {
			array = append(array, i+1, k+1)
			break
		}
	}
	return array
}
