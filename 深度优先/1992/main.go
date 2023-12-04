package main

import "fmt"

func findNext(land [][]int, marked [][]bool, startX, startY int) (int, int) {
	marked[startX][startY] = true
	r1, b1, r2, b2 := startY, startX, startY, startX
	if startX+1 < len(land) && land[startX+1][startY] == 1 && !marked[startX+1][startY] {
		b1, r1 = findNext(land, marked, startX+1, startY)
	}
	if startY+1 < len(land[0]) && land[startX][startY+1] == 1 && !marked[startX][startY+1] {
		b2, r2 = findNext(land, marked, startX, startY+1)
	}
	return max(b1, b2), max(r1, r2)
}

func findFarmland(land [][]int) [][]int {
	m, n := len(land), len(land[0])
	marked := make([][]bool, m)
	for i := 0; i < m; i++ {
		marked[i] = make([]bool, n)
	}
	var ans [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 1 && !marked[i][j] {
				b, r := findNext(land, marked, i, j)
				ans = append(ans, []int{i, j, b, r})
			}
		}
	}
	return ans
}

func main() {
	land := [][]int{
		{1, 1},
		{0, 0},
	}
	fmt.Println(findFarmland(land))
}
