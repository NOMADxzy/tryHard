package main

import "fmt"

func isSelfCrossing(distance []int) bool {
	for i := 3; i < len(distance); i++ {
		if i-3 >= 0 && distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}
		if i-4 >= 0 && distance[i-1] == distance[i-3] && distance[i]+distance[i-4] >= distance[i-2] {
			return true
		}
		if i-5 >= 0 && distance[i-1] <= distance[i-3] && distance[i]+distance[i-4] >= distance[i-2] && distance[i-2] > distance[i-4] && distance[i-1]+distance[i-5] >= distance[i-3] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isSelfCrossing([]int{2, 1, 1, 2}))
}
