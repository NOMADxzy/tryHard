package main

import (
	"fmt"
	"sort"
)

func smallerCounts(nums []int, val int) int {
	cnt := 0
	for j := len(nums) - 1; j > 0; j-- {
		left, right := 0, j-1
		if nums[j]-nums[left] <= val {
			cnt += j
		} else if nums[j]-nums[right] <= val {
			for left < right {
				mid := (left + right) / 2
				if nums[j]-nums[mid] > val {
					left = mid + 1
				} else {
					right = mid
				}
			}
			cnt += j - right
		}
	}
	return cnt
}

func smallestDistancePair(nums []int, k int) int {
	m := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	left, right := 0, nums[m-1]-nums[0]
	if m*(m-1)/2 < k {
		return -1
	} else if m*(m-1)/2 == k {
		return right
	}
	// smallerCounts(left) = 0
	// smallerCounts(right) = wuqiong

	for left < right {
		mid := (left + right) / 2
		if smallerCounts(nums, mid) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(smallestDistancePair([]int{1, 3, 1}, 2))
}
