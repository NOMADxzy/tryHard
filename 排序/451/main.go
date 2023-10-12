package main

import "fmt"

func frequencySort(s string) string {
	ans := make([]byte, len(s))

	fMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		fMap[s[i]]++
	}
	maxFreq := 0
	for _, f := range fMap {
		if f > maxFreq {
			maxFreq = f
		}
	}
	fArr := make([][]byte, maxFreq+1)
	for c, f := range fMap {
		fArr[f] = append(fArr[f], c)
	}

	idx := 0
	for i := maxFreq; i > 0; i-- {
		if bkt := fArr[i]; bkt != nil {
			for _, c := range bkt {
				for k := 0; k < i; k++ {
					ans[idx] = c
					idx++
				}
			}
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(frequencySort("tree"))
}
