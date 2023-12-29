package main

import "fmt"

func minDominoRotations(tops []int, bottoms []int) int {
	ans := len(tops) + 1
	topArr, botArr, sameArr := [6]int{}, [6]int{}, [6]int{}
	n := len(tops)
	for i := 0; i < n; i++ {
		topArr[tops[i]-1]++
		botArr[bottoms[i]-1]++
		if tops[i] == bottoms[i] {
			sameArr[tops[i]-1]++
		}
	}
	for i := 0; i < 6; i++ {
		if botArr[i]+topArr[i]-sameArr[i] >= n {
			ans = min(ans, n-max(botArr[i], topArr[i]))
		}
	}
	if ans == n+1 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}))
}
