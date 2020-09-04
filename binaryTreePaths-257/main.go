package main

import "strconv"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var paths []string

func main(){
}

func binaryTreePaths(root *TreeNode) []string {
	dfsTree(root,"")
	return paths
}
func dfsTree(root *TreeNode,path string){
	if root !=nil{
		path += strconv.Itoa(root.Val)
		if root.Left ==nil && root.Right ==nil{
			paths = append(paths,path)
		}else{
			path+= "->"
			dfsTree(root.Left,path)
			dfsTree(root.Right,path)
		}
	}
}
