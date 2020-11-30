package _767_reorganize_string

import (
	"fmt"
)

//TODO
func reorganizeString(S string) string {
	if len(S) < 2 {
		return S
	}

	sMap := make(map[rune]int, 0)
	for _, value := range S {
		sMap[value]++
	}
	tag := 0
	for _, value := range sMap {
		if tag < value {
			tag = value
		}
	}
	fmt.Println("tag", tag)

	if len(S)%2 == 0 && tag > len(S)/2 {
		return ""
	}
	if len(S)%2 == 1 && tag > len(S)/2+1 {
		return ""
	}
	arr := make([]rune, len(S))
	t := 0
	for key, _ := range sMap {
		if (len(S)%2 == 1 && tag == len(S)/2) || (len(S)%2 == 1 && tag == len(S)/2) {
			for sMap[key] > 0 {
				if t < len(S) {
					arr[t] = key
					t += 2
					sMap[key]--
				}
			}
		}
		if (len(S)%2 == 1 && tag < len(S)/2) || (len(S)%2 == 1 && tag < len(S)/2) {
			for sMap[key] > 0 {
				if t < len(S) {
					arr[t] = key
					t += 2
					sMap[key]--
				} else {
					t = 1
					arr[t] = key
					t += 2
					sMap[key]--
				}
			}
		}
	}
	fmt.Println("arr", string(arr))
	return string(arr)
}
