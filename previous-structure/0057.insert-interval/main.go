package main

import "fmt"

func main() {
	arr := [][]int{{1, 5}}
	arrIns := []int{5, 7}
	fmt.Println("insert(arr, arrIns)", insert(arr, arrIns))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval} //第一次没有考虑到第一个数组可能是空
	}
	fmt.Println("len(intervals)", len(intervals))
	intervalsArray, intervalsTemp, returnArr, temp, tempP := make([]int, 0), make([][]int, 0), make([]int, 2), make([]int, 2), make([]int, 2)
	for _, i := range intervals {
		intervalsArray = append(intervalsArray, i[0], i[1])
	}
	for i := 0; i < len(intervalsArray); i++ {
		if newInterval[0] == intervalsArray[i] {
			if i%2 == 0 {
				temp[0] = i
				returnArr[0] = intervalsArray[i]
			} else {
				temp[0] = i - 1
				returnArr[0] = intervalsArray[i-1]
			}
		}
		if i+1 < len(intervalsArray) { //i+1 超过边界
			if intervalsArray[i] < newInterval[0] && intervalsArray[i+1] > newInterval[0] {
				if i%2 == 0 {
					temp[0] = i
					returnArr[0] = intervalsArray[i]
				} else {
					temp[0] = i + 1
					returnArr[0] = newInterval[0]
				}
			}

		}

		if newInterval[1] == intervalsArray[i] {
			if i%2 == 0 {
				temp[1] = i + 1
				returnArr[1] = intervalsArray[i+1]
			} else {
				temp[1] = i
				returnArr[1] = intervalsArray[i]
			}
		}
		if i+1 < len(intervalsArray) {
			if intervalsArray[i] < newInterval[1] && intervalsArray[i+1] > newInterval[1] {
				if i%2 == 0 {
					temp[1] = i + 1
					returnArr[1] = intervalsArray[i+1]
				} else {
					temp[1] = i
					returnArr[1] = newInterval[1]
				}
			}
		}

		if newInterval[1] > intervalsArray[len(intervalsArray)-1] {
			temp[1] = len(intervalsArray) - 1
			returnArr[1] = newInterval[1] //第一版没考虑边界
		}
	}
	tempP[0] = temp[0] / 2
	tempP[1] = (temp[1] - 1) / 2
	tag := false
	for i := 0; i < len(intervals); i++ {
		if i < tempP[0] || i > tempP[1] {
			intervalsTemp = append(intervalsTemp, intervals[i])
		} else if tag == false {
			intervalsTemp = append(intervalsTemp, returnArr)
			tag = true
		}
	}

	fmt.Println("temp", temp)
	fmt.Println("tempP", tempP)
	fmt.Println("returnArr", returnArr)

	return intervalsTemp
}
