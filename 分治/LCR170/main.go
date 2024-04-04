package main

import "fmt"

func mSort(arr []int) ([]int, int) { // 分治思想
	if len(arr) <= 1 { // 只要长度超过1就继续分
		return arr, 0
	}
	mid := len(arr) / 2
	left, right := arr[:mid], arr[mid:]
	arr1, cnt1 := mSort(left)  // i,j都在左边的逆序对数量
	arr2, cnt2 := mSort(right) // i,j都在右边的逆序对数量

	newArr, cnt3 := merge(arr1, arr2) // i在左边,j在右边的逆序对数量
	return newArr, cnt1 + cnt2 + cnt3
}

func merge(left, right []int) (newArr []int, count int) {
	m, n := len(left), len(right)
	newArr = make([]int, m+n)
	var i, j int
	for i < m && j < n {
		if left[i] <= right[j] { // right[j] 没有 对应的 逆序对
			newArr[i+j] = left[i]
			i++
		} else {
			count += m - i // right[j]与每一个left[i] 都形成一个 逆序对
			newArr[i+j] = right[j]
			j++
		}
	}
	for i < m { // 剩余元素加到末尾
		newArr[i+j] = left[i]
		i++
	}
	for j < n {
		newArr[i+j] = right[j]
		j++
	}
	return newArr, count
}

func reversePairs(record []int) int {
	_, cnt := mSort(record)
	return cnt
}

func main() {
	fmt.Println(reversePairs([]int{4, 3, 2, 1}))
}
