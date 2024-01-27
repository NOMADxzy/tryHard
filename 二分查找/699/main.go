package main

import "fmt"

func fallingSquares(positions [][]int) []int {
	hists := [][]int{{0, 1 << 31, 0}}
	ans := make([]int, len(positions))
	preMax := 0
	for i := 0; i < len(positions); i++ {
		targetLeft, targetRight, curHeight := positions[i][0], positions[i][0]+positions[i][1], positions[i][1]
		lb, rb := 0, 0
		left, right := 0, len(hists)-1
		for left < right {
			mid := (left + right) / 2
			if hists[mid][1] <= targetLeft {
				left = mid + 1
			} else {
				right = mid
			}
		}
		lb = right
		left, right = 0, len(hists)-1
		for left < right {
			mid := (left + right + 1) / 2
			if hists[mid][0] >= targetRight {
				right = mid - 1
			} else {
				left = mid
			}
		}
		rb = left
		maxVal := 0
		for j := lb; j <= rb; j++ {
			maxVal = max(maxVal, hists[j][2])
		}
		newLen := lb + 1 + 1 + len(hists) - rb
		newHist := make([][]int, newLen)
		copy(newHist, hists[:lb])
		pos := lb
		if hists[lb][0] < targetLeft {
			newHist[pos] = []int{hists[lb][0], targetLeft, hists[lb][2]}
			pos++
		} else {
			newLen--
		}
		newHist[pos] = []int{targetLeft, targetRight, maxVal + curHeight}
		pos++
		if hists[rb][1] > targetRight {
			newHist[pos] = []int{targetRight, hists[rb][1], hists[rb][2]}
			pos++
		} else {
			newLen--
		}
		copy(newHist[pos:], hists[rb+1:])
		ans[i] = max(preMax, maxVal+curHeight)
		if ans[i] > preMax {
			preMax = ans[i]
		}
		hists = newHist[:newLen]
	}
	return ans
}

func main() {
	fmt.Println(fallingSquares([][]int{{1, 5}, {2, 2}, {7, 5}}))
}
