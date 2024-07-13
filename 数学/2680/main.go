package main

import "fmt"

func maximumOr(nums []int, k int) int64 {
	m := len(nums)
	pre, post := make([]int, m), make([]int, m)
	pre[0], post[m-1] = nums[0], nums[m-1]
	for i := 1; i < m-1; i++ {
		pre[i] = pre[i-1] | nums[i]
		post[m-i-1] = post[m-i] | nums[m-i-1]
	}
	ans := int64(0)
	b := int64(1 << k)
	for i := 0; i < m; i++ {
		tmp1, tmp2 := int64(0), int64(0)
		if i > 0 {
			tmp1 = int64(pre[i-1])
		}
		if i < m-1 {
			tmp2 = int64(post[i+1])
		}
		cur := int64(nums[i])*b | tmp1 | tmp2
		ans = max(ans, cur)
	}
	return ans
}
func main() {
	fmt.Println(maximumOr([]int{8, 1, 2}, 2))
}
