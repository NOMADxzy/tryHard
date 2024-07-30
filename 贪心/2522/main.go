package main

import "fmt"

func minimumPartition(s string, k int) int {
	m := len(s)
	acc := 0
	cnt := 0
	for i := 0; i < m; i++ {
		cur := int(s[i] - '0')
		if acc*10+cur <= k {
			acc = acc*10 + cur
		} else {
			if acc == 0 {
				return -1
			} else {
				cnt++
				if cur > k {
					return -1
				}
				acc = cur
			}
		}
	}
	if acc > 0 {
		cnt++
	}
	return cnt
}

func main() {
	fmt.Println(minimumPartition("165462", 60))
}
