package main

import (
	"fmt"
	"sort"
)

type pair struct {
	val byte
	cnt int
}

func longestDiverseString(a int, b int, c int) string {
	arr := []*pair{{'a', a}, {'b', b}, {'c', c}}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].cnt > arr[j].cnt
	})
	drop := 0
	if v := a + b + c - arr[0].cnt; 2*(v+1) < arr[0].cnt {
		drop = arr[0].cnt - 2*(v+1)
		arr[0].cnt -= drop
	}

	getRemain := func(limit byte) string {
		sum := arr[0].cnt + arr[1].cnt + arr[2].cnt
		max_idx := 0
		if arr[0].val == limit {
			max_idx = 1
		}
		if arr[1].val != limit && arr[1].cnt > arr[max_idx].cnt {
			max_idx = 1
		}
		if arr[2].val != limit && arr[2].cnt > arr[max_idx].cnt {
			max_idx = 2
		}
		if arr[max_idx].cnt == 0 {
			return ""
		}
		if arr[max_idx].cnt*2 < sum || arr[max_idx].cnt == 1 {
			arr[max_idx].cnt--
			return string(arr[max_idx].val)
		} else {
			arr[max_idx].cnt -= 2
			return string(arr[max_idx].val) + string(arr[max_idx].val)
		}
	}
	ans := ""
	limit := byte('d')
	for {
		r := getRemain(limit)
		if r == "" {
			break
		}
		ans += r
		limit = r[0]
	}
	return ans
}

func main() {
	fmt.Println(longestDiverseString(3, 3, 3))
}
