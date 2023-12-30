package main

import "fmt"

func pathInZigZagTree(label int) []int {
	var layers [][]int
	lastEnd := 0
	lay := 0
	for i := 1; ; i = 2*i + 1 {
		layers = append(layers, []int{lastEnd + 1, i})
		lastEnd = i
		lay++
		if i >= label {
			break
		}
	}
	ans := make([]int, lay)
	pos := label
	for lay > 0 {
		ans[lay-1] = pos
		if lay == 1 {
			break
		}
		if lay%2 == 0 {
			pos = layers[lay-1][1] + layers[lay-1][0] - pos
			pos = pos / 2
			lay--
		} else {
			pos = pos / 2
			lay--
			pos = layers[lay-1][1] + layers[lay-1][0] - pos
		}
	}
	return ans
}

func main() {
	fmt.Println(pathInZigZagTree(14))
}
