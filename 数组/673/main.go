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

//仅适合互不相等的数组
func findNumberOfLIS1(nums []int) int {

	maxN := -1000000
	minN := 1000000
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxN {
			maxN = nums[i]
		}
		if nums[i] < minN {
			minN = nums[i]
		}
	}

	countSubArr := make([]int, len(nums))
	for i := 0; i < len(countSubArr); i++ {
		countSubArr[i] = 1
	}

	epoch := 1
	var longestCount int
	for epoch < len(nums) { //开始计算2、3、4。。。长度子串的数目
		tree := make([]int, maxN)
		cnt := 0
		for i := epoch - 1; i < len(nums); i++ {
			res := query(tree, nums[i])
			cnt += res
			update(tree, nums[i], countSubArr[i])
			countSubArr[i] = res
		}
		if cnt == 0 {
			break
		}
		longestCount = cnt
		epoch++
	}
	return longestCount
}

func findNumberOfLIS(nums []int) int {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 1
	dp[0][1] = 1

	for i := 1; i < len(nums); i++ {
		max_l, max_c := 0, 1

		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[j][0] > max_l {
					max_l = dp[j][0]
					max_c = dp[j][1]
				} else if dp[j][0] == max_l {
					max_c += dp[j][1]
				}
			}
		}
		dp[i][0] = max_l + 1
		dp[i][1] = max_c
	}
	cnt := 0
	max := 0
	for i := 0; i < len(dp); i++ {
		if dp[i][0] > max {
			max = dp[i][0]
			cnt = dp[i][1]
		} else if dp[i][0] == max {
			cnt += dp[i][1]
		}
	}
	return cnt
}

func main() {
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2}))
}
