package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	id        int
	start     bool
	timestamp int
}

func findTime(logs []*Pair, pos int, ans *[]int) (int, int) {
	id := logs[pos].id
	p := pos + 1
	childT := 0
	for logs[p].id != id || logs[p].start {
		np, t := findTime(logs, p, ans)
		childT += t
		p = np + 1
	}
	(*ans)[id] += logs[p].timestamp - logs[pos].timestamp + 1 - childT
	return p, logs[p].timestamp - logs[pos].timestamp + 1
}

func exclusiveTime(n int, logs []string) []int {
	ans := make([]int, n)
	fLogs := make([]*Pair, len(logs))
	for i, log := range logs {
		logSplit := strings.Split(log, ":")
		id, _ := strconv.Atoi(logSplit[0])
		timestamp, _ := strconv.Atoi(logSplit[2])
		start := false
		if logSplit[1] == "start" {
			start = true
		}
		fLogs[i] = &Pair{id, start, timestamp}
	}
	sort.Slice(fLogs, func(i, j int) bool {
		return fLogs[i].timestamp < fLogs[j].timestamp
	})
	for i := 0; i < len(fLogs); i++ {
		i, _ = findTime(fLogs, i, &ans)
	}
	return ans
}

func main() {
	fmt.Println(exclusiveTime(1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}))
}
