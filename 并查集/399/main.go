package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := len(equations)

	keyStringMap := map[string]int{}
	cnt := 0
	for i := 0; i < m; i++ {
		if _, ok := keyStringMap[equations[i][0]]; !ok {
			keyStringMap[equations[i][0]] = cnt
			cnt++
		}
		if _, ok := keyStringMap[equations[i][1]]; !ok {
			keyStringMap[equations[i][1]] = cnt
			cnt++
		}
	}

	parents := make([]int, cnt)
	weights := make([]float64, cnt)
	for i := 0; i < len(parents); i++ {
		parents[i] = i
		weights[i] = 1
	}

	var find func(parents []int, index int) (int, float64)
	find = func(parents []int, index int) (int, float64) {
		if parents[index] != index {
			r, v := find(parents, parents[index])
			return r, weights[index] * v
		} else {
			return index, 1
		}
	}
	union := func(parents []int, index1, index2 int, w float64) {
		root1, _ := find(parents, index1)
		root2, _ := find(parents, index2)
		parents[root1] = root2
		weights[root1] = w
	}

	for i, equation := range equations {
		a, b, w := keyStringMap[equation[0]], keyStringMap[equation[1]], values[i]
		rootA, w1 := find(parents, a)
		rootB, w2 := find(parents, b)
		union(parents, rootA, rootB, (w2*w)/w1)
	}
	ans := make([]float64, len(queries))
	for i, query := range queries {
		ans[i] = -1.0
		idx1, ok1 := keyStringMap[query[0]]
		idx2, ok2 := keyStringMap[query[1]]
		if ok1 && ok2 {
			root1, w1 := find(parents, idx1)
			root2, w2 := find(parents, idx2)
			if root1 == root2 {
				ans[i] = w1 / w2
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"},
		{"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))
}
