package main

import "fmt"

func nextPermutation(nums []int) {
	m := len(nums)

	i := m - 2
	for i = m - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		t := m - 1
		for s := i + 1; s < m; s++ {
			if nums[s] <= nums[i] {
				t = s - 1
				break
			}
		}
		nums[i], nums[t] = nums[t], nums[i]
	}

	for j := i + 1; j < i+1+(m-i)/2; j++ {
		nums[j], nums[i+m-j] = nums[i+m-j], nums[j]
	}
}

func main() {
	arr := []int{5, 4, 7, 5, 3, 2}
	nextPermutation(arr)
	fmt.Println(arr)
}
