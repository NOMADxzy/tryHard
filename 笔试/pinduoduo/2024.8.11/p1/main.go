package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 1 3 5 7 9
// 2 4 6 8 10
func solve(details [][]int) int {
	sort.Slice(details, func(i, j int) bool {
		return details[i][0] < details[j][0]
	})
	ans := 0
	for i := 0; i < len(details); i++ {
		detail := details[i]
		startDay, interval := detail[1], detail[2]
		if ans <= startDay {
			ans = startDay
		} else {
			minVal := ans + 1
			ans = (ans/interval)*interval + startDay
			if ans < minVal {
				ans += interval
			}
		}

	}
	return ans
}

func main() {
	var m int
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	s := in.Text()
	m, _ = strconv.Atoi(s)
	details := make([][]int, m)
	for i := 0; i < m; i++ {
		in.Scan()
		s = in.Text()
		splits := strings.Split(s, " ")
		p, _ := strconv.Atoi(splits[0])
		startDay, _ := strconv.Atoi(splits[1])
		interval, _ := strconv.Atoi(splits[2])
		details[i] = []int{p, startDay, interval}
	}
	ans := solve(details)
	fmt.Println(ans)
}
