package main

import "fmt"

func goodDaysToRobBank(security []int, time int) []int {
	m := len(security)
	if len(security) == 1 {
		if time == 0 {
			return []int{0}
		} else {
			return []int{}
		}
	}
	nextSmaller := make([]int, m)
	preSmaller := make([]int, m)
	for i := 1; i < m; i++ {
		if security[i] <= security[i-1] {
			preSmaller[i] = preSmaller[i-1] + 1
		}
	}
	for i := m - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			nextSmaller[i] = nextSmaller[i+1] + 1
		}
	}
	var ans []int
	for i := time; i <= m-1-time; i++ {
		if preSmaller[i] >= time && nextSmaller[i] >= time {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	fmt.Println(goodDaysToRobBank([]int{1, 1, 1, 1, 1}, 1))
}
