package main

import "fmt"

func beautifulIndices(s string, a string, b string, k int) []int {
	index := func(a byte) int {
		return int(a - 'a')
	}
	var listA, listB []int
	sA, sB := 1, 1
	for i := 0; i < len(a); i++ {
		sA *= 26
	}
	for i := 0; i < len(b); i++ {
		sB *= 26
	}
	hashA := 0
	hashB := 0
	for i := 0; i < len(a); i++ {
		hashA = hashA*26 + index(a[i])
	}
	for i := 0; i < len(b); i++ {
		hashB = hashB*26 + index(b[i])
	}
	hash1, hash2 := 0, 0
	for i := 0; i < len(s); i++ {
		hash1 = hash1*26 + index(s[i])
		hash2 = hash2*26 + index(s[i])
		if i >= len(a) {
			hash1 -= sA * index(s[i-len(a)])
		}
		if i >= len(b) {
			hash2 -= sB * index(s[i-len(b)])
		}

		if i >= len(a)-1 && hash1 == hashA {
			listA = append(listA, i-len(a)+1)
		}
		if i >= len(b)-1 && hash2 == hashB {
			listB = append(listB, i-len(b)+1)
		}
	}
	var ans []int
	check := func(a, b int) bool {
		delta := max(a-b, b-a)
		return delta <= k
	}
	if len(listB) == 0 {
		return ans
	}
	for i := 0; i < len(listA); i++ {
		target := listA[i]
		left, right := 0, len(listB)-1
		if listB[left] > target && !check(listB[left], target) ||
			listB[right] < target && !check(listB[right], target) {
			continue
		} else if check(listB[left], target) || check(listB[right], target) {
			ans = append(ans, listA[i])
			continue
		}
		for left < right {
			mid := (left + right) / 2
			if check(listB[mid], target) {
				ans = append(ans, listA[i])
				break
			} else if listB[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(beautifulIndices("lahhnlwx", "hhnlw", "ty", 6))
}
