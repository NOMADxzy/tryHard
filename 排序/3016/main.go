package main

import (
	"fmt"
	"sort"
)

func minimumPushes(word string) int {
	cnts := map[byte]int{}
	for i := 0; i < len(word); i++ {
		cnts[word[i]]++
	}
	var arr []int
	for _, n := range cnts {
		arr = append(arr, n)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	//buts := []int{8, 8, 8, 2}
	ans := 0
	for i := 0; i < len(arr); i++ {
		if i < 8 {
			ans += 1 * arr[i]
		} else if i < 16 {
			ans += 2 * arr[i]
		} else if i < 24 {
			ans += 3 * arr[i]
		} else {
			ans += 4 * arr[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumPushes("xyzxyzxyzxyz"))
}
