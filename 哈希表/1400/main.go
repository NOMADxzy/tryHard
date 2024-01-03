package main

import "fmt"

func canConstruct(s string, k int) bool {
	cnts := map[byte]int{}
	for i := 0; i < len(s); i++ {
		cnts[s[i]]++
	}
	remain := 0
	for _, cnt := range cnts {
		remain += cnt % 2
	}
	return remain <= k && len(s) >= k
}

func main() {
	fmt.Println(canConstruct("annabelle", 30))
}
