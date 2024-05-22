package main

import "fmt"

func oneEditAway(first string, second string) bool { // 思想：从两侧匹配的字母数 >= 较长字符串的长度-1
	if len(first) > len(second) {
		first, second = second, first
	}
	cnts := 0
	leftIdx := 0
	for leftIdx < len(first) && first[leftIdx] == second[leftIdx] {
		leftIdx++
		cnts++
	}
	rightIdx := len(first) - 1
	for i := len(second) - 1; i >= 0 && rightIdx > leftIdx-1 && first[rightIdx] == second[i]; i-- {
		rightIdx--
		cnts++
	}
	return cnts >= len(second)-1
}

func main() {
	fmt.Println(oneEditAway("ab", "ac"))
}
