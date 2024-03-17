package main

import "fmt"

func solve(arr []int, k int) int {
	n := len(arr)
	//cnt2, cnt5 := 0, 0
	total2, total5 := 0, 0
	pairs := make([][2]int, n)
	for i := 0; i < len(arr); i++ {
		num := arr[i]
		cnt2, cnt5 := 0, 0
		for num%2 == 0 {
			cnt2++
			num /= 2
		}
		for num%5 == 0 {
			cnt5++
			num /= 5
		}
		pairs[i][0], pairs[i][1] = cnt2, cnt5
		total2 += cnt2
		total5 += cnt5
	}
	l, r := 0, 0
	maxDel2, maxDel5 := total2-k, total5-k
	if maxDel2 < 0 || maxDel5 < 0 {
		return 0
	}
	ans := 0
	sumDel2, sumDel5 := 0, 0
	for r < n {
		sumDel2 += pairs[r][0]
		sumDel5 += pairs[r][1]
		for sumDel2 > maxDel2 || sumDel5 > maxDel5 {
			sumDel2 -= pairs[l][0]
			sumDel5 -= pairs[l][1]
			l++
		}
		ans += r - l + 1
		r++
	}
	return ans
}

func main() {
	n, k := 0, 0
	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&k)
	arr := make([]int, n)
	var tmp int
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&tmp)
		arr[i] = tmp
	}
	fmt.Println(solve(arr, k))
}
