package main

import (
	"fmt"
)

func main() {
	fmt.Println("---", masterMind("YBBY", "GYYB"))
}
func masterMind(solution string, guess string) []int {
	hit, pseudoHit := 0, 0
	var arrayKeyTemp []int
	for i := 0; i < len(solution); i++ {
		if string(solution[i]) == string(guess[i]) {
			hit++
			arrayKeyTemp = append(arrayKeyTemp, i)
		}
	}

	tag := 0
	fmt.Println(arrayKeyTemp, "arrayKeyTemp")
	solutionArray := map[string]int{}
	fmt.Println("init value ----:", solutionArray["KK"])
	for i := 0; i < len(solution); i++ {
		for j := 0; j < len(arrayKeyTemp); j++ {
			if i != arrayKeyTemp[j] && (tag < len(arrayKeyTemp)+1) {
				tag++
			}
		}
		if tag == len(arrayKeyTemp) {
			if solutionArray[string(solution[i])] > 0 {
				solutionArray[string(solution[i])] += 1
			} else {
				solutionArray[string(solution[i])] = 1
			}
		}
		tag = 0
	}
	fmt.Println(solutionArray, "solutionArray")

	notHitArray := notHit(guess, arrayKeyTemp)
	for i := 0; i < len(notHitArray); i++ {
		if solutionArray[notHitArray[i]] > 0 {
			pseudoHit += 1
			solutionArray[notHitArray[i]] -= 1
		}
	}
	var array []int
	array = append(array, hit)
	array = append(array, pseudoHit)
	return array
}

func notHit(guess string, arrayHit []int) []string {
	array := []string{}
	fmt.Println(arrayHit, "arrayHit")
	tag := 0
	for i := 0; i < len(guess); i++ {
		for j := 0; j < len(arrayHit); j++ {
			if i != arrayHit[j] && (tag < len(arrayHit)+1) {
				tag++
			}
		}
		if tag == len(arrayHit) {
			array = append(array, string(guess[i]))
		}
		tag = 0
	}
	fmt.Println(array, "array00")
	return array
}

//another man's Way
//masterMind https://leetcode-cn.com/problems/master-mind-lcci/solution/liang-ci-bian-li-by-triste_24/

//func masterMind(solution string, guess string) []int {
//	solutionMap := make(map[byte]int, 0)
//
//	ret := []int{0, 0}
//	for i := 0; i < len(solution); i++ {
//		if solution[i] == guess[i] {
//			//猜中
//			ret[0]++
//		} else {
//			//非猜中，用来统计伪猜中
//			solutionMap[solution[i]]++
//		}
//	}
//	for i := 0; i < len(guess); i++ {
//		if solution[i] != guess[i] {
//			//统计伪猜中
//			if solutionMap[guess[i]] > 0 {
//				//消耗掉一个
//				solutionMap[guess[i]]--
//				ret[1]++
//			}
//		}
//	}
//	return ret
//}
