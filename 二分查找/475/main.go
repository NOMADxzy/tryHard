package main

import (
	"fmt"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Slice(houses, func(i, j int) bool {
		return houses[i] < houses[j]
	})
	sort.Slice(heaters, func(i, j int) bool {
		return heaters[i] < heaters[j]
	})
	ans := 0
	leftLastHeaterIdx := 0
	for _, house := range houses {
		if heaters[leftLastHeaterIdx] > house {
			if d := heaters[leftLastHeaterIdx] - house; d > ans {
				ans = d
			}
		} else if heaters[len(heaters)-1] <= house {
			if d := house - heaters[len(heaters)-1]; d > ans {
				ans = d
			}
			leftLastHeaterIdx = len(heaters) - 1
		} else {
			left, right := leftLastHeaterIdx, len(heaters)-1
			for left < right {
				mid := (left + right) / 2
				if heaters[mid] <= house {
					left = mid + 1
				} else {
					right = mid
				}
			}
			left--
			ans = max(ans, min(house-heaters[left], heaters[right]-house))
			leftLastHeaterIdx = left
		}
	}
	return ans
}

func main() {
	fmt.Println(findRadius([]int{1, 2, 3, 4}, []int{1, 4}))
}
