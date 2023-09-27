package main

import "fmt"

func totalFruit(fruits []int) int {
	l, r := 0, 0

	cnt1, cnt2 := []int{fruits[0], 1}, []int{-1, 0}
	for r = 1; r < len(fruits); r++ {
		if fruits[r] != cnt1[0] {
			cnt2[0] = fruits[r]
			cnt2[1] = 1
			r++
			break
		} else {
			cnt1[1]++
		}
	}
	if r == len(fruits) {
		return len(fruits)
	}
	maxLen := r - l

	for ; r < len(fruits); r++ {
		if fruits[r] == cnt1[0] {
			cnt1[1]++
		} else if fruits[r] == cnt2[0] {
			cnt2[1]++
		} else {
			for l < len(fruits) && cnt1[1] > 0 && cnt2[1] > 0 {
				if fruits[l] == cnt1[0] {
					cnt1[1]--
				} else if fruits[l] == cnt2[0] {
					cnt2[1]--
				}
				l++
			}
			if l == len(fruits) {
				break
			}
			if cnt1[1] == 0 {
				cnt1[0] = fruits[r]
				cnt1[1] = 1
			} else if cnt2[1] == 0 {
				cnt2[0] = fruits[r]
				cnt2[1] = 1
			}
		}
		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}
	return maxLen
}

func main() {
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))
}
