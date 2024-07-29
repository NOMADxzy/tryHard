package main

func numberOfWays(s string) int64 {
	m := len(s)
	acc0, acc1 := make([]int, m), make([]int, m)
	if s[0] == '0' {
		acc0[0] = 1
	} else {
		acc1[0] = 1
	}
	for i := 1; i < m; i++ {
		acc0[i] = acc0[i-1]
		acc1[i] = acc1[i-1]
		if s[i] == '0' {
			acc0[i]++
		} else {
			acc1[i]++
		}
	}
	ans := int64(0)
	for i := 1; i < m-1; i++ {
		if s[i] == '0' {
			ans += int64(acc1[i-1]) * int64(acc1[m-1]-acc1[i])
		} else {
			ans += int64(acc0[i-1]) * int64(acc0[m-1]-acc0[i])
		}
	}
	return ans
}

func main() {
	println(numberOfWays("001101"))
}
