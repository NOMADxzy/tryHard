package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func compute(a int, b int) int {
	s := strconv.Itoa(b)
	min10pow := int(math.Pow10(len(s)))
	return a*min10pow + b
}

func bigger(a int, b int) bool {
	s1 := strconv.Itoa(a)
	s2 := strconv.Itoa(b)
	if s1[0] > s2[0] {
		return true
	} else if s1[0] < s2[0] {
		return false
	} else {
		r1 := compute(a, b)
		r2 := compute(b, a)
		if r1 > r2 {
			return true
		} else {
			return false
		}
	}
}

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a := nums[i]
		b := nums[j]
		return bigger(a, b)
	})

	if nums[0] == 0 {
		return "0"
	}

	res := ""
	for i := 0; i < len(nums); i++ {
		res += strconv.Itoa(nums[i])
	}
	return res
}

func main() {
	fmt.Println(largestNumber([]int{10, 2, 9, 39, 17}))
}
