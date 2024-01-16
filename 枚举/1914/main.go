package main

import "fmt"

func rotateGrid(grid [][]int, k int) [][]int {
	layer := 0
	m, n := len(grid), len(grid[0])
	maxLayer := min((m+1)/2, (n+1)/2)
	for layer < maxLayer {
		var arr []int
		x, y := layer, layer
		for x < m-layer-1 {
			arr = append(arr, grid[x][y])
			x++
		}
		for y < n-layer-1 {
			arr = append(arr, grid[x][y])
			y++
		}
		for x > layer {
			arr = append(arr, grid[x][y])
			x--
		}
		for y > layer {
			arr = append(arr, grid[x][y])
			y--
		}
		if len(arr) == 0 {
			break
		}
		d := k % len(arr)
		idx := 0
		for x < m-layer-1 {
			pos := idx - d
			if pos < 0 {
				pos += len(arr)
			}
			grid[x][y] = arr[pos]
			x++
			idx++
		}
		for y < n-layer-1 {
			pos := idx - d
			if pos < 0 {
				pos += len(arr)
			}
			grid[x][y] = arr[pos]
			y++
			idx++
		}
		for x > layer {
			pos := idx - d
			if pos < 0 {
				pos += len(arr)
			}
			grid[x][y] = arr[pos]
			x--
			idx++
		}
		for y > layer {
			pos := idx - d
			if pos < 0 {
				pos += len(arr)
			}
			grid[x][y] = arr[pos]
			y--
			idx++
		}
		layer++
	}
	return grid
}

func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	fmt.Println(rotateGrid(mat, 1))
}
