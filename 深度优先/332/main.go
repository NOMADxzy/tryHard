package main

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	n := len(tickets)
	path := []string{"JFK"} //初始化第一站
	used := make([]bool, n)
	sort.Sort(ticketsSli(tickets))
	var backtracking func() bool //因为已经字母排序过了，所以找到的第一条就是结果，需要通过bool返回值快速退出
	backtracking = func() bool {
		if len(path) == n+1 {
			return true
		}
		for i := 0; i < n; i++ {
			if used[i] || tickets[i][0] != path[len(path)-1] { //需要找没用过的车票，以及车票起点与path最后一个节点相同的车票
				continue
			}
			used[i] = true
			path = append(path, tickets[i][1])
			if backtracking() { //找到一条结果即快速退出
				return true
			}
			path = path[:len(path)-1]
			used[i] = false
		}
		return false
	}
	backtracking()
	return path
}

type ticketsSli [][]string

func (this ticketsSli) Less(i, j int) bool {
	if this[i][0] == this[j][0] {
		return this[i][1] < this[j][1]
	}
	return this[i][0] < this[j][0]
}
func (this ticketsSli) Len() int {
	return len(this)
}
func (this ticketsSli) Swap(i, j int) {
	(this)[i], (this)[j] = (this)[j], (this)[i]
}

func main() {
	fmt.Println(findItinerary([][]string{{"EZE", "AXA"}, {"TIA", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}, {"ANU", "EZE"}, {"TIA", "ANU"}, {"AXA", "TIA"}, {"TIA", "JFK"}, {"ANU", "TIA"}, {"JFK", "TIA"}}))
}
