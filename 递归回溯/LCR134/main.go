package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var dfs func(x float64, n int) float64
	dfs = func(x float64, n int) float64 {
		if n == 1 {
			return x
		} else if n == -1 {
			return 1 / x
		}
		ret := float64(1)
		if n%2 == 1 {
			n--
			ret = x
		} else if n%2 == -1 {
			n--
			ret = x
		}
		v := dfs(x, n/2)
		ret *= v * v
		return ret
	}
	return dfs(x, n)
}

func main() {
	fmt.Println(myPow(2, 10))
}
