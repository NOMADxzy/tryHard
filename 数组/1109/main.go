package main

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	deferArr := make([]int, n)
	ans := make([]int, n)

	for i := 0; i < len(bookings); i++ {
		bk := bookings[i]
		bk[0], bk[1] = bk[0]-1, bk[1]-1
		deferArr[bk[0]] += bk[2]
		if bk[1]+1 < n {
			deferArr[bk[1]+1] -= bk[2]
		}
	}

	for i := 0; i < len(ans); i++ {
		if i == 0 {
			ans[i] = deferArr[i]
		} else {
			ans[i] = ans[i-1] + deferArr[i]
		}
	}
	return ans
}

func main() {
	bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	fmt.Println(corpFlightBookings(bookings, 5))
}
