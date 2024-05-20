package main

import "fmt"

// 00 25 50 75
func minimumOperations(num string) int {
	tails := [][]int{
		{0, 0},
		{5, 2},
		{0, 5},
		{5, 7},
	}
	ptrs := make([]int, 4)
	contains0 := false
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] == '0' {
			contains0 = true
		}
		for j := 0; j < 4; j++ {
			if int(num[i]-'0') == tails[j][ptrs[j]] {
				ptrs[j]++
				if ptrs[j] == 2 {
					return len(num) - i - 2
				}
			}
		}
	}
	ans := len(num)
	if contains0 {
		ans--
	}
	return ans
}

func main() {
	fmt.Println(minimumOperations("2245047"))
}
