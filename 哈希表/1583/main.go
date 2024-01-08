package main

import "fmt"

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	marks := make([]bool, n)
	idxs := map[int]int{}
	ranks := make([]map[int]int, n)
	for i, preference := range preferences {
		ranks[i] = map[int]int{}
		for j := 0; j < len(preference); j++ {
			ranks[i][preference[j]] = j
		}
	}
	for i, pair := range pairs {
		idxs[pair[0]] = i
		idxs[pair[1]] = i
	}
	for _, pair := range pairs {
		x, y := pair[0], pair[1]
		if preferences[x][0] != y {
			for j := 0; j < n-1 && preferences[x][j] != y && !marks[x]; j++ {
				target := preferences[x][j]
				u, v := pairs[idxs[target]][0], pairs[idxs[target]][1]
				if u != target {
					u, v = v, u
				}
				if ranks[u][x] < ranks[u][v] {
					marks[x] = true
				}
			}
		}
		if preferences[y][0] != x {
			for j := 0; j < n-1 && preferences[y][j] != x && !marks[y]; j++ {
				target := preferences[y][j]
				u, v := pairs[idxs[target]][0], pairs[idxs[target]][1]
				if u != target {
					u, v = v, u
				}
				if ranks[u][y] < ranks[u][v] {
					marks[y] = true
				}
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		if marks[i] {
			ans++
		}
	}
	return ans
}

func main() {
	prefer := [][]int{{1, 2, 3}, {3, 2, 0}, {3, 1, 0}, {1, 2, 0}}
	fmt.Println(unhappyFriends(4, prefer, [][]int{{0, 1}, {2, 3}}))
}
