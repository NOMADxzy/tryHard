package main

import "fmt"

func lexicalOrder(n int) []int {
	arr := make([]int, n)
	idx := 0
	k := 1

	for k <= n {
		arr[idx] = k
		idx++
		for k*10 <= n {
			k *= 10
			arr[idx] = k
			idx++
		}
		if k+1 > n || k%10 == 9 {
			k /= 10
			for k%10 == 9 {
				k /= 10
			}
			if k == 0 {
				break
			}
			k++
		} else {
			k += 1
		}
	}

	return arr
}

func main() {
	fmt.Println(lexicalOrder(132))
}
