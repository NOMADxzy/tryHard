package main

import "fmt"

func solve1(arr []string) int {
	cnt := 0
	m, n := len(arr), len(arr[0])
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	target := "tencent"
	hist := map[int]int{}
	var dfs func(x, y int, pos int) int
	dfs = func(x, y int, pos int) int {
		key := 1001*1001*pos + 1001*x + y
		if v, ok := hist[key]; ok {
			return v
		}
		cnt = 0
		if arr[x][y] == target[pos] {
			if pos == 6 {
				cnt = 1
			} else {
				for _, dir := range dirs {
					nx, ny := x+dir[0], y+dir[1]
					if nx >= 0 && nx < m && ny >= 0 && ny < n {
						nextCnts := dfs(nx, ny, pos+1)
						cnt += nextCnts
					}
				}
			}
		}
		hist[key] = cnt
		return cnt
	}
	ans := 0
	for pos := 6; pos >= 0; pos-- {
		ans = 0
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				ans += dfs(i, j, pos)
			}
		}
	}

	return ans
}

func solve(arr []string) int {
	m, n := len(arr), len(arr[0])
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	target := "tencent"
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if target[len(target)-1] == arr[i][j] {
				dp[i][j] = 1
			}
		}
	}

	ans := 0
	for pos := 5; pos >= 0; pos-- {
		ans = 0
		newDp := make([][]int, m)
		for i := 0; i < m; i++ {
			newDp[i] = make([]int, n)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if arr[i][j] == target[pos] {
					acc := 0
					for _, dir := range dirs {
						nx, ny := i+dir[0], j+dir[1]
						if nx >= 0 && nx < m && ny >= 0 && ny < n {
							acc += dp[nx][ny]
						}
					}
					newDp[i][j] = acc
					ans += acc
				}
			}
		}
		dp = newDp
	}

	return ans
}

func main() {
	var m, n int
	_, _ = fmt.Scan(&m, &n)
	arr := make([]string, m)
	for i := 0; i < m; i++ {
		_, _ = fmt.Scan(&arr[i])
	}
	fmt.Println(solve(arr))
}
