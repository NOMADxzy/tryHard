package main

import "fmt"

func maxIceCream(costs []int, coins int) int {
	min, max := costs[0], costs[0]
	m := len(costs)
	for i := 0; i < m; i++ {
		if costs[i] < min {
			min = costs[i]
		} else if costs[i] > max {
			max = costs[i]
		}
	}

	n := max - min + 1
	bkts := make([]int, n)

	for i := 0; i < m; i++ {
		bkts[costs[i]-min]++
	}

	ans := 0
	idx := 0
	for {
		for idx < n && bkts[idx] == 0 {
			idx++
		}
		if idx == n || coins < min+idx {
			break
		}
		bkts[idx]--
		coins -= min + idx
		ans++
	}
	return ans
}

func main() {
	fmt.Println(maxIceCream([]int{1, 3, 2, 4, 1}, 7))
}
