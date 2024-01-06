package main

import "fmt"

func minDays(bloomDay []int, m int, k int) int {
	left, right := 1, 1
	for i := 0; i < len(bloomDay); i++ {
		if bloomDay[i] > right {
			right = bloomDay[i]
		}
	}

	getFlower := func(d int) int {
		cnt := 0
		for i := 0; i < len(bloomDay); {
			for i < len(bloomDay) && bloomDay[i] > d {
				i++
			}
			j := i
			for j < len(bloomDay) && j < i+k {
				if bloomDay[j] > d {
					break
				}
				j++
			}
			if j == i+k {
				cnt++
			}
			i = j
		}
		return cnt
	}
	if getFlower(right) < m {
		return -1
	}
	for left < right {
		mid := (left + right) / 2
		if getFlower(mid) < m {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

func main() {
	fmt.Println(minDays([]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2))
}
