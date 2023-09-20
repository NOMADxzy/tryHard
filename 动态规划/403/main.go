package main

import (
	"fmt"
	"strconv"
)

func move(river map[int]bool, pos int, k int, maxLen int, stateMap map[string]bool) bool {
	if pos == maxLen {
		return true
	}
	keystring := strconv.Itoa(pos) + "_" + strconv.Itoa(k)
	if stateMap[keystring] {
		return false
	}
	for i := k - 1; i < k+2; i++ {
		if i > 0 && pos+i <= maxLen && river[pos+i] && move(river, pos+i, i, maxLen, stateMap) {
			return true
		}
	}
	stateMap[keystring] = true
	return false
}

func canCross(stones []int) bool {
	//river := make([]bool, stones[len(stones)-1]+1)
	river := map[int]bool{}
	stateMap := map[string]bool{}
	for i := 0; i < len(stones); i++ {
		river[stones[i]] = true
	}
	if stones[1]-stones[0] != 1 {
		return false
	}
	return move(river, 1, 1, stones[len(stones)-1], stateMap)
}

func main() {
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
}
