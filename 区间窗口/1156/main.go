package main

import "fmt"

func maxRepOpt1(text string) int {
	maxSeqLen := make([]int, 26)
	cnts := make([]int, 26)
	for i := 0; i < len(text); i++ {
		cnts[int(text[i]-'a')]++
	}

	for i := 0; i < 26; i++ {
		if cnts[i] == 1 {
			continue
		}
		cur := byte('a' + i)
		left, right := 0, 0
		limit := false

		for ; right < len(text); right++ {
			if text[right] != cur {
				if !limit {
					limit = true
				} else {
					for text[left] == cur {
						left++
					}
					left++
				}
			}
			if right-left+1 > maxSeqLen[int(cur-'a')] {
				maxSeqLen[int(cur-'a')] = right - left + 1
			}
		}
	}

	ans := 1
	for i := 0; i < 26; i++ {
		v := maxSeqLen[i]
		if cnts[i] < v {
			v--
		}
		if v > ans {
			ans = v
		}
	}
	return ans
}

func main() {
	fmt.Println(maxRepOpt1("ababa"))
}
