package main

import (
	"fmt"
	"sort"
)

func countLatticePoints(circles [][]int) int {
	lines := make([][][]int, 201)
	for _, circle := range circles {
		x, y, r := circle[0], circle[1], circle[2]
		for i := y - r; i <= y+r; i++ {
			left, right := 0, r
			if i == y-r || i == r-y {
				right = left
			}
			for left < right {
				mid := (left + right) / 2
				if mid*mid+(y-i)*(y-i) < r*r {
					left = mid + 1
				} else {
					right = mid
				}
			}
			j1, j2 := x-right-1, x+right
			if right*right+(y-i)*(y-i) > r*r {
				j1++
				j2--
			}
			lines[i] = append(lines[i], []int{j1, j2})
		}
	}
	ans := 0
	for _, line := range lines {
		sort.Slice(line, func(i, j int) bool {
			return line[i][0] < line[j][0]
		})
		for i := 0; i < len(line); i++ {
			start := line[i][0]
			preRight := line[i][1]
			for i+1 < len(line) && line[i+1][0] <= preRight {
				i++
				preRight = max(preRight, line[i][1])
			}
			ans += preRight - start
		}
	}
	return ans
}

func countLatticePoints1(circles [][]int) (ans int) { // 暴力方法

	// 按半径从大到小排序，这样能更早遇到包含 (x,y) 的圆
	sort.Slice(circles, func(i, j int) bool { return circles[i][2] > circles[j][2] })

	for x := 0; x <= 200; x++ {
		for y := 0; y <= 200; y++ {
			for _, c := range circles {

				// 判断 (x,y) 是否在圆 c 内
				if (x-c[0])*(x-c[0])+(y-c[1])*(y-c[1]) <= c[2]*c[2] {
					ans++
					break
				}
			}
		}
	}
	return

}

func main() {
	fmt.Println(countLatticePoints([][]int{{8, 9, 6}, {9, 8, 4}}))
}
