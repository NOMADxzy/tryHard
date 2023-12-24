package main

import "fmt"

func prisonAfterNDays(cells []int, n int) []int {
	state := 0
	mask := 1
	masks := make([]int, 8)
	for i := 0; i < 8; i++ {
		masks[i] = mask
		state += cells[i] * mask
		mask *= 2
	}
	T := 0
	var hist []int
	hist = append(hist, state)
	start := 0
	stateMap := map[int]int{state: 0}
	for i := 1; i < 100; i++ {
		newState := 0
		for j := 1; j < 7; j++ {
			if state&masks[j-1] == 0 && state&masks[j+1] == 0 || state&masks[j-1] > 0 && state&masks[j+1] > 0 {
				newState += masks[j]
			}
		}
		if v, ok := stateMap[newState]; ok {
			start = v
			T = i - v
			break
		}
		stateMap[newState] = i
		hist = append(hist, newState)
		state = newState
	}
	endState := 0
	if n < start {
		endState = hist[n]
	} else {
		idx := n % T
		for idx < start {
			idx += T
		}
		endState = hist[idx]
	}
	ans := make([]int, 8)
	for i := 0; i < 8; i++ {
		if endState&masks[i] > 0 {
			ans[i] = 1
		}
	}
	return ans
}

func main() {
	fmt.Println(prisonAfterNDays([]int{1, 1, 0, 0, 0, 0, 1, 1}, 7))
}
