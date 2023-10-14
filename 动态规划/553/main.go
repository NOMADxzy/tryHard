package main

import (
	"fmt"
	"strconv"
)

func addBracks(bracks *[][]int, left, right int, isMax bool, dpMaxPos, dpMinPos [][]int) {
	if left >= right {
		return
	}
	var mid int
	if isMax {
		mid = dpMaxPos[left][right]
		addBracks(bracks, left, mid, true, dpMaxPos, dpMinPos)
		addBracks(bracks, mid+1, right, false, dpMaxPos, dpMinPos)
	} else {
		mid = dpMinPos[left][right]
		addBracks(bracks, left, mid, false, dpMaxPos, dpMinPos)
		addBracks(bracks, mid+1, right, true, dpMaxPos, dpMinPos)
	}
	if right > mid+1 {
		*bracks = append(*bracks, []int{mid + 1, right})
	}
}

func optimalDivision(nums []int) string {
	m := len(nums)
	dpMax := make([][]float64, m)
	dpMaxPos := make([][]int, m)
	for i := 0; i < m; i++ {
		dpMax[i] = make([]float64, m)
		dpMaxPos[i] = make([]int, m)
	}
	dpMin := make([][]float64, m)
	dpMinPos := make([][]int, m)
	for i := 0; i < m; i++ {
		dpMin[i] = make([]float64, m)
		dpMinPos[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		dpMax[i][i] = float64(nums[i])
		dpMin[i][i] = float64(nums[i])
	}
	for l := 2; l <= m; l++ {
		for i := 0; i < m-l+1; i++ {
			left, right := i, i+l-1
			minVal, maxVal := dpMin[left][left]/dpMax[left+1][right], dpMax[left][left]/dpMin[left+1][right]
			minPos, maxPos := left, left
			for mid := left + 1; mid < right; mid++ {
				newMinVal, newMaxVal := dpMin[left][mid]/dpMax[mid+1][right], dpMax[left][mid]/dpMin[mid+1][right]
				if newMinVal < minVal {
					minVal = newMinVal
					minPos = mid
				}
				if newMaxVal > maxVal {
					maxVal = newMaxVal
					maxPos = mid
				}
			}
			dpMax[left][right], dpMin[left][right] = maxVal, minVal
			dpMaxPos[left][right], dpMinPos[left][right] = maxPos, minPos
		}
	}

	var bracks [][]int
	addBracks(&bracks, 0, m-1, true, dpMaxPos, dpMinPos)
	LeftBrackCntOfEachPos := make([]int, m)
	RightBrackCntOfEachPos := make([]int, m)
	for _, brack := range bracks {
		LeftBrackCntOfEachPos[brack[0]]++
		RightBrackCntOfEachPos[brack[1]]++
	}
	s := ""
	for i := 0; i < LeftBrackCntOfEachPos[0]; i++ {
		s += "("
	}
	for i := 0; i < m; i++ {
		s += strconv.Itoa(nums[i])
		for j := 0; j < RightBrackCntOfEachPos[i]; j++ {
			s += ")"
		}
		if i == m-1 {
			break
		}
		s += "/"
		for j := 0; j < LeftBrackCntOfEachPos[i+1]; j++ {
			s += "("
		}
	}
	return s
}

func main() {
	fmt.Println(optimalDivision([]int{1000, 100, 10, 2}))
}
