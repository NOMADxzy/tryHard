package main

import "fmt"

func averageWaitingTime(customers [][]int) float64 {
	n := len(customers)
	ans := float64(0)
	ans += float64(customers[0][1])
	customers[0][1] = customers[0][0] + customers[0][1]
	for i := 1; i < n; i++ {
		customers[i][1] = max(customers[i][0], customers[i-1][1]) + customers[i][1]
		ans += float64(customers[i][1] - customers[i][0])
	}
	return ans / float64(n)
}

func main() {
	fmt.Println(averageWaitingTime([][]int{{1, 2}, {2, 5}, {4, 3}}))
}
