package main

import (
	"fmt"
	"sort"
)

func lastNonEmptyString(s string) string {
	type Info struct {
		val       byte
		cnt       int
		lastPlace int
	}
	var arr []*Info
	var left []*Info
	arr = make([]*Info, 26)
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if arr[idx] == nil {
			arr[idx] = &Info{
				s[i],
				1,
				i,
			}
		} else {
			info := arr[idx]
			info.lastPlace = i
			info.cnt++
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != nil {
			left = append(left, arr[i])
		}
	}
	arr = left
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].cnt > arr[j].cnt
	})
	var tmp []*Info
	maxTime := arr[0].cnt
	for i := 0; i < len(arr) && arr[i].cnt == maxTime; i++ {
		tmp = append(tmp, arr[i])
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].lastPlace < tmp[j].lastPlace
	})
	ans := ""
	for _, info := range tmp {
		ans += string(info.val)
	}
	return ans
}

func main() {
	fmt.Println(lastNonEmptyString("aabcbbca"))
}
