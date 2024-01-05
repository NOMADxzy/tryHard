package main

import "fmt"

func maxDiff(num int) int {
	w := 0
	maxV, minV := num, num
	for i := num; i > 0; i /= 10 {
		w++
	}
	masks := make([]int, w)
	mask := 1
	arr := make([]int, w)
	for i := 0; i < w; i++ {
		arr[w-1-i] = num % 10
		masks[w-1-i] = mask
		mask *= 10
		num /= 10
	}
	for i := 0; i < w; i++ {
		if arr[i] < 9 {
			for j := i; j < w; j++ {
				if arr[j] == arr[i] {
					maxV += masks[j] * (9 - arr[i])
				}
			}
			break
		}
	}
	not1 := false
	for i := 0; i < w; i++ {
		if i == 0 {
			if arr[i] > 1 {
				for j := i; j < w; j++ {
					if arr[j] == arr[i] {
						minV -= masks[j] * (arr[i] - 1)
					}
				}
				break
			} else {
				not1 = true
			}
		} else if i > 0 && arr[i] > 0 {
			if not1 && arr[i] == 1 {
				continue
			}
			for j := i; j < w; j++ {
				if arr[j] == arr[i] {
					minV -= masks[j] * (arr[i] - 0)
				}
			}
			break
		}
	}
	return maxV - minV
}

func main() {
	fmt.Println(maxDiff(111))
}
