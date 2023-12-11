package main

import "fmt"

func GetNextBigger(num []int) (nextBigger []int) {
	nextBigger = make([]int, len(num))
	var decStack []int
	for i := 0; i < len(num); i++ {
		for len(decStack) > 0 && num[decStack[len(decStack)-1]] < num[i] {
			idx := decStack[len(decStack)-1]
			decStack = decStack[:len(decStack)-1]
			nextBigger[idx] = i
		}
		nextBigger[i] = -1
		decStack = append(decStack, i)
	}
	return
}

func findNext(nums1, nums2, nextB1, nextB2 []int, pos1, pos2 int, shouldRemain int) (ans []int) {
	remain := len(nums1) + len(nums2) - pos1 - pos2
	if shouldRemain == 0 {
		return []int{}
	}
	best1, best2 := -1, -1
	best1idx, best2idx := pos1, pos2
	if pos1 < len(nums1) {
		var i int
		for i = pos1; nextB1[i] >= 0 && nextB1[i] <= pos1+remain-shouldRemain; i = nextB1[i] {
		}
		best1 = nums1[i]
		best1idx = i
	}
	if pos2 < len(nums2) {
		var i int
		for i = pos2; nextB2[i] >= 0 && nextB2[i] <= pos2+remain-shouldRemain; i = nextB2[i] {
		}
		best2 = nums2[i]
		best2idx = i
	}
	curVal := max(best1, best2)
	if best1 > best2 {
		ans = append(findNext(nums1, nums2, nextB1, nextB2, best1idx+1, pos2, shouldRemain-1), curVal)
	} else if best2 > best1 {
		ans = append(findNext(nums1, nums2, nextB1, nextB2, pos1, best2idx+1, shouldRemain-1), curVal)
	} else {
		nextAns := findNext(nums1, nums2, nextB1, nextB2, best1idx+1, pos2, shouldRemain-1)
		nextAns1 := findNext(nums1, nums2, nextB1, nextB2, pos1, best2idx+1, shouldRemain-1)
		choose1 := false
		for i := shouldRemain - 2; i >= 0; i-- {
			if nextAns1[i] != nextAns[i] {
				choose1 = nextAns1[i] > nextAns[i]
				break
			}
		}
		if choose1 {
			ans = append(nextAns1, curVal)
		} else {
			ans = append(nextAns, curVal)
		}
	}
	return ans
}

func maxNumber1(nums1 []int, nums2 []int, k int) []int {
	nextB1, nextB2 := GetNextBigger(nums1), GetNextBigger(nums2)
	ans := findNext(nums1, nums2, nextB1, nextB2, 0, 0, k)
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}

func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	for i := range merged {
		if lexicographicalLess(a, b) {
			merged[i], b = b[0], b[1:]
		} else {
			merged[i], a = a[0], a[1:]
		}
	}
	return merged
}

func choose(nums, nextBigger []int, k int) (ans []int) {
	i := 0
	for len(ans) < k {
		for nextBigger[i] > 0 && len(nums)-nextBigger[i]+len(ans) >= k {
			i = nextBigger[i]
		}
		ans = append(ans, nums[i])
		i++
	}
	return
}

func maxNumber(nums1, nums2 []int, k int) (res []int) {
	nextB1, nextB2 := GetNextBigger(nums1), GetNextBigger(nums2)
	leftLen := min(len(nums1), k)
	for i := k - min(len(nums2), k); i <= leftLen; i++ {
		s1, s2 := choose(nums1, nextB1, i), choose(nums2, nextB2, k-i)
		ans := merge(s1, s2)
		if lexicographicalLess(res, ans) {
			res = ans
		}
	}
	return
}

func main() {
	nums1 := []int{2, 5, 6, 4, 4, 0}
	nums2 := []int{7, 3, 8, 0, 6, 5, 7, 6, 2}
	fmt.Println(maxNumber(nums1, nums2, 15))
}
