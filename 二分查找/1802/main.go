package main

import "fmt"

func maxValue(n int, index int, maxSum int) int {
	if n == 1 {
		return maxSum
	}
	left, right := 1, maxSum
	minSum := func(val int) int {
		leftCnt, rightCnt := max(0, min(index, val-1)), max(0, min(n-1-index, val-1))
		sumLeft := (val - 1 + val - leftCnt) * leftCnt / 2
		sumRight := (val - 1 + val - rightCnt) * rightCnt / 2
		ans := sumRight + sumLeft + val
		if leftCnt < index {
			ans += index - leftCnt
		}
		if rightCnt < n-1-index {
			ans += n - 1 - index - rightCnt
		}
		return ans
	}
	for left < right {
		mid := (left + right) / 2
		if minSum(mid) <= maxSum {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
}

func main() {
	fmt.Println(maxValue(4, 0, 4))
}
