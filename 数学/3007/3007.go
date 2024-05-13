package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func findMaximumNumber(k int64, x int) int64 {
	ans := sort.Search(int(k+1)<<x, func(num int) bool {
		num++
		n := bits.Len(uint(num))
		memo := make([][]int, n)
		for i := range memo {
			memo[i] = make([]int, n+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}
		var dfs func(int, int, bool) int
		dfs = func(i, cnt1 int, limitHigh bool) (res int) {
			if i < 0 {
				return cnt1
			}
			if !limitHigh {
				p := &memo[i][cnt1]
				if *p >= 0 {
					return *p
				}
				defer func() { *p = res }()
			}
			up := 1
			if limitHigh {
				up = num >> i & 1
			}
			for d := 0; d <= up; d++ {
				c := cnt1
				if d == 1 && (i+1)%x == 0 {
					c++
				}
				res += dfs(i-1, c, limitHigh && d == up)
			}
			return
		}
		return dfs(n-1, 0, true) > int(k)
	})
	return int64(ans)
}

func findMaximumNumber1(k int64, x int) int64 {
	ans := sort.Search(int(k+1)<<x, func(num int) bool {
		n := bits.Len(uint(num + 1)) // 补一个0
		memo := make([][]int, n)
		for i := range memo {
			memo[i] = make([]int, n+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}
		var dfs func(int, int, bool) int
		dfs = func(i int, cnt1 int, limit bool) (res int) {
			if i < 0 {
				return cnt1
			}
			{
				key1 := i
				if limit {
					key1 *= -1
				}
				key2 := cnt1
				val := &memo[key1][key2]
				if *val >= 0 {
					return *val
				}
				defer func() { *val = res }()
			}
			upVal := 1
			if limit {
				upVal = num >> i & 1
			}
			for j := 0; j <= upVal; j++ {
				new1cnt := 0
				if j == 1 && (i+1)%x == 0 {
					new1cnt = 1
				}
				res += dfs(i-1, cnt1+new1cnt, limit && j == upVal)
			}
			return
		}
		return dfs(n-1, 0, true) > int(k)
	})
	return int64(ans)
}

func main() {
	fmt.Println(findMaximumNumber(9, 1))
}
