package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	a :=&ListNode{Val: 8, Next:nil}
	//b :=&ListNode{Val: 7, Next:nil}
	b :=&ListNode{Val: 7, Next:&ListNode{Val:4, Next:nil}}
	fmt.Println("output:",addTwoNumbers(a,b))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumber(l1, l2, 0)
}

func addTwoNumber(l1 *ListNode, l2 *ListNode, add int) *ListNode {
	if l1 == nil && l2 ==nil && add == 0 {
		return nil
	}
	if l1 != nil{
		add += l1.Val
		l1 = l1.Next
	}
	if l2 != nil{
		add += l2.Val
		l2 = l2.Next
	}
	node := ListNode{
		Val: add%10,
		Next: addTwoNumber(l1, l2,add/10),
	}
	return &node
}