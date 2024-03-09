package main

import "fmt"

// 0000001001
func secondsToRemoveOccurrences(s string) int {
	var times []int
	times = append(times, 0)
	shouldPlace := 0 // 这个1最终应该所在的位置
	//last1Place := 1
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			var timeCur int
			delta := i - shouldPlace //
			if delta == 0 {
				timeCur = delta // 不需要再换了
			} else {
				timeCur = max(times[len(times)-1]+1, delta)
			}

			shouldPlace++
			times = append(times, timeCur)
		}
	}
	return times[len(times)-1]
}

func main() {
	fmt.Println(secondsToRemoveOccurrences("111000"))
}
