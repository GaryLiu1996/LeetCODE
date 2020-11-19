package main

import (
	"fmt"
	"strings"
)

func main() {
	str := ""
	fmt.Println("	lengthOfLongestSubstring(str)", lengthOfLongestSubstring(str))
}

// //my way
//func lengthOfLongestSubstring(s string) int {
//	array := make([]string,len(s))
//	for i := 0; i <len(s); i++{
//		array[i]= s[i:i+1]
//	}
//
//	start,end:=0,1;
//
//	if len(s)== 0{
//		return 0
//	}
//	for ;end < len(s);{
//		if strings.Index(s[start:end],array[end]) == -1{
//			tag := 0
//			for i := start; i < end; i++{
//				if strings.Index(s[start:end],array[i])!=-1{
//					tag+=strings.Count(s[start:end],array[i])
//				}
//			}
//			if tag==len(s[start:end]){
//				end++
//			}else{
//				start++
//				end++
//			}
//		}else{
//			temp := strings.Index(s[start:end],array[end])+1
//			start += temp
//			end+= temp
//		}
//	}
//	return end-start
//}

//example way
func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ { //字符串start部分可跳跃，字符串END部分不可跳跃
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}
