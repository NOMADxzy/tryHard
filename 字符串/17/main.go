package main

import "fmt"

func backTrack(strs *[]string, cur string, digits string, pos int, phoneMap map[byte][]byte) {
	if pos == len(digits) {
		*strs = append(*strs, cur)
	} else {
		elems := phoneMap[digits[pos]]
		for i := 0; i < len(elems); i++ {
			backTrack(strs, cur+string(elems[i]), digits, pos+1, phoneMap)
		}
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	phoneMap := map[byte][]byte{}
	var strs []string
	//strs = append(strs, "")

	for i := 2; i <= 7; i++ {
		k := i - 2
		phoneMap['0'+byte(i)] = []byte{'a' + byte(k*3), 'a' + byte(k*3+1), 'a' + byte(k*3+2)}
	}
	phoneMap['7'] = append(phoneMap['7'], 's')
	phoneMap['8'] = []byte{'t', 'u', 'v'}
	phoneMap['9'] = []byte{'w', 'x', 'y', 'z'}

	backTrack(&strs, "", digits, 0, phoneMap)

	return strs
}

func main() {
	fmt.Println(letterCombinations(""))
}
