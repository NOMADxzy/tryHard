package main

import "fmt"

func maxSumDivThree(nums []int) int {
	m := len(nums)
	states := make([]int, 3)

	states[nums[0]%3] = nums[0]

	for i := 1; i < m; i++ {
		newStates := make([]int, 3)
		for j := 0; j < 3; j++ {
			idx := (nums[i]%3 + j) % 3
			if states[j] == 0 {
				newStates[idx] = states[idx]
				idx = nums[i] % 3
			}

			if states[j]+nums[i] > states[idx] && states[j]+nums[i] > newStates[idx] {
				newStates[idx] = states[j] + nums[i]
			} else if states[j]+nums[i] > newStates[idx] {
				newStates[idx] = states[idx]
			}
		}

		states = newStates
	}
	return states[0]
}

func main() {
	fmt.Println(maxSumDivThree([]int{3}))
}
