package main

import "fmt"

// 分治，无法产生k=n-1的情况 arrange(n, k)
func arrange(cnt, k int) []int {
	if cnt == 1 {
		return []int{1}
	} else if k <= 1 {
		a := make([]int, cnt)
		for i := 0; i < cnt; i++ {
			a[i] = i + 1
		}
		return a
	}
	lcnt, rcnt := cnt/2, cnt/2
	if cnt%2 == 1 {
		lcnt++
	}
	lk := k / 2
	if k%2 != 0 {
		lk++
	}
	rk := k - lk
	arr1, arr2 := arrange(lcnt, lk), arrange(rcnt, rk)
	for i := 0; i < len(arr1); i++ {
		arr1[i] = 2*arr1[i] - 1
	}
	for i := 0; i < len(arr2); i++ {
		arr2[i] = 2 * arr2[i]
		arr1 = append(arr1, arr2[i])
	}
	return arr1
}

// 构造
func constructArray(n int, k int) []int {
	ans := make([]int, n)
	idx := 0
	leftMax := (n + 1) / 2
	left, right := 1, n
	cnt := 0
	for cnt < k {
		if idx%2 == 0 {
			ans[idx] = left
			left++
		} else {
			ans[idx] = right
			right--
		}
		cnt++
		idx++
	}
	if idx%2 == 0 {
		for right > leftMax {
			ans[idx] = right
			right--
			idx++
		}
		for i := leftMax; i >= left; i-- {
			ans[idx] = i
			idx++
		}
	} else {
		for left <= leftMax {
			ans[idx] = left
			left++
			idx++
		}
		for i := leftMax + 1; i <= right; i++ {
			ans[idx] = i
			idx++
		}
	}
	return ans
}

func main() {
	fmt.Println(constructArray(5, 4))
}
