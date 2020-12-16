package _290_word_pattern

import (
	"strings"
)

//My solution
//执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
//内存消耗： 1.9 MB, 在所有 Go 提交中击败了 64.96% 的用户
func wordPattern(pattern string, s string) bool {
	mapTemp := map[string]byte{}
	pMapTemp := map[int32]int{}
	words := strings.Split(s, " ")
	//fmt.Println("-", words,words[1])
	if len(pattern) != len(words) {
		return false
	}
	for _, v := range pattern {
		pMapTemp[v] = 1
	}

	for i := 0; i < len(pattern); i++ {
		if mapTemp[words[i]] == byte(0) && pMapTemp[int32(pattern[i])] == 1 {
			mapTemp[words[i]] = pattern[i]
			pMapTemp[int32(pattern[i])] = 0
		} else {
			if mapTemp[words[i]] != pattern[i] {
				return false
			}
		}
	}
	return true
}
