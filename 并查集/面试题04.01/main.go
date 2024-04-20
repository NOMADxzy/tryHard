package main

import "fmt"

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	// 设置arrive记录已经到达过的点, 下次遇到此点直接跳过
	arrive := make(map[int]bool)
	adjustGraph := make([][]int, n)
	// 如果不转成邻接表就要每次遍历整个graph列表, 判断是否为要遍历的下一个点, 从而导致超时
	// 转成邻接表后指定遍历下一个点, 减少复杂度.
	// 同时邻接表也可以用map套列表的形式保存, 这里用的列表套列表的形式保存的, 也就是用索引代替了map的键
	for _, link := range graph {
		adjustGraph[link[0]] = append(adjustGraph[link[0]], link[1])
	}

	res := false
	var dfs func(next int)
	dfs = func(next int) {
		if res {
			return
		}
		arrive[next] = true
		for _, v := range adjustGraph[next] {
			_, exists := arrive[v]
			if exists {
				continue
			}
			if v == target {
				res = true
				return
			}
			dfs(v)
		}
	}
	dfs(start)
	return res

}

func findWhetherExistsPath1(n int, graph [][]int, start int, target int) bool {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if parents[i] != i {
			parents[i] = find(parents[i])
		}
		return parents[i]
	}

	union := func(i1, i2 int) {
		parents[find(i1)] = parents[i2]
	}

	for _, edge := range graph {
		if edge[0] > edge[1] {
			edge[0], edge[1] = edge[1], edge[0]
		}
		union(edge[0], edge[1])
	}

	return find(start) == find(target)
}

func main() {
	graph := [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 2}, {1, 1}, {2, 1}}
	fmt.Println(findWhetherExistsPath(3, graph, 0, 2))
}
