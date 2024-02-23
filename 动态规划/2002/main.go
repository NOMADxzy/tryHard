package main

import "fmt"

func maxProduct(s string) int {
	hist := map[int]int{}
	var valids []int
	check := func(state int) int {
		if v, ok := hist[state]; ok {
			return v
		}
		str := ""
		tmp := state
		for j := 0; tmp > 0; j++ {
			if tmp%2 == 1 {
				str += string(s[j])
			}
			tmp /= 2
		}
		for i := 0; i < len(str)/2; i++ {
			if str[i] != str[len(str)-1-i] {
				hist[state] = 0
				return 0
			}
		}
		hist[state] = len(str)
		if len(str) > 0 {
			valids = append(valids, state)
		}
		return len(str)
	}
	ans := 0

	for state := 1; state < 1<<len(s); state++ {
		check(state)
	}
	for _, state1 := range valids {
		for _, state2 := range valids {
			if state2&state1 == 0 {
				ans = max(ans, check(state1)*check(state2))
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxProduct("leetcodecom"))
}
