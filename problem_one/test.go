package main

import "fmt"

func main() {
	var targetArray []int
	targetArray = append(targetArray,1,2,3,4,5,6,7,8)
	twoSum(targetArray,4)
	fmt.Println("targetArray", twoSum(targetArray,5))
	fmt.Println("targetArray", targetArray[twoSum(targetArray,5)[0]],targetArray[twoSum(targetArray,4)[1]])
}

func twoSum(nums []int, target int) []int {
	result :=[]int{}//创建一个返回用的result
	m:=make(map[int]int)//创建一个key和value都是int类型的map,python可以理解为字典,键值对
	for index,value:=range nums{//循环
		key,exist :=m[target-value]//返回索引和是否存在
		fmt.Println(key,exist,value)
		if exist{//如果为true,也就是在map中存在的话
			result = append(result,key)//把在map中的存在的的位置放到数组中
			result = append(result,index)//在把当前遍历的index放到数组中
		}
		m[value] = index//把当前的遍历的值作为键,把索引作为值(全部放到字典中)
	}
	return result

}