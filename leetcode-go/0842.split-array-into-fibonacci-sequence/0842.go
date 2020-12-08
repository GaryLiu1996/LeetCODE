package _842_split_array_into_fibonacci_sequence

import (
	"math"
	"strconv"
	"strings"
)

func splitIntoFibonacci(S string) []int {
	if len(S) < 3 {
		return []int{}
	}
	res, isComplete := []int{}, false
	for firstEnd := 0; firstEnd < len(S)/2; firstEnd++ {
		if isComplete == true {
			break
		}
		if S[0] == '0' && firstEnd > 0 {
			break
		}
		first, _ := strconv.Atoi(S[:firstEnd+1]) //这个区间是前闭后开的
		if first > (1<<31 - 1) {
			break
		}
		for secondEnd := firstEnd + 1; int(math.Max(float64(firstEnd), float64(secondEnd-firstEnd))) < len(S)-firstEnd; secondEnd++ { //<=
			if S[firstEnd+1] == '0' && secondEnd-firstEnd > 1 {
				break
			}
			second, _ := strconv.Atoi(S[firstEnd+1 : secondEnd+1])
			if second > (1<<31 - 1) {
				break
			}
			recursiveCheck(S, first, second, secondEnd+1, &res, &isComplete) //设置标记位和返回参数,指到地址，修改内容，直接得到res
			if isComplete == true {
				break
			}
		}

	}
	return res
}

func recursiveCheck(S string, x1 int, x2 int, left int, res *[]int, isComplete *bool) {
	if x1 > (1<<31-1) || x2 > (1<<31-1) {
		return
	}
	if left == len(S) && len(*res) > 0 {
		if !*isComplete {
			*isComplete = true
			*res = append(*res, x1)
			*res = append(*res, x2)
		}
		return
	}
	if strings.HasPrefix(S[left:], strconv.Itoa(x1+x2)) && !*isComplete { //HasPrefix:判断字符串是不是以xxx为开头
		*res = append(*res, x1)
		recursiveCheck(S, x2, x1+x2, left+len(strconv.Itoa(x1+x2)), res, isComplete)
		return
	}
	if len(*res) > 0 && !*isComplete {
		*res = (*res)[:len(*res)-1]
	}
	return
}
