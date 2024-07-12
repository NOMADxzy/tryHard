package main

import "fmt"

func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	ans := make([]int, length)
	getIdxAndVal := func(x, y int) (int, int) {
		idx, val := 0, 0
		idx = (y*w + x) / 32
		val = (y*w + x) % 32
		return idx, 31 - val
	}
	idx1, delta1 := getIdxAndVal(x1, y)
	idx2, delta2 := getIdxAndVal(x2, y)

	if idx1 == idx2 {
		val := 0
		for i := delta1; i >= delta2; i-- {
			val += 1 << i
		}
		ans[idx1] = val
	} else {
		if delta1 == 31 {
			ans[idx1] = -1
		} else {
			for i := delta1; i >= 0; i-- {
				ans[idx1] += 1 << i
			}
		}

		for i := idx1 + 1; i < idx2; i++ {
			ans[i] = -1
		}

		if delta2 == 31 {
			ans[idx2] = 1
		} else {
			for i := delta2 - 1; i >= 0; i-- {
				ans[idx2] += -1 << i
			}
		}
		ans[idx2] += -1
	}

	return ans
}

func main() {
	fmt.Println(drawLine(2, 64, 0, 1, 0))
}
