package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var paths []string

func main() {
	treeEx := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println("paths:", binaryTreePaths(treeEx))
}

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	dfsTree(root, "")
	return paths
}
func dfsTree(root *TreeNode, path string) {
	if root != nil {
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, path)
		} else {
			path += "->"
			dfsTree(root.Left, path)
			dfsTree(root.Right, path)
		}
	}
}
