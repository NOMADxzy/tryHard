package main

import "fmt"

func arrayNesting(nums []int) int {
	m := len(nums)
	cnts := map[int]int{}

	for i := 0; i < m; i++ {
		if cnts[nums[i]] == 0 {

			root := i
			cnts[root] += 1
			for cur := nums[root]; cur != root; {
				cnts[root]++
				tmp := cur
				cur = nums[cur]
				nums[tmp] = root
			}
		}
	}

	ans := 0
	for _, cnt := range cnts {
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}

func main() {
	fmt.Println(arrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))
}
