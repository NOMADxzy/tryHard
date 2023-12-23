package main

import "fmt"

func minDeletionSize(strs []string) int {
	n := len(strs[0])
	m := len(strs)
	//maybeDel := make([]bool, n)
	strict := make([]bool, m-1)
	del := 0
loop:
	for i := 0; i < n; i++ {
		newStrict := make([]bool, m-1)
		cnt := 0
		for j := 0; j < m-1; j++ {
			if strs[j][i] < strs[j+1][i] || strict[j] {
				newStrict[j] = true
				cnt++
			} else if strs[j][i] > strs[j+1][i] {
				del++
				continue loop
			}
		}
		if cnt == m-1 {
			break
		}
		strict = newStrict
	}
	return del
}

func main() {
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}
