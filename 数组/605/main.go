package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	m := len(flowerbed)
	for i := 0; i < m; i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == m-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}
