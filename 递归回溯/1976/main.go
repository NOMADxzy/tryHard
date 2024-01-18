package main

import "fmt"

type Node struct {
	pos  int
	cost int64
}

func countPaths(n int, roads [][]int) int {
	var remain []*Node
	var upCheck func(i int)
	var downCheck func(i int)
	upCheck = func(i int) {
		f := (i - 1) / 2
		if f >= 0 && remain[i].cost < remain[f].cost {
			remain[i], remain[f] = remain[f], remain[i]
			upCheck(f)
		}
	}
	downCheck = func(i int) {
		left, right := 2*i+1, 2*i+2
		best := i
		if left < len(remain) && remain[left].cost < remain[best].cost {
			best = left
		}
		if right < len(remain) && remain[right].cost < remain[best].cost {
			best = right
		}
		if best != i {
			remain[best], remain[i] = remain[i], remain[best]
			downCheck(best)
		}
	}
	graph := make([][][2]int, n)
	for _, road := range roads {
		d := road[2]
		graph[road[0]] = append(graph[road[0]], [2]int{road[1], d})
		graph[road[1]] = append(graph[road[1]], [2]int{road[0], d})
	}
	hist := make([][]int64, n)
	for i := 0; i < n; i++ {
		hist[i] = make([]int64, 2)
		hist[i][0] = 1 << 62
	}
	hist[0][1] = 1
	S := make([]bool, n)
	MOD := 1000000007
	remain = append(remain, &Node{0, 0})
	for len(remain) > 0 && !S[n-1] {
		for S[remain[0].pos] {
			remain[0] = remain[len(remain)-1]
			remain = remain[:len(remain)-1]
			downCheck(0)
		}
		cur := remain[0]
		remain[0] = remain[len(remain)-1]
		remain = remain[:len(remain)-1]
		downCheck(0)

		S[cur.pos] = true
		preCost := cur.cost
		preWays := hist[cur.pos][1]
		for _, pair := range graph[cur.pos] {
			np, nCost := pair[0], int64(pair[1])
			if S[np] {
				continue
			}
			if v := preCost + nCost; v < hist[np][0] {
				hist[np][0] = v
				hist[np][1] = preWays
				remain = append(remain, &Node{np, v})
				upCheck(len(remain) - 1)
			} else if v == hist[np][0] {
				hist[np][1] = (hist[np][1] + preWays) % int64(MOD)
			}
		}
	}
	return int(hist[n-1][1])
}

func main() {
	fmt.Println(countPaths(7, [][]int{{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}}))
}
