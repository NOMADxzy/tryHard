package main

import "fmt"

//暴力递归，超时
func dfs(count []int, pos int, numToChar [][]byte, curMap, kMap map[byte]int, cnt, total int) bool {
	if pos == len(count) {
		return cnt == total
	}

	var c int
	for c = 0; ; c++ {
		if dfs(count, pos+1, numToChar, curMap, kMap, cnt, total) {
			count[pos] = c
			return true
		}
		chars := numToChar[pos]
		var i int
		for i = 0; i < len(chars); i++ {
			curMap[chars[i]]++
			if curMap[chars[i]] > kMap[chars[i]] {
				break
			}
		}
		if i < len(chars) {
			for ; i >= 0; i-- {
				curMap[chars[i]]--
			}
			break
		} else {
			cnt += len(chars)
		}
	}
	chars := numToChar[pos]
	for i := 0; i < len(chars); i++ {
		curMap[chars[i]] -= c
	}
	return false
}

func originalDigits(s string) string {
	kMap := map[byte]int{}
	numToChar := [][]byte{{'z', 'e', 'r', 'o'}, {'o', 'n', 'e'}, {'t', 'w', 'o'}, {'t', 'h', 'r', 'e', 'e'}, {'f', 'o', 'u', 'r'}, {'f', 'i', 'v', 'e'},
		{'s', 'i', 'x'}, {'s', 'e', 'v', 'e', 'n'}, {'e', 'i', 'g', 'h', 't'}, {'n', 'i', 'n', 'e'}}
	count := make([]int, 10)
	curMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		kMap[s[i]]++
	}
	b := dfs(count, 0, numToChar, curMap, kMap, 0, len(s))
	fmt.Println(b)
	str := ""
	for i := 0; i < len(count); i++ {
		c := count[i]
		b := '0' + byte(i)
		for j := 0; j < c; j++ {
			str += string(b)
		}
	}
	return str
}

func originalDigits1(s string) string {
	kMap := map[byte]int{}
	count := make([]int, 10)
	for i := 0; i < len(s); i++ {
		kMap[s[i]]++
	}
	count[0] = kMap['z']
	count[2] = kMap['w']
	count[4] = kMap['u']
	count[6] = kMap['x']
	count[8] = kMap['g']

	count[3] = kMap['h'] - count[8]
	count[5] = kMap['f'] - count[4]
	count[7] = kMap['s'] - count[6]

	count[1] = kMap['o'] - count[0] - count[2] - count[4]
	count[9] = kMap['i'] - count[5] - count[6] - count[8]

	str := ""
	for i := 0; i < len(count); i++ {
		c := count[i]
		b := '0' + byte(i)
		for j := 0; j < c; j++ {
			str += string(b)
		}
	}
	return str
}

func main() {
	fmt.Println(originalDigits1("fviefurozerozerofuronineeithg"))
}
