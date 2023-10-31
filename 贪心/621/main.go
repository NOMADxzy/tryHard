package main

import (
	"fmt"
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	//kMap := map[byte]int{}
	m := len(tasks)

	cnts := make([]int, 26)
	for i := 0; i < m; i++ {
		idx := tasks[i] - 'A'
		cnts[idx]++
	}
	TASKs := make([]byte, 26)
	for i := 0; i < 26; i++ {
		TASKs[i] = byte(i + 'A')
	}

	sort.Slice(TASKs, func(i, j int) bool {
		return cnts[TASKs[i]-'A'] > cnts[TASKs[j]-'A']
	})

	lastPos := make([]int, 26)
	for i := 0; i < 26; i++ {
		lastPos[i] = -(1 << 31)
	}

	var i int
	for i = 0; cnts[TASKs[0]-'A'] > 0; i++ {
		for c := 0; cnts[TASKs[c]-'A'] > 0 && c < 26; c++ {
			if id := TASKs[c] - 'A'; lastPos[id]+n < i {
				cnts[id]--
				lastPos[id] = i
				break
			}
		}
		sort.Slice(TASKs, func(i, j int) bool {
			return cnts[TASKs[i]-'A'] > cnts[TASKs[j]-'A']
		})
	}
	return i
}

func main() {
	fmt.Println(leastInterval([]byte{'A'}, 10))
}
