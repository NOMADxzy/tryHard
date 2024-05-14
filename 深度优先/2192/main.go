package main

import "fmt"

// 暴力
func getAncestors1(n int, edges [][]int) [][]int {
	isAncestors := make([][]bool, n)
	for i := 0; i < n; i++ {
		isAncestors[i] = make([]bool, n)
	}
	nexts := make([][]int, n)
	degrees := make([]int, n)
	for _, edge := range edges {
		nexts[edge[0]] = append(nexts[edge[0]], edge[1])
		degrees[edge[1]]++
	}
	var dfs func(int, []bool)
	dfs = func(pos int, prePoints []bool) {
		for i := 0; i < len(prePoints); i++ {
			if prePoints[i] {
				isAncestors[pos][i] = true
			}
		}
		if len(nexts[pos]) == 0 {
			return
		}
		tmp := prePoints[pos]
		prePoints[pos] = true
		for _, np := range nexts[pos] {
			dfs(np, prePoints)
		}
		prePoints[pos] = tmp
	}
	for i := 0; i < n; i++ {
		if degrees[i] == 0 {
			dfs(i, make([]bool, n))
		}
	}
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isAncestors[i][j] {
				ans[i] = append(ans[i], j)
			}
		}
	}
	return ans
}

func getAncestors(n int, edges [][]int) [][]int {
	pres := make([][]int, n)
	inDegrees := make([]int, n)
	outDegrees := make([]int, n)
	for _, edge := range edges {
		pres[edge[1]] = append(pres[edge[1]], edge[0])
		inDegrees[edge[1]]++
		outDegrees[edge[0]]++
	}
	hist := map[int][]bool{}
	var dfs func(pos int) []bool
	ans := make([][]int, n)
	dfs = func(pos int) []bool {
		if inDegrees[pos] == 0 {
			return make([]bool, n)
		}
		if v, ok := hist[pos]; ok {
			return v
		}
		marks := make([]bool, n)
		for _, pp := range pres[pos] {
			preRes := dfs(pp)
			for i := 0; i < len(preRes); i++ {
				if preRes[i] {
					marks[i] = true
				}
			}
			marks[pp] = true
		}
		hist[pos] = marks
		for i := 0; i < n; i++ {
			if marks[i] {
				ans[pos] = append(ans[pos], i)
			}
		}
		return marks
	}
	for i := 0; i < n; i++ {
		if outDegrees[i] == 0 {
			dfs(i)
		}
	}
	return ans
}

func main() {
	edges := [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}
	//edges := [][]int{{0, 1}}
	fmt.Println(getAncestors(8, edges))

}
