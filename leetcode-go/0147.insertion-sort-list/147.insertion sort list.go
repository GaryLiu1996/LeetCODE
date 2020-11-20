package _147_insertion_sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

//模仿标准答案
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummyHead := &ListNode{Next: head}
	last, curr := head, head.Next
	for curr != nil {
		if last.Val <= curr.Val {
			last = last.Next
		} else {
			prev := dummyHead //每次都要从虚拟头指针开始遍历
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			last.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = last.Next //重新赋值curr的位置
	}
	return dummyHead.Next
}

////标准答案
//func insertionSortList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	dummyHead := &ListNode{Next: head}
//	lastSorted, curr := head, head.Next
//	for curr != nil {
//		if lastSorted.Val <= curr.Val {
//			lastSorted = lastSorted.Next
//		} else {
//			prev := dummyHead
//			for prev.Next.Val <= curr.Val {
//				prev = prev.Next
//			}
//			lastSorted.Next = curr.Next
//			curr.Next = prev.Next
//			prev.Next = curr
//		}
//		curr = lastSorted.Next//重新赋值curr的位置
//	}
//	return dummyHead.Next
//}

////My Wrong solution
//func insertionSortList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//	prev := &ListNode{Next:head}
//	previousNext := prev.Next
//	lastNode, currentNode := head, head.Next
//	for currentNode != nil{
//		if lastNode.Val>currentNode.Val{
//			for previousNext != nil{
//				if previousNext.Val >currentNode.Val{
//					lastNode.Next = currentNode.Next
//					currentNode.Next = prev.Next
//					prev.Next = currentNode
//					break
//				}else{
//					previousNext = previousNext.Next
//				}
//			}
//		}else{
//			lastNode = lastNode.Next
//		}
//		currentNode = lastNode.Next
//	}
//	arr := []int{}
//	for prev.Next !=nil{
//		arr =append(arr,prev.Next.Val)
//		prev.Next = prev.Next.Next
//	}
//	fmt.Println("arr",arr)
//	return head
//}
