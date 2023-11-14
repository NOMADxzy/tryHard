package main

import "fmt"

func dfs(mark []bool, rooms [][]int, pos int) {
	if mark[pos] {
		return
	}
	mark[pos] = true
	for _, key := range rooms[pos] {
		dfs(mark, rooms, key)
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	mark := make([]bool, n)
	dfs(mark, rooms, 0)
	for i := 0; i < n; i++ {
		if !mark[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
}
