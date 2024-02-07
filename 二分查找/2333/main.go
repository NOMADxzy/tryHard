package main

import (
	"fmt"
	"sort"
)

func upCheck(arr []int, i int) {
	f := (i - 1) / 2
	if f >= 0 && arr[f] < arr[i] {
		arr[i], arr[f] = arr[f], arr[i]
		upCheck(arr, f)
	}
}

func downCheck(arr []int, i int) {
	left, right := 2*i+1, 2*i+2
	largest := i
	if left < len(arr) && arr[left] > arr[largest] {
		largest = left
	}
	if right < len(arr) && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		downCheck(arr, largest)
	}
}

// 优先队列、超时
func minSumSquareDiff1(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	n := len(nums1)
	var deltas []int
	ans := int64(0)

	for i := 0; i < n; i++ {
		d := nums2[i] - nums1[i]
		if d < 0 {
			d = -d
		}
		if d > 0 {
			deltas = append(deltas, d)
			upCheck(deltas, len(deltas)-1)
			ans += int64(d * d)
		}
	}

	// 贪心的每一步选择能降低最多值的操作
	for i := 0; i < k1+k2 && len(deltas) > 0; i++ {
		v := deltas[0]
		ans -= int64(2*v - 1) // v*v - (v-1)*(v-1)

		deltas[0] = deltas[len(deltas)-1]
		deltas = deltas[:len(deltas)-1]
		downCheck(deltas, 0)

		if v-1 > 0 {
			deltas = append(deltas, v-1)
			upCheck(deltas, len(deltas)-1)
		}
	}
	return ans
}

// 二分查找
func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	deltas := make([]int, len(nums1))
	ans := int64(0)

	for i := 0; i < len(nums1); i++ {
		d := nums2[i] - nums1[i]
		if d < 0 {
			d = -d
		}
		deltas[i] = d
	}
	deltas = append(deltas, 0)
	n := len(deltas)
	sort.Slice(deltas, func(i, j int) bool {
		return deltas[i] > deltas[j]
	})

	preSums := make([]int, n)
	preSums[0] = deltas[0]
	for i := 1; i < n; i++ {
		preSums[i] = preSums[i-1] + deltas[i]
	}
	target := k1 + k2
	left, right := 1, n-1
	if preSums[n-1] <= target {
		return int64(0)
	} else if target == 0 {
		right = 0
	}
	for left < right {
		mid := (left + right) / 2
		lim := mid * deltas[mid]
		if preSums[mid-1]-lim < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if right > 0 {
		leftRemain := preSums[right] - deltas[right] - target
		newVal := leftRemain / right
		up1Num := leftRemain - newVal*right
		ans += int64((newVal+1)*(newVal+1))*int64(up1Num) + int64(newVal*newVal)*int64(right-up1Num)
	}
	for i := right; i < n; i++ {
		ans += int64(deltas[i] * deltas[i])
	}
	return ans
}

func main() {
	fmt.Println(minSumSquareDiff([]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1))
}
