package main

import "fmt"

func maximumBinaryString(binary string) string {
	cnt0, cnt1 := 0, 0
	n := len(binary)
	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		if binary[i] == '0' {
			cnt0++
		} else {
			cnt1++
		}
		ans[i] = binary[i]
	}
	var i int
	for i = 0; i < n && cnt0 > 1; i++ {
		if ans[i] == '0' {
			ans[i] = '1'
			ans[i+1] = '0'
			cnt0--
		}
	}
	if cnt0 == 1 {
		i++
		for i < n {
			ans[i] = '1'
			i++
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(maximumBinaryString("000110"))
}
