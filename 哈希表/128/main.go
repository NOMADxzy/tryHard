package main

import "fmt"

func longestConsecutive(nums []int) int {
	numsMap := map[int]bool{}
	for _, num := range nums {
		numsMap[num] = true
	}

	longest := 0

	for num := range numsMap {
		if !numsMap[num-1] {

			l := 1
			for n := num + 1; ; n++ {
				if numsMap[n] {
					l += 1
				} else {
					break
				}
			}
			if longest < l {
				longest = l
			}
		}
	}
	return longest

}

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
