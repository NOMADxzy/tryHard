package main

import "fmt"

func platesBetweenCandles(s string, queries [][]int) []int {
	var idxs []int
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			idxs = append(idxs, i)
		}
	}
	ans := make([]int, len(queries))
	if len(idxs) < 2 {
		return ans
	}
	for i, query := range queries {
		target1, target2 := query[0], query[1]
		l1, r1 := 0, len(idxs)-1
		if target1 >= idxs[r1] {
			l1 = r1 + 1
		} else if target1 <= idxs[l1] {
		} else {
			for l1 < r1 {
				mid := (l1 + r1) / 2
				if idxs[mid] < target1 {
					l1 = mid + 1
				} else {
					r1 = mid
				}
			}
		}
		l2, r2 := 0, len(idxs)-1
		if target2 <= idxs[l2] {
			r2 = l2 - 1
		} else if target2 >= idxs[r2] {
		} else {
			for l2 < r2 {
				mid := (l2 + r2) / 2
				if idxs[mid] <= target2 {
					l2 = mid + 1
				} else {
					r2 = mid
				}
			}
			r2--
		}
		v := 0
		if r2 > l1 {
			v = idxs[r2] - idxs[l1] - (r2 - l1)
		}
		ans[i] = v
	}
	return ans
}

func main() {
	fmt.Println(platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}}))
}
