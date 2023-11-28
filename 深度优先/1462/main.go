package main

import "fmt"

func findNext(isPre [][]int, pres [][]int, pos, target int) bool {
	if isPre[target][pos] != 0 {
		return isPre[target][pos] > 0
	}
	for _, p := range pres[pos] {
		if findNext(isPre, pres, p, target) {
			isPre[target][pos] = 1
			return true
		}
	}
	isPre[target][pos] = -1
	return false
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	pres := make([][]int, numCourses)
	isPre := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		isPre[i] = make([]int, numCourses)
		isPre[i][i] = 1
	}

	for _, pr := range prerequisites {
		pres[pr[1]] = append(pres[pr[1]], pr[0])
		isPre[pr[0]][pr[1]] = 1
	}

	ans := make([]bool, len(queries))
	for i, query := range queries {
		if findNext(isPre, pres, query[1], query[0]) {
			ans[i] = true
		}
	}
	return ans
}

func main() {
	fmt.Println(checkIfPrerequisite(3, [][]int{{1, 2}, {1, 0}, {2, 0}}, [][]int{{1, 0}, {1, 2}}))
}
