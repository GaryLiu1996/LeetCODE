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
