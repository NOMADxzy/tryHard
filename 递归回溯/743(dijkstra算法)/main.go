package main

import "fmt"

func networkDelayTime(times [][]int, n int, k int) int {
	vis := make([]bool, n)
	mat := make([][]int, n)
	MAX_INT := 1<<31 - 1
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				mat[i][j] = MAX_INT
			}
		}
	}
	k--
	for _, time := range times {
		mat[time[0]-1][time[1]-1] = time[2]
	}

	//vis[k] = true
	for i := 0; i < n; i++ {
		bestDist := MAX_INT
		bestPoint := -1
		for j := 0; j < n; j++ {
			if !vis[j] && mat[k][j] < bestDist {
				bestPoint = j
				bestDist = mat[k][j]
			}
		}
		if bestPoint < 0 {
			break
		}
		//
		for j := 0; j < n; j++ {
			if !vis[j] && bestDist+mat[bestPoint][j] < mat[k][j] {
				mat[k][j] = bestDist + mat[bestPoint][j]
			}
		}
		vis[bestPoint] = true
	}
	ans := 0
	for i := 0; i < n; i++ {
		if mat[k][i] > ans {
			ans = mat[k][i]
		}
	}
	if ans == MAX_INT {
		return -1
	}
	return ans
}

func main() {
	edges := [][]int{{1, 2, 5}, {1, 3, 2}, {2, 4, 1}, {2, 5, 6}, {3, 4, 6}, {3, 6, 8}, {4, 5, 1}, {4, 6, 2}, {5, 7, 7}, {6, 7, 3}}
	fmt.Println(networkDelayTime(edges, 7, 1))
}
