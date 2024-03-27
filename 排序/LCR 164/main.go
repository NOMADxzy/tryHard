package main

import (
	"fmt"
	"sort"
	"strconv"
)

func crackPassword(password []int) string {
	compare := func(a, b int) bool {
		aStr, bStr := strconv.Itoa(a), strconv.Itoa(b)
		return aStr+bStr < bStr+aStr
	}
	sort.Slice(password, func(i, j int) bool {
		return compare(password[i], password[j])
	})
	ans := ""
	for i := 0; i < len(password); i++ {
		ans += strconv.Itoa(password[i])
	}
	return ans
}

func main() {
	fmt.Println(crackPassword([]int{15, 8, 7}))
}
