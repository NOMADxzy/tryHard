package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	m, n := len(baseCosts), len(toppingCosts)

	var vals []int
	vals = append(vals, 0)
	valSet := map[int]bool{0: true}

	for i := 0; i < n; i++ {
		for _, val := range vals {
			if newVal := val + toppingCosts[i]; !valSet[newVal] {
				valSet[newVal] = true
				vals = append(vals, newVal)
			}
		}
	}
	sort.Slice(vals, func(i, j int) bool {
		return vals[i] < vals[j]
	})
	ans := 10000000000

	for i := 0; i < m; i++ {
		T := target - baseCosts[i]
		if T <= 0 {
			if -T < abs(ans) {
				ans = -T
			}
			continue
		} else {
			for _, firstVal := range vals {
				if firstVal >= T {
					if firstVal-T < abs(ans) {
						ans = firstVal - T
					}
					break
				} else {
					newTarget := T - firstVal
					left, right := 0, len(vals)-1
					if vals[right] <= newTarget {
						if newTarget-vals[right] <= abs(ans) {
							ans = vals[right] - newTarget
						}
						continue
					}
					for left < right {
						mid := (left + right) / 2
						if vals[mid] < newTarget {
							left = mid + 1
						} else {
							right = mid
						}
					}
					if newTarget-vals[left-1] <= abs(ans) {
						ans = vals[left-1] - newTarget
					}
					if vals[right]-newTarget < abs(ans) {
						ans = vals[right] - newTarget
					}
				}
			}
		}
	}
	return ans + target
}

func main() {
	fmt.Println(closestCost([]int{5, 77, 38, 61, 97}, []int{62, 7, 100, 30, 16, 84}, 73))
}
