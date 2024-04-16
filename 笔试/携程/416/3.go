package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution3(arr1, arr2, arr3 []int) int {
	m := len(arr1)

	cnts := map[int]int{}
	for _, val := range arr3 {
		cnts[val]++
	}

	ans := 0
	for i := 0; i < m; i++ {
		target := arr2[i] + arr1[i]
		if cnts[target] > 0 {
			cnts[target]--
			ans++
		}
	}
	return ans
}

func main() {
	var n int
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	n, _ = strconv.Atoi(input.Text())
	arr1, arr2, arr3 := make([]int, n), make([]int, n), make([]int, n)
	input.Scan()
	tmps := strings.Split(input.Text(), " ")
	for i := 0; i < n; i++ {
		val, _ := strconv.Atoi(tmps[i])
		arr1[i] = val
	}
	input.Scan()
	tmps = strings.Split(input.Text(), " ")
	for i := 0; i < n; i++ {
		val, _ := strconv.Atoi(tmps[i])
		arr2[i] = val
	}
	input.Scan()
	tmps = strings.Split(input.Text(), " ")
	for i := 0; i < n; i++ {
		val, _ := strconv.Atoi(tmps[i])
		arr3[i] = val
	}

	fmt.Println(solution3(arr1, arr2, arr3))
}
