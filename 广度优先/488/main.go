package main

import "fmt"

type Pair struct {
	desk string
	hand [5]int
}

func findMinStep(board string, hand string) int {
	balls := []byte{'R', 'Y', 'B', 'G', 'W'}
	handMap := map[byte]int{'R': 0, 'Y': 1, 'B': 2, 'G': 3, 'W': 4}
	initHand := [5]int{}
	for i := 0; i < len(hand); i++ {
		initHand[handMap[hand[i]]]++
	}
	var queue []*Pair
	queue = append(queue, &Pair{board, initHand})
	ans := 0
	for ; len(queue) > 0; ans++ {
		var nextQ []*Pair
		for len(queue) > 0 {
			pa := queue[0]
			desk, hands := pa.desk, pa.hand
			if len(desk) == 0 {
				return ans
			}
			queue = queue[1:]
			for i := 0; i < 5; i++ {
				if hands[i] > 0 {
					ball := balls[i]
					for j := 0; j < len(desk); j++ {
						if j == len(desk)-1 || desk[j+1] != ball {
							if handMap[desk[j]] == i {
								newDesk := desk[:j+1] + string(ball) + desk[j+1:]
								trackback := true
								for trackback {
									trackback = false
									for l := 0; l < len(newDesk)-2; l++ {
										if newDesk[l] == newDesk[l+1] && newDesk[l] == newDesk[l+2] {
											right := l + 3
											for right < len(newDesk) && newDesk[right] == newDesk[l] {
												right++
											}
											newDesk = newDesk[:l] + newDesk[right:]
											trackback = true
										}
									}
								}

								newHands := [5]int{}
								for id := 0; id < 5; id++ {
									newHands[id] = hands[id]
								}
								newHands[i]--
								nextQ = append(nextQ, &Pair{newDesk, newHands})
							} else if j < len(desk)-1 && desk[j] == desk[j+1] {
								newDesk := desk[:j+1] + string(ball) + desk[j+1:]
								newHands := [5]int{}
								for id := 0; id < 5; id++ {
									newHands[id] = hands[id]
								}
								newHands[i]--
								nextQ = append(nextQ, &Pair{newDesk, newHands})
							}
						}
					}
				}
			}
		}
		queue = nextQ
	}
	return -1
}

func main() {
	fmt.Println(findMinStep("WWRRBBWW", "WRBRW"))
}
