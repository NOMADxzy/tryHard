package main

import "fmt"

func balancedString(s string) int {
	kMap := make([]int, 4)
	idx := map[byte]int{'Q': 0, 'W': 1, 'E': 2, 'R': 3}
	for i := 0; i < len(s); i++ {
		kMap[idx[s[i]]]++
	}
	already := true
	for _, cnt := range kMap {
		if cnt != len(s)/4 {
			already = false
		}
	}
	if already {
		return 0
	}
	//kMap[s[0]] --
	ans := len(s)
	l, r := 0, 0
	for ; r < len(s); r++ {
		kMap[idx[s[r]]]--
		maxVal := 0
		sum := 0
		for _, cnt := range kMap {
			if cnt > maxVal {
				maxVal = cnt
			}
			sum += cnt
		}
		need := maxVal*4 - sum
		for need <= r-l+1 {
			ans = min(ans, r-l+1)
			kMap[idx[s[l]]]++
			l++
			maxVal := 0
			sum := 0
			for _, cnt := range kMap {
				if cnt > maxVal {
					maxVal = cnt
				}
				sum += cnt
			}
			need = maxVal*4 - sum
		}
	}
	return ans
}

func main() {
	fmt.Println(balancedString("QQQQ"))
}
