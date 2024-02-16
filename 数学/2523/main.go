package main

import "fmt"

func closestPrimes(left int, right int) []int {

}

// 埃筛 超时
func closestPrimes1(left int, right int) []int {
	mark := [1e6 + 1]bool{}           // 预分配降低耗时 TODO
	altoNums := make([]int, 0, 78500) // 降低耗时
	pre := -1
	ans := []int{-1, -1}
	for i := 2; i <= right; i++ {
		if !mark[i] {
			altoNums = append(altoNums, i)
			if i >= left {
				if pre > 0 {
					if ans[0] == -1 || i-pre < ans[1]-ans[0] {
						ans[0], ans[1] = pre, i
					}
				}
				pre = i
			}
		}
		for _, x := range altoNums {
			if x*i > right {
				break
			}
			mark[x*i] = true
			if i%x == 0 {
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(closestPrimes(19, 31))
}
