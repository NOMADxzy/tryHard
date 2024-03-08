package main

import "fmt"

func addSpaces(s string, spaces []int) string {
	idx1 := 0
	idx2 := 0
	arr := make([]byte, len(s)+len(spaces))
	for i := 0; i < len(arr); i++ {
		if idx2 < len(spaces) && idx1 == spaces[idx2] {
			arr[i] = ' '
			idx2++
		} else {
			arr[i] = s[idx1]
			idx1++
		}
	}
	return string(arr)
}

func main() {
	fmt.Println(addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15}))
}
