package main

import "fmt"

func compare(a string, b string) int {
	if len(a) < len(b) {
		return -1
	} else if len(a) == len(b) {
		if a > b {
			return 1
		} else if a == b {
			return 0
		} else {
			return -1
		}
	} else {
		return 1
	}
}

func partQuickSort(nums []string, start, end int, k int) {
	if end == start {
		return
	}
	guard := nums[start]
	l, r := start, end
	for l < r {
		for l < r && compare(nums[r], guard) <= 0 {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
		for l < r && compare(nums[l], guard) >= 0 {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	if l < k {
		partQuickSort(nums, l+1, end, k)
	} else if l > k {
		partQuickSort(nums, start, l-1, k)
	}
}

func kthLargestNumber(nums []string, k int) string {
	partQuickSort(nums, 0, len(nums)-1, k-1) //TODO 注意第k大对应 k-1 位置处
	return nums[k-1]
}

func main() {
	fmt.Println(kthLargestNumber([]string{"3", "6", "7", "10"}, 4))
}
