package main

import "fmt"

func minFlips(target string) int {
	ans := 0
	//cur := 0
	for i := 0; i < len(target); i++ {
		if target[i] == '0' {
			if ans%2 == 1 {
				//cur ++
				ans++
			}
		} else if ans%2 == 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(minFlips("10111"))
}
