package main

import "fmt"

func reorderedPowerOf2(n int) bool {
	powArr := make([]int, 10)
	pow := 1
	for i := 0; i < len(powArr); i++ {
		powArr[i] = pow
		pow *= 10
	}

	num := 1
	nMap := map[int]bool{}
	for i := 0; i < 32; i++ {
		key := 0
		tmp := num
		for tmp > 0 {
			key += powArr[tmp%10]
			tmp /= 10
		}
		nMap[key] = true
		num = num << 1
	}

	key := 0
	for n > 0 {
		key += powArr[n%10]
		n /= 10
	}
	return nMap[key]
}

func main() {
	fmt.Println(reorderedPowerOf2(10))
}
