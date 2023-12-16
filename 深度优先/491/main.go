package main

import "fmt"

// 深度优先搜索
func dfs(nums []int, pos int, acc []int, ans *[][]int) {
	if len(acc) > 1 {
		tmp := make([]int, len(acc))
		copy(tmp, acc)
		*ans = append(*ans, tmp)
	}
	visited := map[int]bool{} //去重
	for i := pos + 1; i < len(nums); i++ {
		if !visited[nums[i]] && (pos < 0 || nums[i] >= nums[pos]) {
			visited[nums[i]] = true
			dfs(nums, i, append(acc, nums[i]), ans)
		}
	}
}

func findSubsequences(nums []int) [][]int {
	var ans [][]int
	dfs(nums, -1, []int{}, &ans)
	return ans
}

// 动态规划，失败
func findSubsequences1(nums []int) [][]int {
	m := len(nums)
	res := make([][][]int, m)

	for i := 1; i < m; i++ {
		var curAns [][]int
		for j := i - 1; j >= 0; j-- {
			if nums[i] >= nums[j] {
				curAns = append(curAns, []int{nums[j], nums[i]})
				for _, preAns := range res[j] {
					tmp := make([]int, len(preAns)+1)
					copy(tmp, preAns)
					tmp[len(tmp)-1] = nums[i]
					curAns = append(curAns, tmp)
				}
			}
			if nums[i] == nums[j] {
				break
			}
		}
		res[i] = curAns
	}
	var ans [][]int
	for _, seqs := range res {
		for _, seq := range seqs {
			ans = append(ans, seq)
		}
	}
	return ans
}

func main() {
	a1 := findSubsequences([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	fmt.Println(len(a1))
}
