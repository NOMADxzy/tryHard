package main

import "fmt"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var i, j int
	m1, m2 := len(firstList), len(secondList)
	var ans [][]int
	for i < m1 && j < m2 {
		a, b := firstList[i], secondList[j]
		if a[0] > b[1] {
			j++
		} else if b[1] <= a[1] {
			ans = append(ans, []int{max(a[0], b[0]), b[1]})
			j++
		} else if b[0] > a[1] {
			i++
		} else {
			ans = append(ans, []int{max(a[0], b[0]), a[1]})
			i++
		}
	}
	return ans
}

func main() {
	first := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	second := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	fmt.Println(intervalIntersection(first, second))
}
