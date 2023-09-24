package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	left, right := 0, 0
	s1Map := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]]++
	}
	tmpMap := map[byte]int{}
	cnt := 0

	for right < len(s2) {
		b := s2[right]
		tmpMap[b]++
		cnt++
		if tmpMap[b] > s1Map[b] {
			for s2[left] != b {
				tmpMap[s2[left]]--
				left++
				cnt--
			}
			tmpMap[b]--
			left++
			cnt--
		}
		if cnt == len(s1) {
			return true
		}
		right++
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("abi", "eidbaooo"))
}
