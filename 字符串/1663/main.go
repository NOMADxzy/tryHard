package main

import (
	"fmt"
	"strconv"
)

//递归超时
func jumpNext(target int, pos int, size int, cache map[string]bool) []int {
	keyString := strconv.Itoa(target) + "_" + strconv.Itoa(pos)
	if cache[keyString] {
		return nil
	}

	if target > 26*(size-pos) {
		return nil
	}
	if pos == size-1 {
		return []int{target}
	}

	for i := 1; i <= 26; i++ {
		tail := jumpNext(target-i, pos+1, size, cache)
		if tail != nil {
			return append(tail, i)
		}
	}
	cache[keyString] = true
	return nil
}

func getSmallestString_(n int, k int) string {
	cache := map[string]bool{}
	arr := jumpNext(k, 0, n, cache)
	str := ""
	for i := len(arr) - 1; i >= 0; i-- {
		str += string('a' + arr[i] - 1)
	}
	return str
}

//二分查找, 超时, 法二
func getSmallestString(n int, k int) string {
	left := 0
	right := n
	ans := make([]byte, n)

	remain := n - left
	if remain < (k-left)/26 || k < n {
		return ""
	}
	if k == n {
		for i := 0; i < n; i++ {
			ans[i] = 'a'
		}
		return string(ans)
	}

	for left < right {
		mid := (left + right) / 2
		if (n-mid)*26 >= (k - mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	countA := left - 1
	countZ := n - left
	mid := k - countA - countZ*26

	idx := 0
	for ; idx < countA; idx++ {
		ans[idx] = 'a'
	}
	ans[idx] = byte('a' + mid - 1)

	for i := idx + 1; i < n; i++ {
		ans[i] = 'z'
	}
	return string(ans)
}

//法3
func getSmallestString3(n int, k int) string {
	ans := make([]byte, n)
	mark := false
	for i := 0; i < n; i++ {
		if mark {
			ans[i] = 'z'
			continue
		}
		if (n-i-1)*26 >= k-i {
			ans[i] = 'a'
		} else if !mark {
			t := k - i - 26*(n-i-1)
			ans[i] = byte('a' + t - 1)
			mark = true
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(getSmallestString(5, 73))
}
