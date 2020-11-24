package _222_count_complete_tree_nodes

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

////递归 ： 没有利用到完全二叉树性质
//func countNodes(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return 1 + countNodes(root.Left) + countNodes(root.Right)
//}

//利用到完全二叉树性质+递归 ：
//注意： 算N次方用左移右移1<<N 是2的N次方。
//执行结果：
//通过
//显示详情
//执行用时：
//16 ms
//, 在所有 Go 提交中击败了
//92.44%
//的用户
//内存消耗：
//7.1 MB
//, 在所有 Go 提交中击败了
//68.75%
//的用户
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	i, j := 0, 0
	lRoot, rRoot := root, root //注意:lRoot和rRoot要区别开，否则会出现错误，root被修改后再继续遍历一次，出现错误。
	for lRoot != nil {
		i++
		lRoot = lRoot.Left
	}
	for rRoot != nil {
		j++
		rRoot = rRoot.Right
	}
	if i == j {
		return 1<<i - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
