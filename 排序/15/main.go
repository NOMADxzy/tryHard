package main

import "fmt"

func partSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	val := nums[l]

	for i < j {
		for ; i < j && nums[j] >= val; j-- {
		}
		nums[i], nums[j] = nums[j], nums[i]

		for ; i < j && nums[i] <= val; i++ {
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	partSort(nums, l, i-1)
	partSort(nums, i+1, r)
}

func quickSort(nums []int) {
	partSort(nums, 0, len(nums)-1)
}

func threeSum(nums []int) [][]int {
	var arrs [][]int
	size := len(nums)

	quickSort(nums)
	var exists []int

	for i := 0; i < size-2; i++ {
		j, k := i+1, size-1
		val := nums[i]

		if i == 0 || nums[i] != nums[i-1] {
			var empty []int
			exists = empty
		}
		for j < k {
			if nums[j]+nums[k]+val > 0 {
				k--
			} else if nums[j]+nums[k]+val < 0 {
				j++
			} else {
				arr := []int{nums[i], nums[j], nums[k]}

				p := 0
				for ; p < len(exists); p++ {
					if exists[p] == nums[k] {
						break
					}
				}
				if p == len(exists) {
					arrs = append(arrs, arr)
					exists = append(exists, nums[k])
				}
				k--
				j++
			}
		}
	}
	return arrs
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
