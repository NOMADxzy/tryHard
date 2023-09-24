package main

import "fmt"

func lowBit(x int) int {
	return x & (-x)
}

func update(tree []int, x int, val int) {
	for x < len(tree) {
		tree[x] += val
		x += lowBit(x)
	}
}

func query(tree []int, x int) int {
	sum := 0
	for x > 0 {
		sum += tree[x]
		x -= lowBit(x)
	}
	return sum
}

func numTeams(rating []int) int {

	maxN := 1
	for i := 0; i < len(rating); i++ {
		if rating[i] > maxN {
			maxN = rating[i]
		}
	}
	maxN++
	tree1 := make([]int, maxN)
	tree2 := make([]int, maxN)

	for i := 0; i < len(rating); i++ {
		update(tree2, rating[i], 1)
	}

	cnt := 0
	for i := 0; i < len(rating); i++ {
		x := rating[i]
		update(tree2, x, -1)
		leftLower := query(tree1, x)
		leftLarger := i - leftLower
		rightLower := query(tree2, x)
		rightLarger := len(rating) - i - 1 - rightLower

		cnt += leftLower*rightLarger + leftLarger*rightLower
		update(tree1, x, 1)
	}
	return cnt
}

func main() {
	fmt.Println(numTeams([]int{2, 5, 3, 4, 1}))
}

//朴素做法
func numTeams1(rating []int) int {

	cnt := 0
	for i := 0; i < len(rating); i++ {
		x := rating[i]
		leftLower, leftLarger, rightLower, rightLarger := 0, 0, 0, 0

		for j := 0; j < i; j++ {
			if rating[j] < x {
				leftLower++
			} else {
				leftLarger++
			}
		}
		for j := i + 1; j < len(rating); j++ {
			if rating[j] < x {
				rightLower++
			} else {
				rightLarger++
			}
		}
		cnt += leftLower*rightLarger + leftLarger*rightLower
	}
	return cnt
}
