package main

import "fmt"

func numPairsDivisibleBy60(time []int) int {
	vMap := map[int]int{}
	ans := 0
	for i := 0; i < len(time); i++ {
		modVal := time[i] % 60
		target := 60 - modVal
		if target == 60 {
			target = 0
		}
		ans += vMap[target]
		vMap[modVal]++
	}
	return ans
}

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{60, 60}))
}
