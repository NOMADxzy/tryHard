package main

import "fmt"

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	arr := make([]int, 4)
	reward := 0
	best := 0
	bestIdx := -1
	cptStopReward := func() int {
		i := 0
		for i < 4 && arr[i] == 0 {
			i++
		}
		return reward - (4-i)*runningCost
	}
	cur, wait := 0, 0
	for i := 0; i < len(customers) || wait > 0; i++ {
		if i < len(customers) {
			cur = customers[i]
		} else {
			cur = 0
		}
		v := min(4, cur+wait)
		wait = cur + wait - v
		copy(arr[1:], arr)
		reward += v*boardingCost - runningCost
		if r := cptStopReward(); r > best {
			best = r
			bestIdx = i + 1
		}
	}
	return bestIdx
}

func main() {
	fmt.Println(minOperationsMaxProfit([]int{8, 3}, 5, 6))
}
