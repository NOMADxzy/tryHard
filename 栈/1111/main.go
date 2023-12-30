package main

import "fmt"

func maxDepthAfterSplit(seq string) []int {
	var lstack, rstack int
	ans := make([]int, len(seq))
	var lmax, rmax int
	for i := 0; i < len(seq); i++ {
		c := seq[i]
		if c == '(' {
			if lstack <= rstack {
				lstack++
				lmax = max(lmax, lstack)
			} else {
				rstack++
				rmax = max(rmax, rstack)
				ans[i] = 1
			}
		} else {
			if lstack < rstack {
				rstack--
				ans[i] = 1
			} else {
				lstack--
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxDepthAfterSplit("()(())()"))
}
