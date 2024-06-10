package main

import "fmt"

func upCheck(pairs [][]int, i int) {
	f := (i - 1) / 2
	for i > 0 && pairs[i][0] < pairs[f][0] {
		pairs[i], pairs[f] = pairs[f], pairs[i]
		i = f
		f = (i - 1) / 2
	}
}

func downCheck(pairs [][]int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	smallest := i
	if left < len(pairs) && pairs[left][0] < pairs[smallest][0] {
		smallest = left
	}
	if right < len(pairs) && pairs[right][0] < pairs[smallest][0] {
		smallest = right
	}
	if smallest != i {
		pairs[i], pairs[smallest] = pairs[smallest], pairs[i]
		downCheck(pairs, smallest)
	}
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m := len(nums1)
	ptrs := make([]int, m)
	var pairs [][]int
	for i := 0; i < m; i++ {
		pairs = append(pairs, []int{nums1[i] + nums2[ptrs[i]], i})
	}
	var ans [][]int
	for len(ans) < k && len(pairs) > 0 {
		t := pairs[0][1]
		ans = append(ans, []int{nums1[t], nums2[ptrs[t]]})
		pairs[0] = pairs[len(pairs)-1]
		downCheck(pairs, 0)
		if ptrs[t]+1 < len(nums2) {
			ptrs[t]++
			newPair := []int{nums1[t] + nums2[ptrs[t]], t}
			pairs[len(pairs)-1] = newPair
			upCheck(pairs, len(pairs)-1)
		} else {
			pairs = pairs[:len(pairs)-1]
		}
	}
	return ans
}

func main() {
	fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
}
