package main

import "fmt"

func nthSuperUglyNumber(n int, primes []int) int {
	pointers := make([]int, len(primes))
	for i := 0; i < len(pointers); i++ {
		pointers[i] = 1
	}
	var uglyNums []int
	uglyNums = append(uglyNums, 1)
	cnt := 1

	max_int := 1
	for i := 0; i < 30; i++ {
		max_int = max_int<<1 + 1
	}

	for cnt < n {

		min_new_ugly := max_int
		min_pointer_idx := -1
		for i := 0; i < len(pointers); i++ {
			if uglyNums[pointers[i]-1] <= max_int/primes[i] && uglyNums[pointers[i]-1]*primes[i] < min_new_ugly {
				min_new_ugly = uglyNums[pointers[i]-1] * primes[i]
				min_pointer_idx = i
			}
		}
		pointers[min_pointer_idx]++
		if min_new_ugly > uglyNums[len(uglyNums)-1] {
			uglyNums = append(uglyNums, min_new_ugly)
			cnt++
		}
	}
	return uglyNums[len(uglyNums)-1]
}

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}
