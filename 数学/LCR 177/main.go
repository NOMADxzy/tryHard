package main

import "fmt"

func sockCollocation(sockets []int) []int {
	xorResult := 0
	for i := 0; i < len(sockets); i++ {
		xorResult ^= sockets[i]
	}
	flag := (xorResult) & (-xorResult)
	ans1, ans2 := 0, 0
	for i := 0; i < len(sockets); i++ {
		if sockets[i]&flag == 0 {
			ans1 ^= sockets[i]
		} else {
			ans2 ^= sockets[i]
		}
	}
	return []int{ans1, ans2}
}

func main() {
	fmt.Println(sockCollocation([]int{4, 5, 2, 4, 6, 6}))
}
