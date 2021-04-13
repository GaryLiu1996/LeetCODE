package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("--", decodeString("3[a2[c]]"))
}

type Stack1 struct {
	Number   int
	Response string
}

func decodeString(s string) string {
	num := 0
	res := ""
	var stack []Stack1
	for _, v := range s {

		if string(v) <= "9" && string(v) >= "0" {
			value, _ := strconv.Atoi(string(v))
			num = value + num*10
		} else if string(v) == "[" {
			stack = append(stack, Stack1{num, res})
			num = 0
			res = ""
		} else if string(v) == "]" {
			tmp := ""
			for j := 0; j < stack[len(stack)-1].Number; j++ {
				tmp += res
			}
			res = stack[len(stack)-1].Response + tmp
			stack = stack[:len(stack)-1]
		} else {
			res += string(v)
		}
	}
	return res
}
