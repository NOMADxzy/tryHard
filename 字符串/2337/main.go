package main

import "fmt"

func canChange(start string, target string) bool {
	size := len(start)

	i, j := 0, 0

	for {
		for i < size && start[i] == '_' {
			i++
		}
		for j < size && target[j] == '_' {
			j++
		}
		if i < size && j < size {
			if start[i] == 'L' {
				if target[j] != 'L' {
					return false
				} else if j > i {
					return false
				} else {
					i++
					j++
				}
			} else {
				if target[j] != 'R' {
					return false
				} else if j < i {
					return false
				} else {
					i++
					j++
				}
			}
		} else if i == size {
			for ; j < size; j++ {
				if target[j] != '_' {
					return false
				}
			}
			return true
		} else if j == size {
			for ; i < size; i++ {
				if start[i] != '_' {
					return false
				}
			}
			return true
		} else {
			return true
		}
	}
}

func main() {
	fmt.Println(canChange("_R", "R_"))
}
