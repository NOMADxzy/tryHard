package main

import "fmt"

func subarrayBitwiseORs(arr []int) int {
	m := len(arr)

	nMap := map[int]bool{arr[0]: true}
	ans := 1

	//dp := make([]int, m)
	//dp[0] = arr[0]

	lastRes := make([]int, 32)
	lastRes[0] = arr[0]
	//lastLen := 1

	for i := 1; i < m; i++ {
		tmpMap := map[int]bool{}
		newRes := make([]int, 32)
		idx := 0
		for j := 0; j < 32; j++ {
			val := lastRes[j] | arr[i]
			if !nMap[val] {
				ans++
				nMap[val] = true
			}
			if !tmpMap[val] {
				tmpMap[val] = true
				newRes[idx] = val
				idx++
			}
		}
		lastRes = newRes
	}
	return ans
}

func main() {
	fmt.Println(subarrayBitwiseORs([]int{1, 2, 4}))
}
