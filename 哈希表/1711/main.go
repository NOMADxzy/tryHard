package main

import "fmt"

func countPairs(deliciousness []int) int {
	cnts := map[int]int{}
	maxVal := deliciousness[0]
	for i := 0; i < len(deliciousness); i++ {
		cnts[deliciousness[i]]++
		maxVal = max(maxVal, deliciousness[i])
	}
	var sums []int
	v := 1
	for v <= 2*maxVal {
		sums = append(sums, v)
		v *= 2
	}
	MOD := 1000000007
	ans := 0
	for i := 0; i < len(deliciousness)-1; i++ {
		cnts[deliciousness[i]]--
		for _, sum := range sums {
			ans = (ans + cnts[sum-deliciousness[i]]) % MOD
		}
	}
	return ans
}

func main() {
	fmt.Println(countPairs([]int{1, 1, 1, 3, 3, 3, 7}))
}
