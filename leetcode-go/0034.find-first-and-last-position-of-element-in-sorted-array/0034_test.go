package _034_find_first_and_last_position_of_element_in_sorted_array

import (
	"fmt"
	"testing"
)

func Test_0034(t *testing.T) {
	target := 11
	arr := []int{5, 7, 7, 8, 8, 10}
	fmt.Println("--", searchRange(arr, target))
}
