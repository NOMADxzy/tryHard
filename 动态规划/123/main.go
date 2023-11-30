package main

import "fmt"

func maxProfit(prices []int) int {
	m := len(prices)
	leftMin, rightMax := make([]int, m), make([]int, m)
	leftMin[0] = prices[0]
	rightMax[m-1] = prices[m-1]

	for i := 1; i < m; i++ {
		leftMin[i] = leftMin[i-1]
		if prices[i] < leftMin[i] {
			leftMin[i] = prices[i]
		}
		i_ := m - 1 - i
		rightMax[i_] = rightMax[i_+1]
		if prices[i_] > rightMax[i_] {
			rightMax[i_] = prices[i_]
		}
	}
	benifit := make([][]int, m)
	for i := 0; i < m; i++ {
		benifit[i] = make([]int, 2)
		benifit[i][0] = prices[i] - leftMin[i]
		benifit[i][1] = rightMax[i] - prices[i]
	}

	for i := 1; i < m; i++ {
		if benifit[i-1][0] > benifit[i][0] {
			benifit[i][0] = benifit[i-1][0]
		}
		i_ := m - 1 - i
		if benifit[i_+1][1] > benifit[i_][1] {
			benifit[i_][1] = benifit[i_+1][1]
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		if benifit[i][0]+benifit[i][1] > ans {
			ans = benifit[i][0] + benifit[i][1]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
