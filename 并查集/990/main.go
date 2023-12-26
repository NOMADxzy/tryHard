package main

import "fmt"

func equationsPossible(equations []string) bool {
	parents := make([]int, 26)
	for i := 0; i < 26; i++ {
		parents[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		if parents[i] != i {
			parents[i] = find(parents[i])
		}
		return parents[i]
	}
	union := func(i1, i2 int) {
		parents[find(i1)] = find(i2)
	}

	for _, equation := range equations {
		if equation[1] == '=' {
			a, b := int(equation[0]-'a'), int(equation[3]-'a')
			union(a, b)
		}
	}
	for _, equation := range equations {
		if equation[1] == '!' {
			a, b := int(equation[0]-'a'), int(equation[3]-'a')
			if find(a) == find(b) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(equationsPossible([]string{"a==b", "b!=c", "c==a"}))
}
