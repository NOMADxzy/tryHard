package main

import "fmt"

func groupThePeople(groupSizes []int) [][]int {
	n := len(groupSizes)

	roots := map[int][]int{}
	for i := 0; i < n; i++ {
		if r, ok := roots[groupSizes[i]]; !ok {
			roots[groupSizes[i]] = []int{i}
		} else {
			roots[groupSizes[i]] = append(r, i)
		}
	}
	var ans [][]int
	for k, vals := range roots {
		for i := 0; i < len(vals); {
			arr := make([]int, k)
			for j := 0; j < k; j++ {
				idx := vals[i+j]
				arr[j] = idx
			}
			i += k
			ans = append(ans, arr)
		}
	}
	return ans
}

func main() {
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
}
