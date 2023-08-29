package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func recur(citations []int, l int, r int, pos int, H *int) {
	if pos > r || l > pos {
		return
	}
	a := citations[pos]
	b := len(citations) - pos
	if *H < min(a, b) {
		*H = min(a, b)
	}

	if a < b {
		recur(citations, pos+1, r, (r+pos+1)/2, H)
	} else if a > b {
		recur(citations, l, pos-1, (pos+l-1)/2, H)
	} else {
		return
	}
}

func hIndex(citations []int) int {
	var H int
	recur(citations, 0, len(citations)-1, len(citations)/2, &H)
	return H
}

func main() {
	fmt.Println(hIndex([]int{0, 0, 1}))
}
