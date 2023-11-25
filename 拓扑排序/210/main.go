package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	nexts := make([][]int, numCourses)
	var queue []int
	inDegrees := make([]int, numCourses)
	for _, pr := range prerequisites {
		nexts[pr[1]] = append(nexts[pr[1]], pr[0])
		inDegrees[pr[0]]++
	}
	for i := 0; i < numCourses; i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	var ans []int
	for len(queue) > 0 && inDegrees[queue[0]] == 0 {
		p := queue[0]
		ans = append(ans, p)
		queue = queue[1:]
		for _, np := range nexts[p] {
			inDegrees[np]--
			if inDegrees[np] == 0 {
				queue = append(queue, np)
			}
		}
	}
	if len(ans) < numCourses {
		return []int{}
	}
	return ans
}

func main() {
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
