package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	maxTime := 0
	for i := 0; i < len(piles); i++ {
		if piles[i] > r {
			r = piles[i]
		}
		maxTime += piles[i]
	}
	if maxTime <= h {
		return l
	}
	if len(piles) > h {
		return -1
	}

	for l+1 < r {
		mid := (l + r) / 2
		time := 0
		for i := 0; i < len(piles); i++ {
			if piles[i]%mid == 0 {
				time += piles[i] / mid
			} else {
				time += piles[i]/mid + 1
			}

		}
		if time <= h {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func main() {
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
}
