package main

import (
	"fmt"
	"sort"
)

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	sorted := make([]int, len(nums1))
	copy(sorted, nums1)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	MOD := 1000000007
	bestDec := 0
	place := -1
	sum := 0
	for i := 0; i < len(nums2); i++ {
		target := nums2[i]
		pre := max(nums2[i]-nums1[i], nums1[i]-nums2[i])
		if pre <= bestDec {
			//sum = (sum + pre) % MOD
			continue
		}
		left, right := 0, len(sorted)-1
		var delta int
		if sorted[left] >= target {
			delta = sorted[left] - target
		} else if sorted[right] <= target {
			delta = target - sorted[right]
		} else {
			for left < right {
				mid := (left + right) / 2
				if sorted[mid] < target {
					left = mid + 1
				} else {
					right = mid
				}
			}
			//left-- //TODO 易错，left有可能为0
			if left > 0 {
				left--
			}
			delta = min(target-sorted[left], sorted[right]-target)
		}
		if pre-delta > bestDec {
			bestDec = pre - delta
			place = i
		}
		bestDec = max(bestDec, pre-delta)
		//sum = (sum + pre) % MOD
	}
	for i := 0; i < len(nums2); i++ {
		v := max(nums2[i]-nums1[i], nums1[i]-nums2[i])
		if i == place {
			v -= bestDec
		}
		sum = (sum + v) % MOD
	}
	return sum
}

func main() {
	fmt.Println(minAbsoluteSumDiff([]int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4}))
}
