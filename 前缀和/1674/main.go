package main

import "fmt"

func minMoves(nums []int, limit int) int {
	m := len(nums)
	defArr := make([]int, 2*limit+2)
	total := m / 2
	zeroMap := make([]int, 2*limit+2)

	for i := 0; i < m/2; i++ {
		a, b := nums[i], nums[m-1-i]
		if a > b {
			a, b = b, a
		}
		defArr[a+1] += 1
		defArr[b+limit+1] -= 1
		zeroMap[a+b]++
	}
	ans := m

	for i := 1; i < len(defArr); i++ {
		defArr[i] += defArr[i-1]
		steps := defArr[i]
		steps += 2 * (total - steps)
		steps -= zeroMap[i]
		if steps < ans {
			ans = steps
		}
	}
	return ans
}

func main() {
	fmt.Println(minMoves([]int{1, 2, 4, 3}, 4))
}
