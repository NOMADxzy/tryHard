package main

import "fmt"

func minimumCost(s string) int64 {
	m := len(s)
	if m == 1 {
		return 0
	}
	dpLeft := make([][2]int64, m)
	dpRight := make([][2]int64, m)
	if s[0] == '0' {
		dpLeft[0][1] = 1
	} else {
		dpLeft[0][0] = 1
	}
	if s[m-1] == '0' {
		dpRight[m-1][1] = 1
	} else {
		dpRight[m-1][0] = 1
	}
	for i := 1; i < m; i++ {
		j := m - 1 - i
		if s[i] == '0' {
			dpLeft[i][0] = dpLeft[i-1][0]
			dpLeft[i][1] = dpLeft[i-1][0] + int64(i+1)
		} else {
			dpLeft[i][1] = dpLeft[i-1][1]
			dpLeft[i][0] = dpLeft[i-1][1] + int64(i+1)
		}
		if s[j] == '0' {
			dpRight[j][0] = dpRight[j+1][0]
			dpRight[j][1] = dpRight[j+1][0] + int64(m-j)
		} else {
			dpRight[j][1] = dpRight[j+1][1]
			dpRight[j][0] = dpRight[j+1][1] + int64(m-j)
		}
	}
	ans := int64(1) << 61
	for i := -1; i < m; i++ {
		costLeft0, costRight0, costLeft1, costRight1 := int64(0), int64(0), int64(0), int64(0)
		if i >= 0 {
			costLeft0 = dpLeft[i][0]
			costLeft1 = dpLeft[i][1]
		}
		if i+1 < m {
			costRight0 = dpRight[i+1][0]
			costRight1 = dpRight[i+1][1]
		}
		cur := min(costLeft0+costRight0, costLeft1+costRight1)
		ans = min(ans, cur)
	}
	return ans
}

func main() {
	fmt.Println(minimumCost("010101"))
}
