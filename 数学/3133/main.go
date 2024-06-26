package main

import "fmt"

func minEnd(n int, x int) int64 {
	mask := int64(1)
	val := 1
	for mask <= 1<<61 {
		if mask&int64(x) == 0 {
			val *= 2
		}
		mask <<= 1
	}
	mask /= 2
	val /= 2
	ans := int64(x)
	rank := 0
	for mask > 0 { // 0100
		if ans&mask == 0 {
			if rank+val+1 == n {
				ans += mask
				rank += val
				break
			} else if rank+val+1 < n {
				ans += mask
				rank += val
			}
			val /= 2
		}
		mask /= 2
	}
	return ans
}

func main() {
	fmt.Println(minEnd(4, 3))
}
