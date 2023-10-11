package main

import "fmt"

func waysToSplit(nums []int) int {
	MOD_NUM := 1000000007
	m := len(nums)
	ans := 0
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < m; i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	leftLimit := sums[m-1] / 3
	//1,2,2,2,5,0
	for idx := 1; idx < m-1 && sums[idx-1] <= leftLimit; idx++ {
		leftSum := sums[idx-1]
		l, r := idx, m-2
		if sums[l]-leftSum < leftSum {
			left, right := idx, r
			if sums[right]-leftSum < leftSum {
				continue
			}
			for left < right {
				mid := (left + right) / 2
				if sums[mid]-leftSum < leftSum {
					left = mid + 1
				} else {
					right = mid
				}
			}
			l = right
		}
		if sums[m-1]-sums[r] < sums[r]-leftSum {
			left, right := idx, r
			if sums[m-1]-sums[left] < sums[left]-leftSum {
				continue
			}
			for left < right {
				mid := (left + right) / 2
				if sums[m-1]-sums[mid] >= sums[mid]-leftSum {
					left = mid + 1
				} else {
					right = mid
				}
			}
			r = left - 1
		}
		if l <= r {
			ans = (ans + (r-l+1)%MOD_NUM) % MOD_NUM
		}
	}
	return ans
}

func main() {
	fmt.Println(waysToSplit([]int{1, 2, 2, 2, 5, 0}))
}
