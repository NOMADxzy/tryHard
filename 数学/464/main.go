package main

import (
	"fmt"
)

func definitelyWin(choosable []int, marked int, desire int, acc int, visited map[int]int) bool {
	if visited[marked] == 1 {
		return false
	} else if visited[marked] == 2 {
		return true
	}
	mask := 1
	for i := 0; i < len(choosable); i++ {
		if marked&mask == 0 {
			if acc+choosable[i] >= desire {
				visited[marked] = 2
				return true
			} else {

				if !definitelyWin(choosable, marked+mask, desire, acc+choosable[i], visited) {
					visited[marked] = 2
					return true
				}
			}
		}
		mask *= 2
	}
	visited[marked] = 1
	return false
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	marked := 0
	choosable := make([]int, maxChoosableInteger)
	maxTotal := 0
	for i := 0; i < maxChoosableInteger; i++ {
		choosable[i] = i + 1
		maxTotal += i + 1
	}
	if maxTotal < desiredTotal {
		return false
	}

	visited := map[int]int{}

	return definitelyWin(choosable, marked, desiredTotal, 0, visited)
}

func main() {
	fmt.Println(canIWin(5, 79))
}

//1,2,3,4,5,6,7,8,
