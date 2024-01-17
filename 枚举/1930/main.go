package main

import "fmt"

func countPalindromicSubsequence(s string) int {
	idxs := make([][]int, 26)
	for i := 0; i < len(s); i++ {
		id := int(s[i] - 'a')
		idxs[id] = append(idxs[id], i)
	}
	ans := 0
	for i := 0; i < 26; i++ {
		if len(idxs[i]) == 0 {
			continue
		}
		cur := idxs[i]
		//if len(cur)>2{
		//	ans ++
		//}
		for j := 0; j < 26; j++ {
			if len(idxs[j]) < 2 || idxs[j][0] > cur[len(cur)-1] || idxs[j][len(idxs[j])-1] < cur[0] {
				continue
			} else {
				targetL, targetR := idxs[j][0], idxs[j][len(idxs[j])-1]
				left, right := 0, len(cur)-1
				if cur[left] > targetL && cur[left] < targetR || cur[right] > targetL && cur[right] < targetR {
					ans++
					continue
				}
				for left < right {
					mid := (left + right) / 2
					if cur[mid] > targetL && cur[mid] < targetR {
						ans++
						break
					} else if cur[mid] <= targetL {
						left = mid + 1
					} else {
						right = mid
					}
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countPalindromicSubsequence("tlpjzdmtwderpkpmgoyrcxttiheassztncqvnfjeyxxp"))
}
