package main

import (
	"fmt"
	"strconv"
)

func cptVal(w int, c int64) int64 {
	r := int64(0)
	for i := 0; i < w; i++ {
		if 1<<62/c < r { //溢出了，返回个很大的值即可
			return 1 << 62
		}
		r = c*r + 1
	}
	return r
}

func smallestGoodBase(n string) string {
	target := int64(0)
	for i := 0; i < len(n); i++ {
		target = 10*target + int64(n[i]-'0')
	}
	ans := int64(0)
	for w := 61; w >= 1; w-- {
		minVal := cptVal(w, 2)
		if minVal > target {
			continue
		} else if minVal == target {
			return "2"
		} else {
			left, right := int64(2), target
			for left < right {
				mid := (left + right) / 2
				val := cptVal(w, mid)
				if val < target {
					left = mid + 1
				} else {
					right = mid
				}
			}
			if cptVal(w, left) == target {
				ans = left
				break
			}
		}
	}
	ans_str := ""
	for ans > 0 {
		ans_str = strconv.Itoa(int(ans%10)) + ans_str
		ans /= 10
	}
	return ans_str
}

func main() {
	fmt.Println(smallestGoodBase("194148123592167400"))
	//fmt.Println(1<<62 > int64(194148123592167400))
	//a := int64(579043)
	//b := int64(194148123592167400)
	//x := int64(0)
	//w := 0
	//for b > x {
	//	x = a*x + 1
	//	w++
	//}
	//fmt.Println(w)
}

//2401 + 343 + 49 + 7 + 1
