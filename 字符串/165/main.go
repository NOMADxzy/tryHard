package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	strs1 := strings.Split(version1, ".")
	strs2 := strings.Split(version2, ".")

	for i := 0; i < len(strs2) || i < len(strs1); i++ {
		cur1, cur2 := "0", "0"
		if i < len(strs1) {
			cur1 = strs1[i]
			for len(cur1) > 1 && cur1[0] == '0' {
				cur1 = cur1[1:]
			}
		}
		if i < len(strs2) {
			cur2 = strs2[i]
			for len(cur2) > 1 && cur2[0] == '0' {
				cur2 = cur2[1:]
			}
		}
		val1, _ := strconv.Atoi(cur1)
		val2, _ := strconv.Atoi(cur2)
		if val1 < val2 {
			return -1
		} else if val1 > val2 {
			return 1
		}
	}
	return 0
}

func main() {
	fmt.Println(compareVersion("1", "1.001"))
}
