package main

import (
	"fmt"
	"sort"
)

func better(arr [][]int, i, j int, Bigger bool) bool {
	if Bigger && arr[i][0] > arr[j][0] || !Bigger && arr[i][0] < arr[j][0] {
		return true
	}
	return false
}

func upCheck(arr [][]int, i int, Bigger bool) {
	f := (i - 1) / 2
	if f >= 0 && better(arr, i, f, Bigger) {
		arr[f], arr[i] = arr[i], arr[f]
		upCheck(arr, f, Bigger)
	}
}

func downCheck(arr [][]int, i int, Bigger bool) {
	left := 2*i + 1
	right := 2*i + 2
	best := i
	if left < len(arr) && better(arr, left, best, Bigger) {
		best = left
	}
	if right < len(arr) && better(arr, right, best, Bigger) {
		best = right
	}
	if best != i {
		arr[best], arr[i] = arr[i], arr[best]
		downCheck(arr, best, Bigger)
	}
}

func medianSlidingWindow(nums []int, k int) []float64 {
	ans := make([]float64, len(nums)-k+1)
	if k == 1 {
		for i := 0; i < len(nums); i++ {
			ans[i] = float64(nums[i])
		}
		return ans
	}
	var smallHeap, largeHeap [][]int
	leftSize, rightSize := k/2, k/2
	if k%2 == 1 {
		leftSize++
	}
	win := nums[:k]
	winIdx := make([]int, len(win))
	for i := 0; i < k; i++ {
		winIdx[i] = i
	}
	sort.Slice(winIdx, func(i, j int) bool {
		return win[winIdx[i]] < win[winIdx[j]]
	})

	makeBalance := func(lcnt, rcnt, lb int, belongToleft []bool) {
		if lcnt < leftSize {
			for smallHeap[0][1] < lb {
				smallHeap[0] = smallHeap[len(smallHeap)-1]
				smallHeap = smallHeap[:len(smallHeap)-1]
				downCheck(smallHeap, 0, false)
			}
			e := smallHeap[0]
			smallHeap[0] = smallHeap[len(smallHeap)-1]
			smallHeap = smallHeap[:len(smallHeap)-1]
			downCheck(smallHeap, 0, false)

			largeHeap = append(largeHeap, e)
			belongToleft[e[1]] = true
			upCheck(largeHeap, len(largeHeap)-1, true)
		} else if rcnt < rightSize {
			for largeHeap[0][1] < lb {
				largeHeap[0] = largeHeap[len(largeHeap)-1]
				largeHeap = largeHeap[:len(largeHeap)-1]
				downCheck(largeHeap, 0, true)
			}
			e := largeHeap[0]
			largeHeap[0] = largeHeap[len(largeHeap)-1]
			largeHeap = largeHeap[:len(largeHeap)-1]
			downCheck(largeHeap, 0, true)

			smallHeap = append(smallHeap, e)
			belongToleft[e[1]] = false
			upCheck(smallHeap, len(smallHeap)-1, false)
		}
		for largeHeap[0][1] < lb {
			largeHeap[0] = largeHeap[len(largeHeap)-1]
			largeHeap = largeHeap[:len(largeHeap)-1]
			downCheck(largeHeap, 0, true)
		}
		for smallHeap[0][1] < lb {
			smallHeap[0] = smallHeap[len(smallHeap)-1]
			smallHeap = smallHeap[:len(smallHeap)-1]
			downCheck(smallHeap, 0, false)
		}
	}

	belongToLeft := make([]bool, len(nums))
	for i := 0; i < k; i++ {
		if i < leftSize {
			largeHeap = append(largeHeap, []int{nums[winIdx[i]], winIdx[i]})
			upCheck(largeHeap, len(largeHeap)-1, true)
			belongToLeft[winIdx[i]] = true
		} else {
			smallHeap = append(smallHeap, []int{nums[winIdx[i]], winIdx[i]})
			upCheck(smallHeap, len(smallHeap)-1, false)
		}
	}
	lcnt, rcnt := leftSize, rightSize
	for lb := 0; lb <= len(nums)-k; lb++ {
		if k%2 == 0 {
			ans[lb] = float64(largeHeap[0][0]+smallHeap[0][0]) / 2
		} else {
			ans[lb] = float64(largeHeap[0][0])
		}

		if lb < len(nums)-k {
			addVal := nums[lb+k]
			if belongToLeft[lb] {
				lcnt--
			} else {
				rcnt--
			}
			if addVal > largeHeap[0][0] {
				smallHeap = append(smallHeap, []int{nums[lb+k], lb + k})
				upCheck(smallHeap, len(smallHeap)-1, false)
				rcnt++
			} else {
				largeHeap = append(largeHeap, []int{nums[lb+k], lb + k})
				upCheck(largeHeap, len(largeHeap)-1, true)
				belongToLeft[lb+k] = true
				lcnt++
			}
			makeBalance(lcnt, rcnt, lb+1, belongToLeft)
			lcnt, rcnt = leftSize, rightSize
		}
	}
	return ans
}

func main() {
	fmt.Println(medianSlidingWindow([]int{7, 0, 3, 9, 9, 9, 1, 7, 2, 3}, 6))
}
