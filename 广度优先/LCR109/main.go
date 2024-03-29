package main

import "fmt"

func openLock(deadends []string, target string) int {
	isDead := map[string]bool{}
	for _, deadend := range deadends {
		isDead[deadend] = true
	}
	var leftQueue, rightQueue []string
	leftVisited, rightVisited := map[string]bool{"0000": true}, map[string]bool{target: true}
	leftQueue = append(leftQueue, "0000")
	rightQueue = append(rightQueue, target)
	if isDead["0000"] || isDead[target] {
		return -1
	}

	step := 0
	toCharArr := func(s string) []byte {
		res := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			res[i] = s[i]
		}
		return res
	}
loop:
	for len(leftQueue) > 0 && len(rightQueue) > 0 {
		if len(leftQueue) <= len(rightQueue) {
			var nextQueue []string
			for len(leftQueue) > 0 {
				e := leftQueue[0]
				leftQueue = leftQueue[1:]
				if rightVisited[e] {
					break loop
				}
				eArr := toCharArr(e)
				for i := 0; i < 4; i++ {
					origin := eArr[i]
					eArr[i] = byte('0' + (origin-'0'+1)%10)
					if !leftVisited[string(eArr)] && !isDead[string(eArr)] {
						nextQueue = append(nextQueue, string(eArr))
						leftVisited[string(eArr)] = true
					}
					eArr[i] = byte('0' + (origin-'0'+9)%10)
					if !leftVisited[string(eArr)] && !isDead[string(eArr)] {
						nextQueue = append(nextQueue, string(eArr))
						leftVisited[string(eArr)] = true
					}
					eArr[i] = origin
				}
			}
			leftQueue = nextQueue
		} else {
			var nextQueue []string
			for len(rightQueue) > 0 {
				e := rightQueue[0]
				rightQueue = rightQueue[1:]
				if leftVisited[e] {
					break loop
				}
				eArr := toCharArr(e)
				for i := 0; i < 4; i++ {
					origin := eArr[i]
					eArr[i] = byte('0' + (origin-'0'+1)%10)
					if !rightVisited[string(eArr)] && !isDead[string(eArr)] {
						nextQueue = append(nextQueue, string(eArr))
						rightVisited[string(eArr)] = true
					}
					eArr[i] = byte('0' + (origin-'0'+9)%10)
					if !rightVisited[string(eArr)] && !isDead[string(eArr)] {
						nextQueue = append(nextQueue, string(eArr))
						rightVisited[string(eArr)] = true
					}
					eArr[i] = origin
				}
			}
			rightQueue = nextQueue
		}
		if leftQueue == nil || rightQueue == nil {
			return -1
		}
		step++
	}
	return step
}

func main() {
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
}
