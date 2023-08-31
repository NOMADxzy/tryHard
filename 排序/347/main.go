package main

import "fmt"

func BuildHeap(nums []int, length int, kmap map[int]int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		check(nums, i, length, kmap)
	}
}

func check(nums []int, i int, length int, kmap map[int]int) {
	left := 2*i + 1
	right := 2*i + 2
	smallest := i
	if left < length && kmap[nums[left]] < kmap[nums[smallest]] {
		smallest = left
	}
	if right < length && kmap[nums[right]] < kmap[nums[smallest]] {
		smallest = right
	}
	if smallest != i {
		swap(nums, smallest, i)
		check(nums, smallest, length, kmap)
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func topKFrequent(nums []int, k int) []int {
	Kmap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		Kmap[nums[i]] += 1
	}
	var topK []int

	for key, val := range Kmap {
		if len(topK) < k {
			topK = append(topK, key)
			if len(topK) == k {
				BuildHeap(topK, k, Kmap)
			}
		} else {
			if val > Kmap[topK[0]] {
				topK[0] = key
				check(topK, 0, k, Kmap)
			}
		}
	}
	return topK
}

func main() {
	fmt.Println(topKFrequent([]int{-1, -1}, 1))
}
