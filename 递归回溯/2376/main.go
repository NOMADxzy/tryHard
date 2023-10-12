package main

import "fmt"

func findNext(mark int, limit []int, isPreLimit bool, pos int, size int, history map[int]int) int {
	keytype := 1 // TODO 注意isPrelimit也是状态之一
	if !isPreLimit {
		keytype = -1
	}
	if r, ok := history[keytype*mark]; ok {
		return r
	}
	if pos == size {
		return 1
	}
	mask := 1
	ans := 0
	for i := 0; i < 10; i++ {
		if !(pos == 0 && i == 0) && mark&mask == 0 {
			if isPreLimit && i > limit[size-1-pos] {
				break
			} else if isPreLimit && i == limit[size-1-pos] {
				ans += findNext(mark+mask, limit, true, pos+1, size, history)
			} else {
				ans += findNext(mark+mask, limit, false, pos+1, size, history)
			}
		}
		mask *= 2
	}
	history[keytype*mark] = ans
	return ans
}

func countSpecialNumbers(n int) int {
	w := 0
	var limit []int
	for i := n; i > 0; i /= 10 {
		w++
		limit = append(limit, i%10)
	}
	ans := 0
	for i := 1; i < w; i++ {
		cur_ans, t := 9, 10
		for j := 1; j < i; j++ {
			cur_ans *= t - j
		}
		ans += cur_ans
	}
	ans += findNext(0, limit, true, 0, w, map[int]int{})
	return ans
}

func main() {
	fmt.Println(countSpecialNumbers(420))
}
