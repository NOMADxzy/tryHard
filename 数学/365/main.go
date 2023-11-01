package main

import (
	"fmt"
)

type State struct {
	x int
	y int
}

func move(a int, b int, aMax int, bMax int) []State {
	var states []State
	if a > 0 {
		states = append(states, State{0, b})
	}
	if b > 0 {
		states = append(states, State{a, 0})
	}
	if a < aMax {
		states = append(states, State{aMax, b})
		if b > 0 {
			if a+b <= aMax {
				states = append(states, State{a + b, 0})
			} else {
				states = append(states, State{aMax, a + b - aMax})
			}
		}
	}
	if b < bMax {
		states = append(states, State{a, bMax})
		if a > 0 {
			if a+b <= bMax {
				states = append(states, State{0, a + b})
			} else {
				states = append(states, State{a + b - bMax, bMax})
			}
		}
	}
	return states

}

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	var queue []State
	stateMap := map[State]bool{}

	queue = append(queue, State{0, 0})
	stateMap[State{0, 0}] = true

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.x == targetCapacity || state.y == targetCapacity || state.x+state.y == targetCapacity {
			return true
		}

		states := move(state.x, state.y, jug1Capacity, jug2Capacity)
		for i := 0; i < len(states); i++ {
			st := states[i]
			if !stateMap[st] {
				stateMap[st] = true
				queue = append(queue, st)
			}
		}
	}
	return false
}

func main() {
	fmt.Println(canMeasureWater(104707, 104711, 1))
}
