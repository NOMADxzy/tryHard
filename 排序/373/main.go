package main

import "fmt"

func buildHeap(arr [][]int) {
	l := len(arr)

	for i := (l - 2) / 2; i >= 0; i-- {
		check(arr, i, l)
	}
}

func sum(pair []int) int {
	return pair[0] + pair[1]
}

func check(arr [][]int, i int, size int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < size && sum(arr[left]) > sum(arr[largest]) {
		largest = left
	}
	if right < size && sum(arr[right]) > sum(arr[largest]) {
		largest = right
	}
	if i != largest {
		arr[i], arr[largest] = arr[largest], arr[i]
		check(arr, largest, size)
	}
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	var hp [][]int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if len(hp) < k {
				hp = append(hp, []int{nums1[i], nums2[j]})
				if len(hp) == k {
					buildHeap(hp)
				}
			} else {
				if sum(hp[0]) > nums1[i]+nums2[j] {
					hp[0][0], hp[0][1] = nums1[i], nums2[j]
					check(hp, 0, k)
				} else {
					break // 后面的肯定 > 当前元素 > 堆顶
				}
			}
		}
	}

	return hp
}

func main() {
	fmt.Println(kSmallestPairs([]int{-10, -4, 0, 0, 6}, []int{3, 5, 6, 7, 8, 100}, 10))
}
