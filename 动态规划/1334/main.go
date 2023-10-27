package main

import "fmt"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {

	INF := 10000000000
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = INF
		}
	}
	for _, edge := range edges {
		dp[edge[1]][edge[0]] = edge[2]
		dp[edge[0]][edge[1]] = edge[2]
	}

	flag := true
	for epoch := 1; epoch <= n && flag; epoch++ {
		flag = false
		for _, edge := range edges {
			f, t, w := edge[0], edge[1], edge[2]
			for i := 0; i < n; i++ {
				if i != f && i != t {
					if w+dp[t][i] < dp[f][i] {
						dp[i][f] = w + dp[t][i]
						dp[f][i] = w + dp[t][i]
						flag = true
					}
					if w+dp[f][i] < dp[t][i] {
						dp[t][i] = w + dp[f][i]
						dp[i][t] = w + dp[f][i]
						flag = true
					}
				}
			}
		}
	}
	neighs := make([]int, n)
	best := n
	bestIdx := n - 1
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if dp[i][j] <= distanceThreshold {
				neighs[i]++
				neighs[j]++
			}
		}
	}
	for i := 0; i < len(neighs); i++ {
		if neighs[i] < best || neighs[i] == best && i > bestIdx {
			bestIdx = i
			best = neighs[i]
		}
	}
	return bestIdx
}

func main() {
	edges := [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
	fmt.Println(findTheCity(4, edges, 4))
}
