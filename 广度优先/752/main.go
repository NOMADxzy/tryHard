package main

import (
	"fmt"
)

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	deadStates := map[string]bool{}
	for _, deadend := range deadends {
		deadStates[deadend] = true
	}
	if deadStates["0000"] {
		return -1
	}
	var left, right [][]byte
	leftVisited, rightVisited := map[string]bool{}, map[string]bool{}
	t := make([]byte, 4)
	for i := 0; i < 4; i++ {
		t[i] = target[i]
	}
	right = append(right, t)
	rightVisited[string(t)] = true
	t = make([]byte, 4)
	for i := 0; i < 4; i++ {
		t[i] = '0'
	}
	leftVisited[string(t)] = true
	left = append(left, t)
	step := 0
	for len(left) > 0 && len(right) > 0 {
		step++
		if len(left) <= len(right) {
			var newLeft [][]byte
			for len(left) > 0 {
				e := left[0]
				left = left[1:]
				for i := 0; i < 4; i++ {
					origin := e[i]
					e[i] = origin + 1
					if e[i] > '9' {
						e[i] = '0'
					}
					se := string(e)
					if !leftVisited[se] && !deadStates[se] {
						if rightVisited[se] {
							return step
						} else {
							leftVisited[se] = true
							tmp := make([]byte, 4)
							copy(tmp, e)
							newLeft = append(newLeft, tmp)
						}
					}
					e[i] = origin - 1
					if e[i] < '0' {
						e[i] = '9'
					}
					se = string(e)
					if !leftVisited[se] && !deadStates[se] {
						if rightVisited[se] {
							return step
						} else {
							leftVisited[se] = true
							tmp := make([]byte, 4)
							copy(tmp, e)
							newLeft = append(newLeft, tmp)
						}
					}
					e[i] = origin
				}
			}
			left = newLeft
		} else {
			var newRight [][]byte
			for len(right) > 0 {
				e := right[0]
				right = right[1:]
				for i := 0; i < 4; i++ {
					origin := e[i]
					e[i] = origin + 1
					if e[i] > '9' {
						e[i] = '0'
					}
					se := string(e)
					if !rightVisited[se] && !deadStates[se] {
						if leftVisited[se] {
							return step
						} else {
							rightVisited[se] = true
							tmp := make([]byte, 4)
							copy(tmp, e)
							newRight = append(newRight, tmp)
						}
					}
					e[i] = origin - 1
					if e[i] < '0' {
						e[i] = '9'
					}
					se = string(e)
					if !rightVisited[se] && !deadStates[se] {
						if leftVisited[se] {
							return step
						} else {
							rightVisited[se] = true
							tmp := make([]byte, 4)
							copy(tmp, e)
							newRight = append(newRight, tmp)
						}
					}
					e[i] = origin
				}
			}
			right = newRight
		}
	}
	return -1
}

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
