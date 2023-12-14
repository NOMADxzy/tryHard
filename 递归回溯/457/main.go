package main

import "fmt"

func circularArrayLoop(nums []int) bool {
	m := len(nums)
	visited := make([]bool, m)

	var check func(pos, root int, first bool) bool
	check = func(pos int, root int, first bool) bool {
		if !first && pos == root {
			return true
		}
		if visited[pos] {
			return false
		}
		visited[pos] = true
		np := pos + nums[pos]
		for np >= m {
			np -= m
		}
		for np < 0 {
			np += m
		}
		if np != pos && (nums[pos]*nums[np] > 0) {
			return check(np, root, false)
		}
		return false
	}
	for i := 0; i < m; i++ {
		if check(i, i, true) {
			return true
		}
		clear(visited)
	}
	return false
}

func main() {
	fmt.Println(circularArrayLoop([]int{2, -1, 1, -2, -2}))
}
