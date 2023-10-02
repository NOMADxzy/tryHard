package main

import "fmt"

func find(k int, size int) byte {
	if size == 1 {
		return '0'
	}
	mid := size/2 + 1
	if k == mid {
		return '1'
	} else if k < mid {
		return find(k, size/2)
	} else {
		if find(1+size-k, size/2) == '1' {
			return '0'
		} else {
			return '1'
		}
	}

}

func findKthBit(n int, k int) byte {
	initLength := 1
	for i := 1; i < n; i++ {
		initLength = initLength*2 + 1
	}
	return find(k, initLength)
}

func main() {
	fmt.Println(findKthBit(4, 11))
}
