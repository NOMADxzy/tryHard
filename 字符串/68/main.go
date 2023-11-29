package main

import "fmt"

func fullJustify(words []string, maxWidth int) []string {
	m := len(words)
	var spaces [][]int
	var cnts []int

	cur := 0
	for cur < m {
		right := cur + 1
		sumLen := len(words[cur])
		for right < m && sumLen+len(words[right])+right-cur <= maxWidth {
			sumLen += len(words[right])
			right++
		}
		var sp []int
		if right < m {
			numSpaces := right - cur - 1
			if numSpaces == 0 {
				numSpaces = 1
			}
			sp = make([]int, numSpaces)
			more := (maxWidth - sumLen) % numSpaces
			for i := 0; i < numSpaces; i++ {
				sp[i] = (maxWidth - sumLen) / numSpaces
				if i < more {
					sp[i]++
				}
			}
		} else {
			numSpaces := right - cur
			sp = make([]int, numSpaces)
			for i := 0; i < numSpaces; i++ {
				sp[i] = 1
				if i == numSpaces-1 {
					sp[i] = maxWidth - sumLen - numSpaces + 1
				}
			}
		}
		spaces = append(spaces, sp)
		cnts = append(cnts, right-cur)
		cur = right
	}
	ans := make([]string, len(spaces))
	wordIdx := 0
	for i, sp := range spaces {
		str := ""
		for j := 0; j < cnts[i]; j++ {
			str += words[wordIdx]
			if j < len(sp) {
				for k := 0; k < sp[j]; k++ {
					str += " "
				}
			}
			wordIdx++
		}
		ans[i] = str
	}
	return ans
}

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
}
