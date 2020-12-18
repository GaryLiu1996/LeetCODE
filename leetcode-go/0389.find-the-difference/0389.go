package _389_find_the_difference

import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
	sMap := map[int32]int{}
	for _, v := range s {
		sMap[v]++
	}
	var result int32
	for _, v := range t {
		if sMap[v] > 0 {
			sMap[v]--
		} else if sMap[v] == 0 {
			result = v
		}
	}
	fmt.Println("-", byte(result))
	return byte(result)
}
