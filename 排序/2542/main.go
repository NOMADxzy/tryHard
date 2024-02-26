package main

import (
	"fmt"
	"sort"
)

func upCheck(arr []int, i int) {
	f := (i - 1) / 2
	if f >= 0 && arr[i] < arr[f] {
		arr[i], arr[f] = arr[f], arr[i]
		upCheck(arr, f)
	}
}

func downCheck(arr []int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	smallest := i
	if left < len(arr) && arr[left] < arr[smallest] {
		smallest = left
	}
	if right < len(arr) && arr[right] < arr[smallest] {
		smallest = right
	}
	if smallest != i {
		arr[smallest], arr[i] = arr[i], arr[smallest]
		downCheck(arr, smallest)
	}
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	m := len(nums1)
	pairs := make([][2]int, m)
	for i := 0; i < m; i++ {
		pairs[i] = [2]int{nums1[i], nums2[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	sum := int64(0)
	ans := int64(0)
	var topK []int
	for i := m - 1; i >= 0; i-- {
		if len(topK) < k {
			topK = append(topK, pairs[i][0])
			sum += int64(pairs[i][0])
			upCheck(topK, len(topK)-1)
		} else if pairs[i][0] > topK[0] {
			sum += int64(pairs[i][0] - topK[0])
			topK[0] = pairs[i][0]
			downCheck(topK, 0)
		}
		if len(topK) == k {
			ans = max(ans, sum*int64(pairs[i][1]))
		}
	}
	return ans
}

func main() {
	fmt.Println(maxScore([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 1))
}
