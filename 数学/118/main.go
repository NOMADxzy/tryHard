package main

import "fmt"

func generate(numRows int) [][]int {
	var triangle [][]int

	triangle = append(triangle, []int{1})

	for i := 1; i < numRows; i++ {
		layer := make([]int, i+1)
		for j := 0; j < (i+2)/2; j++ {
			pre := 0
			if j > 0 {
				pre = triangle[len(triangle)-1][j-1]
			}
			sum := pre + triangle[len(triangle)-1][j]
			layer[j] = sum
			layer[i-j] = sum
		}
		triangle = append(triangle, layer)
	}
	return triangle
}

func main() {
	fmt.Println(generate(5))
}
