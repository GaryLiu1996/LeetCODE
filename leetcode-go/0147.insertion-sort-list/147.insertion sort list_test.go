package _147_insertion_sort_list

import "testing"

func Test_MoveZeroes(t *testing.T) {
	arr := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	insertionSortList(arr)
	t.Log("123")
}
