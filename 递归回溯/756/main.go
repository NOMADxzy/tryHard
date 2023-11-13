package main

import "fmt"

func findChoices(res *[]string, bMap map[string][]byte, bottom string, cur []byte, pos int) {
	if pos == len(cur) {
		*res = append(*res, string(cur))
		return
	}
	for _, b := range bMap[bottom[pos:pos+2]] {
		cur[pos] = b
		findChoices(res, bMap, bottom, cur, pos+1)
	}
}

func findNextRow(bottom string, bMap map[string][]byte) bool {
	if len(bottom) == 1 {
		return true
	}
	var choices []string
	findChoices(&choices, bMap, bottom, make([]byte, len(bottom)-1), 0)
	for _, nextRow := range choices {
		if findNextRow(nextRow, bMap) {
			return true
		}
	}
	return false
}

func pyramidTransition(bottom string, allowed []string) bool {
	bMap := map[string][]byte{}
	for _, s := range allowed {
		bMap[s[:2]] = append(bMap[s[:2]], s[2])
	}
	return findNextRow(bottom, bMap)
}

func main() {
	allowed := []string{"AAB", "AAC", "BCD", "BBE", "DEF"}
	fmt.Println(pyramidTransition("AAAA", allowed))
}
