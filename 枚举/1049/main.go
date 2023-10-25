package main

import "fmt"

func lastStoneWeightII(stones []int) int {
	m := len(stones)

	var allRes []int

	allRes = append(allRes, stones[0])
	allRes = append(allRes, -stones[0])

	for i := 1; i < m; i++ {
		set := map[int]bool{}
		var newRes []int
		for j := 0; j < len(allRes); j++ {
			if val := allRes[j] + stones[i]; !set[val] {
				newRes = append(newRes, val)
				set[val] = true
			}
			if val := allRes[j] - stones[i]; !set[val] {
				newRes = append(newRes, val)
				set[val] = true
			}
		}
		allRes = newRes
	}
	ans := 101
	for i := 0; i < len(allRes); i++ {
		if allRes[i] >= 0 && allRes[i] < ans {
			ans = allRes[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
}
