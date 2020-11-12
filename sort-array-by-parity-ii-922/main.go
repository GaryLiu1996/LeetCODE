package main

import "fmt"

func main() {
	arr := []int{4, 2, 5, 7}
	fmt.Println("arr", sortArrayByParityII(arr))
}

//way1 : 双指针
func sortArrayByParityII(A []int) []int {
	i, j := 0, 1
	arr := make([]int, len(A))
	fmt.Println("arr1", arr)
	for n := 0; n < len(A); n++ {
		if A[n]%2 == 0 {
			arr[i] = A[n]
			i += 2
		} else {
			arr[j] = A[n]
			j += 2
		}
	}
	fmt.Println("arr2", arr)
	return arr
}

//way2 : 两次遍历
//func sortArrayByParityII(a []int) []int {
//	ans := make([]int, len(a))
//	i := 0
//	for _, v := range a {
//		if v%2 == 0 {
//			ans[i] = v
//			i += 2
//		}
//	}
//	i = 1
//	for _, v := range a {
//		if v%2 == 1 {
//			ans[i] = v
//			i += 2
//		}
//	}
//	return ans
//}
