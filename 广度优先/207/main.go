package main

import (
	"fmt"
	"sort"
)

func downCheck(pres, idxs []int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	smallest := i
	if left < len(idxs) && pres[idxs[left]] < pres[idxs[smallest]] {
		smallest = left
	}
	if right < len(idxs) && pres[idxs[left]] < pres[idxs[smallest]] {
		smallest = right
	}
	if smallest != i {
		idxs[smallest], idxs[i] = idxs[i], idxs[smallest]
		downCheck(pres, idxs, smallest)
	}
}

func upCheck(pres, idxs []int, i int) {
	f := (i - 1) / 2
	if f >= 0 && pres[idxs[i]] < pres[idxs[f]] {
		idxs[i], idxs[f] = idxs[f], idxs[i]
		upCheck(pres, idxs, f)
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	nexts := make([][]int, numCourses)
	presCnt := make([]int, numCourses)
	idxs := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		idxs[i] = i
	}
	for _, p := range prerequisites {
		nexts[p[1]] = append(nexts[p[1]], p[0])
		presCnt[p[0]] += 1
	}
	sort.Slice(idxs, func(i, j int) bool {
		return presCnt[idxs[i]] < presCnt[idxs[j]]
	})
	for len(idxs) > 0 && presCnt[idxs[0]] == 0 {
		pos := idxs[0]
		idxs[0] = idxs[len(idxs)-1]
		idxs = idxs[:len(idxs)-1]
		downCheck(presCnt, idxs, 0)
		for _, p := range nexts[pos] {
			presCnt[p]--
		}
		sort.Slice(idxs, func(i, j int) bool {
			return presCnt[idxs[i]] < presCnt[idxs[j]]
		})
	}
	return len(idxs) == 0

}

func main() {
	fmt.Println(canFinish(2, [][]int{{0, 1}}))
}
