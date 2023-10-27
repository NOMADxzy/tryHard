package main

import "fmt"

func numSplits(s string) int {
	m := len(s)
	dpLeft, dpRight := make([]int, m), make([]int, m)
	dpLeft[0] = 1
	dpRight[m-1] = 1

	mapLeft := map[byte]bool{s[0]: true}
	mapRight := map[byte]bool{s[m-1]: true}

	for i := 1; i < m; i++ {
		i_ := m - 1 - i
		dpLeft[i] = dpLeft[i-1]
		dpRight[i_] = dpRight[i_+1]
		if !mapLeft[s[i]] {
			dpLeft[i]++
			mapLeft[s[i]] = true
		}
		if !mapRight[s[i_]] {
			dpRight[i_]++
			mapRight[s[i_]] = true
		}
	}
	ans := 0
	for i := 0; i < m-1 && dpLeft[i] <= dpRight[i+1]; i++ {
		if dpLeft[i] == dpRight[i+1] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numSplits("aacaba"))
}
