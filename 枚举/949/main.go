package main

import (
	"fmt"
	"strconv"
)

func combine(res *[]string, arr []int, mark int, s string, pos int) {
	if pos == 4 {
		*res = append(*res, s)
	}
	mask := 1
	for i := 0; i < len(arr); i++ {
		if mask&mark == 0 {
			combine(res, arr, mark+mask, s+strconv.Itoa(arr[i]), pos+1)
		}
		mask *= 2
	}
}

func largestTimeFromDigits(arr []int) string {
	var res []string

	combine(&res, arr, 0, "", 0)

	s := "0000"
	flag := false
	for i := 0; i < len(res); i++ {
		r := res[i]
		hour, _ := strconv.Atoi(r[0:2])
		minute, _ := strconv.Atoi(r[2:])
		if hour < 24 && minute < 59 && r >= s {
			s = r
			flag = true
		}
	}
	if !flag {
		return ""
	}

	return s[0:2] + ":" + s[2:]
}

func main() {
	fmt.Println(largestTimeFromDigits([]int{0, 0, 0, 0}))
}
