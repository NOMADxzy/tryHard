package main

import "fmt"

func maxScore1(nums []int, x int) int64 {
	last0, last1 := -1, -1
	m := len(nums)
	dp := make([]int64, m)
	dp[0] = int64(nums[0])
	if nums[0]%2 == 0 {
		last0 = 0
	} else {
		last1 = 0
	}
	ans := dp[0]
	for i := 1; i < m; i++ {
		if nums[i]%2 == 0 {
			dp[i] = -1 << 60
			if last0 != -1 {
				dp[i] = dp[last0] + int64(nums[i])
			}
			if last1 != -1 {
				dp[i] = max(dp[i], dp[last1]+int64(nums[i]-x))
			}
			last0 = i
		} else {
			dp[i] = -1 << 60
			if last1 != -1 {
				dp[i] = dp[last1] + int64(nums[i])
			}
			if last0 != -1 {
				dp[i] = max(dp[i], dp[last0]+int64(nums[i]-x))
			}
			last1 = i
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func maxScore(nums []int, x int) int64 {
	m := len(nums)
	hasOdd, hasEven := false, false
	lastOddDp, lastEvenDp := int64(0), int64(0)
	if nums[0]%2 == 0 {
		hasEven = true
		lastEvenDp = int64(nums[0])
	} else {
		hasOdd = true
		lastOddDp = int64(nums[0])
	}
	for i := 1; i < m; i++ {
		if nums[i]%2 == 0 {
			if hasEven {
				lastEvenDp += int64(nums[i])
			}
			if hasOdd {
				newVal := lastOddDp + int64(nums[i]-x)
				if hasEven {
					lastEvenDp = max(lastEvenDp, newVal)
				} else {
					lastEvenDp = newVal
				}
			}
			hasEven = true
		} else {
			if hasOdd {
				lastOddDp += int64(nums[i])
			}
			if hasEven {
				newVal := lastEvenDp + int64(nums[i]-x)
				if hasOdd {
					lastOddDp = max(lastOddDp, newVal)
				} else {
					lastOddDp = newVal
				}
			}
			hasOdd = true
		}
	}
	if !hasOdd {
		return lastEvenDp
	} else if !hasEven {
		return lastOddDp
	} else {
		return max(lastOddDp, lastEvenDp)
	}
}

func main() {
	fmt.Println(maxScore1([]int{2, 3, 6, 1, 9, 2}, 5))
}
