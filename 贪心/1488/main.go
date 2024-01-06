package main

import "fmt"

func avoidFlood(rains []int) []int {
	var options []int
	ans := make([]int, len(rains))
	hMap := map[int]int{}
	for i := 0; i < len(rains); i++ {
		if rains[i] == 0 {
			options = append(options, i)
			continue
		} else if _, ok := hMap[rains[i]]; !ok {
			hMap[rains[i]] = i
		} else if len(options) == 0 {
			return []int{}
		} else {
			R := hMap[rains[i]]
			left, right := 0, len(options)-1
			if options[left] > R {
				ans[options[left]] = rains[i]
				hMap[rains[i]] = i
				options = options[1:]
			} else if options[right] < R {
				return []int{}
			} else {
				for left < right {
					mid := (left + right) / 2
					if options[mid] < R {
						left = mid + 1
					} else {
						right = mid
					}
				}
				ans[options[right]] = rains[i]
				hMap[rains[i]] = i
				copy(options[right:], options[right+1:])
				options = options[:len(options)-1]
			}
		}
		ans[i] = -1
	}
	for _, idx := range options {
		ans[idx] = 1
	}
	return ans
}

func main() {
	fmt.Println(avoidFlood([]int{2, 3, 0, 0, 3, 1, 0, 1, 0, 2, 2}))
}

// 1 0 0 2 3 0 2 3 1
