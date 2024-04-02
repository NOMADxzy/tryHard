package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(words []string, things [][]string) []string {
	exists := map[string]bool{}
	for _, word := range words {
		exists[word] = true
	}
	m := len(things)
	ans := make([]string, m)
	cnts := map[string]int{}
	for i := 0; i < m; i++ {
		curThing := things[i]
		cnt := 0
		for j := 1; j < len(curThing); j++ {
			if exists[curThing[j]] {
				cnt++
			}
		}
		ans[i] = curThing[0]
		cnts[curThing[0]] = cnt
	}
	sort.SliceStable(ans, func(i, j int) bool {
		return cnts[ans[i]] > cnts[ans[j]]
	})
	return ans
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	//var thing string
	var m, _ int
	input.Scan()
	tmps := strings.Split(input.Text(), " ")
	m, _ = strconv.Atoi(tmps[0])
	//n, _ = strconv.Atoi(tmps[1])
	//things := make([][]string, m)

	input.Scan()
	words := strings.Split(input.Text(), " ")
	exists := map[string]bool{}
	for _, word := range words {
		exists[word] = true
	}

	ans := make([]string, m)
	cnts := map[string]int{}

	for i := 0; i < m; i++ {
		//_, _ = fmt.Scan(&thing, &tmp)
		input.Scan()
		tmps = strings.Split(input.Text(), " ")
		input.Scan()
		strs := strings.Split(input.Text(), " ")
		cnt := 0
		for _, str := range strs {
			if exists[str] {
				cnt++
			}
		}
		ans[i] = tmps[0]
		cnts[tmps[0]] = cnt
	}
	sort.SliceStable(ans, func(i, j int) bool {
		return cnts[ans[i]] > cnts[ans[j]]
	})

	//ans := solve(words, things)
	for _, an := range ans {
		fmt.Println(an)
	}
}
