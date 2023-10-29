package main

import "fmt"

func minFlips(s string) int {
	m := len(s)

	dpLeft := make([][]int, m)
	dpRight := make([][]int, m)

	for i := 0; i < m; i++ {
		dpLeft[i] = make([]int, 2)
		dpRight[i] = make([]int, 2)
	}

	if s[0] == '0' {
		dpLeft[0][1] = 1
	} else {
		dpLeft[0][0] = 1
	}

	for i := 1; i < m; i++ {
		ans0, ans1 := 0, 0
		if s[i] == '0' {
			ans0 = dpLeft[i-1][1]
			ans1 = 1 + dpLeft[i-1][0]
		} else {
			ans1 = dpLeft[i-1][0]
			ans0 = 1 + dpLeft[i-1][1]
		}
		dpLeft[i][0], dpLeft[i][1] = ans0, ans1
	}

	if s[m-1] == '0' {
		dpRight[m-1][1] = 1
	} else {
		dpRight[m-1][0] = 1
	}
	for i := m - 2; i >= 0; i-- {
		ans0, ans1 := 0, 0
		if s[i] == '0' {
			ans0 = dpRight[i+1][1]
			ans1 = 1 + dpRight[i+1][0]
		} else {
			ans1 = dpRight[i+1][0]
			ans0 = 1 + dpRight[i+1][1]
		}
		dpRight[i][0], dpRight[i][1] = ans0, ans1
	}

	best := dpLeft[m-1][0]
	if dpLeft[m-1][1] < best {
		best = dpLeft[m-1][1]
	}

	for i := 0; i < m-1; i++ {
		var val1, val2 int
		firstSame := i%2 == 0
		secondSame := (m-i-2)%2 == 0
		beSame := false
		if firstSame && !secondSame || !firstSame && secondSame {
			beSame = true
		}
		if beSame {
			val1 = dpLeft[i][1] + dpRight[i+1][1]
			val2 = dpLeft[i][0] + dpRight[i+1][0]
		} else {
			val1 = dpLeft[i][0] + dpRight[i+1][1]
			val2 = dpLeft[i][1] + dpRight[i+1][0]
		}
		if val1 < best {
			best = val1
		}
		if val2 < best {
			best = val2
		}
	}
	return best
}

func main() {
	fmt.Println(minFlips("1"))
}
