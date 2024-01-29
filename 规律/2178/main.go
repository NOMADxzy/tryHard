package main

import "fmt"

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 == 1 {
		return []int64{}
	}
	acc := int64(0)
	cur := int64(2)
	var ans []int64
	for acc < finalSum {
		if acc+cur <= finalSum {
			ans = append(ans, cur)
			acc += cur
			cur += 2
		} else {
			idx := int(acc+cur-finalSum)/2 - 1
			acc += cur - ans[idx]
			ans[idx] = cur
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumEvenSplit(4))
}
