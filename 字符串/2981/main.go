package main

import "fmt"

func maximumLength(s string) int {
	type Info struct {
		Len  int
		Freq int
	}
	push := func(infos []*Info, size int) {
		for i := 0; i < 3; i++ {
			if infos[i] == nil {
				infos[i] = &Info{size, 1}
				return
			} else if infos[i].Len == size {
				infos[i].Freq++
				return
			} else if infos[i].Len < size {
				newInfo := &Info{size, 1}
				for j := 2; j > i; j-- {
					infos[j] = infos[j-1]
				}
				infos[i] = newInfo
				return
			}
		}
	}
	cptMax := func(infos []*Info) int {
		if infos[0] == nil {
			return -1
		}
		l := infos[0].Len
		f := infos[0].Freq
		sum := 0
		res := -1
		for i := l; i > l-3 && i > 0; i-- {
			sum = f * (l - i + 1)
			if sum >= 3 {
				res = i
				break
			}
		}
		if infos[1] != nil {
			res = max(res, infos[1].Len)
		}
		return res
	}

	infos := make([][]*Info, 26)
	for i := 0; i < 26; i++ {
		infos[i] = make([]*Info, 3)
	}
	for i := 0; i < 26; i++ {
		b := byte('a' + i)
		l, r := 0, 0
		for r < len(s) {
			for l < len(s) && s[l] != b {
				l++
			}
			if l == len(s) {
				break
			}
			r = l
			for r < len(s) && s[r] == b {
				r++
			}
			delta := r - l
			push(infos[i], delta)
			l = r
		}

	}
	maxVal := -1
	for i := 0; i < 26; i++ {
		for j := 0; j < 3; j++ {
			maxVal = max(maxVal, cptMax(infos[i]))
		}
	}
	return maxVal
}

func main() {
	fmt.Println(maximumLength("abcccabacc"))
}
