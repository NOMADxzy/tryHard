package main

import "fmt"

func numMatchingSubseq(s string, words []string) int {
	cArr := make([][]int, 26)
	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		cArr[idx] = append(cArr[idx], i)
	}
	ans := 0

	for _, word := range words {
		if len(word) > len(s) {
			continue
		}
		P := 0
		i := 0
		cArrLeftIdx := make([]int, 26)
		for ; i < len(word); i++ {
			idx := int(word[i] - 'a')
			if cArrLeftIdx[idx] >= len(cArr[idx]) {
				break
			}
			valids := cArr[idx][cArrLeftIdx[idx]:]
			if valids[0] >= P {
				P = valids[0] + 1
				cArrLeftIdx[idx]++
			} else if valids[len(valids)-1] < P {
				break
			} else {
				left, right := 0, len(valids)-1
				for left < right {
					mid := (left + right) / 2
					if valids[mid] < P {
						left = mid + 1
					} else {
						right = mid
					}
				}
				P = valids[right] + 1
				cArrLeftIdx[idx] += right + 1
			}
		}
		if i == len(word) {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numMatchingSubseq("dsahjpjauf", []string{"a", "d", "b", "c"}))

}
