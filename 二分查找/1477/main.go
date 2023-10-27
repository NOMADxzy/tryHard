package main

import "fmt"

func minSumOfLengths(arr []int, target int) int {
	m := len(arr)
	dpLeft, dpRight := make([]int, m), make([]int, m)

	sums := make([]int, m)
	sums[0] = arr[0]
	for i := 1; i < m; i++ {
		sums[i] = sums[i-1] + arr[i]
	}

	for rb := 0; rb < m; rb++ {
		if sums[rb] < target || arr[rb] > target {
			//dpRight[rb] = -1
		} else if arr[rb] == target {
			dpRight[rb] = 1
			dpLeft[rb] = 1
		} else {
			left, right := 0, rb
			for left < right {
				mid := (left + right + 1) / 2
				val := sums[rb] - sums[mid-1]
				if val >= target {
					left = mid
				} else {
					right = mid - 1
				}
			}
			if sums[rb] == target || left > 0 && sums[rb]-sums[left-1] == target {
				length := rb - left + 1
				dpLeft[rb] = length
				dpRight[left] = length
			}
		}
	}
	for i := 1; i < m; i++ {
		i_ := m - 1 - i
		if dpLeft[i-1] > 0 && (dpLeft[i] == 0 || dpLeft[i-1] < dpLeft[i]) {
			dpLeft[i] = dpLeft[i-1]
		}
		if dpRight[i_+1] > 0 && (dpRight[i_] == 0 || dpRight[i_+1] < dpRight[i_]) {
			dpRight[i_] = dpRight[i_+1]
		}
	}
	ans := m + 1
	for i := 0; i < m-1; i++ {
		if dpLeft[i] == 0 {
			continue
		} else if dpRight[i+1] == 0 {
			break
		} else if val := dpLeft[i] + dpRight[i+1]; val < ans {
			ans = val
		}
	}
	if ans > m {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minSumOfLengths([]int{3, 1, 1, 1, 5, 1, 2, 1}, 3))
}
