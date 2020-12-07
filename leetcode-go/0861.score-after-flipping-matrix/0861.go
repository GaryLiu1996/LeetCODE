package _861_score_after_flipping_matrix

import "fmt"

//My solution
func matrixScore(A [][]int) int {
	for _, v := range A {
		if v[0] == 0 {
			for i := 0; i < len(v); i++ {
				switch v[i] {
				case 0:
					v[i] = 1
					break
				case 1:
					v[i] = 0
					break
				}
			}
		}
	}
	for i := 0; i < len(A[0]); i++ {
		A0, A1 := 0, 0
		for j := 0; j < len(A); j++ {
			if A[j][i] == 1 {
				A1++
			}
			if A[j][i] == 0 {
				A0++
			}
		}
		if A0 > A1 {
			for j := 0; j < len(A); j++ {
				switch A[j][i] {
				case 0:
					A[j][i] = 1
					break
				case 1:
					A[j][i] = 0
					break
				}
			}
		}
	}
	sum := 0
	for _, v := range A {
		for k, value := range v {
			sum += value << (len(v) - 1 - k)
		}
	}
	fmt.Println("---", sum)
	return sum
}

//官方题解
//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//2.2 MB
//, 在所有 Go 提交中击败了
//86.67%
//的用户
//func matrixScore(a [][]int) int {
//    m, n := len(a), len(a[0])
//    ans := 1 << (n - 1) * m
//    for j := 1; j < n; j++ {
//        ones := 0
//        for _, row := range a {
//            if row[j] == row[0] {
//                ones++
//            }
//        }
//        if ones < m-ones {
//            ones = m - ones
//        }
//        ans += 1 << (n - 1 - j) * ones
//    }
//    return ans
//}
//
//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/score-after-flipping-matrix/solution/fan-zhuan-ju-zhen-hou-de-de-fen-by-leetc-cxma/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
