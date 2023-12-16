package main

import "fmt"

func magicalString(n int) int {
	seq := []int{1}
	pos := 0
	cnt1 := 0
	for ; pos < n; pos++ {
		if seq[pos] == 1 {
			cnt1++
		} else {
			seq = append(seq, seq[len(seq)-1])
		}
		seq = append(seq, 3-seq[len(seq)-1])
	}
	return cnt1
}

func main() {
	fmt.Println(magicalString(6))
}
