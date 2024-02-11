package main

import "fmt"

func verifyTreeOrder(postorder []int) bool {
	var judge func(arr []int) bool
	judge = func(arr []int) bool {
		if len(arr) <= 1 {
			return true
		}
		val := arr[len(arr)-1]
		mid := -1
		for mid+1 < len(arr) && arr[mid+1] < val {
			mid++
		}
		for i := mid + 1; i < len(arr)-1; i++ {
			if arr[i] <= val {
				return false
			}
		}
		return judge(arr[:mid+1]) && judge(arr[mid+1:len(arr)-1])
	}
	return judge(postorder)
}

func main() {
	fmt.Println(verifyTreeOrder([]int{8, 6, 9}))
}
