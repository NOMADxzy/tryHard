package main

import (
	"fmt"
	"sort"
)

// 动态规划 超时
func maximumTastiness1(price []int, k int) int {
	sort.Slice(price, func(i, j int) bool {
		return price[i] < price[j]
	})
	m := len(price)
	maxTastiness := make([]int, m)
	for i := 1; i < m; i++ {
		maxTastiness[i] = price[i] - price[0]
	}
	for t := 3; t <= k; t++ {
		nextMaxTastiness := make([]int, m)
		for i := t - 1; i < m; i++ {
			left, right := t-2, i-1
			if price[i]-price[right] >= maxTastiness[right] {
				left = right
			} else {
				for left < right {
					mid := (left + right) / 2
					if price[i]-price[mid] >= maxTastiness[mid] {
						left = mid + 1
					} else {
						right = mid
					}
				}
				left--
			}
			nextMaxTastiness[i] = max(min(price[i]-price[left], maxTastiness[left]),
				min(price[i]-price[right], maxTastiness[right]))
		}
		maxTastiness = nextMaxTastiness
	}
	return maxTastiness[m-1]
}

func maximumTastiness(price []int, k int) int {
	sort.Slice(price, func(i, j int) bool {
		return price[i] < price[j]
	})

	valid := func(val int) bool {
		cnt := 1
		pre := 0
		for i := 1; i < len(price) && cnt < k; i++ {
			if price[i]-price[pre] >= val {
				pre = i
				cnt++
			}
		}
		return cnt >= k
	}

	left, right := 0, price[len(price)-1]-price[0]
	if valid(right) {
		return right
	}
	for left < right {
		mid := (left + right) / 2
		if valid(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
}

func main() {
	fmt.Println(maximumTastiness([]int{13, 5, 1, 8, 21, 2}, 3))
}
