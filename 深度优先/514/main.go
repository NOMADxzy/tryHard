package main

import "fmt"

func findRotateSteps(ring string, key string) int {
	hist := map[int]int{}
	idxs := map[byte][]int{}
	for i := 0; i < len(ring); i++ {
		idxs[ring[i]] = append(idxs[ring[i]], i)
	}
	var dfs func(pos, keyPos int) int
	dfs = func(pos, keyPos int) int {
		if keyPos == len(key) {
			return 0
		}
		hashKey := pos*101 + keyPos
		if hist[hashKey] > 0 {
			return hist[hashKey]
		}
		validPos := idxs[key[keyPos]]
		best := 1 << 31
		for _, np := range validPos {
			var delta int
			if np > pos {
				delta = min(np-pos, pos+len(ring)-np)
			} else {
				delta = min(pos-np, np+len(ring)-pos)
			}
			best = min(best, delta+1+dfs(np, keyPos+1))
		}
		hist[hashKey] = best
		return best
	}
	return dfs(0, 0)
}

func main() {
	fmt.Println(findRotateSteps("godding", "gd"))
}
