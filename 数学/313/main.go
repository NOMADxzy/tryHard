package main

import "fmt"

func downCheck(arr []int, i int, size int) {
	l := 2*i + 1
	r := 2*i + 2
	smallest := i
	if l < size && arr[l] < arr[smallest] {
		smallest = l
	}
	if r < size && arr[r] < arr[smallest] {
		smallest = r
	}
	if smallest != i {
		arr[smallest], arr[i] = arr[i], arr[smallest]
		downCheck(arr, smallest, size)
	}
}

func upCheck(arr []int, i int) {
	f := (i - 1) / 2
	if f >= 0 && arr[i] < arr[f] {
		arr[i], arr[f] = arr[f], arr[i]
		upCheck(arr, f)
	}
}

func nthSuperUglyNumber(n int, primes []int) int {
	var priorityQueue []int
	priorityQueue = append(priorityQueue, 1)

	max_int := 1
	for i := 0; i < 30; i++ {
		max_int = max_int<<1 + 1
	}

	var e int
	cnt := 0
	for cnt < n {
		e = priorityQueue[0]
		cnt++
		priorityQueue[0] = priorityQueue[len(priorityQueue)-1]
		priorityQueue = priorityQueue[:len(priorityQueue)-1]
		downCheck(priorityQueue, 0, len(priorityQueue))
		for _, prime := range primes {
			if e <= max_int/prime {
				priorityQueue = append(priorityQueue, e*prime)
				upCheck(priorityQueue, len(priorityQueue)-1)
				if e%prime == 0 {
					break
				}
			}
		}
	}
	return e
}

func main() {
	fmt.Println(nthSuperUglyNumber(10, []int{2, 7, 13, 19}))
}
