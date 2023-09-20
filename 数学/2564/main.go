package main

import "fmt"

func substringXorQueries(s string, queries [][]int) [][]int {
	m := len(s)
	var results [][]int
	valMap := map[int][]int{}

	for i := 0; i < m; i++ {
		if s[i] == '0' {
			if valMap[0] == nil {
				valMap[0] = []int{i, i}
			}
			continue
		}
		val := 0
		for r := i; r < i+30 && r < len(s); r++ {
			val = val*2 + int(s[r]-'0')
			if valMap[val] == nil {
				valMap[val] = []int{i, r}
			}
		}
	}

	for i := 0; i < len(queries); i++ {
		query := queries[i]
		target := query[1] ^ query[0]
		if valMap[target] == nil {
			results = append(results, []int{-1, -1})
		} else {
			results = append(results, []int{valMap[target][0], valMap[target][1]})
		}
	}
	return results
}

func main() {
	fmt.Println(substringXorQueries("111010110", [][]int{{4, 2}, {3, 3}, {1, 2}}))
}
