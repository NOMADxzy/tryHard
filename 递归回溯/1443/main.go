package main

import "fmt"

func cntNext(pos int, childs map[int][]int, hasApple []bool, mark []bool) (bool, int) {
	mark[pos] = true
	if childs[pos] == nil {
		return hasApple[pos], 0
	}
	curHasApple := hasApple[pos]
	totalTime := 0
	for _, c := range childs[pos] {
		if !mark[c] {
			ok, time := cntNext(c, childs, hasApple, mark)
			if ok {
				curHasApple = true
				totalTime += 2 + time
			}
		}
	}
	return curHasApple, totalTime
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	childs := map[int][]int{}
	for _, edge := range edges {
		childs[edge[0]] = append(childs[edge[0]], edge[1])
		childs[edge[1]] = append(childs[edge[1]], edge[0])
	}
	ok, ans := cntNext(0, childs, hasApple, make([]bool, n))
	if !ok {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, true, false, false, true, false}))
}
