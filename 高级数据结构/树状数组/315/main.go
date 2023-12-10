package main

import "fmt"

func lowBit(x int) int {
	return x & (-x)
}

func query(tree []int, x int) int {
	sum := 0
	for x > 0 {
		sum += tree[x]
		x -= lowBit(x)
	}
	return sum
}

func update(tree []int, x, val int) {
	for x < len(tree) {
		tree[x] += val
		x += lowBit(x)
	}
}

func countSmaller(nums []int) []int {
	MAX_NUM := 20002
	delta := 10001
	m := len(nums)
	ans := make([]int, m)
	tree := make([]int, MAX_NUM)

	for i := m - 1; i >= 0; i-- {
		x := nums[i] + delta
		ans[i] = query(tree, x-1)
		update(tree, x, 1)
	}
	return ans
}

func main() {
	fmt.Println(countSmaller([]int{-1, -1}))
}
