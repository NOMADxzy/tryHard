package main

import (
	"fmt"
	"sort"
)

func videoStitching(clips [][]int, time int) int {
	MAXINT := 10000000000
	m := len(clips)
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 2)
	}
	if clips[0][0] > 0 {
		return -1
	}
	dp[0][0], dp[0][1] = clips[0][1], 1
	maxRight := clips[0][1]
	ans := MAXINT
	if clips[0][1] >= time {
		return 1
	}

	for i := 1; i < m; i++ {
		if clips[i][0] > maxRight {
			break
		}
		if clips[i][1] <= maxRight {
			continue
		}
		maxRight = clips[i][1]
		val, cnt := clips[i][1], MAXINT
		if clips[i][0] == 0 {
			cnt = 0
		} else {
			for j := i - 1; j >= 0; j-- {
				if dp[j][0] >= clips[i][0] && dp[j][1] < cnt {
					cnt = dp[j][1]
				}
			}
		}
		dp[i][0], dp[i][1] = val, cnt+1
		if dp[i][0] >= time && dp[i][1] < ans {
			ans = dp[i][1]
		}
	}
	if ans == MAXINT {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(videoStitching([][]int{{17, 18}, {25, 26}, {16, 21}, {3, 3}, {19, 23}, {1, 5}, {0, 2}, {9, 20}, {5, 17}, {8, 10}}, 15))
}
