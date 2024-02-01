package main

import "fmt"

// 根据贪心思想，优先选择靠前匹配上的字符；记录s1第
// 一个匹配上的idx，存在某一位置i，区间[idx, i]长度是len(s1)的倍数，s2在其中出现了m次，这个区间就是循环结
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	map2 := map[byte]bool{}
	for i := 0; i < len(s2); i++ {
		map2[s2[i]] = true
	}
	tmp := ""
	for i := 0; i < len(s1); i++ {
		if map2[s1[i]] {
			tmp += string(s1[i])
		}
	}
	s1 = tmp
	var i, j int
	firstIdx := -1
	cnt2 := 0
	for i < len(s1)*n1 {
		if s1[(i+len(s1))%len(s1)] == s2[j] {
			if firstIdx == -1 {
				firstIdx = i
			}
			j++
			if j == len(s2) {
				j = 0
				cnt2 += 1
			}
		}
		if j == 0 && firstIdx >= 0 && (i-firstIdx+1)%len(s1) == 0 {
			break
		}
		i++
	}
	if cnt2 == 0 {
		return 0
	}
	var nodeNum int
	if i == len(s1)*n1 {
		i--
		j = 0
	} else {
		nodeNum = (n1*len(s1) - firstIdx) / (i - firstIdx + 1)
	}
	remainLen := n1*len(s1) - firstIdx - nodeNum*(i-firstIdx+1)
	ans := nodeNum * cnt2 / n2
	remaincnt2 := nodeNum*cnt2 - ans*n2
	for k := firstIdx; k < firstIdx+remainLen; k++ {
		trueIdx := (k + len(s1)) % len(s1)
		if s1[trueIdx] == s2[j] {
			j++
		}
		if j == len(s2) {
			j = 0
			remaincnt2++
		}
	}
	ans += remaincnt2 / n2
	return ans
}

func main() {
	fmt.Println(getMaxRepetitions("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 1000000, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 1000000))
}
