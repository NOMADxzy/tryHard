package main

import (
	"fmt"
	"sort"
)

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	n := len(friends)
	maxInt := 1 << 31
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = maxInt
		}
	}

	for i, fs := range friends {
		for _, f := range fs {
			if f > i {
				dp[i][f] = 1
				dp[f][i] = 1
			}
		}
	}

	for epoch := 1; epoch < level; epoch++ {
		for i, fs := range friends {
			for _, f := range fs {
				if f > i {
					for c := 0; c < n; c++ {
						if c != f && c != i && 1+dp[f][c] < dp[i][c] {
							dp[i][c] = 1 + dp[f][c]
							dp[c][i] = 1 + dp[f][c]
						}
						if c != f && c != i && 1+dp[i][c] < dp[f][c] {
							dp[f][c] = 1 + dp[i][c]
							dp[c][f] = 1 + dp[i][c]
						}
					}
				}
			}
		}
	}

	vMap := map[string]int{}
	var ans []string
	for j := 0; j < n; j++ {
		if dp[id][j] == level {

			for _, v := range watchedVideos[j] {
				if vMap[v] == 0 {
					ans = append(ans, v)
				}
				vMap[v]++
			}
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return vMap[ans[i]] < vMap[ans[j]] || vMap[ans[i]] == vMap[ans[j]] && ans[i] < ans[j]
	})
	return ans

}

func main() {
	wvs := [][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}
	fs := [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}

	fmt.Println(watchedVideosByFriends(wvs, fs, 0, 2))
}
