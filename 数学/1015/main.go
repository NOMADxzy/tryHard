package main

import "fmt"

func smallestRepunitDivByK(k int) int {
	target := 0
	lastModVal := 0
	modValMap := map[int]bool{0: true}
	var i int
	for i = 1; ; i++ {
		target = 10*target + 1
		modVal := (((10%k)*lastModVal)%k + 1) % k
		lastModVal = modVal
		if modValMap[modVal] {
			if modVal == 0 {
				return i
			}
			break
		}
		modValMap[modVal] = true
	}
	return -1
}

func main() {
	fmt.Println(smallestRepunitDivByK(3))
}
