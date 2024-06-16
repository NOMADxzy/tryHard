package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegrees := make([]int, numCourses)
	nexts := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		inDegrees[prerequisite[0]]++
		nexts[prerequisite[1]] = append(nexts[prerequisite[1]], prerequisite[0])
	}
	var starts []int
	for i, degree := range inDegrees {
		if degree == 0 {
			starts = append(starts, i)
		}
	}
	var ans []int
	var dfs func(p int)
	dfs = func(p int) {
		ans = append(ans, p)
		inDegrees[p]-- // 防止重复访问
		for _, np := range nexts[p] {
			inDegrees[np]--
			if inDegrees[np] == 0 {
				dfs(np)
			}
		}
	}
	for i := 0; i < len(starts); i++ {
		dfs(starts[i])
	}
	if len(ans) < numCourses {
		return []int{}
	}
	return ans
}

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
}
