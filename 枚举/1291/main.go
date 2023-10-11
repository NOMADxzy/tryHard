package main

import "fmt"

func sequentialDigits(low int, high int) []int {
	var ans []int
	startWidth := 0
	endWidth := 0
	for i := low; i > 0; i /= 10 {
		startWidth++
	}
	for i := high; i > 0; i /= 10 {
		endWidth++
	}

	for w := startWidth; w <= endWidth; w++ {
		n := 1
		d := 1
		for i := 1; i < w; i++ {
			n = n*10 + (i + 1)
			d = d*10 + 1
		}
		for i := 0; i < 10-w; i++ {
			if n >= low && n <= high {
				ans = append(ans, n)
			} else if n > high {
				break
			}
			n += d
		}
	}
	return ans
}

func main() {
	fmt.Println(sequentialDigits(1000, 13000))
}
