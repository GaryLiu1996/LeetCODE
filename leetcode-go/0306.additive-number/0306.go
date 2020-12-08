package _306_additive_number

import (
	"math"
	"strconv"
	"strings"
)

//halforst 霜神代码
//执行用时：
//0 ms
//, 在所有 Go 提交中击败了
//100.00%
//的用户
//内存消耗：
//1.9 MB
//, 在所有 Go 提交中击败了
//97.26%
//的用户
func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	for firstEnd := 0; firstEnd < len(num)/2; firstEnd++ {
		if num[0] == '0' && firstEnd > 0 {
			break
		}
		first, _ := strconv.Atoi(num[0 : firstEnd+1])                                                                                    //string->int
		for secondEnd := firstEnd + 1; int(math.Max(float64(firstEnd), float64(secondEnd-firstEnd))) < len(num)-secondEnd; secondEnd++ { //前两个数字的位数应该小于等于剩余数字位数<=,但是我觉得不需要有等号
			if num[firstEnd+1] == '0' && secondEnd-firstEnd > 1 {
				break
			}
			second, _ := strconv.Atoi(num[firstEnd+1 : secondEnd+1])
			if recursiveCheck(num, first, second, secondEnd+1) {
				return true
			}
		}
	}
	return false
}

//Propagate for rest of the string
func recursiveCheck(num string, x1 int, x2 int, left int) bool {
	if left == len(num) {
		return true
	}
	if strings.HasPrefix(num[left:], strconv.Itoa(x1+x2)) { //int->string
		return recursiveCheck(num, x2, x1+x2, left+len(strconv.Itoa(x1+x2)))
	}
	return false
}
