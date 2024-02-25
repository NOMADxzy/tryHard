package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findHighAccessEmployees(access_times [][]string) []string {
	times := map[string][]string{}
	for _, pair := range access_times {
		times[pair[0]] = append(times[pair[0]], pair[1])
	}
	getVal := func(a string) int {
		var a1, a2 int
		a1, _ = strconv.Atoi(a[:2])
		a2, _ = strconv.Atoi(a[2:])
		return a1*60 + a2
	}
	var ans []string
	for man, ts := range times {
		sort.Slice(ts, func(i, j int) bool {
			return ts[i] < ts[j]
		})
		var stack []int
		for i := 0; i < len(ts); i++ {
			val := getVal(ts[i])
			if len(stack) == 0 || val-stack[0] < 60 {
				stack = append(stack, val)
			} else {
				for j := 0; j < len(stack) && val-stack[j] >= 60; j++ {
					stack = stack[1:]
				}
				stack = append(stack, val)
			}
			if len(stack) == 3 {
				ans = append(ans, man)
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findHighAccessEmployees([][]string{{"a", "0549"}, {"b", "0457"}, {"a", "0532"}, {"a", "0621"}, {"b", "0540"}}))
}
