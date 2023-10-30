package main

import "fmt"

//标答
func findPairs(nums []int, k int) int {
	visited := map[int]bool{}
	res := map[int]bool{}
	for _, num := range nums {
		if _, ok := visited[num-k]; ok {
			res[num-k] = true
		}
		if _, ok := visited[num+k]; ok {
			res[num] = true
		}
		visited[num] = true
	}
	return len(res)
}

func findPairs1(nums []int, k int) int {
	m := len(nums)
	ans := 0
	nMap := map[int]bool{}
	min := nums[0]
	max := nums[0]
	samePair := 0
	sameMap := map[int]bool{}

	for i := 0; i < m; i++ {
		if !nMap[nums[i]] {
			nMap[nums[i]] = true
		} else if !sameMap[nums[i]] {
			samePair++
			sameMap[nums[i]] = true
		}
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	if k == 0 {
		return samePair
	}

	for i := min; i < min+k; i++ {
		chainLen := 0
		for val := i; val < max; val += k {
			if nMap[val] && nMap[val+k] {
				chainLen++
			}
		}
		ans += chainLen
	}
	return ans
}

func main() {
	fmt.Println(findPairs([]int{1, 2}, 1))
}
