package main

import "fmt"

func countTriplets(arr []int) int {
	hMap := map[int][]int{0: {-1}, arr[0]: {0}}
	last := arr[0]
	ans := 0
	for i := 1; i < len(arr); i++ {
		t := last ^ arr[i]
		if len(hMap[t]) > 0 {
			for _, p := range hMap[t] {
				ans += i - p - 1
			}
		}
		hMap[t] = append(hMap[t], i)
		last = t
	}
	return ans
}

func main() {
	fmt.Println(countTriplets([]int{7, 11, 12, 9, 5, 2, 7, 17, 22}))
}
