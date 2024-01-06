package main

import "fmt"

func countSubTrees(n int, edges [][]int, labels string) []int {
	nexts := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		nexts[a] = append(nexts[a], b)
		nexts[b] = append(nexts[b], a)
	}
	ans := make([]int, n)
	var dfs func(pos, pre int) [26]int
	dfs = func(pos, pre int) [26]int {
		res := [26]int{}
		res[labels[pos]-'a']++
		for _, np := range nexts[pos] {
			if np != pre {
				nres := dfs(np, pos)
				for i := 0; i < 26; i++ {
					res[i] += nres[i]
				}
			}
		}
		ans[pos] = res[labels[pos]-'a']
		return res
	}
	dfs(0, -1)
	return ans
}

func main() {
	edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	fmt.Println(countSubTrees(7, edges, "abaedcd"))
}
