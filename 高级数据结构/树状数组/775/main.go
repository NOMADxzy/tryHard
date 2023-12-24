package main

import "fmt"

func isIdealPermutation(nums []int) bool {
	n := len(nums)
	lowBit := func(x int) int {
		return x & -x
	}
	tree := make([]int, n+1)
	update := func(x, val int) {
		for x < len(tree) {
			tree[x] += val
			x += lowBit(x)
		}
	}
	query := func(x int) int {
		sum := 0
		for x > 0 {
			sum += tree[x]
			x -= lowBit(x)
		}
		return sum
	}
	rightCnt := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		c := query(nums[i])
		rightCnt[i] = c
		update(nums[i]+1, 1)
	}
	cnt, cntGlobal := 0, 0
	for i := 0; i < n-1; i++ {
		cnt += rightCnt[i]
		if nums[i] > nums[i+1] {
			cntGlobal++
		}
	}
	return cnt == cntGlobal
}

func main() {
	fmt.Println(isIdealPermutation([]int{0, 1}))
}
