package main

import "fmt"

func removeDuplicates(s string, k int) string {
	var cstack []byte
	var vstack []int
	for i := 0; i < len(s); i++ {
		if len(cstack) > 0 && s[i] == cstack[len(cstack)-1] {
			if vstack[len(vstack)-1] == k-1 {
				cstack = cstack[:len(cstack)-k+1]
				vstack = vstack[:len(vstack)-k+1]
			} else {
				cstack = append(cstack, s[i])
				vstack = append(vstack, vstack[len(vstack)-1]+1)
			}
		} else {
			cstack = append(cstack, s[i])
			vstack = append(vstack, 1)
		}
	}
	return string(cstack)
}

func main() {
	fmt.Println(removeDuplicates("pbbcggttciiippooaais", 2))
}
