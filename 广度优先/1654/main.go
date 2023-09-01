package main

import (
	"container/list"
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Pos struct {
	p         int
	onlyRight bool
	step      int
}

func minimumJumps1(forbidden []int, a int, b int, x int) int {
	upperBound := x
	if a > b {
		upperBound = x + b
	} else if a < b {
		f := 0
		for i := 0; i < len(forbidden); i++ {
			if forbidden[i] > f {
				f = forbidden[i]
			}
		}
		upperBound = max(f+a+b, x)
	}

	forbiddenMap := map[int]bool{}
	visited := map[int]bool{}
	visited[0] = true
	for i := 0; i < len(forbidden); i++ {
		forbiddenMap[forbidden[i]] = true
	}

	var queue list.List
	queue.PushBack(&Pos{0, false, 0})
	for {
		if queue.Len() == 0 {
			break
		}
		P := queue.Remove(queue.Front()).(*Pos)
		step := P.step
		fmt.Println(step)

		if P.p == x {
			return P.step
		} else {
			direction := 1
			if P.onlyRight {
				direction = -1
			}

			if !forbiddenMap[P.p+a] && P.p+a >= 0 && P.p+a <= upperBound && !visited[direction*(P.p+a)] {
				visited[direction*(P.p+a)] = true
				queue.PushBack(&Pos{P.p + a, false, P.step + 1})
			}
			if !P.onlyRight {
				if !forbiddenMap[P.p-b] && P.p-b >= 0 && P.p-b <= upperBound && !visited[direction*(P.p-b)] {
					visited[direction*(P.p-b)] = true
					queue.PushBack(&Pos{P.p - b, true, P.step + 1})
				}
			}
		}
	}

	return -1
}

func minimumJumps(forbidden []int, a int, b int, x int) int {
	lower := 0
	maxVal := 0
	for _, val := range forbidden {
		maxVal = max(maxVal, val)
	}
	upper := max(maxVal+a, x) + b
	q := [][3]int{[3]int{0, 1, 0}}
	visited := make(map[int]bool)
	forbiddenSet := make(map[int]bool)
	visited[0] = true

	for _, position := range forbidden {
		forbiddenSet[position] = true
	}
	for len(q) > 0 {
		position, direction, step := q[0][0], q[0][1], q[0][2]
		q = q[1:]
		if position == x {
			return step
		}
		nextPosition, nextDirection := position+a, 1
		_, ok1 := visited[nextPosition*nextDirection]
		_, ok2 := forbiddenSet[nextPosition]
		if lower <= nextPosition && nextPosition <= upper && !ok1 && !ok2 {
			visited[nextPosition*nextDirection] = true
			q = append(q, [3]int{nextPosition, nextDirection, step + 1})
		}
		if direction == 1 {
			nextPosition, nextDirection := position-b, -1
			_, ok1 := visited[nextPosition*nextDirection]
			_, ok2 := forbiddenSet[nextPosition]
			if lower <= nextPosition && nextPosition <= upper && !ok1 && !ok2 {
				visited[nextPosition*nextDirection] = true
				q = append(q, [3]int{nextPosition, nextDirection, step + 1})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(minimumJumps([]int{549, 693, 456, 1814, 1609}, 748, 889, 545))
}
