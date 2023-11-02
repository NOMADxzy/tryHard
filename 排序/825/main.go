package main

import (
	"fmt"
	"sort"
)

func numFriendRequests(ages []int) int {
	m := len(ages)
	sort.Slice(ages, func(i, j int) bool {
		return ages[i] < ages[j]
	})
	fages := make([]int, m)
	for i := 0; i < m; i++ {
		fages[i] = ages[i]/2 + 7
	}

	ans := 0
	for yidx := 0; yidx < m; yidx++ {
		y := ages[yidx]
		var R int
		left, right := 0, m-1
		if fages[left] >= y {
			R = -1
		} else if fages[right] < y {
			R = right
		} else {
			for left < right {
				mid := (left + right) / 2
				if fages[mid] < y {
					left = mid + 1
				} else {
					right = mid
				}
			}
			R = right - 1
		}
		var L int
		left, right = 0, m-1
		if ages[left] >= y {
			L = left
		} else if ages[right] < y {
			L = -1
		} else {
			for left < right {
				mid := (left + right) / 2
				if ages[mid] < y {
					left = mid + 1
				} else {
					right = mid
				}
			}
			L = right
		}
		if L == -1 || R == -1 || L > R {
			continue
		}
		if y > 100 {
			left, right = L, R
			if ages[left] >= 100 {
				L = left
			} else if ages[right] < 100 {
				continue
			} else {
				for left < right {
					mid := (left + right) / 2
					if ages[mid] < 100 {
						left = mid + 1
					} else {
						right = mid
					}
				}
				L = right
			}
		}
		ans += R - L + 1
		if yidx >= L && yidx <= R {
			ans--
		}
	}
	return ans
}

func main() {
	fmt.Println(numFriendRequests([]int{16, 16}))
}
