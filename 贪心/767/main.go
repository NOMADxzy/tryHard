package main

import "fmt"

func buildHeap(heap [][]int) {
	length := len(heap)
	for i := length/2 - 1; i >= 0; i-- {
		downCheck(heap, i)
	}
}

func downCheck(heap [][]int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	biggest := i
	if left < len(heap) && heap[left][1] > heap[biggest][1] {
		biggest = left
	}
	if right < len(heap) && heap[right][1] > heap[biggest][1] {
		biggest = right
	}
	if biggest != i {
		heap[biggest], heap[i] = heap[i], heap[biggest]
		downCheck(heap, biggest)
	}
}

func reorganizeString(s string) string {
	sMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	if len(sMap) == 1 && len(s) > 1 {
		return ""
	}
	pairs := make([][]int, len(sMap))
	id := 0
	for c, v := range sMap {
		pairs[id] = []int{int(c), v}
		id++
	}
	buildHeap(pairs)
	newS := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if i == 0 || pairs[0][0] != int(newS[i-1]) {
			newS[i] = byte(pairs[0][0])
			pairs[0][1]--
			downCheck(pairs, 0)
		} else {
			idx := 1
			if len(pairs) > 2 && pairs[2][1] > pairs[1][1] {
				idx = 2
			}
			if pairs[idx][1] == 0 {
				return ""
			}
			newS[i] = byte(pairs[idx][0])
			pairs[idx][1]--
			downCheck(pairs, idx)
		}
	}
	return string(newS)
}

func main() {
	fmt.Println(reorganizeString("bbbb"))
}
