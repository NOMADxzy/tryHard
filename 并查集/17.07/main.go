package main

import (
	"fmt"
	"strconv"
	"strings"
)

func trulyMostPopular(names []string, synonyms []string) []string {
	cnts := map[string]int{}
	var ans []string
	parents := map[string]string{}
	var find func(s string) string
	find = func(s string) string {
		if parents[s] != s {
			if parents[s] == "" {
				parents[s] = s
			} else {
				parents[s] = find(parents[s])
			}
		}
		return parents[s]
	}
	union := func(s1, s2 string) {
		root1, root2 := find(s1), find(s2)
		if root1 < root2 {
			root1, root2 = root2, root1
		}
		parents[root1] = root2
	}
	for _, synonym := range synonyms {
		tmp := strings.Split(synonym, ",")
		name1, name2 := tmp[0][1:], tmp[1][:len(tmp[1])-1]
		union(name1, name2)
	}
	for _, name := range names {
		tmp := strings.Split(name, "(")
		na := tmp[0]
		v, _ := strconv.Atoi(tmp[1][:len(tmp[1])-1])
		cnts[find(na)] += v
	}
	for k, v := range cnts {
		ans = append(ans, k+"("+strconv.Itoa(v)+")")
	}
	return ans
}

func main() {
	names := []string{"John(15)", "Jon(12)", "Chris(13)", "Kris(4)", "Christopher(19)"}
	synonyms := []string{"(Jon,John)", "(John,Johnny)", "(Chris,Kris)", "(Chris,Christopher)"}
	fmt.Println(trulyMostPopular(names, synonyms))
}
