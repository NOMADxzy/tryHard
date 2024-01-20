package main

import (
	"fmt"
	"strconv"
)

func stoneGameIX(stones []int) bool {
	hist := map[string]int{}
	cnt0, cnt1, cnt2 := 0, 0, 0
	for i := 0; i < len(stones); i++ {
		if stones[i]%3 == 0 {
			cnt0++
		} else if stones[i]%3 == 1 {
			cnt1++
		} else {
			cnt2++
		}
	}
	cnt0 %= 2
	var Loss func(acc int, cnt0, cnt1, cnt2 int) bool
	Loss = func(acc int, cnt0, cnt1, cnt2 int) bool {
		key := strconv.Itoa(cnt0) + "_" + strconv.Itoa(cnt1) + "_" + strconv.Itoa(cnt2)
		if hist[key] != 0 {
			return hist[key] > 0
		}
		if cnt0 == 0 && cnt1 == 0 && cnt2 == 0 {
			return len(stones)%2 == 0
		}
		if cnt0 > 0 && acc+0 != 3 {
			if Loss((acc+0)%3, cnt0-1, cnt1, cnt2) {
				hist[key] = -1
				return false
			}
		}
		if cnt1 > 0 && acc+1 != 3 {
			if Loss((acc+1)%3, cnt0, cnt1-1, cnt2) {
				hist[key] = -1
				return false
			}
		}
		if cnt2 > 0 && acc+2 != 3 {
			if Loss((acc+2)%3, cnt0, cnt1, cnt2-1) {
				hist[key] = -1
				return false
			}
		}
		hist[key] = 1
		return true
	}
	isLoss := Loss(0, cnt0, cnt1, cnt2)
	return !isLoss
}

func main() {
	fmt.Println(stoneGameIX([]int{2, 1, 2, 1, 0}))
}

//0 1 1 2
