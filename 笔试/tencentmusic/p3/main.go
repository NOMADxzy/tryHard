package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getMethod(head *ListNode, colors string) int {
	MOD := 1000000007
	m := len(colors)
	dp := make([]int, 2)
	dp[0] = 1
	cur := head
	for i := 0; i < m; i++ {
		if colors[i] == 'R' {
			if cur.Val%2 == 0 {
				continue
			} else {
				dp[0], dp[1] = dp[1], dp[0]
			}
		} else {
			newDp := make([]int, 2)
			if cur.Val%2 == 0 {
				newDp[0] = (dp[0] * 2) % MOD
				newDp[1] = (dp[1] * 2) % MOD
			} else {
				newDp[0] = (dp[0] + dp[1]) % MOD
				newDp[1] = (dp[1] + dp[0]) % MOD
			}
			dp = newDp
		}
		cur = cur.Next
	}
	return dp[0]
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Println(getMethod(head, "RWW"))
}
