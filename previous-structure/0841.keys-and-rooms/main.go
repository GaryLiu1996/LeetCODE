package main

import "fmt"

var (
	num int
	vis []bool
)

func main() {
	//rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	rooms := [][]int{{1}, {2}, {3}, {}}
	fmt.Println("result:", bfs(rooms))
	//fmt.Println("result:", canVisitAllRooms(rooms))
}

//func canVisitAllRooms(rooms [][]int) bool {
//	num = 0
//	n := len(rooms)
//	vis = make([]bool, n)
//	deepSearch(rooms, 0)
//	return n == num
//}
//
//func deepSearch(rooms [][]int, x int) {
//	vis[x] = true
//	num++
//	for _, value := range rooms[x] {
//		if vis[value] == false {
//			deepSearch(rooms, value)
//		}
//	}
//}

func bfs(rooms [][]int) bool {
	n := len(rooms)
	num := 0
	vis := make([]bool, n)
	vis[0] = true
	queen := []int{}
	queen = append(queen, 0)
	for i := 0; i < len(queen); i++ {
		x := queen[i]
		num++
		for _, j := range rooms[x] {
			if vis[j] != true {
				vis[j] = true
				queen = append(queen, j)
			}
		}
	}
	return num == n
}
