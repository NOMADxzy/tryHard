package main

import "fmt"

func sumFourDivisors(nums []int) int {
	maxVal := nums[0]
	for i := 0; i < len(nums); i++ {
		maxVal = max(maxVal, nums[i])
	}
	compose := make([]bool, maxVal+1)
	for i := 2; i <= maxVal; i++ {
		if !compose[i] {
			for j := i * 2; j <= maxVal; j += i {
				compose[j] = true
			}
		}
	}
	var composeList []int
	for v, _ := range compose {
		if v >= 2 && !compose[v] {
			composeList = append(composeList, v)
		}
	}
	check := func(num int) int {
		if !compose[num] {
			return 0
		}
		for i := 0; i < len(composeList); i++ {
			x := composeList[i]
			if x*x >= num {
				break
			}
			if num%x == 0 && (!compose[num/x] || num/x == x*x) {
				return 1 + x + num/x + num
			}
		}
		return 0
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += check(nums[i])
	}
	return ans
}

func main() {
	fmt.Println(sumFourDivisors([]int{21, 4, 7}))
}
