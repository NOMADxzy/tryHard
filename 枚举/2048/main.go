package main

import "fmt"

func findCombines(res *[][]int, cur []int, limit int, mark int) {
	if limit == 0 {
		*res = append(*res, cur)
		return
	}
	mask := 1
	for i := 1; i <= limit; i++ {
		if mask&mark == 0 {
			findCombines(res, append(cur, i), limit-i, mark+mask)
		}
		mask *= 2
	}
}

func findNums(min *int, arr []int, acc, mark, bottom, remain int) {
	if remain == 0 {
		if acc > bottom && acc < *min {
			*min = acc
		}
	}

	mask := 1
	for i := 0; i < len(arr); i++ {
		if mask&mark == 0 {
			findNums(min, arr, acc*10+arr[i], mark+mask, bottom, remain-1)
		}
		mask *= 2
	}
}

func findTheNum(limit int, pos int) []int {
	if limit == 0 {
		return []int{}
	} else if limit < pos {
		return nil
	}
	for i := pos; i <= limit; i++ {
		arr := findTheNum(limit-i, i+1)
		if arr != nil {
			for j := 0; j < i; j++ {
				arr = append(arr, i)
			}
			return arr
		}
	}
	return nil
}

func nextBeautifulNumber(n int) int {
	w := 0
	for i := n; i > 0; i /= 10 {
		w++
	}
	var combines [][]int
	var cur []int
	findCombines(&combines, cur, w, 0)

	arr := make([]int, w)
	min := 10000000
	for _, cb := range combines {
		idx := 0
		for i := 0; i < len(cb); i++ {
			for j := 0; j < cb[i]; j++ {
				arr[idx] = cb[i]
				idx++
			}
		}
		findNums(&min, arr, 0, 0, n, w)
	}
	if min == 10000000 {
		arr1 := findTheNum(w+1, 1)
		min = 0
		for i := len(arr1) - 1; i >= 0; i-- {
			min = min*10 + arr1[i]
		}
	}
	return min
}

func main() {
	fmt.Println(nextBeautifulNumber(1334))
}
