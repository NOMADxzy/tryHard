package main

import "fmt"

// TODO 注意不要数组越界访问，利用前期结果节省时间
func jumpAll(nums []int, pos int) bool {
	if pos >= len(nums)-1 {
		return true
	}
	if nums[pos] == -1 {
		return false
	}
	for i := pos + nums[pos]; i > pos; i-- {
		if jumpAll(nums, i) {
			return true
		} else {
			nums[i] = -1 //拉黑这个位置，提高效率
		}
	}
	return false
}

func canJump(nums []int) bool {
	return jumpAll(nums, 0)
}

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
