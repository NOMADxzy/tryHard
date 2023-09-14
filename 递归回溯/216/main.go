package main

import "fmt"

func find(combines *[][]int, combine []int, cur int, target int, limit int) {
	if target == 0 && len(combine) == limit {
		newArr := make([]int, limit)
		for i := 0; i < len(combine); i++ {
			newArr[i] = combine[i]
		}
		*combines = append(*combines, newArr)
	} else {
		if len(combine) >= limit {
			return
		}
		for i := cur; i < 10 && i <= target; i++ {
			find(combines, append(combine, i), i+1, target-i, limit)
		}
	}
}

func combinationSum3(k int, n int) [][]int {
	var combines [][]int
	find(&combines, []int{}, 1, n, k)
	return combines
}

func main() {
	fmt.Println(combinationSum3(8, 36))
}
