package main

import "fmt"

func removeKdigits(num string, k int) string {
	m := len(num)
	if k == m {
		return "0"
	}
	var decrementStack [][]int
	nextSmaller := make([]int, len(num))
	for i := 0; i < len(num); i++ {
		c := num[i]
		x := int(c - '0')
		for len(decrementStack) > 0 && decrementStack[len(decrementStack)-1][0] > x {
			e := decrementStack[len(decrementStack)-1]
			nextSmaller[e[1]] = i
			decrementStack = decrementStack[:len(decrementStack)-1]
		}
		decrementStack = append(decrementStack, []int{x, i})
	}
	for len(decrementStack) > 0 {
		e := decrementStack[len(decrementStack)-1]
		nextSmaller[e[1]] = -1
		decrementStack = decrementStack[:len(decrementStack)-1]
	}
	ansLen := len(num) - k
	ans := make([]byte, ansLen)
	p := 0
	for i := 0; i < m && p < ansLen; {
		if m-i == ansLen-p {
			for ; i < m; i++ {
				ans[p] = num[i]
				p++
			}
		}
		if nextSmaller[i] < 0 {
			ans[p] = num[i]
			p++
			i++
		} else if m-nextSmaller[i] > ansLen-p {
			i = nextSmaller[i]
		} else if m-nextSmaller[i] == ansLen-p {
			for j := nextSmaller[i]; j < m; j++ {
				ans[p] = num[j]
				p++
			}
			break
		} else {
			ans[p] = num[i]
			p++
			i++
		}
	}
	not0idx := 0
	for ans[not0idx] == '0' {
		not0idx++
		if not0idx == ansLen {
			return "0"
		}
	}
	return string(ans[not0idx:])
}

func main() {
	fmt.Println(removeKdigits("1432219", 3))
}
