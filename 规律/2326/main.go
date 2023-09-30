package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextDir(dir []int) {
	x, y := dir[0], dir[1]
	if x == 0 && y > 0 {
		dir[0] = y
		dir[1] = 0
	} else if x > 0 && y == 0 {
		dir[0] = 0
		dir[1] = -x
	} else if x == 0 && y < 0 {
		dir[0] = y
		dir[1] = 0
	} else {
		dir[0] = 0
		dir[1] = -x
	}
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	p := head
	dir := []int{0, 1}
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = -1
		}
	}

	x, y := 0, 0
	for p != nil {
		ans[x][y] = p.Val
		if x+dir[0] < 0 || x+dir[0] >= m || y+dir[1] < 0 || y+dir[1] >= n || ans[x+dir[0]][y+dir[1]] != -1 {
			nextDir(dir)
		}
		x, y = x+dir[0], y+dir[1]
		p = p.Next
	}
	return ans
}

func main() {
	head := &ListNode{3, &ListNode{0, &ListNode{2, &ListNode{6, &ListNode{8,
		&ListNode{1, &ListNode{7, &ListNode{9, &ListNode{4, &ListNode{2,
			&ListNode{5, &ListNode{5, &ListNode{0, nil}}}}}}}}}}}}}
	fmt.Println(spiralMatrix(3, 5, head))
}
