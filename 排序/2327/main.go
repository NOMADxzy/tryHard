package main

import "fmt"

func peopleAwareOfSecret(n int, delay int, forget int) int {
	MOD := 1000000007
	//var queue []int
	queue := make([]int, forget)
	queue[len(queue)-1] = 1

	for i := 1; i < n; i++ {
		queue = queue[1:]
		todayKnow := 0
		for j := 0; j < forget-delay; j++ {
			todayKnow = (todayKnow + queue[j]) % MOD
		}
		queue = append(queue, todayKnow)
	}
	ans := 0
	for _, e := range queue {
		ans = (ans + e) % MOD
	}
	return ans
}

func main() {
	fmt.Println(peopleAwareOfSecret(6, 2, 4))
}
