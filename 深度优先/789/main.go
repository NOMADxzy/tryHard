package main

import "fmt"

func getDist(x1, y1, x2, y2 int) int {
	d1 := max(x1-x2, x2-x1)
	d2 := max(y1-y2, y2-y1)
	return d1 + d2
}

func escapeGhosts(ghosts [][]int, target []int) bool {
	dirs := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	var dfs func(x, y, step, maxStep int) bool
	dfs = func(x, y, step, maxStep int) bool {
		if x == target[0] && y == target[1] {
			return true
		} else if step+getDist(x, y, target[0], target[1]) > maxStep { //剪枝
			return false
		}
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			valid := true
			for _, ghost := range ghosts {
				if getDist(ghost[0], ghost[1], nx, ny) <= step+1 {
					valid = false
					break
				}
			}
			if valid && dfs(nx, ny, step+1, maxStep) {
				return true
			}
		}
		return false
	}
	maxStep := 1 << 31
	for _, ghost := range ghosts {
		maxStep = min(maxStep, getDist(ghost[0], ghost[1], target[0], target[1])-1)
		if ghost[0] == 0 && ghost[1] == 0 {
			return false
		}
	}
	return dfs(0, 0, 0, maxStep)
}

func main() {
	fmt.Println(escapeGhosts([][]int{{-26, 46}, {56, -87}, {-36, 9}, {95, -97}, {-5, 1}, {84, 83}, {13, -61}, {-72, 93}, {96, 66}, {-73, -27}}, []int{0, 1}))
}
