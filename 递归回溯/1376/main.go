package main

import "fmt"

func informNext(id int, informTime []int, childs [][]int) int {
	if len(childs[id]) == 0 {
		return 0
	}
	var maxTime int

	for _, c := range childs[id] {
		if t := informNext(c, informTime, childs); t > maxTime {
			maxTime = t
		}
	}

	return maxTime + informTime[id]
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	childs := make([][]int, n)
	for i, v := range manager {
		if v >= 0 {
			childs[v] = append(childs[v], i)
		}
	}
	return informNext(headID, informTime, childs)
}

func main() {
	fmt.Println(numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}))
}
