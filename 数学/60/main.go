package main

import "fmt"

func findNext(ans []byte, table []int, k, pos int, mark int) bool {
	mask := 1
	rank := 0
	for i := 1; i <= len(ans); i++ {
		if mark&mask == 0 {
			if pos == len(ans)-1 {
				ans[pos] = byte('0' + i)
				return true
			}
			if rank+table[pos+1] >= k {
				ans[pos] = byte('0' + i)
				if findNext(ans, table, k-rank, pos+1, mark+mask) {
					return true
				}
			}
			rank += table[pos+1]
		}
		mask *= 2
	}
	return false
}

func getPermutation(n int, k int) string {
	table := make([]int, n)
	table[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		table[i] = (n - i) * table[i+1]
	}
	ans := make([]byte, n)
	findNext(ans, table, k, 0, 0)
	return string(ans)
}

func main() {
	fmt.Println(getPermutation(3, 4))
}
