package main

import "fmt"

func canChoose(groups [][]int, nums []int) bool {
	var i int

	for _, group := range groups {
		var j int
		for j = 0; j < len(group) && i < len(nums); i++ {
			if group[j] != nums[i] {
				i -= j
				j = 0
			} else {
				j++
			}
		}
		if j < len(group) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canChoose([][]int{{1, -1, -1}, {3, -2, 1}}, []int{1, -1, -1, 1, -1, 3, -2, 1}))
}
