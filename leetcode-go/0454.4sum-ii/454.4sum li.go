package _454_4sum_ii

import "fmt"

//My solution
//执行用时：
//68 ms
//, 在所有 Go 提交中击败了
//64.90%
//的用户
//内存消耗：
//22.7 MB
//, 在所有 Go 提交中击败了
//34.25%
//的用户
func fourSumCount(A []int, B []int, C []int, D []int) int {
	countAB := make(map[int]int, 0)
	count := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			countAB[A[i]+B[j]]++
		}
	}
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			count += countAB[-C[i]-D[j]]
		}
	}
	fmt.Println("123", count)
	return count
}

////官方解法：用for range遍历的，感觉for range更适合本道题
//执行用时：
//72 ms
//, 在所有 Go 提交中击败了
//50.99%
//的用户
//内存消耗：
//22.7 MB
//, 在所有 Go 提交中击败了
//38.36%
//的用户
//func fourSumCount(a, b, c, d []int) (ans int) {
//	countAB := map[int]int{}
//	for _, v := range a {
//		for _, w := range b {
//			countAB[v+w]++
//		}
//	}
//	for _, v := range c {
//		for _, w := range d {
//			ans += countAB[-v-w]
//		}
//	}
//	return
//}
