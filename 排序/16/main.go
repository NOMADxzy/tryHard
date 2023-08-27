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

func threeSumClosest(nums []int, target int) int {
	quickSort(nums)
	size := len(nums)
	var delta int
	dist := nums[0] + nums[1] + nums[len(nums)-1] - target

	for i := 0; i < size-2; i++ {
		targetBC := target - nums[i]
		j, k := i+1, size-1

		for j < k {
			if nums[j]+nums[k] > targetBC {
				delta = nums[j] + nums[k] - targetBC
				if dist < 0 && delta < -dist || dist > 0 && delta < dist {
					dist = delta
				}
				k--
			} else if nums[j]+nums[k] < targetBC {
				delta = nums[j] + nums[k] - targetBC
				if dist < 0 && delta > dist || dist > 0 && -delta < dist {
					dist = delta
				}
				j++
			} else {
				return target
			}
		}
	}
	return target + dist
}

func main() {
	fmt.Println(threeSumClosest([]int{0, 0, 4, 0}, 1))
}
