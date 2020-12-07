package _118_pascals_triangle

import "fmt"

func generate(numRows int) [][]int {
	ans := make([][]int, numRows) // 先给外层下定义
	fmt.Println("ans", ans)       //[[] [] []]
	fmt.Println("ans", ans[0])    //[]
	for i := range ans {
		ans[i] = make([]int, i+1) //内层定义长度
		ans[i][0] = 1             //首是1不需要加入循环
		ans[i][i] = 1             //尾是1不需要加入循环
		for j := 1; j < i; j++ {  //保证 i>=2时（第三层），才需要计算。
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}
