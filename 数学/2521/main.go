package main

import "fmt"

func distinctPrimeFactors(nums []int) int {
	set := map[int]bool{}

	for _, num := range nums {
		for i := 2; i <= num; {
			if num%i == 0 {
				set[i] = true
				num /= i
			} else {
				i++
			}
		}
	}
	ans := 0
	for _, _ = range set {
		ans++
	}
	return ans
}

func main() {
	fmt.Println(distinctPrimeFactors([]int{2, 4, 3, 7, 10, 6}))
}
