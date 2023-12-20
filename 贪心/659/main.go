package main

import "fmt"

func upCheck(arr []int, i int) {
	f := (i - 1) / 2
	if f >= 0 && arr[i] < arr[f] {
		arr[i], arr[f] = arr[f], arr[i]
		upCheck(arr, f)
	}
}

func downCheck(arr []int, i int) {
	left, right := 2*i+1, 2*i+2
	smallest := i
	if left < len(arr) && arr[left] < arr[smallest] {
		smallest = left
	}
	if right < len(arr) && arr[right] < arr[smallest] {
		smallest = right
	}
	if smallest != i {
		arr[smallest], arr[i] = arr[i], arr[smallest]
		downCheck(arr, smallest)
	}
}

func isPossible(nums []int) bool {
	tMap := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		t := nums[i]
		cnt := 1
		if len(tMap[t]) > 0 {
			cnt += tMap[t][0]
			tMap[t][0] = tMap[t][len(tMap[t])-1]
			tMap[t] = tMap[t][:len(tMap[t])-1]
			downCheck(tMap[t], 0)
		}
		tMap[t+1] = append(tMap[t+1], cnt)
		upCheck(tMap[t+1], len(tMap[t+1])-1)
	}
	for _, q := range tMap {
		for _, cnt := range q {
			if cnt < 3 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 5}))
}
