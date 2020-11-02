package main

import "fmt"

func main() {
	arr1, arr2 := []int{1, 2, 2, 1}, []int{2, 2}
	intersection(arr1, arr2)
}

func intersection(nums1 []int, nums2 []int) []int {
	num1Map, num2Map := map[int]int{}, map[int]int{}
	for i := 0; i < len(nums1); i++ {
		if num1Map[nums1[i]] != 1 {
			num1Map[nums1[i]] = 1
		}
	}
	for i := 0; i < len(nums2); i++ {
		if num2Map[nums2[i]] != 1 {
			num2Map[nums2[i]] = 1
		}
	}
	var array []int
	for i := range num1Map {
		if num2Map[i] == 1 {
			array = append(array, i)
		}
	}

	fmt.Println(array, "array")
	return array
}

//another way

//func intersection(nums1 []int, nums2 []int) []int {
//	hash := make(map[int]bool)
//	res := make([]int, 0)
//
//	for _, v := range nums1 {
//		hash[v] = true
//	}
//
//	for _, v := range nums2 {
//		if hash[v] == true {
//			res = append(res, v)
//			hash[v] = false
//		}
//	}
//	return res
//}

//作者：qi-ang-qi-ang
//链接：https://leetcode-cn.com/problems/intersection-of-two-arrays/solution/jian-dan-ming-bai-shuang-bai-jie-fa-by-qi-ang-q-15/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
