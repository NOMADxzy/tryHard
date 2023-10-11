package main

import "fmt"

func getSumAbsoluteDifferences(nums []int) []int {
	m := len(nums)
	sums := make([]int, m)
	for i := 1; i < m; i++ {
		sums[i] = sums[i-1] + nums[i] - nums[i-1]
	}
	//sumsums := make([]int, m)
	//for i := 1; i < m; i++ {
	//	sumsums[i] = sumsums[i-1] + sums[i]
	//}

	ans := make([]int, m)
	for i := 0; i < m; i++ {
		s := 0
		if i > 0 {
			s = ans[i-1]
			d := nums[i] - nums[i-1]
			s += i * d
			s -= (m - i) * d
		} else {
			for j := 0; j < m; j++ {
				if i > j {
					s += sums[i] - sums[j]
				} else if i < j {
					s += sums[j] - sums[i]
				}
			}
		}
		ans[i] = s
	}
	return ans
}

func main() {
	fmt.Println(getSumAbsoluteDifferences([]int{2, 3, 5}))
}
