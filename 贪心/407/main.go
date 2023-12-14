package main

import "fmt"

func upCheck(arr [][]int, i int) {
	f := (i - 1) / 2
	if f >= 0 && arr[i][0] < arr[f][0] {
		arr[i], arr[f] = arr[f], arr[i]
		upCheck(arr, f)
	}
}

func downCheck(arr [][]int, i int) {
	left, right := 2*i+1, 2*i+2
	smallest := i
	if left < len(arr) && arr[left][0] < arr[smallest][0] {
		smallest = left
	}
	if right < len(arr) && arr[right][0] < arr[smallest][0] {
		smallest = right
	}
	if smallest != i {
		arr[smallest], arr[i] = arr[i], arr[smallest]
		downCheck(arr, smallest)
	}
}

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	var queue [][]int
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 && !visited[i][j] {
				queue = append(queue, []int{heightMap[i][j], i, j})
				visited[i][j] = true
				upCheck(queue, len(queue)-1)
			}
		}
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ans := 0

	for len(queue) > 0 {
		pos := queue[0]
		h := pos[0]
		queue[0] = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		downCheck(queue, 0)

		for _, dir := range dirs {
			nx, ny := pos[1]+dir[0], pos[2]+dir[1]
			if nx > 0 && nx < m-1 && ny > 0 && ny < n-1 && !visited[nx][ny] {
				if val := h - heightMap[nx][ny]; val > 0 {
					ans += val
				}
				queue = append(queue, []int{max(heightMap[nx][ny], h), nx, ny})
				visited[nx][ny] = true
				upCheck(queue, len(queue)-1)
			}
		}
	}
	return ans
}

func main() {
	//height := [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 10}, {11, 12, 13}, {14, 15, 16}}
	//height := [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}
	height := [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}
	fmt.Println(trapRainWater(height))
}
