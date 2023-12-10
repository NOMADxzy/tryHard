package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findMultiVal(num string) ([]string, []int) {
	curVal, _ := strconv.Atoi(num)
	if len(num) < 2 {
		return []string{num}, []int{curVal}
	}
	var strs []string
	var vals []int
	strs = append(strs, num)
	vals = append(vals, curVal)
	for l := 1; l < len(num); l++ {
		val, _ := strconv.Atoi(num[:l])
		str := num[:l] + "*"
		nextStrs, nextVals := findMultiVal(num[l:])
		for i := 0; i < len(nextVals); i++ {
			strs = append(strs, str+nextStrs[i])
			vals = append(vals, val*nextVals[i])
		}
	}
	return strs, vals
}

func findNext(num string, target int, isNeg bool) []string {
	var ans []string
	if val, ok := strconv.Atoi(num); ok == nil {
		if isNeg {
			val = -val
		}
		if val == target {
			ans = append(ans, num)
		}
	}
	for l := 1; l <= len(num); l++ {
		strs, vals := findMultiVal(num[:l])
		for i := 0; i < len(strs); i++ {
			val, str := vals[i], strs[i]
			if isNeg {
				val = -val
			}
			if l == len(num) {
				if val == target {
					ans = append(ans, str)
				}
			} else {
				for _, symb := range []string{"+", "-"} {
					newTarget := target - val
					curAns := str + symb
					nextAnses := findNext(num[l:], newTarget, symb == "-")
					for _, nextAns := range nextAnses {
						ans = append(ans, curAns+nextAns)
					}
				}
			}

		}
	}
	return ans
}

func addOperators(num string, target int) []string {
	ans0 := findNext(num, target, false)
	var ans []string
	resMap := map[string]bool{}
	for _, a := range ans0 {
		ok := true
		if !resMap[a] {
			na := make([]byte, len(a))
			for i := 0; i < len(a); i++ {
				if a[i] >= '0' && a[i] <= '9' {
					na[i] = a[i]
				} else {
					na[i] = '#'
				}
			}
			sps := strings.Split(string(na), "#")
			for _, sp := range sps {
				if len(sp) > 1 && sp[0] == '0' {
					ok = false
				}
			}
			if ok {
				ans = append(ans, a)
			}
		}
		resMap[a] = true
	}
	return ans
}

func addOperators1(num string, target int) (ans []string) {
	n := len(num)
	var backtrack func(expr []byte, i, res, mul int)
	backtrack = func(expr []byte, i, res, mul int) {
		if i == n {
			if res == target {
				ans = append(ans, string(expr))
			}
			return
		}
		signIndex := len(expr)
		if i > 0 {
			expr = append(expr, 0) // 占位，下面填充符号
		}
		// 枚举截取的数字长度（取多少位），注意数字可以是单个 0 但不能有前导零
		for j, val := i, 0; j < n && (j == i || num[i] != '0'); j++ {
			val = val*10 + int(num[j]-'0')
			expr = append(expr, num[j])
			if i == 0 { // 表达式开头不能添加符号
				backtrack(expr, j+1, val, val)
			} else { // 枚举符号
				expr[signIndex] = '+'
				backtrack(expr, j+1, res+val, val)
				expr[signIndex] = '-'
				backtrack(expr, j+1, res-val, -val)
				expr[signIndex] = '*'
				backtrack(expr, j+1, res-mul+mul*val, mul*val)
			}
		}
	}
	backtrack(make([]byte, 0, n*2-1), 0, 0, 0)
	return
}

func main() {
	strs1 := addOperators("123456789", 45)
	strs2 := addOperators1("123456789", 45)
	map1, map2 := map[string]bool{}, map[string]bool{}
	for _, str := range strs1 {
		map1[str] = true
	}
	for _, str := range strs2 {
		map2[str] = true
	}
	for _, str := range strs1 {
		if map2[str] {
			delete(map2, str)
		} else {
			fmt.Println("wrong")
		}
	}
	fmt.Println("ok")
}
