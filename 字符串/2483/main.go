package main

import "fmt"

func bestClosingTime(customers string) int {
	cost1, cost2 := 0, 0
	for i := 0; i < len(customers); i++ {
		if customers[i] == 'Y' {
			cost2++
		}
	}
	minCost := cost2
	ans := 0
	for i := 0; i < len(customers); i++ {
		if customers[i] == 'N' {
			cost1++
		} else {
			cost2--
		}
		if v := cost2 + cost1; v < minCost {
			minCost = v
			ans = i + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(bestClosingTime("YYNY"))
}
