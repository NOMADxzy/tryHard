package main

import "fmt"

func findMinFibonacciNumbers(k int) int {
	arr := []int{1}
	a, b := 1, 1
	for b < k {
		arr = append(arr, b)
		tmp := b
		b = a + b
		a = tmp
	}
	arr = append(arr, b)
	ans := 0
	remain := k
	R := len(arr) - 1
	for remain > 0 {
		ans++
		left, right := 0, R
		if arr[right] == remain {
			break
		}
		for left < right {
			mid := (left + right) / 2
			if arr[mid] <= remain {
				left = mid + 1
			} else {
				right = mid
			}
		}
		left--
		remain -= arr[left]
		R = left
	}
	return ans
}

func main() {
	fmt.Println(findMinFibonacciNumbers(1))
}
