package main

import "fmt"

func canArrange(arr []int, k int) bool {
	mods := make([]int, k)
	for i := 0; i < len(arr); i++ {
		v := arr[i] % k
		if v < 0 {
			v += k
		}
		mods[v]++
	}
	if mods[0]%2 == 1 {
		return false
	}
	for i := 1; i <= (len(mods)-1)/2; i++ {
		if mods[i] != mods[len(mods)-i] {
			return false
		}
	}
	if len(mods)%2 == 0 && mods[len(mods)/2] == 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canArrange([]int{-1, 1, -2, 2, -3, 3, -4, 4}, 3))
}
