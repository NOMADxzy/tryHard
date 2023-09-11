package main

import (
	"fmt"
	"sort"
)

func find(candidates []int, target int, pos int) [][]int {
	if target == 0 {
		return [][]int{{}}
	} else {
		var combines [][]int
		for i := pos; i < len(candidates); i++ {
			if candidates[i] <= target {
				childs := find(candidates, target-candidates[i], i)
				for j := 0; j < len(childs); j++ {
					combines = append(combines, append(childs[j], candidates[i]))
				}
			}
		}
		return combines
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	return find(candidates, target, 0)
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 9))
}
