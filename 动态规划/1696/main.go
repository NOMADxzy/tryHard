package main

import "fmt"

func upCheck(arr []int, dp []int, i int) {
	f := (i - 1) / 2
	if f >= 0 && dp[arr[f]] < dp[arr[i]] {
		arr[f], arr[i] = arr[i], arr[f]
		upCheck(arr, dp, f)
	}
}

func downCheck(arr []int, dp []int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < len(arr) && dp[arr[left]] > dp[arr[largest]] {
		largest = left
	}
	if right < len(arr) && dp[arr[right]] > dp[arr[largest]] {
		largest = right
	}
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		downCheck(arr, dp, largest)
	}
}

func maxResult(nums []int, k int) int {
	MIN := -10000000000
	m := len(nums)

	dp := make([]int, m)
	dp[0] = nums[0]
	var priorQueue []int
	priorQueue = append(priorQueue, 0)

	for i := 1; i < m; i++ {

		for priorQueue[0] < i-k {
			dp[priorQueue[0]] = MIN
			downCheck(priorQueue, dp, 0)
		}

		dp[i] = dp[priorQueue[0]] + nums[i]

		priorQueue = append(priorQueue, i)
		upCheck(priorQueue, dp, len(priorQueue)-1)
	}
	return dp[m-1]
}

func main() {
	fmt.Println(maxResult([]int{100, -100, -300, -300, -300, -100, 100}, 4))
}
