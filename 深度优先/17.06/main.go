package main

import "fmt"

func numberOf2sInRange(n int) int {
	mask := 1
	for mask <= n {
		mask *= 10
	}
	mask /= 10

	var dfs func(mask int, pre2cnt int, preLimit bool) int
	dfs = func(mask int, pre2cnt int, preLimit bool) int {
		if mask == 0 {
			return pre2cnt
		}
		topVal := (n / mask) % 10
		curIs2, notLimited, limited := 0, 0, 0
		for i := 0; i <= 9; i++ {
			if preLimit && i > topVal {
				break
			}

			if i == 2 {
				curIs2++
			} else if preLimit && i == topVal {
				limited++
			} else {
				notLimited++
			}

		}
		total := 0
		if curIs2 > 0 {
			total += dfs(mask/10, pre2cnt+1, preLimit && topVal == 2)
		}
		if limited > 0 {
			total += dfs(mask/10, pre2cnt, true)
		}
		if notLimited > 0 {
			total += notLimited * dfs(mask/10, pre2cnt, false)
		}
		return total
	}
	return dfs(mask, 0, true)
}

func main() {
	fmt.Println(numberOf2sInRange(29))
}
