package main

import (
	"fmt"
	"sort"
)

func maximumSwap(num int) int {
	var nArr []int
	for i := num; i > 0; i /= 10 {
		nArr = append(nArr, i%10)
	}
	for i := 0; i < len(nArr)/2; i++ {
		nArr[i], nArr[len(nArr)-1-i] = nArr[len(nArr)-1-i], nArr[i]
	}
	sortedArr := make([]int, len(nArr))
	copy(sortedArr, nArr)
	sort.Slice(sortedArr, func(i, j int) bool {
		return sortedArr[i] > sortedArr[j]
	})
	for i := 0; i < len(nArr); i++ {
		if sortedArr[i] > nArr[i] {
			tmp := nArr[i]
			nArr[i] = sortedArr[i]
			for j := len(nArr) - 1; j > i; j-- {
				if nArr[j] == nArr[i] {
					nArr[j] = tmp
					break
				}
			}
			break
		}
	}
	sum := 0
	for i := 0; i < len(nArr); i++ {
		sum = sum*10 + nArr[i]
	}
	return sum
}

func main() {
	fmt.Println(maximumSwap(72366))
}
