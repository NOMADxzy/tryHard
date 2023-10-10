package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	accs := make([]int, len(arr))
	accs[0] = arr[0]
	for i := 1; i < len(accs); i++ {
		accs[i] = arr[i] ^ accs[i-1]
	}

	ans := make([]int, len(queries))
	for i := 0; i < len(ans); i++ {
		if queries[i][0] > 0 {
			ans[i] = accs[queries[i][1]] ^ accs[queries[i][0]-1]
		} else {
			ans[i] = accs[queries[i][1]]
		}
	}
	return ans
}

func main() {
	fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
}
