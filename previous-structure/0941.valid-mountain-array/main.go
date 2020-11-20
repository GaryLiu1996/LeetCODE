package main

import "fmt"

func main() {
	arr := []int{9, 8, 7, 6, 5}
	fmt.Println(validMountainArray(arr))
}

//validMountainArray: way1
func validMountainArray(A []int) bool {
	i, j := 0, len(A)-2
	if len(A) < 3 {
		return false
	}
	for i < len(A)-1 {
		if A[i] < A[i+1] {
			i++
		} else if A[j] > A[j+1] && j > 0 {
			j--
		} else if i == j+1 {
			return true
		} else {
			break
		}
	}
	return false
}

////validMountainArray: way2:双指针
//func validMountainArray(A []int) bool {
//	n := len(A)
//	// if n < 3 {
//	// 	return false
//	// }
//	i, j := 0, n-1
//	for i+1 < n && A[i] < A[i+1] {
//		i++
//	}
//	for j-1 >= 0 && A[j-1] > A[j] {
//		j--
//	}
//	if i != 0 && i == j && j != n-1 {
//		return true
//	}
//	return false
//}
//作者：xiao_ben_zhu
//链接：https://leetcode-cn.com/problems/valid-mountain-array/solution/tu-jie-941-you-xiao-de-shan-mai-shu-zu-shuang-zhi-/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

////validMountainArray: way3:线性扫描
//func validMountainArray(a []int) bool {
//	i, n := 0, len(a)
//
//	// 递增扫描
//	for ; i+1 < n && a[i] < a[i+1]; i++ {
//	}
//
//	// 最高点不能是数组的第一个位置或最后一个位置
//	if i == 0 || i == n-1 {
//		return false
//	}
//
//	// 递减扫描
//	for ; i+1 < n && a[i] > a[i+1]; i++ {
//	}
//
//	return i == n-1
//}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/valid-mountain-array/solution/you-xiao-de-shan-mai-shu-zu-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
