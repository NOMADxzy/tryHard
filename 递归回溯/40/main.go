package main

import (
	"fmt"
	"sort"
)

func dfs(nums []int, pos int, target int, kmap map[int]int) [][]int {
	if target == 0 {
		return [][]int{{}}
	} else if pos == len(nums) {
		return nil
	} else if nums[pos] > target {
		return nil
	} else if nums[pos] == target {
		return [][]int{{target}}
	}

	K := kmap[nums[pos]]
	var lists [][]int

	for k := 0; k <= K; k++ {
		r := dfs(nums, pos+K, target-k*nums[pos], kmap)
		for i := 0; i < len(r); i++ {
			if k == 0 {
				lists = append(lists, r[i])
			} else {
				for j := 0; j < k; j++ {
					r[i] = append(r[i], nums[pos])
				}
				lists = append(lists, r[i])
			}

		}
	}
	return lists
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] <= candidates[j]
	})
	kmap := map[int]int{}
	for i := 0; i < len(candidates); i++ {
		kmap[candidates[i]] += 1
	}
	return dfs(candidates, 0, target, kmap)
}

func main() {
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
