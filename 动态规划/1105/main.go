package main

import "fmt"

func minHeightShelves(books [][]int, shelfWidth int) int {
	m := len(books)

	dp := make([]int, m)
	dp[0] = books[0][1]

	for i := 1; i < m; i++ {
		book := books[i]
		accWid := book[0]
		best := dp[i-1] + book[1]
		maxHeight := book[1]
		for j := i - 1; j >= 0; j-- {
			accWid += books[j][0]
			if books[j][1] > maxHeight {
				maxHeight = books[j][1]
			}
			if accWid <= shelfWidth {
				h := maxHeight
				if j > 0 {
					h += dp[j-1]
				}
				if h < best {
					best = h
				}
			} else {
				break
			}
		}
		dp[i] = best
	}
	return dp[m-1]
}

func main() {
	fmt.Println(minHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4))
}
