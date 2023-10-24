package main

import "fmt"

func longestArithSeqLength(nums []int) int {
	m := len(nums)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}

	ans := 2
	dp[0][1] = 2

	//idxMap := map[int][]int{nums[0]: {0}, nums[1]: {1}} //TODO 易错点，num[0]可能等于num[1]
	var idxMap map[int][]int
	if nums[0] == nums[1] {
		idxMap = map[int][]int{nums[0]: {0, 1}}
	} else {
		idxMap = map[int][]int{nums[0]: {0}, nums[1]: {1}}
	}

	for i := 2; i < m; i++ {
		for j := i - 1; j > 0; j-- {
			dp[j][i] = 2
			kPoses := idxMap[2*nums[j]-nums[i]]
			for _, k := range kPoses {
				if k < j && nums[i]-nums[j] == nums[j]-nums[k] && dp[k][j]+1 > dp[j][i] {
					dp[j][i] = dp[k][j] + 1
					if dp[j][i] > ans {
						ans = dp[j][i]
					}
				}
			}
		}
		dp[0][i] = 2
		idxMap[nums[i]] = append(idxMap[nums[i]], i)
	}
	return ans
}

func main() {
	fmt.Println(longestArithSeqLength([]int{1, 1, 9, 1}))
}
